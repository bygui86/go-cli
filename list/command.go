package list

import (
	"github.com/urfave/cli"
)

func BuildCommand() *cli.Command {

	return &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list files and/or folders in current directory",
		Subcommands: []cli.Command{
			{
				Name:  "files",
				Usage: "List all files in a directory",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "path, p",
						Usage:    "`PATH` to scan for files",
						EnvVar:   "GOCLI_LIST_PATH,LIST_PATH",
						Value:    ".",
						Required: false,
					},
					cli.BoolFlag{
						Name:     "show-hidden, d",
						Usage:    "show hidden files",
						EnvVar:   "GOCLI_LIST_HIDDEN,LIST_HIDDEN",
						Required: false,
						// Value field not available, but `FALSE` per default
					},
				},
				Action: listFiles(),
			},
			{
				Name:  "folders",
				Usage: "List all folders in a directory",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "path, p",
						Usage:    "`PATH` to scan for folders",
						EnvVar:   "GOCLI_LIST_PATH,LIST_PATH",
						Value:    ".",
						Required: false,
					},
					cli.BoolFlag{
						Name:     "show-hidden, d",
						Usage:    "show hidden files",
						EnvVar:   "GOCLI_LIST_HIDDEN,LIST_HIDDEN",
						Required: false,
					},
				},
				Action: listFolders(),
			},
			{
				Name:  "all",
				Usage: "List all files and folders in a directory",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "path, p",
						Usage:    "`PATH` to scan for folders",
						EnvVar:   "GOCLI_LIST_PATH,LIST_PATH",
						Value:    ".",
						Required: false,
					},
				},
				Action: listAll(),
			},
		},
		// Custom shell autocompletion
		// BashComplete: autocompletion.BuildShellCompletionFunc([]string{"files", "folders", "all", "help"}),
	}
}
