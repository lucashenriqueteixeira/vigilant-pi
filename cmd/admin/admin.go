package admin

import (
	"github.com/urfave/cli"
)

// Command ...
func Command() cli.Command {
	return cli.Command{
		Name:        "admin",
		Description: "starts a web admin interface",
		Action: func(c *cli.Context) error {

			return nil
		},
	}
}
