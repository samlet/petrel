package alfin

import (
	"context"
	"github.com/beevik/etree"
	"github.com/samlet/petrel/alfin/modules/workload"
	"github.com/samlet/petrel/alfin/modules/workload/ent"
	"log"
	"strconv"
	"testing"
)

func TestXmlSeed(t *testing.T) {
	xmlSeed := `<entity-engine-xml>
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

func TestWorkloadSeed(t *testing.T) {
	xmlSeed := `<entity-engine-xml>
	<WorkloadType workloadTypeId="REAL_WORLD" description="Real World"/>
    <WorkloadType workloadTypeId="MADE_UP" description="Made Up"/>
    <WorkloadType workloadTypeId="CONTRIVED" description="Contrived" parentTypeId="MADE_UP"/>
    <WorkloadType workloadTypeId="INSPIRED" description="Inspired" parentTypeId="MADE_UP"/>

    <WorkloadFeatureApplType workloadFeatureApplTypeId="REQUIRED" description="Required"/>
    <WorkloadFeatureApplType workloadFeatureApplTypeId="DESIRED" description="Desired"/>
    <WorkloadFeatureApplType workloadFeatureApplTypeId="NOT_ALLOWED" description="Not Allowed"/>

    <StatusType description="Workload" statusTypeId="WORKLOAD_STATUS" hasTable="N"/>
    <StatusItem description="In Design" sequenceId="01" statusCode="IN_DESIGN" statusId="WLST_IN_DESIGN" statusTypeId="WORKLOAD_STATUS"/>
    <StatusItem description="Defined" sequenceId="02" statusCode="DEFINED" statusId="WLST_DEFINED" statusTypeId="WORKLOAD_STATUS"/>
    <StatusItem description="Approved" sequenceId="03" statusCode="APPROVED" statusId="WLST_APPROVED" statusTypeId="WORKLOAD_STATUS"/>
    <StatusItem description="Implemented" sequenceId="04" statusCode="IMPLEMENTED" statusId="WLST_IMPLEMENTED" statusTypeId="WORKLOAD_STATUS"/>
    <StatusItem description="Tested" sequenceId="05" statusCode="TESTED" statusId="WLST_TESTED" statusTypeId="WORKLOAD_STATUS"/>
    <StatusItem description="Complete" sequenceId="06" statusCode="COMPLETE" statusId="WLST_COMPLETE" statusTypeId="WORKLOAD_STATUS"/>
    <StatusItem description="Cancelled" sequenceId="99" statusCode="CANCELLED" statusId="WLST_CANCELLED" statusTypeId="WORKLOAD_STATUS"/>
    <StatusValidChange statusId="WLST_IN_DESIGN" statusIdTo="WLST_DEFINED" transitionName="Definition Complete"/>
    <StatusValidChange statusId="WLST_DEFINED" statusIdTo="WLST_APPROVED" transitionName="Approve"/>
    <StatusValidChange statusId="WLST_APPROVED" statusIdTo="WLST_IMPLEMENTED" transitionName="Implementation Complete"/>
    <StatusValidChange statusId="WLST_IMPLEMENTED" statusIdTo="WLST_TESTED" transitionName="Testing Complete"/>
    <StatusValidChange statusId="WLST_TESTED" statusIdTo="WLST_COMPLETE" transitionName="Workload Completed"/>
    <StatusValidChange statusId="WLST_IN_DESIGN" statusIdTo="WLST_CANCELLED" transitionName="Cancel Workload"/>
    <StatusValidChange statusId="WLST_DEFINED" statusIdTo="WLST_CANCELLED" transitionName="Cancel Workload"/>
    <StatusValidChange statusId="WLST_APPROVED" statusIdTo="WLST_CANCELLED" transitionName="Cancel Workload"/>
    <StatusValidChange statusId="WLST_IMPLEMENTED" statusIdTo="WLST_CANCELLED" transitionName="Cancel Workload"/>
    <StatusValidChange statusId="WLST_TESTED" statusIdTo="WLST_CANCELLED" transitionName="Cancel Workload"/>

    <EnumerationType description="Workload Feature Source" enumTypeId="WORKL_FEAT_SOURCE" hasTable="N"/>
    <Enumeration description="Customer" enumCode="CUSTOMER" enumId="WLFTSRC_CUSTOMER" sequenceId="01" enumTypeId="WORKL_FEAT_SOURCE"/>
    <Enumeration description="Partner" enumCode="PARTNER" enumId="WLFTSRC_PARTNER" sequenceId="02" enumTypeId="WORKL_FEAT_SOURCE"/>
    <Enumeration description="Employee" enumCode="EMPLOYEE" enumId="WLFTSRC_EMPLOYEE" sequenceId="03" enumTypeId="WORKL_FEAT_SOURCE"/>

    <Workload workloadId="WL01" workloadName="Workload 1" workloadTypeId="CONTRIVED" statusId="WLST_IN_DESIGN"/>
    <Workload workloadId="WL02" workloadName="Workload 2" workloadTypeId="CONTRIVED" statusId="WLST_IN_DESIGN"/>
    <Workload workloadId="WL03" workloadName="Workload 3" workloadTypeId="CONTRIVED" statusId="WLST_IN_DESIGN"/>
    <Workload workloadId="WL04" workloadName="Workload 4" workloadTypeId="REAL_WORLD" statusId="WLST_IN_DESIGN"/>
    <Workload workloadId="WL05" workloadName="Workload 5" workloadTypeId="REAL_WORLD" statusId="WLST_IN_DESIGN"/>
    <Workload workloadId="WL06" workloadName="Workload 6" workloadTypeId="MADE_UP" statusId="WLST_IN_DESIGN"/>
    <Workload workloadId="WL07" workloadName="Workload 7" workloadTypeId="MADE_UP" statusId="WLST_IN_DESIGN"/>
    <Workload workloadId="WL08" workloadName="Workload 8" workloadTypeId="MADE_UP" statusId="WLST_IN_DESIGN"/>
    <Workload workloadId="WL09" workloadName="Workload 9" workloadTypeId="MADE_UP" statusId="WLST_IN_DESIGN"/>
    <Workload workloadId="WL10" workloadName="Workload 10" workloadTypeId="MADE_UP" statusId="WLST_IN_DESIGN"/>
    <Workload workloadId="WL11" workloadName="Workload 11" workloadTypeId="INSPIRED" statusId="WLST_IN_DESIGN" description="mMy inspired workload 11 description"/>
    <Workload workloadId="WL12" workloadName="Workload 12" workloadTypeId="INSPIRED" statusId="WLST_IN_DESIGN" description="mMy inspired workload 12 description"/>
    <create>
        <Workload workloadId="WL13" workloadName="Workload 13" workloadTypeId="INSPIRED" statusId="WLST_IN_DESIGN"/>
        <Workload workloadId="WL14" workloadName="Workload 14" workloadTypeId="INSPIRED" statusId="WLST_IN_DESIGN"/>
    </create>
    <create-replace>
        <Workload workloadId="WL12" workloadName="Workload 12 after replace" workloadTypeId="INSPIRED" statusId="WLST_IN_DESIGN"/>
    </create-replace>
    <create-update>
        <Workload workloadId="WL12" workloadName="Workload 11 after update" workloadTypeId="INSPIRED" statusId="WLST_IN_DESIGN"/>
    </create-update>
    <delete>
        <Workload workloadId="WL09"/>
    </delete>
    <Workload workloadId="WL10" workloadName="Workload 10 after update"/>

    <WorkloadItem workloadId="WL01" workloadItemSeqId="00001" description="WL1-001" amount="10"/>
    <WorkloadItem workloadId="WL01" workloadItemSeqId="00002" description="WL1-002" amount="20"/>
    <WorkloadItem workloadId="WL02" workloadItemSeqId="00001" description="WL2-001" amount="10"/>
    <WorkloadItem workloadId="WL02" workloadItemSeqId="00002" description="WL2-002" amount="20"/>
    <WorkloadItem workloadId="WL02" workloadItemSeqId="00003" description="WL2-003" amount="30"/>

    <WorkloadStatus workloadId="WL01" statusDate="2010-01-02 00:00:00" statusEndDate="2011-01-02 00:00:00" statusId="WLST_IN_DESIGN"/>
    <WorkloadStatus workloadId="WL01" statusDate="2011-01-02 00:00:01" statusEndDate="2012-01-02 00:00:00" statusId="WLST_DEFINED"/>
    <WorkloadStatus workloadId="WL01" statusDate="2012-01-02 00:00:00" statusEndDate="2013-01-02 00:00:00" statusId="WLST_APPROVED"/>
    <WorkloadStatus workloadId="WL02" statusDate="2010-01-02 00:00:00" statusEndDate="2011-01-02 00:00:00" statusId="WLST_IN_DESIGN"/>
    <WorkloadStatus workloadId="WL02" statusDate="2011-01-02 00:00:01" statusEndDate="2012-01-02 00:00:00" statusId="WLST_DEFINED"/>
    <WorkloadStatus workloadId="WL02" statusDate="2012-01-02 00:00:00" statusEndDate="2013-01-02 00:00:00" statusId="WLST_APPROVED"/>
</entity-engine-xml>`

	ctx := context.Background()
	store, err := workload.NewDevStore(ctx)
	if err != nil {
		panic(err)
	}
	ctx = context.WithValue(ctx, WorkloadStoreKey, store)
	ctx=context.WithValue(ctx, SeedsKey, NewSeedElements())

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
		switch node.Tag {
		case "WorkloadType":
			newWorkloadType(ctx, node)
		case "WorkloadItem":
			val, err := newWorkloadItem(ctx, node)
			if err != nil {
				panic(err)
			}
			log.Printf(".. store %s", val)
		default:
			log.Printf("Cannot handle entity %s", node.Tag)
		}
	}
}

func newWorkloadItem(ctx context.Context, node *etree.Element) (*ent.WorkloadItem, error) {
	store:=ctx.Value(WorkloadStoreKey).(*workload.WorkloadStore)
	seeds:=ctx.Value(SeedsKey).(*SeedElements)
	seqId, err := strconv.Atoi(node.SelectAttrValue("workloadItemSeqId", "0"))
	if err != nil {
		return nil, err
	}
	item, err := store.Client().WorkloadItem.Create().
		SetWorkloadItemSeqID(seqId).
		SetDescription(node.SelectAttrValue("description", "")).
		Save(ctx)
	if err != nil {
		log.Fatalf("save fail: %v", err)
	}
	itemId:=node.SelectAttrValue("workloadId","")+":"+node.SelectAttrValue("workloadItemSeqId","")
	seeds.Put(itemId, item)
	return item, nil
}

func newWorkloadType(ctx context.Context, node *etree.Element) (*ent.WorkloadType, error) {
	store:=ctx.Value(WorkloadStoreKey).(*workload.WorkloadStore)
	seeds:=ctx.Value(SeedsKey).(*SeedElements)
	recId:=node.SelectAttrValue("workloadTypeId", "")

	rec,err:= store.Client().WorkloadType.Create().
		SetDescription(node.SelectAttrValue("description", "")).
		Save(ctx)
	if err != nil {
		log.Fatalf(" fail: %v", err)
	}
	seeds.Put(recId, rec)
	return rec,nil
}

