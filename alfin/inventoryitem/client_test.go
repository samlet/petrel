package inventoryitem

import (
	"fmt"
	"github.com/samlet/petrel/alfin"
	"github.com/samlet/petrel/services"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInventoryItemClient_CreateInventoryItem(t *testing.T) {
	client := newClient("")
	rs, err := client.CreateInventoryItem(&alfin.CreateInventoryItemParams{
		LocationSeqId:       "TLTLTLUL01",
		DatetimeReceived:    services.Timestamp{time.Now()},
		OwnerPartyId:        "Company",
		CurrencyUomId:       "USD",
		FacilityId:          "WebStoreWarehouse",
		UnitCost:            services.Decimal{decimal.NewFromFloat(1.234)},
		ProductId:           "MAT_A_COST",
		InventoryItemTypeId: "NON_SERIAL_INV_ITEM",
	})

	assert.Nil(t, err)
	assert.Equal(t, 200, rs.LastResponse.StatusCode)
	assert.NotNil(t, rs)
	println(rs.InventoryItemId)

	var invitem *services.Response
	invitem, err = FindInvItem(rs.InventoryItemId)
	if err != nil {
		panic(err)
	}
	alfin.Pretty(invitem.Data)
}

func FindInvItem(id string) (*services.Response, error) {
	body, err := services.InvokeService("performFindItem",
		fmt.Sprintf(`{
				"entityName":"InventoryItem",
				"inputFields":{
					"inventoryItemId":"%s"
				}
			}`, id))
	return services.WrapResult(err, body)
}
