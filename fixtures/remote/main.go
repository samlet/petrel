package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"log"
	"os"

	"crypto/sha1"
	"golang.org/x/net/context"
	"hash"
	"net"

	"github.com/samlet/petrel/fixtures/remote/hashes"
	"zombiezen.com/go/capnproto2/rpc"
)

/**
$ just run remote -s boot
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
			case "boot":
				err := boot()
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

// hashFactory is a local implementation of HashFactory.
type hashFactory struct{}

func (hf hashFactory) NewSha1(call hashes.HashFactory_newSha1) error {
	// Create a new locally implemented Hash capability.
	hs := hashes.Hash_ServerToClient(hashServer{sha1.New()})
	// Notice that methods can return other interfaces.
	return call.Results.SetHash(hs)
}

// hashServer is a local implementation of Hash.
type hashServer struct {
	h hash.Hash
}

func (hs hashServer) Write(call hashes.Hash_write) error {
	data, err := call.Params.Data()
	if err != nil {
		return err
	}
	_, err = hs.h.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (hs hashServer) Sum(call hashes.Hash_sum) error {
	s := hs.h.Sum(nil)
	return call.Results.SetHash(s)
}

func server(c net.Conn) error {
	// Create a new locally implemented HashFactory.
	main := hashes.HashFactory_ServerToClient(hashFactory{})
	// Listen for calls, using the HashFactory as the bootstrap interface.
	conn := rpc.NewConn(rpc.StreamTransport(c), rpc.MainInterface(main.Client))
	// Wait for connection to abort.
	err := conn.Wait()
	return err
}

func client(ctx context.Context, c net.Conn) error {
	// Create a connection that we can use to get the HashFactory.
	conn := rpc.NewConn(rpc.StreamTransport(c))
	defer conn.Close()
	// Get the "bootstrap" interface.  This is the capability set with
	// rpc.MainInterface on the remote side.
	hf := hashes.HashFactory{Client: conn.Bootstrap(ctx)}

	// Now we can call methods on hf, and they will be sent over c.
	s := hf.NewSha1(ctx, func(p hashes.HashFactory_newSha1_Params) error {
		return nil
	}).Hash()
	// s refers to a remote Hash.  Method calls are delivered in order.
	s.Write(ctx, func(p hashes.Hash_write_Params) error {
		err := p.SetData([]byte("Hello, "))
		return err
	})
	s.Write(ctx, func(p hashes.Hash_write_Params) error {
		err := p.SetData([]byte("World!"))
		return err
	})
	// Get the sum, waiting for the result.
	result, err := s.Sum(ctx, func(p hashes.Hash_sum_Params) error {
		return nil
	}).Struct()
	if err != nil {
		return err
	}

	// Display the result.
	sha1Val, err := result.Hash()
	if err != nil {
		return err
	}
	fmt.Printf("sha1: %x\n", sha1Val)
	return nil
}

func boot() error {
	c1, c2 := net.Pipe()
	go server(c1)
	err := client(context.Background(), c2)
	return err
}

/**
refs:
	- https://github.com/capnproto/go-capnproto2/wiki/Getting-Started
*/
