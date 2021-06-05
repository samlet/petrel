// nolint
package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/alecthomas/kong"
)

/**
$ just run rsshub run env
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

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

type CLI struct {
	Globals
	Attach AttachCmd `cmd help:"Attach local standard input, output, and error streams to a running container"`
	Run    RunCmd    `cmd help:"Run a command"`
}

func main() {
	cli := CLI{
		Globals: Globals{
			Version: VersionFlag("0.1.1"),
		},
	}

	ctx := kong.Parse(&cli,
		kong.Name("rsshub"),
		kong.Description("A rsshub cli"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": "0.0.1",
		})
	err := ctx.Run(&cli.Globals)
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
	switch cmd.Arg {
	case "env":
		runFile:=getCurrentFilePath()
		runPath:=path.Dir(runFile)
		println(runFile, runPath)
	default:
		println("..")
	}
	return nil
}

func getCurrentFilePath() string {
	_, filePath, _, _ := runtime.Caller(1)
	return filePath
}
