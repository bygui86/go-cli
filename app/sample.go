package app

import (
	"log"
	"os"
	"sort"

	"github.com/bygui86/go-cli/autocompletion"
	"github.com/bygui86/go-cli/greet"
	"github.com/bygui86/go-cli/list"

	"github.com/urfave/cli"
)

type SampleApp struct {
	app *cli.App
}

func Create() *SampleApp {

	app := cli.NewApp()
	addGlobalConfig(app)
	addGlobalFlags(app)
	addBefore(app)
	addBashCompletion(app)
	addCommands(app)
	lastConfig(app)
	return &SampleApp{
		app: app,
	}
}

func addGlobalConfig(app *cli.App) {

	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Version = "0.0.1"
	// app.UseShortOptionHandling = true // ?!
}

func addGlobalFlags(app *cli.App) {

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "config, c",
			Usage:    "Load configuration from `FILE`",
			FilePath: "~/.go-cli/config", // default value set from file (takes precedence over default values set from the environment "EnvVar")
		},
	}
}

// TODO to be implemented
func addBefore(app *cli.App) {

	app.Before = func(c *cli.Context) error {

		// TODO add config-file action
		// try using https://github.com/kelseyhightower/envconfig

		return nil
	}
}

func addBashCompletion(app *cli.App) {

	app.EnableBashCompletion = true

	// Custom shell autocompletion
	// app.BashComplete = autocompletion.BuildShellCompletionFunc([]string{"greet", "list", "autocompletion", "help"})
}

func addCommands(app *cli.App) {

	app.Commands = []cli.Command{
		*greet.BuildCommand(),
		*list.BuildCommand(),
		*autocompletion.BuildCommand(),
	}
}

func lastConfig(app *cli.App) {

	// sorting flags in help section
	sort.Sort(cli.FlagsByName(app.Flags))
	// sorting commands in help section
	sort.Sort(cli.CommandsByName(app.Commands))
}

func (a *SampleApp) Start() {

	err := a.app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
