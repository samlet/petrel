package alfin

import "github.com/araddon/dateparse"

func ToSecs(datetimeStr string) (int64, error) {
	t, err := dateparse.ParseAny(datetimeStr)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

