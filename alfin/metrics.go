package alfin

import (
	"github.com/samlet/petrel/alfin/pb"
	pbtime "google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"time"
)

func dummyData() []*pb.Metric {
	const size = 1000
	out := make([]*pb.Metric, size)
	t := time.Date(2020, 5, 22, 14, 13, 11, 0, time.UTC)
	for i := 0; i < size; i++ {
		m := pb.Metric{
			Time: Timestamp(t),
			Name: "CPU",
			// normally we're below 40% CPU utilization
			Value: rand.Float64() * 40,
		}
		out[i] = &m
		t.Add(time.Second)
	}
	// Create some outliers
	out[7].Value = 97.3
	out[113].Value = 92.1
	out[835].Value = 93.2
	return out
}

// Timestamp converts time.Time to protobuf *Timestamp
func Timestamp(t time.Time) *pbtime.Timestamp {
	return &pbtime.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

