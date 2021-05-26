package main

import (
	"fmt"
	"github.com/fatih/color"
	metagen "github.com/samlet/petrel/meta"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
)

/**
$ just run generator -s type -e Person
$ just run generator -s service-def createPerson
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
			case "type":
				entity := c.String("entity")
				if entity != "" {
					printEntityDef(entity)
				}
			case "service-def":
				if act != "" {
					printServiceDef(act)
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

func printEntityDef(entity string) {
	meta, err := metagen.GetEntityMeta(entity)
	if err != nil {
		panic(err)
	}
	tplData, err := ioutil.ReadFile("incls/entity_def.tmpl")
	if err != nil {
		panic(err)
	}
	result := metagen.GenEntity(&meta.Data.Entity, string(tplData))
	println(result)
}

func printServiceDef(act string) {
	meta, err := metagen.GetServiceMeta(act)
	if err != nil {
		panic(err)
	}
	println(meta.StatusCode)
	for i, rel := range meta.Data.Service.Parameters {
		fmt.Printf("%d. %s\n", i, rel)
	}
}
