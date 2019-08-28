package greet

import (
	"github.com/urfave/cli"
)

func BuildCommand() *cli.Command {

	var name string
	return &cli.Command{
		Name:    "greet",
		Aliases: []string{"g"},
		Usage:   "Greet world or name",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "name, n",                     // comma-separated list for alternate names
				Usage:       "`NAME` to greet",             // back quotes are used to identify the placeholder for the current flag as "--name NAME, -n NAME"
				EnvVar:      "GOCLI_GREET_NAME,GREET_NAME", // comma-separated "cascade" list of environment variables used as alternative to flag
				Value:       "world",                       // will be automatically added to the current "Usage" as "(default: VALUE)"
				Required:    false,
				Destination: &name,
			},
		},
		Action: greet(&name),
	}
}
