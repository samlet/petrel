package alfin

import (
	"github.com/beevik/etree"
	"log"
	"testing"
)

func TestXmlSeed(t *testing.T) {
	xmlSeed:=`<entity-engine-xml>
<!-- Purchase order test data -->
    <!--for jira issue - 1782-->
    <OrderHeader orderId="DEMO10091" orderTypeId="PURCHASE_ORDER" orderName="New Purchase Order" salesChannelEnumId="UNKNWN_SALES_CHANNEL" orderDate="2008-06-10 13:27:07.024" entryDate="2008-06-10 13:27:07.024" visitId="10000" statusId="ORDER_CREATED" createdBy="admin" currencyUom="USD" productStoreId="9000" remainingSubTotal="108.0" grandTotal="108.0" webSiteId="WebStore"/>
    <OrderItem orderId="DEMO10091" orderItemSeqId="00001" orderItemTypeId="PRODUCT_ORDER_ITEM" productId="GZ-2644" prodCatalogId="DemoCatalog" isPromo="N" quantity="5.0" selectedAmount="0.0" unitPrice="21.6" unitListPrice="0.0" isModifiedPrice="N" itemDescription="GZ-2644-5 Round Gizmo" statusId="ITEM_CREATED"/>
    <OrderItemPriceInfo orderItemPriceInfoId="10001" orderId="DEMO10091" orderItemSeqId="00001" description="SupplierProduct [minimumOrderQuantity:5.0, lastPrice: 21.6]"/>
    <OrderItemPriceInfo orderItemPriceInfoId="10002" orderId="DEMO10091" orderItemSeqId="00001" description="SupplierProduct [minimumOrderQuantity:0.0, lastPrice: 24.0]"/>
    <OrderRole orderId="DEMO10091" partyId="Company" roleTypeId="BILL_TO_CUSTOMER"/>
    <OrderRole orderId="DEMO10091" partyId="DemoSupplier" roleTypeId="BILL_FROM_VENDOR"/>
    <OrderRole orderId="DEMO10091" partyId="DemoSupplier" roleTypeId="SHIP_FROM_VENDOR"/>
    <OrderRole orderId="DEMO10091" partyId="DemoSupplier" roleTypeId="SUPPLIER_AGENT"/>
    <OrderItemShipGroup orderId="DEMO10091" shipGroupSeqId="00001" shipmentMethodTypeId="NO_SHIPPING" carrierPartyId="_NA_" carrierRoleTypeId="CARRIER" contactMechId="9200" maySplit="N" isGift="N"/>
    <OrderItemShipGroupAssoc orderId="DEMO10091" orderItemSeqId="00001" shipGroupSeqId="00001" quantity="5.0"/>
</entity-engine-xml>`

	doc := etree.NewDocument()
	if err := doc.ReadFromString(xmlSeed); err != nil {
		panic(err)
	}
	root := doc.SelectElement("entity-engine-xml")
	log.Println("ROOT element:", root.Tag)
	for _, node := range root.ChildElements() {
		log.Println("CHILD element:", node.Tag)
		for _, att := range node.Attr {
			log.Println("\t-", att.Key, att.Value)
		}
	}
}

