package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// EXAMPLE: Append to an existing template
	cli.AppHelpTemplate = fmt.Sprintf(`%s

	WEBSITE: http://awesometown.example.com

	SUPPORT: support@awesometown.example.com

	`, cli.AppHelpTemplate)

	// EXAMPLE: Override a template
	// 	cli.AppHelpTemplate = `NAME:
	//    {{.Name}} - {{.Usage}}
	// USAGE:
	//    {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	//    {{if len .Authors}}
	// AUTHOR:
	//    {{range .Authors}}{{ . }}{{end}}
	//    {{end}}{{if .Commands}}
	// COMMANDS:
	// {{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
	// GLOBAL OPTIONS:
	//    {{range .VisibleFlags}}{{.}}
	//    {{end}}{{end}}{{if .Copyright }}
	// COPYRIGHT:
	//    {{.Copyright}}
	//    {{end}}{{if .Version}}
	// VERSION:
	//    {{.Version}}
	//    {{end}}
	// `

	// EXAMPLE: Replace the `HelpPrinter` func
	// cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
	// 	fmt.Println("Ha HA.  I pwnd the help!!!")
	// }

	cli.HelpFlag = cli.BoolFlag{
		Name:   "halp, haaaaalp",
		Usage:  "HALP",
		EnvVar: "SHOW_HALP,HALPPLZ",
	}

	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
