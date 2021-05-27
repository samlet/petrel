package main

import (
	"fmt"
	"github.com/fatih/color"
	metagen "github.com/samlet/petrel/meta"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/**
$ just run generator -s type -e Person
$ just run generator -s service-def createPerson
$ just run generator -s flow-def Sample downloadFile processFile
$ just run generator -w -s flow-def Dummy downloadFile processFile
$ just dummy  # recreate dummy workflow
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
			&cli.BoolFlag{Name: "write", Aliases: []string{"w"}, Value: false},
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
			case "flow-def":
				err, _ := genFlow(c, act, prompt)
				if err != nil {
					panic(err)
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func genFlow(c *cli.Context, act string, prompt func(format string, a ...interface{})) (error, bool) {
	if c.NArg() < 3 {
		println("Must specific flow name and activities")
		return nil, false
	}
	acts := c.Args().Slice()[1:]
	flowMeta := metagen.CreateFlowMeta(act, acts)

	prompt(".. workflow def %s\n", act)
	workflowCnt, err := metagen.GenFlowFromTemplate(&flowMeta, "incls/workflow_def.tmpl")
	if err != nil {
		panic(err)
	}
	println(workflowCnt)

	prompt(".. workflow test %s\n", act)
	workflowTestCnt, err := metagen.GenFlowFromTemplate(&flowMeta, "incls/workflow_test.tmpl")
	if err != nil {
		panic(err)
	}
	println(workflowTestCnt)

	write_target := c.Bool("write")
	if write_target {
		targetDir := fmt.Sprintf("./routines/%s", strings.ToLower(act))
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			// target dir does not exist
			err = os.Mkdir(targetDir, 0755)
			check(err)
			err := ioutil.WriteFile(filepath.Join(targetDir, "workflow.go"),
				[]byte(workflowCnt), 0644)
			check(err)
			err = ioutil.WriteFile(filepath.Join(targetDir, "workflow_test.go"),
				[]byte(workflowTestCnt), 0644)
			check(err)

			println(".. write to target dir", targetDir)
		} else {
			prompt("!! target dir '%s' is exists\n", targetDir)
		}

	}

	return nil, true
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
