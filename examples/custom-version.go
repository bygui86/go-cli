package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	Revision = "fafafaf"
)

func main() {

	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, V",
		Usage: "print only the version",
	}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s revision=%s\n", c.App.Version, Revision)
	}

	app := cli.NewApp()
	app.Name = "partay"
	app.Version = "19.99.0"
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
