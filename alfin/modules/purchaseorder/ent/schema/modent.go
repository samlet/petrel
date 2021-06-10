package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    // "entgo.io/ent/schema/index"
    "entgo.io/ent/schema/mixin"
    "time"
)


type OrderItemShipGroup struct {
    ent.Schema
}

// Fields of the OrderItemShipGroup.
func (OrderItemShipGroup) Fields() []ent.Field {
    return []ent.Field{
        field.Int("ship_group_seq_id"),
        field.Int("shipment_method_type_id").Optional(),
        field.Int("supplier_party_id").Optional(),
        field.Int("supplier_agreement_id").Optional(),
        field.Int("vendor_party_id").Optional(),
        field.Int("carrier_party_id").Optional(),
        field.Int("carrier_role_type_id").Optional(),
        field.Int("facility_id").Optional(),
        field.Int("contact_mech_id").Optional(),
        field.Int("telecom_contact_mech_id").Optional(),
        field.String("tracking_number").Optional(),
        field.String("shipping_instructions").Optional(),
        field.Enum("may_split").
            Values("Y", "N", "-").Optional(),
        field.String("gift_message").Optional(),
        field.Enum("is_gift").
            Values("Y", "N", "-").Optional(),
        field.Time("ship_after_date").
            Default(time.Now).Optional(),
        field.Time("ship_by_date").
            Default(time.Now).Optional(),
        field.Time("estimated_ship_date").
            Default(time.Now).Optional(),
        field.Time("estimated_delivery_date").
            Default(time.Now).Optional(),
    }
}

func (OrderItemShipGroup) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
func (OrderItemShipGroup) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("order_id", "ship_group_seq_id").
            Unique(),
    }
}

*/



// Edges of the OrderItemShipGroup.
//  one: OrderHeader
//  one: SupplierParty
//  one: SupplierAgreement
//  one: VendorParty
//  one: CarrierShipmentMethod
//  one: CarrierParty
//  one: CarrierPartyRole
//  one: Facility
//  one: ShipmentMethodType
//  one: ContactMech
//  one: PostalAddress
//  one: TelecomContactMech
//  one: TelecomTelecomNumber
//  many: OrderAdjustment
//  many: FromOrderItemAssoc
//  many: ToOrderItemAssoc
//  many: OrderItemShipGroupAssoc
//  many: OrderItemShipGrpInvRes
//  many: OrderPaymentPreference
//  many: PrimaryPicklistBin
//  many: PicklistItem
//  many: PrimaryShipment
func (OrderItemShipGroup) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("order_header", OrderHeader.Type).Ref("order_item_ship_groups").
                    // Bind the "orderId" field to this edge.
                    // Field("order_id").
                    Unique(),
                edge.To("order_item_ship_group_assocs", OrderItemShipGroupAssoc.Type),
                edge.To("primary_shipments", Shipment.Type),
    }
}


type ShipmentItem struct {
    ent.Schema
}

// Fields of the ShipmentItem.
func (ShipmentItem) Fields() []ent.Field {
    return []ent.Field{
        field.Int("shipment_item_seq_id"),
        field.Int("product_id").Optional(),
        field.Float("quantity").Optional(),
        field.String("shipment_content_description").Optional(),
    }
}

func (ShipmentItem) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
func (ShipmentItem) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("shipment_id", "shipment_item_seq_id").
            Unique(),
    }
}

*/



// Edges of the ShipmentItem.
//  one: Shipment
//  one: Product
//  many: ItemIssuance
//  many: OrderShipment
//  many: ReturnItemShipment
//  many: ShipmentItemBilling
//  many: ShipmentItemFeature
//  many: ShipmentPackageContent
//  many: ShipmentReceipt
//  many: ShippingDocument
func (ShipmentItem) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("shipment", Shipment.Type).Ref("shipment_items").
                    // Bind the "shipmentId" field to this edge.
                    // Field("shipment_id").
                    Unique(),
                edge.To("item_issuances", ItemIssuance.Type),
    }
}


type ItemIssuance struct {
    ent.Schema
}

// Fields of the ItemIssuance.
func (ItemIssuance) Fields() []ent.Field {
    return []ent.Field{
        field.Int("order_item_seq_id").Optional(),
        field.Int("ship_group_seq_id").Optional(),
        field.Int("inventory_item_id").Optional(),
        field.Int("shipment_item_seq_id").Optional(),
        field.Int("fixed_asset_id").Optional(),
        field.Int("maint_hist_seq_id").Optional(),
        field.Time("issued_date_time").
            Default(time.Now).Optional(),
        field.String("issued_by_user_login_id").Optional(),
        field.Float("quantity").Optional(),
        field.Float("cancel_quantity").Optional(),
    }
}

