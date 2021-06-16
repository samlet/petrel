package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateShipmentGatewayFedex(ctx context.Context) error {
	log.Println("ShipmentGatewayFedex creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentGatewayFedex

	c, err = client.ShipmentGatewayFedex.Create().SetStringRef("fedex_config__shipmentgatewayfedex").
		SetConnectURL("https://gatewaybeta.fedex.com/GatewayDC").
		SetConnectSoapURL("https://gatewaybeta.fedex.com:443/web-services").
		SetConnectTimeout(60).
		SetLabelImageType("PNG").
		SetDefaultDropoffType("REGULARPICKUP").
		SetDefaultPackagingType("YOURPACKNG").
		SetTemplateShipment("component://product/template/shipment/FedexShipRequestTemplate.xml.ftl").
		SetTemplateSubscription("component://product/template/shipment/FedexSubscriptionRequestTemplate.xml.ftl").
		SetRateEstimateTemplate("component://product/template/shipment/FedexRateEstimateRequestTemplate.xml.ftl").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fedex_config__shipmentgatewayfedex: %v", err)
		return err
	}
	cache.Put("fedex_config__shipmentgatewayfedex", c)

	return nil
}
