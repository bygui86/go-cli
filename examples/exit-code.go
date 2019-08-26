package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "ginger-crouton",
			Usage: "Add ginger croutons to the soup",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		if !ctx.Bool("ginger-crouton") {
			return cli.NewExitError("Ginger croutons are not in the soup", 86)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