func (ItemIssuance) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
*/



// Edges of the ItemIssuance.
//  one: InventoryItem
//  one-nofk: OrderItemShipGrpInvRes
//  one-nofk: Shipment
//  one: ShipmentItem
//  one: FixedAssetMaint
//  one-nofk: OrderHeader
//  one: OrderItem
//  one: IssuedByUserLogin
//  many: InventoryItemDetail
//  many: InventoryTransfer
//  many: ItemIssuanceRole
//  many: OrderItemBilling
func (ItemIssuance) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("shipment", Shipment.Type).Ref("item_issuances").
                    // Bind the "shipmentId" field to this edge.
                    // Field("shipment_id").
                    Unique(),
                edge.From("shipment_item", ShipmentItem.Type).Ref("item_issuances").
                    // Bind the "shipmentId" field to this edge.
                    // Field("shipment_id").
                    Unique(),
                edge.From("order_header", OrderHeader.Type).Ref("item_issuances").
                    // Bind the "orderId" field to this edge.
                    // Field("order_id").
                    Unique(),
                edge.From("order_item", OrderItem.Type).Ref("item_issuances").
                    // Bind the "orderId" field to this edge.
                    // Field("order_id").
                    Unique(),
    }
}


type OrderStatus struct {
    ent.Schema
}

// Fields of the OrderStatus.
func (OrderStatus) Fields() []ent.Field {
    return []ent.Field{
        field.Int("status_id").Optional(),
        field.Int("order_item_seq_id").Optional(),
        field.Int("order_payment_preference_id").Optional(),
        field.Time("status_datetime").
            Default(time.Now).Optional(),
        field.String("status_user_login").Optional(),
        field.String("change_reason").Optional(),
    }
}

func (OrderStatus) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
*/



// Edges of the OrderStatus.
//  one: StatusItem
//  one: OrderHeader
//  one-nofk: OrderItem
//  one-nofk: OrderPaymentPreference
//  one: UserLogin
func (OrderStatus) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("order_header", OrderHeader.Type).Ref("order_statuses").
                    // Bind the "orderId" field to this edge.
                    // Field("order_id").
                    Unique(),
                edge.From("order_item", OrderItem.Type).Ref("order_statuses").
                    // Bind the "orderId" field to this edge.
                    // Field("order_id").
                    Unique(),
    }
}


type Shipment struct {
    ent.Schema
}

// Fields of the Shipment.
func (Shipment) Fields() []ent.Field {
    return []ent.Field{
        field.Int("shipment_type_id").Optional(),
        field.Int("status_id").Optional(),
        field.Int("primary_return_id").Optional(),
        field.Int("primary_ship_group_seq_id").Optional(),
        field.Int("picklist_bin_id").Optional(),
        field.Time("estimated_ready_date").
            Default(time.Now).Optional(),
        field.Time("estimated_ship_date").
            Default(time.Now).Optional(),
        field.Int("estimated_ship_work_eff_id").Optional(),
        field.Time("estimated_arrival_date").
            Default(time.Now).Optional(),
        field.Int("estimated_arrival_work_eff_id").Optional(),
        field.Time("latest_cancel_date").
            Default(time.Now).Optional(),
        field.Float("estimated_ship_cost").Optional(),
        field.Int("currency_uom_id").Optional(),
        field.String("handling_instructions").Optional(),
        field.Int("origin_facility_id").Optional(),
        field.Int("destination_facility_id").Optional(),
        field.Int("origin_contact_mech_id").Optional(),
        field.Int("origin_telecom_number_id").Optional(),
        field.Int("destination_contact_mech_id").Optional(),
        field.Int("destination_telecom_number_id").Optional(),
        field.Int("party_id_to").Optional(),
        field.Int("party_id_from").Optional(),
        field.Float("additional_shipping_charge").Optional(),
        field.String("addtl_shipping_charge_desc").Optional(),
        field.Time("created_date").
            Default(time.Now).Optional(),
        field.String("created_by_user_login").Optional(),
        field.Time("last_modified_date").
            Default(time.Now).Optional(),
        field.String("last_modified_by_user_login").Optional(),
    }
}

func (Shipment) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
*/



