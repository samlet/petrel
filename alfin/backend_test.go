package alfin

import (
	"context"
	"github.com/samlet/petrel/alfin/pb"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestBackend_EntityInfo(t *testing.T) {
	addr := "localhost:9999"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewOutliersClient(conn)
	Do(err, client)
}

func Do(err error, client pb.OutliersClient) {
	req := &pb.EntityInfoRequest{
		Name: "Person",
	}

	resp, err := client.GetEntityInfo(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response: %v", resp)
}

