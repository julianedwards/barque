package operations

import (
	"context"
	"time"

	"github.com/evergreen-ci/barque"
	"github.com/evergreen-ci/barque/model"
	"github.com/mongodb/grip"
	"github.com/mongodb/grip/level"
	"github.com/mongodb/grip/send"
	"github.com/pkg/errors"
)

const (
	loggingBufferCount    = 100
	loggingBufferDuration = 20 * time.Second
)

func setupLogging(ctx context.Context, env barque.Environment) error {
	conf, err := model.FindConfiguration(ctx, env)
	if err != nil {
		return errors.Wrap(err, "problem finding configuration")
	}

	var senders []send.Sender

	logLevel := grip.GetSender().Level()

	fallback, err := send.NewErrorLogger("barque.error", logLevel)
	if err != nil {
		return errors.Wrap(err, "problem configuring err fallback logger")
	}

	sender, err := send.MakeDefaultSystem()
	if err != nil {
		return errors.WithStack(err)
	}

	senders = append(senders, sender)

	if conf.Splunk.Populated() {
		sender, err = send.NewSplunkLogger("barque", conf.Splunk, logLevel)
		if err != nil {
			return errors.Wrap(err, "problem building splunk logger")
		}
		if err = sender.SetErrorHandler(send.ErrorHandlerFromSender(fallback)); err != nil {
			return errors.Wrap(err, "problem configuring error handler")
		}

		senders = append(senders, send.NewBufferedSender(sender, loggingBufferDuration, loggingBufferCount))
	}
	if conf.Slack.Options != nil {
		if err = conf.Slack.Options.Validate(); err != nil {
			return errors.Wrap(err, "non-nil slack configuration is not valid")
		}

		if conf.Slack.Token == "" || conf.Slack.Level == "" {
			return errors.Wrap(err, "must specify slack token and threshold")
		}

		lvl := send.LevelInfo{
			Default:   logLevel.Default,
			Threshold: level.FromString(conf.Slack.Level),
		}

		sender, err = send.NewSlackLogger(conf.Slack.Options, conf.Slack.Token, lvl)
		if err != nil {
			return errors.Wrap(err, "problem constructing slack alert logger")
		}
		if err = sender.SetErrorHandler(send.ErrorHandlerFromSender(fallback)); err != nil {
			return errors.Wrap(err, "problem configuring error handler")
		}

		senders = append(senders, send.NewBufferedSender(sender, loggingBufferDuration, loggingBufferCount))
	}

	return errors.WithStack(grip.SetSender(send.NewConfiguredMultiSender(senders...)))
}
