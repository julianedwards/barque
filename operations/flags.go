package operations

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

////////////////////////////////////////////////////////////////////////
//
// Flag Name Constants

const (
	configFlag     = "config"
	numWorkersFlag = "workers"
	dbURIFlag      = "dbUri"
	dbNameFlag     = "dbName"

	flagNameflag = "flag"
)

////////////////////////////////////////////////////////////////////////
//
// Utility Functions

func joinFlagNames(ids ...string) string { return strings.Join(ids, ", ") }

func mergeFlags(in ...[]cli.Flag) []cli.Flag {
	out := []cli.Flag{}

	for idx := range in {
		out = append(out, in[idx]...)
	}

	return out
}

////////////////////////////////////////////////////////////////////////
//
// Flag Groups

func dbFlags(flags ...cli.Flag) []cli.Flag {
	return append(flags,
		cli.StringFlag{
			Name:   dbURIFlag,
			Usage:  "specify a mongodb connection string",
			Value:  "mongodb://localhost:27017",
			EnvVar: "BARQUE_MONGODB_URL",
		},
		cli.StringFlag{
			Name:   dbNameFlag,
			Usage:  "specify a database name to use",
			Value:  "barque",
			EnvVar: "BARQUE_DATABASE_NAME",
		})
}

func addModifyFeatureFlagFlags(flags ...cli.Flag) []cli.Flag {
	return append(flags, cli.StringFlag{
		Name:  flagNameflag,
		Usage: "specify the name of the flag to set",
	})
}

func setFlagOrFirstPositional(name string) cli.BeforeFunc {
	return func(c *cli.Context) error {
		val := c.String(name)
		if val == "" {
			if c.NArg() != 1 {
				return errors.Errorf("must specify exactly one positional argument for '%s'", name)
			}

			val = c.Args().Get(0)
		}

		return c.Set(name, val)
	}
}

func baseFlags(flags ...cli.Flag) []cli.Flag {
	return append(flags,
		cli.IntFlag{
			Name:  numWorkersFlag,
			Usage: "specify the number of worker jobs this process will have",
			Value: 2,
		})
}
