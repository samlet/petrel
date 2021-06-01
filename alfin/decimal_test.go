package alfin

import (
	"encoding/json"
	"fmt"
	"github.com/samlet/petrel/services"
	"github.com/shopspring/decimal"
	"testing"
)

type Tax struct {
	Amount services.Decimal `json:"amount"`
}

func TestDecimalEncoding(t *testing.T) {
	jsonStr := `{ "amount": "4242.44" }`
	var m Tax
	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		//panic(err)
	}
	fmt.Println(m)

	jsonStr = `{ "amount": 4242.44 }`
	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}
	fmt.Println(m)
}

func TestDecimal(t *testing.T) {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875
}
