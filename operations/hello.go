package operations

import (
	"github.com/mongodb/grip"
	"github.com/urfave/cli"
)

func Hello() cli.Command {
	return cli.Command{
		Name: "hello",
		Action: func(c *cli.Context) error {
			grip.Info("hello world!")
			return nil
		},
	}
}
