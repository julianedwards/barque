package operations

import (
	"context"

	"github.com/evergreen-ci/barque"
	"github.com/mongodb/grip"
	"github.com/urfave/cli"
)

func Service() cli.Command {
	return cli.Command{
		Name:  "service",
		Usage: "run the barque service",
		Flags: mergeFlags(baseFlags(), dbFlags(
			cli.StringFlag{},
		)),
		Action: func(c *cli.Context) error {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			conf := &barque.Configuration{}

			env, err := barque.NewEnvironment(ctx, conf)
			grip.EmergencyFatal(err)
			barque.SetEnvironment(env)

			return nil
		},
	}
}
