package alfin

import (
	"github.com/araddon/dateparse"
	"log"
	"time"
)

func ToSecs(datetimeStr string) (int64, error) {
	t, err := dateparse.ParseAny(datetimeStr)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func ParseDateTime(datetimeStr string) time.Time {
	t, err := dateparse.ParseAny(datetimeStr)
	if err != nil {
		log.Fatalf(" fail: %v", err)
	}
	return t
}