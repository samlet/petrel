package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"go/build"
	"log"
	"os"
)

/**
$ just gen parser -s env
$ just gen parser -s sources ./codegen/parser
$ just gen parser -s sources ./fixtures/remote/
$ just gen parser -s interfaces
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
			switch service {
			case "sources":
				if act == "" {
					act = "./codegen/parser"
				}
				pkgInfo, err := build.ImportDir(act, 0)
				if err != nil {
					log.Fatal(err)
				}
				for i, file := range pkgInfo.GoFiles {
					fmt.Printf("%d. %s\n", i, file)
				}
			case "interfaces":
				if act == "" {
					act = "./codegen/parser"
				}
				pkgInfo, err := build.ImportDir(act, 0)
				if err != nil {
					log.Fatal(err)
				}
				GetInterfaces(act, pkgInfo.GoFiles)
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
