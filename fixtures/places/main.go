package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

/**
$ just run places -s seed
$ just run places -s allPerson
$ just run places -s findPerson Jason
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
			act := ""
			if c.NArg() > 0 {
				act = c.Args().Get(0)
				prompt(".. act is %s\n", act)
			}

			placesDb, err := NewPlacesDb()
			if err != nil {
				panic(err)
			}

			switch service {
			case "seed":
				placesDb.Seed()
			case "allPerson":
				rs, err := placesDb.AllPerson()
				if err != nil {
					panic(err)
				}
				for n, r := range rs {
					fmt.Printf("%d. %s\n", n, r)
				}
			case "findPerson":
				if act != "" {
					p := Person{}
					err = placesDb.db.Get(&p, "SELECT * FROM person WHERE first_name=$1", act)
					fmt.Printf("%#v\n", p)
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