// Edges of the Shipment.
//  one: ShipmentType
//  one: StatusItem
//  one: EstimatedShipWorkEffort
//  one: EstimatedArrivalWorkEffort
//  one: CurrencyUom
//  one: OriginFacility
//  one: DestinationFacility
//  one-nofk: OriginContactMech
//  one-nofk: DestContactMech
//  one: OriginPostalAddress
//  one: OriginTelecomNumber
//  one: DestinationPostalAddress
//  one: DestinationTelecomNumber
//  one: PrimaryOrderHeader
//  one: PrimaryReturnHeader
//  one: PicklistBin
//  one-nofk: PrimaryOrderItemShipGroup
//  many: ShipmentTypeAttr
//  one: ToParty
//  one-nofk: ToPerson
//  one-nofk: ToPartyGroup
//  one: FromParty
//  one-nofk: FromPerson
//  one-nofk: FromPartyGroup
//  many: ShipmentManifestView
//  many: AcctgTrans
//  many: ItemIssuance
//  many: OrderShipment
//  many: ReturnItemShipment
//  many: ShipmentAttribute
//  many: ShipmentContactMech
//  many: ShipmentItem
//  many: ShipmentItemBilling
//  many: ShipmentItemFeature
//  many: ShipmentPackage
//  many: ShipmentPackageContent
//  many: ShipmentPackageRouteSeg
//  many: ShipmentReceipt
//  many: ShipmentRouteSegment
//  many: ShipmentStatus
func (Shipment) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("primary_order_header", OrderHeader.Type).Ref("primary_shipments").
                    // Bind the "primaryOrderId" field to this edge.
                    // Field("primary_order_id").
                    Unique(),
                edge.From("primary_order_item_ship_group", OrderItemShipGroup.Type).Ref("primary_shipments").
                    // Bind the "primaryOrderId" field to this edge.
                    // Field("primary_order_id").
                    Unique(),
                edge.To("item_issuances", ItemIssuance.Type),
                edge.To("shipment_items", ShipmentItem.Type),
    }
}


type OrderItemShipGroupAssoc struct {
    ent.Schema
}

// Fields of the OrderItemShipGroupAssoc.
func (OrderItemShipGroupAssoc) Fields() []ent.Field {
    return []ent.Field{
        field.Int("order_item_seq_id"),
        field.Int("ship_group_seq_id"),
        field.Float("quantity").Optional(),
        field.Float("cancel_quantity").Optional(),
    }
}

func (OrderItemShipGroupAssoc) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
func (OrderItemShipGroupAssoc) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("order_id", "order_item_seq_id", "ship_group_seq_id").
            Unique(),
    }
}

*/



// Edges of the OrderItemShipGroupAssoc.
//  one: OrderHeader
//  one: OrderItem
//  one: OrderItemShipGroup
//  many: OrderAdjustment
//  many: FromOrderItemAssoc
//  many: ToOrderItemAssoc
//  many: OrderItemShipGrpInvRes
//  many: OrderShipment
//  many: WorkOrderItemFulfillment
func (OrderItemShipGroupAssoc) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("order_header", OrderHeader.Type).Ref("order_item_ship_group_assocs").
                    // Bind the "orderId" field to this edge.
                    // Field("order_id").
                    Unique(),
                edge.From("order_item", OrderItem.Type).Ref("order_item_ship_group_assocs").
                    // Bind the "orderId" field to this edge.
                    // Field("order_id").
                    Unique(),
                edge.From("order_item_ship_group", OrderItemShipGroup.Type).Ref("order_item_ship_group_assocs").
                    // Bind the "orderId" field to this edge.
                    // Field("order_id").
                    Unique(),
    }
}


type OrderRole struct {
    ent.Schema
}

// Fields of the OrderRole.
func (OrderRole) Fields() []ent.Field {
    return []ent.Field{
        field.Int("party_id"),
        field.Int("role_type_id"),
    }
}

func (OrderRole) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
func (OrderRole) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("order_id", "party_id", "role_type_id").
            Unique(),
    }
}

*/



// Edges of the OrderRole.
//  one: OrderHeader
//  one: Party
//  one: PartyRole
//  one-nofk: RoleType
//  many: OrderItem
func (OrderRole) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("order_header", OrderHeader.Type).Ref("order_roles").
                    // Bind the "orderId" field to this edge.
                    // Field("order_id").
                    Unique(),
                edge.To("order_items", OrderItem.Type),
    }
}


type OrderItemPriceInfo struct {
    ent.Schema
}

