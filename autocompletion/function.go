package autocompletion

import (
	"fmt"

	"github.com/urfave/cli"
)

// Used for 'custom shell autocompletion'
func BuildShellCompletionFunc(commands []string) func(c *cli.Context) {

	return func(c *cli.Context) {
		// This will complete if no args are passed
		if c.NArg() > 0 {
			return
		}
		for _, t := range commands {
			fmt.Println(t)
		}
	}
}
