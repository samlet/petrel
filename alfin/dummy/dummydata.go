package dummy

import (
	"github.com/samlet/petrel/alfin"
	"github.com/samlet/petrel/services"
)

// PriceBillingScheme is the list of allowed values for a price's billing scheme.
type PriceBillingScheme string

type Price struct {
	services.APIResource
	Active        bool               `json:"active"`
	BillingScheme PriceBillingScheme `json:"billing_scheme"`
	Created       int64              `json:"created"`
	Currency      alfin.Currency     `json:"currency"`
	Deleted       bool               `json:"deleted"`
	ID            string             `json:"id"`
	Livemode      bool               `json:"livemode"`
	LookupKey     string             `json:"lookup_key"`
	Metadata      map[string]string  `json:"metadata"`
	Nickname      string             `json:"nickname"`
	Object        string             `json:"object"`
}