// Fields of the OrderItemPriceInfo.
func (OrderItemPriceInfo) Fields() []ent.Field {
    return []ent.Field{
        field.Int("order_item_seq_id").Optional(),
        field.Int("product_price_rule_id").Optional(),
        field.Int("product_price_action_seq_id").Optional(),
        field.Float("modify_amount").Optional(),
        field.String("description").Optional(),
        field.String("rate_code").Optional(),
    }
}

func (OrderItemPriceInfo) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
*/



// Edges of the OrderItemPriceInfo.
//  one-nofk: OrderHeader
//  one: OrderItem
//  one-nofk: ProductPriceRule
//  one: ProductPriceAction
func (OrderItemPriceInfo) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("order_header", OrderHeader.Type).Ref("order_item_price_infos").
                    // Bind the "orderId" field to this edge.
                    // Field("order_id").
                    Unique(),
                edge.From("order_item", OrderItem.Type).Ref("order_item_price_infos").
                    // Bind the "orderId" field to this edge.
                    // Field("order_id").
                    Unique(),
    }
}


type Payment struct {
    ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
    return []ent.Field{
        field.Int("payment_type_id").Optional(),
        field.Int("payment_method_type_id").Optional(),
        field.Int("payment_method_id").Optional(),
        field.Int("payment_gateway_response_id").Optional(),
        field.Int("payment_preference_id").Optional(),
        field.Int("party_id_from").Optional(),
        field.Int("party_id_to").Optional(),
        field.Int("role_type_id_to").Optional(),
        field.Int("status_id").Optional(),
        field.Time("effective_date").
            Default(time.Now).Optional(),
        field.String("payment_ref_num").Optional(),
        field.Float("amount").Optional(),
        field.Int("currency_uom_id").Optional(),
        field.String("comments").Optional(),
        field.Int("fin_account_trans_id").Optional(),
        field.Int("override_gl_account_id").Optional(),
        field.Float("actual_currency_amount").Optional(),
        field.Int("actual_currency_uom_id").Optional(),
    }
}

func (Payment) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
*/



// Edges of the Payment.
//  one: PaymentType
//  many: PaymentTypeAttr
//  one: PaymentMethodType
//  one: PaymentMethod
//  one: CurrencyUom
//  one: ActualCurrencyUom
//  one-nofk: CreditCard
//  one-nofk: EftAccount
//  one-nofk: GiftCard
//  one: OrderPaymentPreference
//  one: PaymentGatewayResponse
//  one: FromParty
//  one: ToParty
//  one: ToRoleType
//  one-nofk: ToPartyRole
//  one: StatusItem
//  one: FinAccountTrans
//  one: GlAccount
//  many: AcctgTrans
//  many: Deduction
//  many: PaymentApplication
//  many: ToPaymentApplication
//  many: PaymentAttribute
//  many: PaymentBudgetAllocation
//  many: PaymentContent
//  many: PaymentGroupMember
//  many: PerfReview
//  many: ReturnItemResponse
func (Payment) Edges() []ent.Edge {
    return []ent.Edge{
    }
}


type OrderItem struct {
    ent.Schema
}

// Fields of the OrderItem.
func (OrderItem) Fields() []ent.Field {
    return []ent.Field{
        field.Int("order_item_seq_id"),
        field.Int("external_id").Optional(),
        field.Int("order_item_type_id").Optional(),
        field.Int("order_item_group_seq_id").Optional(),
        field.Enum("is_item_group_primary").
            Values("Y", "N", "-").Optional(),
        field.Int("from_inventory_item_id").Optional(),
        field.Int("budget_id").Optional(),
        field.Int("budget_item_seq_id").Optional(),
        field.Int("product_id").Optional(),
        field.String("supplier_product_id").MaxLen(32).Optional(),
        field.Int("product_feature_id").Optional(),
        field.Int("prod_catalog_id").Optional(),
        field.Int("product_category_id").Optional(),
        field.Enum("is_promo").
            Values("Y", "N", "-").Optional(),
        field.Int("quote_id").Optional(),
        field.Int("quote_item_seq_id").Optional(),
        field.Int("shopping_list_id").Optional(),
        field.Int("shopping_list_item_seq_id").Optional(),
        field.Int("subscription_id").Optional(),
        field.Int("deployment_id").Optional(),
        field.Float("quantity").Optional(),
        field.Float("cancel_quantity").Optional(),
        field.Float("selected_amount").Optional(),
        field.Float("unit_price").Optional(),
        field.Float("unit_list_price").Optional(),
        field.Float("unit_average_cost").Optional(),
        field.Float("unit_recurring_price").Optional(),
        field.Enum("is_modified_price").
            Values("Y", "N", "-").Optional(),
        field.Int("recurring_freq_uom_id").Optional(),
        field.String("item_description").Optional(),
        field.String("comments").Optional(),
        field.Int("corresponding_po_id").Optional(),
        field.Int("status_id").Optional(),
        field.Int("sync_status_id").Optional(),
        field.Time("estimated_ship_date").
            Default(time.Now).Optional(),
        field.Time("estimated_delivery_date").
            Default(time.Now).Optional(),
        field.Time("auto_cancel_date").
            Default(time.Now).Optional(),
        field.Time("dont_cancel_set_date").
            Default(time.Now).Optional(),
        field.String("dont_cancel_set_user_login").Optional(),
        field.Time("ship_before_date").
            Default(time.Now).Optional(),
        field.Time("ship_after_date").
            Default(time.Now).Optional(),
        field.Time("reserve_after_date").
            Default(time.Now).Optional(),
        field.Time("cancel_back_order_date").
            Default(time.Now).Optional(),
        field.Int("override_gl_account_id").Optional(),
        field.Int("sales_opportunity_id").Optional(),
        field.String("change_by_user_login_id").Optional(),
    }
}

