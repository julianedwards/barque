package units

import (
	"context"
	"fmt"
	"time"

	"github.com/evergreen-ci/barque"
	"github.com/evergreen-ci/barque/model"
	"github.com/evergreen-ci/utility"
	"github.com/mongodb/amboy"
	"github.com/mongodb/grip"
	"github.com/pkg/errors"
)

const tsFormat = "2006-01-02.15-04-05"

func StartCrons(ctx context.Context, env barque.Environment) error {
	opts := amboy.QueueOperationConfig{
		ContinueOnError: true,
		LogErrors:       false,
		DebugLogging:    false,
	}

	remote := env.RemoteQueue()
	local := env.LocalQueue()

	amboy.IntervalQueueOperation(ctx, local, time.Minute, time.Now(), opts, func(ctx context.Context, queue amboy.Queue) error {
		conf, err := model.FindConfiguration(ctx, env)
		if err != nil {
			return errors.WithStack(err)
		}

		if conf.Flags.DisableInternalMetricsReporting {
			return nil
		}

		ts := utility.RoundPartOfMinute(0).Format(tsFormat)
		catcher := grip.NewBasicCatcher()
		catcher.Add(queue.Put(ctx, NewSysInfoStatsCollector(fmt.Sprintf("sys-info-stats-%s", ts))))
		catcher.Add(queue.Put(ctx, NewLocalAmboyStatsCollector(env, ts)))
		catcher.Add(queue.Put(ctx, NewJasperManagerCleanup(ts, env)))
		return catcher.Resolve()
	})
	amboy.IntervalQueueOperation(ctx, remote, time.Minute, time.Now(), opts, func(ctx context.Context, queue amboy.Queue) error {
		conf, err := model.FindConfiguration(ctx, env)
		if err != nil {
			return errors.WithStack(err)
		}

		if conf.Flags.DisableInternalMetricsReporting {
			return nil
		}

		return queue.Put(ctx, NewRemoteAmboyStatsCollector(env, utility.RoundPartOfMinute(0).Format(tsFormat)))
	})

	return nil
}
