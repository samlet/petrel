package alfin

import (
	"context"
	"encoding/json"
	"github.com/samlet/petrel/alfin/pb"
	"google.golang.org/grpc"
	"log"
)

type OutliersClient struct{
	Client pb.OutliersClient
}

func NewOutliersClient() (*OutliersClient, error){
	addr := "localhost:9999"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//defer conn.Close()

	client := pb.NewOutliersClient(conn)
	return &OutliersClient{Client: client}, nil
}

func (t OutliersClient) GetEntityMeta(entityName string) (*ModelEntity, error){
	req := &pb.EntityInfoRequest{
		Name: entityName,
	}

	var modelEntity ModelEntity
	resp, err := t.Client.GetEntityInfo(context.Background(), req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err=json.Unmarshal([]byte(resp.Meta), &modelEntity)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		return nil, err
	}
	return &modelEntity, nil
}

