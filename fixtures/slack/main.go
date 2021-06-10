// nolint
package main

import (
	"context"
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
	"log"
	"os"
)

/**
$ just run slack run hi
$ just run slack user U0FCYD892
$ go run *.go user U0FCYD892
$ go run *.go slash
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
	User   UserCmd   `cmd help:"Get user info"`
	Bot BotCmd `cmd help:"Run a bot"`
	Slash SlashCmd `cmd help:"Run slash command"`
}

func main() {
	cli := CLI{
		Globals: Globals{
			Version: VersionFlag("0.1.1"),
		},
	}

	ctx := kong.Parse(&cli,
		kong.Name("slack"),
		kong.Description("A slack cli"),
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

var (
	BotToken  = os.Getenv("SLACK_BOT_TOKEN")
	AppToken  = os.Getenv("SLACK_APP_TOKEN")
	UserToken = os.Getenv("SLACK_USER_TOKEN")
	SigningSecret=os.Getenv("SLACK_SIGNING_SECRET")
)

func (cmd *RunCmd) Run(globals *Globals) error {
	bot := slacker.NewClient(BotToken, AppToken)

	definition := &slacker.CommandDefinition{
		Description: "Ping!",
		Example:     "ping",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong", slacker.WithThreadReply(true))
		},
	}

	bot.Command("ping", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

type UserCmd struct {
	Arg string `arg required`
}

func (cmd *UserCmd) Run(globals *Globals) error {
	api := slack.New(UserToken)
	user, err := api.GetUserInfo(cmd.Arg)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
	return nil
}

type BotCmd struct {
	Arg string `arg optional`
}

func (cmd *BotCmd) Run(globals *Globals) error {
	BotHandler()
	return nil
}

type SlashCmd struct {
	Arg string `arg optional`
}

// Run with: http://80f0d9244ef6.ngrok.io/slash
func (cmd *SlashCmd) Run(globals *Globals) error {
	SlashHandler()
	return nil
}
