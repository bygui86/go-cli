package list

import (
	"github.com/urfave/cli"
)

func BuildCommand() *cli.Command {

	var path string
	var showHidden bool
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
						Name:        "path, p",
						Usage:       "`PATH` to scan for files",
						EnvVar:      "GOCLI_LIST_PATH,LIST_PATH",
						Value:       ".",
						Required:    false,
						Destination: &path,
					},
					cli.BoolFlag{
						Name:        "hidden, d",
						Usage:       "show hidden files",
						EnvVar:      "GOCLI_LIST_HIDDEN,LIST_HIDDEN",
						Required:    false,
						Destination: &showHidden,
						// Value field not available, but `FALSE` per default
					},
				},
				Action: listFiles(&path, &showHidden),
			},
			{
				Name:  "folders",
				Usage: "List all folders in a directory",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "path, p",
						Usage:       "`PATH` to scan for folders",
						EnvVar:      "GOCLI_LIST_PATH,LIST_PATH",
						Value:       ".",
						Required:    false,
						Destination: &path,
					},
					cli.BoolFlag{
						Name:        "hidden, d",
						Usage:       "show hidden files",
						EnvVar:      "GOCLI_LIST_HIDDEN,LIST_HIDDEN",
						Required:    false,
						Destination: &showHidden,
					},
				},
				Action: listFolders(&path, &showHidden),
			},
			{
				Name:  "all",
				Usage: "List all files and folders in a directory",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "path, p",
						Usage:       "`PATH` to scan for folders",
						EnvVar:      "GOCLI_LIST_PATH,LIST_PATH",
						Value:       ".",
						Required:    false,
						Destination: &path,
					},
				},
				Action: listAll(&path),
			},
		},
	}
}
