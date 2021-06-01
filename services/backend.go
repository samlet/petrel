package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

const (
	// APIVersion is the currently supported API version
	APIVersion string = "2021-06-01"
)

type Params struct {
	Metadata map[string]string
	Payload  string
}

func (c *Params) GetParams() *Params {
	return c
}

// ParamsContainer is a general interface for which all parameter structs
// should comply. They achieve this by embedding a Params struct and inheriting
// its implementation of this interface.
type ParamsContainer interface {
	GetParams() *Params
}

// LastResponseSetter defines a type that contains an HTTP response from a Stripe
// API endpoint.
type LastResponseSetter interface {
	SetLastResponse(response *APIResponse)
	SetLastTypedResponse(response *TypedResponse)
}

type Backend interface {
	Call(method, path, key string, params ParamsContainer, v LastResponseSetter) error
}

// NewIdempotencyKey generates a new idempotency key that
// can be used on a request.
func NewIdempotencyKey() string {
	now := time.Now().UnixNano()
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v_%v", now, base64.URLEncoding.EncodeToString(buf)[:6])
}
