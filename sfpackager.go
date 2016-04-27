package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sfpackager"

	app.Commands = []cli.Command{
		cli.Command{
			Name:    "package",
			Aliases: []string{"p"},
			Usage:   "Packages stuff",
			Action:  Package,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "sfprojects, s",
					Value: "",
					Usage: "Path to the root of your solution or a specific .sfproj file",
				},
				cli.StringFlag{
					Name:  "output, o",
					Value: "",
					Usage: "The base folder to package to. This folder will be cleared on run",
				},
			},
		},
	}

	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	app.Run(os.Args)
}

// Exit exit gracefully
var Exit = func(message string) {
	fmt.Printf(message)
	os.Exit(1)
}

/*
-sfprojects
*/
