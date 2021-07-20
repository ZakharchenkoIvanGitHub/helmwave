package action

import "github.com/urfave/cli/v2"

type action interface { //nolint:unused
	Run() error
	Cmd() *cli.Command
}

// toCtx wrapper of urfave v2
func toCtx(a func() error) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		return a()
	}
}