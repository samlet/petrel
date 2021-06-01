package services

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"net/http"
	"time"
)

// APIResponse encapsulates some common features of a response from the
// Stripe API.
type APIResponse struct {
	// Header contain a map of all HTTP header keys to values. Its behavior and
	// caveats are identical to that of http.Header.
	Header http.Header

	// IdempotencyKey contains the idempotency key used with this request.
	// Idempotency keys are a Stripe-specific concept that helps guarantee that
	// requests that fail and need to be retried are not duplicated.
	IdempotencyKey string

	// RawJSON contains the response body as raw bytes.
	RawJSON []byte

	// RequestID contains a string that uniquely identifies the Stripe request.
	// Used for debugging or support purposes.
	RequestID string

	// Status is a status code and message. e.g. "200 OK"
	Status string

	// StatusCode is a status code as integer. e.g. 200
	StatusCode int
}

func newAPIResponse(res *http.Response, resBody []byte) *APIResponse {
	return &APIResponse{
		Header:         res.Header,
		IdempotencyKey: res.Header.Get("Idempotency-Key"),
		RawJSON:        resBody,
		RequestID:      res.Header.Get("Request-Id"),
		Status:         res.Status,
		StatusCode:     res.StatusCode,
	}
}

// APIResource is a type assigned to structs that may come from Alfin API
// endpoints and contains facilities common to all of them.
type APIResource struct {
	LastResponse      *APIResponse   `json:"-"`
	LastTypedResponse *TypedResponse `json:"-"`
}

// SetLastResponse sets the HTTP response that returned the API resource.
func (r *APIResource) SetLastResponse(response *APIResponse) {
	r.LastResponse = response
}

func (r *APIResource) SetLastTypedResponse(response *TypedResponse) {
	r.LastTypedResponse = response
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

type Decimal struct {
	decimal.Decimal
}

func (t Decimal) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *Decimal) UnmarshalJSON(data []byte) error {
	var s float64
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t.Decimal = decimal.NewFromFloat(s)
	return nil
}
