package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/samlet/petrel/alfin"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
)

/**
$ just alfin gentempl -s gen
$ just alfin gentempl -s gen -t alfin/fields.tmpl -i alfin/inventoryitem.json
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
				Value:       "gen",
				Usage:       "service to execute",
				Destination: &service,
			},
			&cli.StringFlag{Name: "template", Aliases: []string{"t"},
				Usage: "template name", Required: true},
			&cli.StringFlag{Name: "input", Aliases: []string{"i"},
				Usage: "input data source", Required: true},
		},
		Action: func(c *cli.Context) error {
			act := "_none_"
			if c.NArg() > 0 {
				act = c.Args().Get(0)
				prompt(".. act is %s\n", act)
			}
			switch service {
			case "gen":
				t, err := ioutil.ReadFile(c.String("template"))
				if err != nil {
					log.Fatal(err)
				}
				d, err := ioutil.ReadFile(c.String("input"))
				if err != nil {
					log.Fatal(err)
				}
				alfin.GenTemplate(string(d), string(t), os.Stdout)
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
