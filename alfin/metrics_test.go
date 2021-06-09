package alfin

import (
	"context"
	"github.com/samlet/petrel/alfin/pb"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestMetrics(t *testing.T) {
	addr := "localhost:9999"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewOutliersClient(conn)
	req := &pb.OutliersRequest{
		Metrics: dummyData(),
	}

	resp, err := client.Detect(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("outliers at: %v", resp.Indices)
}

func BenchmarkClient(b *testing.B) {
	require := require.New(b)

	addr := "localhost:9999"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	require.NoError(err, "connect")
	defer conn.Close()

	client := pb.NewOutliersClient(conn)
	req := &pb.OutliersRequest{
		Metrics: dummyData(),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Detect(context.Background(), req)
		require.NoError(err, "detect")
	}

}
