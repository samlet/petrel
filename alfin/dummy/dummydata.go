package dummy

import "net/http"

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

// APIResource is a type assigned to structs that may come from Stripe API
// endpoints and contains facilities common to all of them.
type APIResource struct {
	LastResponse *APIResponse `json:"-"`
}

// SetLastResponse sets the HTTP response that returned the API resource.
func (r *APIResource) SetLastResponse(response *APIResponse) {
	r.LastResponse = response
}

// PriceBillingScheme is the list of allowed values for a price's billing scheme.
type PriceBillingScheme string

// Currency is the list of supported currencies.
// For more details see https://support.stripe.com/questions/which-currencies-does-stripe-support.
type Currency string

type Price struct {
	APIResource
	Active        bool               `json:"active"`
	BillingScheme PriceBillingScheme `json:"billing_scheme"`
	Created       int64              `json:"created"`
	Currency      Currency           `json:"currency"`
	Deleted       bool               `json:"deleted"`
	ID            string             `json:"id"`
	Livemode      bool               `json:"livemode"`
	LookupKey     string             `json:"lookup_key"`
	Metadata      map[string]string  `json:"metadata"`
	Nickname      string             `json:"nickname"`
	Object        string             `json:"object"`
}
