package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/samlet/petrel/routines/common"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

/**
$ just rt stuffs -s env
$ just rt stuffs -s worker
$ cadence workflow run --tl seedGroup --wt main.SeedWorkflow --et 60 -i '"cadence"'
*/

func main() {
	prompt := color.New(color.FgYellow).PrintfFunc()
	imp := color.New(color.FgGreen).SprintFunc()

	var service string

	var h common.SampleHelper
	h.SetupServiceConfig()

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
			case "worker":
				startWorkers(&h)
				// The workers are supposed to be long running process that should not exit.
				// Use select{} to block indefinitely for samples, you can quit by CMD+C.
				select {}
			case "trigger":
				startSeedWorkflow(&h)
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
