package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/samlet/petrel/alfin"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

/**
$ just run builder -s env

# specific
$ srv resource Example
*/

func main() {
	prompt := color.New(color.FgYellow).PrintfFunc()
	imp := color.New(color.FgGreen).SprintFunc()

	var service string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "service",
				Aliases:     []string{"s"},
				Value:       "find",
				Usage:       "service to execute",
				Destination: &service,
			},
			&cli.IntFlag{Name: "size", Aliases: []string{"m"}, Value: 20},
			&cli.StringFlag{Name: "entity", Aliases: []string{"e"},
				Usage: "entity name",
				Value: "WebSite"},
		},
		Action: func(c *cli.Context) error {
			act := "_none_"
			if c.NArg() > 0 {
				act = c.Args().Get(0)
				prompt(".. act is %s\n", act)
			}
			switch service {
			case "resource":
				if act != "" {
					entName := strings.ToLower(act)
					creator := alfin.Creator{
						PackageName:  entName,
						InputPath:    entName + "_ops.json",
						TemplatePath: "service_impl.tmpl",
						TargetName:   "client.go",
					}
					err := creator.Execute("")
					if err != nil {
						prompt("%s\n", err)
						return nil
					}
				} else {
					prompt("Absent resource name")
				}
			case "env":
				fmt.Println("GOPATH:", os.Getenv("GOPATH"))
			default:
				fmt.Printf("Cannot to execute %s.\n", imp(service))
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
