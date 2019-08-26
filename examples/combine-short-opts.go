package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	// app.UseShortOptionHandling = true // ?!
	app.Commands = []cli.Command{
		{
			Name:  "short",
			Usage: "complete a task on the list",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "serve, s"},
				cli.BoolFlag{Name: "option, o"},
				cli.StringFlag{Name: "message, m"},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("serve:", c.Bool("serve"))
				fmt.Println("option:", c.Bool("option"))
				fmt.Println("message:", c.String("message"))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
