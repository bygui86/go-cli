package app

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

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
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
			// FilePath: "~/.go-cli/config", // default value set from file (takes precedence over default values set from the environment "EnvVar")
			FilePath: "./config", // default value set from file (takes precedence over default values set from the environment "EnvVar")
		},
	}
}

func addBefore(app *cli.App) {

	app.Before = func(c *cli.Context) error {

		wholeConfig := c.String("config")
		if wholeConfig != "" {
			fmt.Printf("[DEBUG] Whole config as string: %s\n", wholeConfig)

			wholeConfigSplit := strings.Fields(wholeConfig)
			fmt.Printf("[DEBUG] Whole config as array: %s\n", wholeConfigSplit)
	
			for _, config := range wholeConfigSplit {
				configSplit := strings.Split(config, "=")
				fmt.Printf("[DEBUG] single config: %v\n", configSplit)
				if len(configSplit) == 2 && configSplit[0] != "" && configSplit[1] != "" {
					fmt.Printf("[DEBUG] env-var set: %s = %s\n", configSplit[0], configSplit[1])
					os.Setenv(configSplit[0], configSplit[1])
				} else {
					fmt.Printf("[ERROR] Config not valid: %s - skipping and taking default...\n", configSplit)
				}
			}
		} else {
			fmt.Println("[DEBUG] Config file not set or empty")
		}

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
