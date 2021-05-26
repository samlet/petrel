package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

/**
$ just run filesys -s env
$ just run filesys -s find .. *.sh
$ just run filesys -s find .. procs*.md
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
			case "find":
				root := os.Getenv("GOPATH")
				ext := "*.md"
				if c.NArg() > 0 {
					root = c.Args().Get(0)
				}
				if c.NArg() > 1 {
					ext = c.Args().Get(1)
				}

				start := time.Now()
				files, err := WalkMatch(root, ext)
				if err != nil {
					panic(err)
				}

				fmt.Printf("total %d files\n", len(files))
				elapsed := time.Since(start)
				log.Printf("time elapsed %s", elapsed)
				for i, f := range files {
					fmt.Printf("%d. %s\n", i, f)
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
