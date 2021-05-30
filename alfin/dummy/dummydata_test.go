package dummy

import (
	"encoding/json"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestPrice_Unmarshal(t *testing.T) {
	priceData := map[string]interface{}{
		"id":     "price_123",
		"object": "price",
		"recurring": map[string]interface{}{
			"aggregate_usage": "last_during_period",
			"interval":        "month",
			"interval_count":  6,
			"usage_type":      "metered",
		},
		"tiers": []map[string]interface{}{
			{
				"flat_amount_decimal": "0.0111111111",
				"up_to":               5,
			},
			{
				"flat_amount_decimal": "0.0222222222",
				"up_to":               10,
			},
			{
				"flat_amount_decimal": "0.0333333333",
			},
		},
		"tiers_mode":          "volume",
		"unit_amount":         0,
		"unit_amount_decimal": "0.0123456789",
	}

	bytes, err := json.Marshal(&priceData)
	assert.NoError(t, err)

	var price Price
	err = json.Unmarshal(bytes, &price)
	assert.NoError(t, err)
}
