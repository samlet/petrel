package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

/**
$ just run distr -s ping
$ just run distr -s value key v1 v2 v3
$ just run distr -s value key  # get value
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
			case "ping":
				distr := NewRedisDistr()
				distr.Ping()
			case "value":
				distr := NewRedisDistr()
				if c.NArg() == 1 {
					val, err := distr.client.Get(act).Result()
					if err != nil {
						fmt.Println(err)
					}

					fmt.Println(val)
				}
				if c.NArg() == 2 {
					err := distr.client.Set(act, c.Args().Get(1), 0).Err()
					// if there has been an error setting the value
					// handle the error
					if err != nil {
						fmt.Println(err)
					}
				}
				if c.NArg() >= 2 {
					valueList, err := json.Marshal(c.Args().Slice()[1:])
					if err != nil {
						panic(err)
					}
					err = distr.client.Set(act, valueList, 0).Err()
					if err != nil {
						fmt.Println(err)
					}
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

/**
refs:
	- https://tutorialedge.net/golang/go-redis-tutorial/
*/
