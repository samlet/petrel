package alfin

import (
	"encoding/json"
	"time"
)

type APIResponse struct {
	// Status is a status code and message. e.g. "200 OK"
	Status string
	// StatusCode is a status code as integer. e.g. 200
	StatusCode int
}

type MetaValue struct {
	Type       string                 `json:"type"`
	EntityName string                 `json:"entityName"`
	Value      map[string]interface{} `json:"value"`
}

type Timestamp struct {
	time.Time
}

type DateTime struct {
	time.Time
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Unix())
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var i int64
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	t.Time = time.Unix(i, 0)
	return nil
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format(time.RFC3339))
}

func (t *DateTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	result, err := time.Parse(time.RFC3339, s)
	if err == nil {
		t.Time = result
	}
	return err
}
