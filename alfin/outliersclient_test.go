package alfin

import (
	"context"
	"fmt"
	"github.com/samlet/petrel/alfin/pb"
	"log"
	"testing"
)

func TestBackend_EntityInfo(t *testing.T) {
	client,err:=NewOutliersClient()
	if err != nil {
		panic(err)
	}
	Do(client.client)
}

func Do(client pb.OutliersClient) {
	req := &pb.EntityInfoRequest{
		Name: "Person",
	}

	resp, err := client.GetEntityInfo(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response: %s (%s)", resp.Name, resp.PackageName)
	println(resp.Meta)
}

func TestOutliersClient_GetEntityMeta(t *testing.T) {
	client,err:=NewOutliersClient()
	if err != nil {
		panic(err)
	}
	model, err:=client.GetEntityMeta("Person")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s: %s\n", model.Name, model.Pks)
}