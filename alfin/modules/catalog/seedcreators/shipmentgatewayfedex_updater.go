package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateShipmentGatewayFedex(ctx context.Context) error {
	log.Println("ShipmentGatewayFedex updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentGatewayFedex
	failures := 0

	c = cache.Get("fedex_config__shipmentgatewayfedex").(*ent.ShipmentGatewayFedex)
	c, err = c.Update().
		SetConnectURL("https://gatewaybeta.fedex.com/GatewayDC").
		SetConnectSoapURL("https://gatewaybeta.fedex.com:443/web-services").
		SetConnectTimeout(60).
		SetLabelImageType("PNG").
		SetDefaultDropoffType("REGULARPICKUP").
		SetDefaultPackagingType("YOURPACKNG").
		SetTemplateShipment("component://product/template/shipment/FedexShipRequestTemplate.xml.ftl").
		SetTemplateSubscription("component://product/template/shipment/FedexSubscriptionRequestTemplate.xml.ftl").
		SetRateEstimateTemplate("component://product/template/shipment/FedexRateEstimateRequestTemplate.xml.ftl").
		SetShipmentGatewayConfig(cache.Get("fedex_config__shipmentgatewayconfig").(*ent.ShipmentGatewayConfig)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fedex_config__shipmentgatewayfedex: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fedex_config__shipmentgatewayfedex", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
