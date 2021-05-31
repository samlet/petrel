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
$ just alfin gentempl -s gen -t alfin/fields.tmpl -i alfin/inventoryitem.json
$ just alfin gentempl -s gen -t alfin/entity_type.tmpl -i alfin/inventoryitem.json
$ just alfin gentempl -s gen -t alfin/entity_type.tmpl -i alfin/inventoryitem.json -o alfin/inventoryitem.go
$ just alfin gentempl -s gen -t alfin/service_intf.tmpl -i alfin/person_ops.json
$ just alfin gentempl -s gen -t alfin/service_intf.tmpl -i alfin/person_ops.json -o alfin/person_ops.go

# on alfin dir
$ go run gentempl/main.go -s gen -t fields.tmpl -i inventoryitem.json
$ gen -t fields.tmpl -i inventoryitem.json
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
			&cli.StringFlag{Name: "output", Aliases: []string{"o"},
				Usage: "target file", Required: false},
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

				if c.IsSet("output") {
					target := c.String("output")
					f, err := os.Create(target)
					check(err)
					defer f.Close()
					alfin.GenTemplate(string(d), string(t), f)
				} else {
					alfin.GenTemplate(string(d), string(t), os.Stdout)
				}
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
