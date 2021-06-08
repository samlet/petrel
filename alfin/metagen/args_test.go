package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/kballard/go-shellquote"
	"testing"
)

func TestArgsParse(t *testing.T) {
	input:=`metagen -D --tls-verify --config .meta run "hello, guys"`
	parts,err:=shellquote.Split(input)
	if err != nil {
		panic(err)
	}
	println(input)
	for _,p := range parts {
		fmt.Printf("%s |", p)
	}
	println()

	DoCmd(parts)
}

func ParseCmd(cli interface{}, args []string, options ...kong.Option) *kong.Context {
	parser, err := kong.New(cli, options...)
	if err != nil {
		panic(err)
	}
	ctx, err := parser.Parse(args[1:])
	parser.FatalIfErrorf(err)
	return ctx
}


func DoCmd(args []string){
	cli := CLI{
		Globals: Globals{
			Version: VersionFlag("0.1.1"),
		},
	}

	ctx := ParseCmd(&cli, args,
		kong.Name("metagen"),
		kong.Description("A metagen cli"),
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