func (OrderItem) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
func (OrderItem) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("order_id", "order_item_seq_id").
            Unique(),
    }
}

*/



// Edges of the OrderItem.
//  one: OrderHeader
//  one: OrderItemType
//  one: OrderItemGroup
//  many: OrderItemTypeAttr
//  one: Product
//  one: FromInventoryItem
//  one: RecurringFreqUom
//  one: StatusItem
//  many: ProductFacilityLocation
//  many: StatusValidChange
//  one: SyncStatusItem
//  one: DontCancelSetUserLogin
//  one: QuoteItem
//  one-nofk: ShoppingListItem
//  one: OverrideGlAccount
//  one: SalesOpportunity
//  one: ChangeByUserLogin
//  many: Order ItemAllocationPlanItem
//  many: FinAccountTrans
//  many: AcquireFixedAsset
//  many: FixedAssetMaintOrder
//  many: GiftCardFulfillment
//  many: ItemIssuance
//  many: OrderAdjustment
//  one-nofk: OrderDeliverySchedule
//  many: FromOrderItemAssoc
//  many: ToOrderItemAssoc
//  many: OrderItemAttribute
//  many: OrderItemBilling
//  many: OrderItemChange
//  many: OrderItemContactMech
//  many: OrderItemGroupOrder
//  many: OrderItemPriceInfo
//  many: OrderItemRole
//  many: OrderItemShipGroupAssoc
//  many: OrderItemShipGrpInvRes
//  many: OrderPaymentPreference
//  many: OrderRequirementCommitment
//  many: OrderShipment
//  many: OrderStatus
//  many: OrderTerm
//  many: PicklistItem
//  many: ProductOrderItem
//  many: EngagementProductOrderItem
//  many: ReturnItem
//  many: ShipmentReceipt
//  many: Subscription
//  many: SurveyResponse
//  many: WorkOrderItemFulfillment
func (OrderItem) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("order_header", OrderHeader.Type).Ref("order_items").
                    // Bind the "orderId" field to this edge.
                    // Field("order_id").
                    Unique(),
                edge.To("item_issuances", ItemIssuance.Type),
                edge.To("order_item_price_infos", OrderItemPriceInfo.Type),
                edge.To("order_item_ship_group_assocs", OrderItemShipGroupAssoc.Type),
                edge.To("order_statuses", OrderStatus.Type),
    }
}


type OrderHeader struct {
    ent.Schema
}

