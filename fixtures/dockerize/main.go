package main

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

/**
$ just run dockerize -s alpine
$ just run dockerize -s start
$ just run dockerize -s start bfirsh/reticulate-splines
$ just run dockerize -s ps
$ just run dockerize -s stop
$ just run dockerize -s logs 09ade4bba3

## references
- https://docs.docker.com/engine/api/sdk/examples/
*/

func main() {
	prompt := color.New(color.FgYellow).PrintfFunc()
	imp := color.New(color.FgGreen).SprintFunc()

	var service string
	ctx := context.Background()
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
			switch service {
			case "alpine":
				runAlpine()
			case "ps":
				ListContainers()
			case "start":
				if act == "" {
					act = "bfirsh/reticulate-splines"
				}
				runBackground(act)
			case "stop":
				StopAllContainers()
			case "logs":
				if act != "" {
					GetLogs(ctx, act)
				} else {
					println("Must specific a container id.")
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
