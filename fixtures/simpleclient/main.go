package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

/*
$ just run simpleclient -s performFindList -m 3
$ just run simpleclient -s find -e Party -m 5
$ srv find -e PartyGroup -m 2 print
*/

func main() {
	// Create a custom print function for convenience
	//red := color.New(color.FgRed).PrintfFunc()
	prompt := color.New(color.FgYellow).PrintfFunc()

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
			if service == "find" {
				result, _ := InvokeService("findCc", fmt.Sprintf(`{
					"entityName":"%s",
					"maxRows":%d
				}`, c.String("entity"), c.Int("size")))
				fmt.Println(result)
			} else if service == "routings" {
				result, _ := FindRoutings(c.Int("size"))
				fmt.Println(result)
			} else {
				fmt.Printf("Cannot to execute %s.\n", service)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
