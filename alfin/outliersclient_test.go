package alfin

import (
	"context"
	"fmt"
	"github.com/samlet/petrel/alfin/pb"
	"log"
	"testing"
)

func TestBackend_EntityInfo(t *testing.T) {
	ctx:=context.Background()
	client,err:=NewOutliersClient()
	if err != nil {
		panic(err)
	}
	Do(client.Client, ctx)
}

func Do(client pb.OutliersClient, ctx context.Context) {
	req := &pb.EntityInfoRequest{
		Name: "Person",
	}

	resp, err := client.GetEntityInfo(ctx, req)
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

func TestOutliersClient_GetEntitiesInPackage(t *testing.T) {
	ctx:=context.Background()
	client,err:=NewOutliersClient()
	if err != nil {
		panic(err)
	}
	model, err:=client.Client.GetEntitiesInPackage(ctx,
		&pb.EntityPackageRequest{Name: "com.bluecc.workload"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s: %s\n", model.Name, model.Entities)
}
