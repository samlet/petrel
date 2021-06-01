package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/samlet/petrel/alfin"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/**
$ just alfin builder -s asset Ballot.json
$ albuilder -s asset Ballot.json

# specific
$ srv resource Example
$ gen -t service_intf.tmpl -i exampleitem_ops.json -o exampleitem_ops.go  # optional
$ srv resource -f Example createExampleStatus createExampleFeature
$ srv resource -f Facility getInventoryAvailableByFacility
$ srv create --help
$ srv create -f -c conf/maint_common.yml
$ srv asset Ballot.json
*/

const (
	AssetFolder    string = "assets"
	TemplateFolder string = "templates"
	ConfFolder     string = "conf"
)

func main() {
	prompt := color.New(color.FgYellow).PrintfFunc()
	imp := color.New(color.FgGreen).SprintFunc()

	var service string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "service",
				Aliases:     []string{"s"},
				Value:       "create",
				Usage:       "service to execute",
				Destination: &service,
			},
			&cli.IntFlag{Name: "size", Aliases: []string{"m"}, Value: 20},
			&cli.StringFlag{Name: "entity", Aliases: []string{"e"},
				Usage: "entity name",
				Value: "WebSite"},
			&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Value: false},
			&cli.StringFlag{Name: "conf", Aliases: []string{"c"},
				Usage: "config file"},
		},
		Action: func(c *cli.Context) error {
			act := ""
			if c.NArg() > 0 {
				act = c.Args().Get(0)
				prompt(".. act is %s\n", act)
			}
			switch service {
			case "resource":
				if act != "" {
					if c.Bool("force") {
						println("force to recreate ..")
						deleteResource(act)
					}
					extra := []string{}
					if c.NArg() > 1 {
						extra = c.Args().Slice()[1:]
					}
					err := genResource(act, extra)
					if err != nil {
						panic(err)
					}
				} else {
					prompt("Absent resource name")
				}
			case "create":
				if c.IsSet("conf") {
					genResourceByConf(
						alfin.ReadMaintConf(c.String("conf")),
						c.Bool("force"))
				} else {
					prompt("Must specific config file")
				}
			case "asset":
				if act != "" {
					assertBox := alfin.NewAssetBox()
					// get file contents as string
					asset, err := assertBox.String(act)
					if err != nil {
						log.Fatal(err)
					}

					var bodySample string
					if len(asset) > 500 {
						bodySample = asset[0:500] + " ..."
					} else {
						bodySample = asset
					}

					println(bodySample)
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

func genMeta(ent string, extra []string) {
	println(".. generate meta")
	extraSrvs := "[" + strings.Join(extra, ",") + "]"
	out, err := alfin.PyGen("service_meta.py", "entity_abi",
		ent, extraSrvs)
	if err != nil {
		panic(err)
	}
	println("PyGen OUT => ", out)
}

func genInterface(act string) error {
	println(".. generate interface")
	entName := strings.ToLower(act)
	creator := alfin.Creator{
		PackageName:  "",
		InputPath:    filepath.Join(AssetFolder, entName+"_ops.json"),
		TemplatePath: filepath.Join(TemplateFolder, "service_intf.tmpl"),
		TargetName:   entName + "_ops.go",
	}

	err := creator.Execute("")
	if err != nil {
		log.Fatalf("%s\n", err)
		return err
	}
	return nil
}

func deleteResource(act string) {
	entName := strings.ToLower(act)
	files := []string{
		filepath.Join(AssetFolder, entName+"_ops.json"),
		entName + "_ops.go",
		filepath.Join(entName, "client.go"),
	}
	for _, f := range files {
		e := os.Remove(f)
		if e != nil {
			log.Println("ignore: ", e)
		}
	}
}

func genResourceByConf(maintConf alfin.MainEntConf, force bool) {
	for _, conf := range maintConf.MainEnts {
		if force {
			deleteResource(conf.Name)
		}
		err := genResource(conf.Name, conf.Extras)
		if err != nil {
			panic(err)
		}
	}
}

func genResource(act string, extra []string) error {
	println(".. generate resource")

	entName := strings.ToLower(act)
	creator := alfin.Creator{
		PackageName:  entName,
		InputPath:    filepath.Join(AssetFolder, entName+"_ops.json"),
		TemplatePath: filepath.Join(TemplateFolder, "service_impl.tmpl"),
		TargetName:   "client.go",
		Extra:        extra,
	}

	err := alfin.EnsureFile(creator.InputPath)
	if err != nil {
		genMeta(act, extra)
	}
	err = alfin.EnsureFile(entName + "_ops.go")
	if err != nil {
		err = genInterface(act)
		if err != nil {
			return err
		}
	}

	err = creator.Execute("")
	if err != nil {
		log.Fatalf("%s\n", err)
		return err
	}
	return nil
}
