// nolint
package main

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/samlet/petrel/alfin"
	log "github.com/sirupsen/logrus"
	"github.com/xlab/treeprint"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kong"
)

/**
$ just cli -l debug meta workeffort WorkEffort
$ just cli meta --show-fields workeffort WorkEffort
$ just cli meta --all-ents workeffort
$ just cli seed-trace workeffort Party partyId DemoCustomer3
$ just cli seed-trace --no-rel workeffort Person partyId DemoCustomer3
$ just cli seed-gen workeffort Party
$ just cli seed-gen --only-creator workeffort StatusItem
*/

type Globals struct {
	Config    string      `help:"Location of client config files" default:"~/.alfin" type:"path"`
	Debug     bool        `short:"D" help:"Enable debug mode"`
	Host      []string    `short:"H" help:"Daemon socket(s) to connect to"`
	LogLevel  string      `short:"l" help:"Set the logging level (debug|info|warn|error|fatal)" default:"info"`
	TLS       bool        `help:"Use TLS; implied by --tls-verify"`
	TLSCACert string      `name:"tls-ca-cert" help:"Trust certs signed only by this CA" default:"~/.alfin/ca.pem" type:"path"`
	TLSCert   string      `help:"Path to TLS certificate file" default:"~/.alfin/cert.pem" type:"path"`
	TLSKey    string      `help:"Path to TLS key file" default:"~/.alfin/key.pem" type:"path"`
	TLSVerify bool        `help:"Use TLS and verify the remote"`
	Version   VersionFlag `name:"version" help:"Print version information and quit"`
}

type CLI struct {
	Globals
	Attach    AttachCmd    `cmd help:"Attach local standard input, output, and error streams to a running container"`
	Run       RunCmd       `cmd help:"Run a command"`
	Meta      MetaCmd      `cmd help:"Show meta-info"`
	SeedTypes SeedTypesCmd `cmd help:"Generate seed types"`
	SeedTrace SeedTraceCmd `cmd help:"Trace a seed record"`
	SeedGen   SeedGenCmd   `cmd help:"Generate seed creator and updater"`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

func main() {
	cli := CLI{
		Globals: Globals{
			Version: VersionFlag("0.1.1"),
		},
	}

	ctx := kong.Parse(&cli,
		kong.Name("cli"),
		kong.Description("A cli cli"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": "0.0.1",
		})

	lvl, err := log.ParseLevel(cli.LogLevel)
	if err != nil {
		log.Fatalf("parse log level fail: %v", err)
	}
	log.SetLevel(lvl)

	err = ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}

type AttachCmd struct {
	DetachKeys string `help:"Override the key sequence for detaching a container"`
	NoStdin    bool   `help:"Do not attach STDIN"`
	SigProxy   bool   `help:"Proxy all received signals to the process" default:"true"`

	Container string `arg required help:"Container ID to attach to."`
}

func (a *AttachCmd) Run(globals *Globals) error {
	fmt.Printf("Config: %s\n", globals.Config)
	fmt.Printf("Attaching to: %v\n", a.Container)
	fmt.Printf("SigProxy: %v\n", a.SigProxy)
	return nil
}

type RunCmd struct {
	Arg string `arg required`
}

func (cmd *RunCmd) Run(globals *Globals) error {
	return nil
}

type MetaCmd struct {
	ShowFields bool   `help:"Show fields" default:"false"`
	AllEnts    bool   `help:"Show all entity meta" default:"false"`
	Pkg        string `arg required`
	Ent        string `arg optional`
}

func (cmd *MetaCmd) Run(globals *Globals) error {
	mani, err := alfin.NewManipulateWithPackage(cmd.Pkg)
	if err != nil {
		panic(err)
	}
	tree := treeprint.New()

	if cmd.AllEnts {
		for _, ent := range mani.Entities().Entities {
			alfin.DisplayEntInfo(ent, tree, cmd.ShowFields)
		}
	} else {
		if len(cmd.Ent) != 0 {
			ent := mani.MustEntity(cmd.Ent)
			alfin.DisplayEntInfo(ent, tree, cmd.ShowFields)
		} else {
			return fmt.Errorf("specify a entity")
		}
	}

	fmt.Println(tree.String())

	return nil
}

type SeedTypesCmd struct {
	Pkg string `arg required`
}

func (cmd *SeedTypesCmd) Run(globals *Globals) error {
	tmpl := "templates/seed_schema.tmpl"
	header := `package seedtypes
import "github.com/samlet/petrel/services"
`

	mani, err := alfin.NewManipulateWithPackage(cmd.Pkg)
	if err != nil {
		panic(err)
	}

	for _, ent := range mani.Entities().Entities {
		filePath := strings.ToLower(ent.Name) + ".go"
		f, err := os.Create(filepath.Join("seedtypes", filePath))
		if err != nil {
			log.Fatalf("create file fail: %v", err)
		}
		_, err = f.WriteString(header)
		if err != nil {
			log.Fatalf("write header fail: %v", err)
		}
		err = alfin.GenModelEntityWithMeta(tmpl, ent, f)
		if err != nil {
			panic(err)
		}
		f.Close()
	}

	return nil
}

type SeedTraceCmd struct {
	NoRel bool   `help:"Not show relations" default:"false"`
	Pkg   string `arg required help:"asset package"`
	Ent   string `arg required`
	Pk    string `arg required`
	PkVal string `arg required`
}

func (cmd *SeedTraceCmd) Run(globals *Globals) error {
	pkg := cmd.Pkg
	entName := cmd.Ent

	xmlSeedFile := filepath.Join("assets", pkg, "seeds.xml")
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(xmlSeedFile); err != nil {
		panic(err)
	}
	root := doc.SelectElement("entity-engine-xml")
	log.Println("ROOT element:", root.Tag)

	mani, err := alfin.NewManipulateWithPackage(pkg)
	if err != nil {
		log.Fatal(err)
	}

	model := mani.MustEntity(entName)
	seedProc := alfin.NewSeedProcessor(mani, !cmd.NoRel)
	elements := seedProc.GetElementsByArgs(doc, entName, cmd.Pk, cmd.PkVal)

	seedProc.ProcessElements(doc, elements, model)
	println("-------------------")
	println(seedProc.Buffer.String())
	return nil
}

type SeedGenCmd struct {
	OnlyCreator bool `help:"Only generate creators" default:"false"`
	Pkg string `arg required`
	Ent string `arg required`
}

func (cmd *SeedGenCmd) Run(globals *Globals) error {
	seedgen := alfin.NewSeedGen(cmd.Pkg, cmd.Ent)
	return seedgen.Generate(cmd.OnlyCreator)
}
