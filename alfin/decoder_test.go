package alfin

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

type BankAccount struct {
	ID            string `json:"id"`
	Object        string `json:"object"`
	RoutingNumber string `json:"routing_number"`
}

type Card struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Last4  string `json:"last4"`
}

// 我们在这里使用指针是因为我们永远不会初始化这两个指针。相反，我们将确定需要使用哪一个，并将其初始化。
//这意味着在确实使用这种类型的代码中，您可能需要编写类似的代码if data.Card != nil { ... }来确定它是否为卡片。
//另外，您也可以将Object属性存储在Data类型本身上。该选择最终取决于您，并且可能需要进行一些较小的代码调整，
//但是此处讨论的总体方法仍然可以使用。
type Data struct {
	*Card
	*BankAccount
}

// Data并不意味着要成为整个JSON结构，而是要表示存储在dataJSON对象的键中的所有内容。
//需要特别注意的是，因为在我们的代码中，我们的Data类型同时具有Card和BankAccount指针，
//但是在JSON中，它们不会是嵌套对象。这意味着我们需要做的第一件事是编写一种MarshalJSON()方法来反映这一点。
func (d Data) MarshalJSON() ([]byte, error) {
	if d.Card != nil {
		return json.Marshal(d.Card)
	} else if d.BankAccount != nil {
		return json.Marshal(d.BankAccount)
	} else {
		return json.Marshal(nil)
	}
}
func (d *Data) UnmarshalJSON(data []byte) error {
	temp := struct {
		Object string `json:"object"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Object == "card" {
		var c Card
		if err := json.Unmarshal(data, &c); err != nil {
			return err
		}
		d.Card = &c
		d.BankAccount = nil
	} else if temp.Object == "bank_account" {
		var ba BankAccount
		if err := json.Unmarshal(data, &ba); err != nil {
			return err
		}
		d.BankAccount = &ba
		d.Card = nil
	} else {
		return errors.New("Invalid object value")
	}
	return nil
}

func TestDecoder(t *testing.T) {
	jsonStr := `
{
  "data": {
    "object": "card",
    "id": "card_123",
    "last4": "4242"
  }
}
`
	var m map[string]Data
	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}
	fmt.Println(m)
	data := m["data"]
	if data.Card != nil {
		fmt.Println(data.Card)
	}
	if data.BankAccount != nil {
		fmt.Println(data.BankAccount)
	}

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