// Fields of the OrderHeader.
func (OrderHeader) Fields() []ent.Field {
    return []ent.Field{
        field.Int("order_type_id").Optional(),
        field.String("order_name").Optional(),
        field.Int("external_id").Optional(),
        field.Int("sales_channel_enum_id").Optional(),
        field.Time("order_date").
            Default(time.Now).Optional(),
        field.Enum("priority").
            Values("Y", "N", "-").Optional(),
        field.Time("entry_date").
            Default(time.Now).Optional(),
        field.Time("pick_sheet_printed_date").
            Default(time.Now).Optional(),
        field.Int("visit_id").Optional(),
        field.Int("status_id").Optional(),
        field.String("created_by").Optional(),
        field.Int("first_attempt_order_id").Optional(),
        field.Int("currency_uom").Optional(),
        field.Int("sync_status_id").Optional(),
        field.Int("billing_account_id").Optional(),
        field.Int("origin_facility_id").Optional(),
        field.Int("web_site_id").Optional(),
        field.Int("product_store_id").Optional(),
        field.Int("agreement_id").Optional(),
        field.String("terminal_id").MaxLen(32).Optional(),
        field.String("transaction_id").MaxLen(32).Optional(),
        field.Int("auto_order_shopping_list_id").Optional(),
        field.Enum("needs_inventory_issuance").
            Values("Y", "N", "-").Optional(),
        field.Enum("is_rush_order").
            Values("Y", "N", "-").Optional(),
        field.String("internal_code").MaxLen(32).Optional(),
        field.Float("remaining_sub_total").Optional(),
        field.Float("grand_total").Optional(),
        field.Enum("is_viewed").
            Values("Y", "N", "-").Optional(),
        field.Enum("invoice_per_shipment").
            Values("Y", "N", "-").Optional(),
    }
}

func (OrderHeader) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
*/



// Edges of the OrderHeader.
//  one: OrderType
//  one: SalesChannelEnumeration
//  one: OriginFacility
//  many: OrderTypeAttr
//  one: BillingAccount
//  one: ProductStore
//  one: AutoOrderShoppingList
//  one: CreatedByUserLogin
//  one: StatusItem
//  one: SyncStatusItem
//  one: Uom
//  one: WebSite
//  many: Order IdAllocationPlanItem
//  many: CommunicationEventOrder
//  many: AcquireFixedAsset
//  many: PurchaseFixedAssetMaint
//  many: FixedAssetMaintOrder
//  many: GiftCardFulfillment
//  many: ItemIssuance
//  many: OrderAdjustment
//  many: OrderAttribute
//  many: OrderContactMech
//  many: OrderContent
//  many: OrderDeliverySchedule
//  many: OrderHeaderNote
//  many: OrderHeaderWorkEffort
//  many: OrderItem
//  many: FromOrderItemAssoc
//  many: ToOrderItemAssoc
//  many: OrderItemBilling
//  many: OrderItemChange
//  many: OrderItemContactMech
//  many: OrderItemGroup
//  many: OrderItemPriceInfo
//  many: OrderItemRole
//  many: OrderItemShipGroup
//  many: OrderItemShipGroupAssoc
//  many: OrderItemShipGrpInvRes
//  many: OrderNotification
//  many: OrderPaymentPreference
//  many: OrderProductPromoCode
//  many: OrderRequirementCommitment
//  many: OrderRole
//  many: OrderShipment
//  many: OrderStatus
//  many: OrderTerm
//  many: PrimaryPicklistBin
//  many: PicklistItem
//  many: PosTerminalLog
//  many: ProductOrderItem
//  many: EngagementProductOrderItem
//  many: ProductPromoUse
//  many: ReturnItem
//  many: ReplacementReturnItemResponse
//  many: PrimaryShipment
//  many: ShipmentReceipt
//  many: Subscription
//  many: SurveyResponse
//  many: TrackingCodeOrder
//  many: TrackingCodeOrderReturn
//  many: WorkOrderItemFulfillment
func (OrderHeader) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("item_issuances", ItemIssuance.Type),
                edge.To("order_contact_meches", OrderContactMech.Type),
                edge.To("order_items", OrderItem.Type),
                edge.To("order_item_price_infos", OrderItemPriceInfo.Type),
                edge.To("order_item_ship_groups", OrderItemShipGroup.Type),
                edge.To("order_item_ship_group_assocs", OrderItemShipGroupAssoc.Type),
                edge.To("order_roles", OrderRole.Type),
                edge.To("order_statuses", OrderStatus.Type),
                edge.To("primary_shipments", Shipment.Type),
    }
}


type OrderContactMech struct {
    ent.Schema
}

// Fields of the OrderContactMech.
func (OrderContactMech) Fields() []ent.Field {
    return []ent.Field{
        field.Int("contact_mech_purpose_type_id"),
        field.Int("contact_mech_id"),
    }
}

func (OrderContactMech) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
func (OrderContactMech) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("order_id", "contact_mech_purpose_type_id", "contact_mech_id").
            Unique(),
    }
}

*/



// Edges of the OrderContactMech.
//  one: OrderHeader
//  one: ContactMech
//  one: ContactMechPurposeType
func (OrderContactMech) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("order_header", OrderHeader.Type).Ref("order_contact_meches").
                    // Bind the "orderId" field to this edge.
                    // Field("order_id").
                    Unique(),
    }
}
