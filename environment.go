package barque

import (
	"context"
	"sync"
	"time"

	"github.com/mongodb/amboy"
	"github.com/mongodb/amboy/reporting"
	"github.com/mongodb/grip"
	"github.com/mongodb/grip/message"
	"github.com/mongodb/jasper"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	envlock   *sync.RWMutex
	globalEnv Environment
)

func init() {
	envlock = &sync.RWMutex{}
	SetEnvironment(&envImpl{name: "init"})
}

func SetEnvironment(env Environment) {
	envlock.Lock()
	defer envlock.Unlock()
	globalEnv = env
}

func GetEnvironment() Environment {
	envlock.RLock()
	defer envlock.RUnlock()
	return globalEnv
}

type Environment interface {
	Context() (context.Context, context.CancelFunc)

	Jasper() jasper.Manager

	Client() *mongo.Client
	DB() *mongo.Database

	LocalQueue() amboy.Queue
	RemoteQueue() amboy.Queue
	QueueGroup() amboy.QueueGroup

	LocalReporter() reporting.Reporter
	RemoteReporter() reporting.Reporter
	GroupReporter() reporting.Reporter

	RegisterCloser(string, bool, CloserFunc)
	Close(context.Context) error
}

type CloserFunc func(context.Context) error

type envImpl struct {
	name           string
	conf           *Configuration
	client         *mongo.Client
	jpm            jasper.Manager
	localQueue     amboy.Queue
	remoteQueue    amboy.Queue
	queueGroup     amboy.QueueGroup
	localReporter  reporting.Reporter
	remoteReporter reporting.Reporter
	groupReporter  reporting.Reporter
	closers        []closerOp
	context        context.Context
	mutex          *sync.RWMutex
}

type closerOp struct {
	name       string
	background bool
	closer     CloserFunc
}

func (e *envImpl) Context() (context.Context, context.CancelFunc) {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	return context.WithCancel(e.context)
}

func (e *envImpl) Jasper() jasper.Manager {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	return e.jpm
}

func (e *envImpl) Client() *mongo.Client {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	return e.client
}

func (e *envImpl) DB() *mongo.Database {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	return e.client.Database(e.conf.DatabaseName)
}

func (e *envImpl) LocalQueue() amboy.Queue {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	return e.localQueue
}

func (e *envImpl) RemoteQueue() amboy.Queue {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	return e.remoteQueue
}

func (e *envImpl) QueueGroup() amboy.QueueGroup {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	return e.queueGroup
}

func (e *envImpl) LocalReporter() reporting.Reporter {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	return e.localReporter
}

func (e *envImpl) RemoteReporter() reporting.Reporter {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	return e.remoteReporter
}

func (e *envImpl) GroupReporter() reporting.Reporter {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	return e.groupReporter
}

func (e *envImpl) RegisterCloser(name string, background bool, fn CloserFunc) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.closers = append(e.closers, closerOp{name: name, background: background, closer: fn})
}

func (e *envImpl) Close(ctx context.Context) error {
	e.RegisterCloser("disconnect-db-client", false, func(ctx context.Context) error {
		return e.client.Disconnect(ctx)
	})

	e.mutex.Lock()
	defer e.mutex.Unlock()

	wg := &sync.WaitGroup{}
	deadline, _ := ctx.Deadline()
	catcher := grip.NewBasicCatcher()
	for idx, closer := range e.closers {
		if closer.background && closer.closer != nil {
			wg.Add(1)
			go func(name string, n int, fn CloserFunc) {
				defer wg.Done()
				grip.Info(message.Fields{
					"message":      "calling closer",
					"index":        n,
					"closer":       name,
					"timeout_secs": time.Until(deadline),
					"deadline":     deadline,
					"background":   true,
				})
				catcher.Add(fn(ctx))
			}(closer.name, idx, closer.closer)
		}
	}

	for idx, closer := range e.closers {
		if !closer.background && closer.closer != nil {
			go func(name string, n int, fn CloserFunc) {
				defer wg.Done()
				grip.Info(message.Fields{
					"message":      "calling closer",
					"index":        n,
					"closer":       name,
					"timeout_secs": time.Until(deadline),
					"deadline":     deadline,
					"background":   false,
				})
				catcher.Add(fn(ctx))
			}(closer.name, idx, closer.closer)
		}
	}
	wg.Wait()
	return catcher.Resolve()
}
