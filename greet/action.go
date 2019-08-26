package greet

import (
	"fmt"

	"github.com/urfave/cli"
)

func greet(name *string) func(c *cli.Context) error {

	return func(c *cli.Context) error {

		// fmt.Printf("[DEBUG] Name: %s\n", *name)

		if name == nil || *name == "" || *name == "world" {
			fmt.Println("Hello world!")
		} else {
			fmt.Println("Hello " + *name + "!")
		}
		return nil
	}
}
