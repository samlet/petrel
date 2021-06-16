package schema

import (
	"entgo.io/ent"
	// "entgo.io/ent/schema/index"
	"entgo.io/contrib/entproto"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/samlet/petrel/alfin/schemamixins"
	"time"
)

type RoleType struct {
	ent.Schema
}

// Fields of the RoleType.
// Unique-Indexes: roleTypeId
func (RoleType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (RoleType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (RoleType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the RoleType.
//  one: ParentRoleType
//  many: AcctgTrans
//  many: AcctgTransEntry
//  many: FromAgreement
//  many: ToAgreement
//  many: AgreementRole
//  many: BillingAccountRole
//  many: ToCommunicationEvent
//  many: FromCommunicationEvent
//  many: CommunicationEventRole
//  many: ContentApproval
//  many: ContentPurposeOperation
//  many: CustRequestParty
//  many: FacilityGroupRole
//  many: FacilityParty
//  many: FinAccountRole
//  many: FixedAsset
//  many: GlAccountOrganization
//  many: GlAccountRole
//  many: Invoice
//  many: InvoiceRole
//  many: MarketingCampaignRole
//  many: OrderItemRole
//  many: OrderRole
//  many: PartyContactMech
//  many: PartyFixedAssetAssignment
//  many: PartyGlAccount
//  many: PartyInvitationRoleAssoc
//  many: PartyNeed
//  many: FromPartyRelationship
//  many: ToPartyRelationship
//  many: ValidFromPartyRelationshipType
//  many: ValidToPartyRelationshipType
//  many: PartyRole
//  many: ToPayment
//  many: PicklistRole
//  many: ProdCatalogRole
//  many: ProductCategoryRole
//  many: UseProductContent
//  many: ProductRole
//  many: ProductStoreGroupRole
//  many: ProductStoreRole
//  many: UseProductSubscriptionResource
//  many: QuoteRole
//  many: ChildRoleType
//  many: RoleTypeAttr
//  many: SalesOpportunityRole
//  many: SegmentGroupRole
//  many: ShipmentCostEstimate
//  many: Subscription
//  many: OriginatedFromSubscription
//  many: TimesheetRole
//  many: ValidContactMechRole
//  many: WebSiteRole
//  many: WorkEffortPartyAssignment
func (RoleType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", RoleType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("order_roles", OrderRole.Type).
			Annotations(entproto.Field(27)),
		// m2o
		edge.To("child_role_types", RoleType.Type).
			Annotations(entproto.Field(48)),
	}
}

type ProductReview struct {
	ent.Schema
}

// Fields of the ProductReview.
// Unique-Indexes: productReviewId
func (ProductReview) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_login_id").Optional().Annotations(entproto.Field(2)),
		field.Enum("posted_anonymous").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(3)),
		field.Time("posted_date_time").
			Default(time.Now).Optional().Annotations(entproto.Field(4)),
		field.Float("product_rating").Optional().Annotations(entproto.Field(5)),
		field.String("product_review").Optional().Annotations(entproto.Field(6)),
	}
}

//* entproto annotations ??
func (ProductReview) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductReview) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductReview.
//  one: ProductStore
//  one: Product
//  one: UserLogin
//  one: StatusItem
func (ProductReview) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("product_store", ProductStore.Type).Ref("product_reviews").
			// Bind the "productStoreId" field to this edge.
			// Field("product_store_id").
			Unique().Annotations(entproto.Field(7)),
		// o2m
		edge.From("product", Product.Type).Ref("product_reviews").
			// Bind the "productId" field to this edge.
			// Field("product_id").
			Unique().Annotations(entproto.Field(8)),
		// o2m
		edge.From("status_item", StatusItem.Type).Ref("product_reviews").
			// Bind the "statusId" field to this edge.
			// Field("status_id").
			Unique().Annotations(entproto.Field(10)),
	}
}

type ProductFeatureType struct {
	ent.Schema
}

// Fields of the ProductFeatureType.
// Unique-Indexes: productFeatureTypeId
func (ProductFeatureType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (ProductFeatureType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductFeatureType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductFeatureType.
//  one: ParentProductFeatureType
//  many: ProductFeature
//  many: ChildProductFeatureType
func (ProductFeatureType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProductFeatureType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("product_features", ProductFeature.Type).
			Annotations(entproto.Field(5)),
		// m2o
		edge.To("child_product_feature_types", ProductFeatureType.Type).
			Annotations(entproto.Field(6)),
	}
}

type VarianceReason struct {
	ent.Schema
}

// Fields of the VarianceReason.
// Unique-Indexes: varianceReasonId
func (VarianceReason) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (VarianceReason) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (VarianceReason) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the VarianceReason.
//  many: InventoryItemVariance
//  many: VarianceReasonGlAccount
func (VarianceReason) Edges() []ent.Edge {
	return []ent.Edge{}
}

type StatusValidChange struct {
	ent.Schema
}

// Fields of the StatusValidChange.
// Unique-Indexes: statusId, statusIdTo
func (StatusValidChange) Fields() []ent.Field {
	return []ent.Field{
		field.String("condition_expression").Optional().Annotations(entproto.Field(2)),
		field.String("transition_name").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (StatusValidChange) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
func (StatusValidChange) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("status_id", "status_id_to").
            Unique(),
    }
}

*/

func (StatusValidChange) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the StatusValidChange.
//  one: MainStatusItem
//  one: ToStatusItem
//  many: OldPicklistStatusHistory
func (StatusValidChange) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("main_status_item", StatusItem.Type).Ref("main_status_valid_changes").
			// Bind the "statusId" field to this edge.
			// Field("status_id").
			Unique().Annotations(entproto.Field(4)),
		// o2m
		edge.From("to_status_item", StatusItem.Type).Ref("to_status_valid_changes").
			// Bind the "statusIdTo" field to this edge.
			// Field("status_id_to").
			Unique().Annotations(entproto.Field(5)),
	}
}

type OrderRole struct {
	ent.Schema
}

// Fields of the OrderRole.
// Unique-Indexes: orderId, partyId, roleTypeId
func (OrderRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int("party_id").Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (OrderRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
func (OrderRole) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("order_id", "party_id", "role_type_id").
            Unique(),
    }
}

*/

func (OrderRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the OrderRole.
//  one: OrderHeader
//  one: Party
//  one: PartyRole
//  one-nofk: RoleType
//  many: OrderItem
func (OrderRole) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("order_header", OrderHeader.Type).Ref("order_roles").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(3)),
		// o2m
		edge.From("role_type", RoleType.Type).Ref("order_roles").
			// Bind the "roleTypeId" field to this edge.
			// Field("role_type_id").
			Unique().Annotations(entproto.Field(6)),
		// m2o
		edge.To("order_items", OrderItem.Type).
			Annotations(entproto.Field(7)),
	}
}

type GoodIdentificationType struct {
	ent.Schema
}

// Fields of the GoodIdentificationType.
// Unique-Indexes: goodIdentificationTypeId
func (GoodIdentificationType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (GoodIdentificationType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (GoodIdentificationType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the GoodIdentificationType.
//  one: ParentGoodIdentificationType
//  many: GoodIdentification
//  many: ChildGoodIdentificationType
func (GoodIdentificationType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", GoodIdentificationType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("child_good_identification_types", GoodIdentificationType.Type).
			Annotations(entproto.Field(6)),
	}
}

type ProductType struct {
	ent.Schema
}

// Fields of the ProductType.
// Unique-Indexes: productTypeId
func (ProductType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("is_physical").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.Enum("is_digital").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(3)),
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(4)),
		field.String("description").Optional().Annotations(entproto.Field(5)),
	}
}

//* entproto annotations ??
func (ProductType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductType.
//  one: ParentProductType
//  many: Product
//  many: ChildProductType
//  many: ProductTypeAttr
func (ProductType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProductType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(6)),
		// m2o
		edge.To("products", Product.Type).
			Annotations(entproto.Field(7)),
		// m2o
		edge.To("child_product_types", ProductType.Type).
			Annotations(entproto.Field(8)),
	}
}

type ContentAssocType struct {
	ent.Schema
}

// Fields of the ContentAssocType.
// Unique-Indexes: contentAssocTypeId
func (ContentAssocType) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (ContentAssocType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ContentAssocType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ContentAssocType.
//  many: ContentAssoc
func (ContentAssocType) Edges() []ent.Edge {
	return []ent.Edge{}
}

type CustomMethod struct {
	ent.Schema
}

// Fields of the CustomMethod.
// Unique-Indexes: customMethodId
func (CustomMethod) Fields() []ent.Field {
	return []ent.Field{
		field.String("custom_method_name").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (CustomMethod) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (CustomMethod) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the CustomMethod.
//  one: CustomMethodType
//  many: Content
//  many: CostComponentCalc
//  many: FixedAssetDepMethod
//  many: InvoicePartyAcctgPreference
//  many: QuotePartyAcctgPreference
//  many: OrderPartyAcctgPreference
//  many: ProductAssoc
//  many: ProductPrice
//  many: ProductPromoAction
//  many: ProductPromoCond
//  many: ProductStorePaymentSetting
//  many: ProductStoreShipmentMeth
//  many: ProductStoreTelecomSetting
//  many: uomCustomMethodUomConversion
//  many: uomCustomMethodUomConversionDated
//  many: WorkEffort
func (CustomMethod) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("custom_method_type", CustomMethodType.Type).Ref("custom_methods").
			// Bind the "customMethodTypeId" field to this edge.
			// Field("custom_method_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("product_assocs", ProductAssoc.Type).
			Annotations(entproto.Field(11)),
		// m2o
		edge.To("product_prices", ProductPrice.Type).
			Annotations(entproto.Field(12)),
	}
}

type StatusType struct {
	ent.Schema
}

// Fields of the StatusType.
// Unique-Indexes: statusTypeId
func (StatusType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (StatusType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (StatusType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the StatusType.
//  one: ParentStatusType
//  many: StatusItem
//  many: ChildStatusType
func (StatusType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", StatusType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("status_items", StatusItem.Type).
			Annotations(entproto.Field(5)),
		// m2o
		edge.To("child_status_types", StatusType.Type).
			Annotations(entproto.Field(6)),
	}
}

type ProductPrice struct {
	ent.Schema
}

// Fields of the ProductPrice.
// Unique-Indexes: productId, productPriceTypeId, productPricePurposeId, currencyUomId, productStoreGroupId, fromDate
func (ProductPrice) Fields() []ent.Field {
	return []ent.Field{
		field.Int("currency_uom_id").Annotations(entproto.Field(2)),
		field.Time("from_date").
			Default(time.Now).Annotations(entproto.Field(3)),
		field.Time("thru_date").
			Default(time.Now).Optional().Annotations(entproto.Field(4)),
		field.Float("price").Optional().Annotations(entproto.Field(5)),
		field.Int("term_uom_id").Optional().Annotations(entproto.Field(6)),
		field.Float("price_without_tax").Optional().Annotations(entproto.Field(7)),
		field.Float("price_with_tax").Optional().Annotations(entproto.Field(8)),
		field.Float("tax_amount").Optional().Annotations(entproto.Field(9)),
		field.Float("tax_percentage").Optional().Annotations(entproto.Field(10)),
		field.Int("tax_auth_party_id").Optional().Annotations(entproto.Field(11)),
		field.Int("tax_auth_geo_id").Optional().Annotations(entproto.Field(12)),
		field.Enum("tax_in_price").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(13)),
		field.Time("created_date").
			Default(time.Now).Optional().Annotations(entproto.Field(14)),
		field.String("created_by_user_login").Optional().Annotations(entproto.Field(15)),
		field.Time("last_modified_date").
			Default(time.Now).Optional().Annotations(entproto.Field(16)),
		field.String("last_modified_by_user_login").Optional().Annotations(entproto.Field(17)),
	}
}

//* entproto annotations ??
func (ProductPrice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
func (ProductPrice) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("product_id", "product_price_type_id", "product_price_purpose_id", "currency_uom_id", "product_store_group_id", "from_date").
            Unique(),
    }
}

*/

func (ProductPrice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductPrice.
//  one: Product
//  one: ProductPriceType
//  one: ProductPricePurpose
//  one: CurrencyUom
//  one: TermUom
//  one: ProductStoreGroup
//  one: CustomMethod
//  one: TaxAuthorityParty
//  one: TaxAuthorityGeo
//  one: CreatedByUserLogin
//  one: LastModifiedByUserLogin
//  many: ProductPriceChange
func (ProductPrice) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("product", Product.Type).Ref("product_prices").
			// Bind the "productId" field to this edge.
			// Field("product_id").
			Unique().Annotations(entproto.Field(18)),
		// o2m
		edge.From("product_price_type", ProductPriceType.Type).Ref("product_prices").
			// Bind the "productPriceTypeId" field to this edge.
			// Field("product_price_type_id").
			Unique().Annotations(entproto.Field(19)),
		// o2m
		edge.From("product_price_purpose", ProductPricePurpose.Type).Ref("product_prices").
			// Bind the "productPricePurposeId" field to this edge.
			// Field("product_price_purpose_id").
			Unique().Annotations(entproto.Field(20)),
		// o2m
		edge.From("product_store_group", ProductStoreGroup.Type).Ref("product_prices").
			// Bind the "productStoreGroupId" field to this edge.
			// Field("product_store_group_id").
			Unique().Annotations(entproto.Field(23)),
		// o2m
		edge.From("custom_method", CustomMethod.Type).Ref("product_prices").
			// Bind the "customPriceCalcService" field to this edge.
			// Field("custom_price_calc_service").
			Unique().Annotations(entproto.Field(24)),
	}
}

type ProdCatalogCategoryType struct {
	ent.Schema
}

// Fields of the ProdCatalogCategoryType.
// Unique-Indexes: prodCatalogCategoryTypeId
func (ProdCatalogCategoryType) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (ProdCatalogCategoryType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProdCatalogCategoryType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProdCatalogCategoryType.
//  one: ParentProdCatalogCategoryType
//  many: ProdCatalogCategory
//  many: ChildProdCatalogCategoryType
func (ProdCatalogCategoryType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProdCatalogCategoryType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(3)),
		// m2o
		edge.To("child_prod_catalog_category_types", ProdCatalogCategoryType.Type).
			Annotations(entproto.Field(5)),
	}
}

type ProductMeterType struct {
	ent.Schema
}

// Fields of the ProductMeterType.
// Unique-Indexes: productMeterTypeId
func (ProductMeterType) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
		field.Int("default_uom_id").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (ProductMeterType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductMeterType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductMeterType.
//  one: DefaultUom
//  many: IntervalFixedAssetMaint
//  many: FixedAssetMeter
//  many: IntervalProductMaint
//  many: ProductMeter
func (ProductMeterType) Edges() []ent.Edge {
	return []ent.Edge{}
}

type ShipmentGatewayUsps struct {
	ent.Schema
}

// Fields of the ShipmentGatewayUsps.
// Unique-Indexes: shipmentGatewayConfigId
func (ShipmentGatewayUsps) Fields() []ent.Field {
	return []ent.Field{
		field.String("connect_url").Optional().Annotations(entproto.Field(2)),
		field.String("connect_url_labels").Optional().Annotations(entproto.Field(3)),
		field.Int("connect_timeout").Optional().Annotations(entproto.Field(4)),
		field.String("access_user_id").Optional().Annotations(entproto.Field(5)),
		field.String("access_password").Optional().Annotations(entproto.Field(6)),
		field.Int("max_estimate_weight").Optional().Annotations(entproto.Field(7)),
		field.String("test").Optional().Annotations(entproto.Field(8)),
	}
}

//* entproto annotations ??
func (ShipmentGatewayUsps) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ShipmentGatewayUsps) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ShipmentGatewayUsps.
//  one: ShipmentGatewayConfig
func (ShipmentGatewayUsps) Edges() []ent.Edge {
	return []ent.Edge{
		// o2o (nofk)
		edge.From("shipment_gateway_config", ShipmentGatewayConfig.Type).
			Ref("shipment_gateway_usps").
			// Bind the "shipmentGatewayConfigId" field to this edge.
			// Field("shipment_gateway_config_id").
			Unique().Annotations(entproto.Field(9)),
	}
}

type OrderContactMech struct {
	ent.Schema
}

// Fields of the OrderContactMech.
// Unique-Indexes: orderId, contactMechPurposeTypeId, contactMechId
func (OrderContactMech) Fields() []ent.Field {
	return []ent.Field{
		field.Int("contact_mech_purpose_type_id").Annotations(entproto.Field(2)),
		field.Int("contact_mech_id").Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (OrderContactMech) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
func (OrderContactMech) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("order_id", "contact_mech_purpose_type_id", "contact_mech_id").
            Unique(),
    }
}

*/

func (OrderContactMech) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the OrderContactMech.
//  one: OrderHeader
//  one: ContactMech
//  one: ContactMechPurposeType
func (OrderContactMech) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("order_header", OrderHeader.Type).Ref("order_contact_meches").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(4)),
	}
}

type ShipmentGatewayConfig struct {
	ent.Schema
}

// Fields of the ShipmentGatewayConfig.
// Unique-Indexes: shipmentGatewayConfigId
func (ShipmentGatewayConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (ShipmentGatewayConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ShipmentGatewayConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ShipmentGatewayConfig.
//  one: ShipmentGatewayConfigType
//  many: ProductStoreShipmentMeth
//  one-nofk: ShipmentGatewayDhl
//  one-nofk: ShipmentGatewayFedex
//  one-nofk: ShipmentGatewayUps
//  one-nofk: ShipmentGatewayUsps
func (ShipmentGatewayConfig) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("shipment_gateway_config_type", ShipmentGatewayConfigType.Type).Ref("shipment_gateway_configs").
			// Bind the "shipmentGatewayConfTypeId" field to this edge.
			// Field("shipment_gateway_conf_type_id").
			Unique().Annotations(entproto.Field(3)),
		// o2o
		edge.To("shipment_gateway_dhl", ShipmentGatewayDhl.Type).
			// Bind the "shipmentGatewayConfigId" field to this edge.
			// Field("shipment_gateway_config_id").
			Unique().Annotations(entproto.Field(5)),
		// o2o
		edge.To("shipment_gateway_fedex", ShipmentGatewayFedex.Type).
			// Bind the "shipmentGatewayConfigId" field to this edge.
			// Field("shipment_gateway_config_id").
			Unique().Annotations(entproto.Field(6)),
		// o2o
		edge.To("shipment_gateway_ups", ShipmentGatewayUps.Type).
			// Bind the "shipmentGatewayConfigId" field to this edge.
			// Field("shipment_gateway_config_id").
			Unique().Annotations(entproto.Field(7)),
		// o2o
		edge.To("shipment_gateway_usps", ShipmentGatewayUsps.Type).
			// Bind the "shipmentGatewayConfigId" field to this edge.
			// Field("shipment_gateway_config_id").
			Unique().Annotations(entproto.Field(8)),
	}
}

type ProductFeatureIactnType struct {
	ent.Schema
}

// Fields of the ProductFeatureIactnType.
// Unique-Indexes: productFeatureIactnTypeId
func (ProductFeatureIactnType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (ProductFeatureIactnType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductFeatureIactnType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductFeatureIactnType.
//  one: ParentProductFeatureIactnType
//  many: ProductFeatureIactn
//  many: ChildProductFeatureIactnType
func (ProductFeatureIactnType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProductFeatureIactnType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("child_product_feature_iactn_types", ProductFeatureIactnType.Type).
			Annotations(entproto.Field(6)),
	}
}

type ProductPriceType struct {
	ent.Schema
}

// Fields of the ProductPriceType.
// Unique-Indexes: productPriceTypeId
func (ProductPriceType) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (ProductPriceType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductPriceType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductPriceType.
//  many: ProductFeaturePrice
//  many: ProductPrice
func (ProductPriceType) Edges() []ent.Edge {
	return []ent.Edge{
		// m2o
		edge.To("product_prices", ProductPrice.Type).
			Annotations(entproto.Field(4)),
	}
}

type FacilityGroupType struct {
	ent.Schema
}

// Fields of the FacilityGroupType.
// Unique-Indexes: facilityGroupTypeId
func (FacilityGroupType) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (FacilityGroupType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (FacilityGroupType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the FacilityGroupType.
//  many: FacilityGroup
func (FacilityGroupType) Edges() []ent.Edge {
	return []ent.Edge{
		// m2o
		edge.To("facility_groups", FacilityGroup.Type).
			Annotations(entproto.Field(3)),
	}
}

type OrderItemShipGroupAssoc struct {
	ent.Schema
}

// Fields of the OrderItemShipGroupAssoc.
// Unique-Indexes: orderId, orderItemSeqId, shipGroupSeqId
func (OrderItemShipGroupAssoc) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_item_seq_id").Annotations(entproto.Field(2)),
		field.Int("ship_group_seq_id").Annotations(entproto.Field(3)),
		field.Float("quantity").Optional().Annotations(entproto.Field(4)),
		field.Float("cancel_quantity").Optional().Annotations(entproto.Field(5)),
	}
}

//* entproto annotations ??
func (OrderItemShipGroupAssoc) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
func (OrderItemShipGroupAssoc) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("order_id", "order_item_seq_id", "ship_group_seq_id").
            Unique(),
    }
}

*/

func (OrderItemShipGroupAssoc) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

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
		// o2m
		edge.From("order_header", OrderHeader.Type).Ref("order_item_ship_group_assocs").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(6)),
		// o2m
		edge.From("order_item", OrderItem.Type).Ref("order_item_ship_group_assocs").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(7)),
		// o2m
		edge.From("order_item_ship_group", OrderItemShipGroup.Type).Ref("order_item_ship_group_assocs").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(8)),
		// m2o
		edge.To("order_adjustments", OrderAdjustment.Type).
			Annotations(entproto.Field(9)),
		// m2o
		edge.To("order_item_ship_grp_inv_res", OrderItemShipGrpInvRes.Type).
			Annotations(entproto.Field(12)),
	}
}

type ProdConfItemContentType struct {
	ent.Schema
}

// Fields of the ProdConfItemContentType.
// Unique-Indexes: confItemContentTypeId
func (ProdConfItemContentType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (ProdConfItemContentType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProdConfItemContentType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProdConfItemContentType.
//  one: ParentProdConfItemContentType
//  many: ProdConfItemContent
//  many: ChildProdConfItemContentType
func (ProdConfItemContentType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProdConfItemContentType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("child_prod_conf_item_content_types", ProdConfItemContentType.Type).
			Annotations(entproto.Field(6)),
	}
}

type CustomMethodType struct {
	ent.Schema
}

// Fields of the CustomMethodType.
// Unique-Indexes: customMethodTypeId
func (CustomMethodType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (CustomMethodType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (CustomMethodType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the CustomMethodType.
//  one: ParentCustomMethodType
//  many: CustomMethod
//  many: ChildCustomMethodType
func (CustomMethodType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", CustomMethodType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("custom_methods", CustomMethod.Type).
			Annotations(entproto.Field(5)),
		// m2o
		edge.To("child_custom_method_types", CustomMethodType.Type).
			Annotations(entproto.Field(6)),
	}
}

type ProductPriceActionType struct {
	ent.Schema
}

// Fields of the ProductPriceActionType.
// Unique-Indexes: productPriceActionTypeId
func (ProductPriceActionType) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (ProductPriceActionType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductPriceActionType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductPriceActionType.
//  many: ProductPriceAction
func (ProductPriceActionType) Edges() []ent.Edge {
	return []ent.Edge{}
}

type ShipmentType struct {
	ent.Schema
}

// Fields of the ShipmentType.
// Unique-Indexes: shipmentTypeId
func (ShipmentType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (ShipmentType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ShipmentType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ShipmentType.
//  one: ParentShipmentType
//  many: Shipment
//  many: ChildShipmentType
//  many: ShipmentTypeAttr
func (ShipmentType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ShipmentType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("child_shipment_types", ShipmentType.Type).
			Annotations(entproto.Field(6)),
	}
}

type StatusItem struct {
	ent.Schema
}

// Fields of the StatusItem.
// Unique-Indexes: statusId
func (StatusItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("status_code").Optional().Annotations(entproto.Field(2)),
		field.Int("sequence_id").Optional().Annotations(entproto.Field(3)),
		field.String("description").Optional().Annotations(entproto.Field(4)),
	}
}

//* entproto annotations ??
func (StatusItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (StatusItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the StatusItem.
//  one: StatusType
//  many: AcctgTrans
//  many: AcctgTransEntry
//  many: Status IdAllocationPlanHeader
//  many: Plan Item Status IdAllocationPlanItem
//  many: BudgetStatus
//  many: CommunicationEvent
//  many: CommunicationEventRole
//  many: ContactListCommStatus
//  many: ContactListParty
//  many: Content
//  many: ApprovalContentApproval
//  many: ContentPurposeOperation
//  many: CustRequest
//  many: CustRequestItem
//  many: CustRequestStatus
//  many: DataResource
//  many: EmplLeave
//  many: EmplPosition
//  many: EmploymentApp
//  many: Example
//  many: ExampleStatus
//  many: FinAccountStatus
//  many: FinAccountTrans
//  many: FixedAssetMaint
//  many: GlReconciliation
//  many: InventoryItem
//  many: InventoryItemStatus
//  many: InventoryTransfer
//  many: Invoice
//  many: InvoiceStatus
//  many: JobSandbox
//  many: MarketingCampaign
//  many: OldPicklistStatusHistory
//  many: ToOldPicklistStatusHistory
//  many: OrderDeliverySchedule
//  many: OrderHeader
//  many: SyncOrderHeader
//  many: OrderItem
//  many: SyncOrderItem
//  many: OrderPaymentPreference
//  many: OrderStatus
//  many: Party
//  many: PartyFixedAssetAssignment
//  many: PartyInvitation
//  many: PartyQual
//  many: VerificationPartyQual
//  many: PartyRelationship
//  many: PartyStatus
//  many: Payment
//  many: Picklist
//  many: PicklistItem
//  many: PicklistStatus
//  many: ToPicklistStatus
//  many: PosTerminalLog
//  many: ProductGroupOrder
//  many: ProductKeyword
//  many: ProductReview
//  many: HeaderApprovedProductStore
//  many: ItemApprovedProductStore
//  many: DigitalItemApprovedProductStore
//  many: HeaderDeclinedProductStore
//  many: ItemDeclinedProductStore
//  many: HeaderCancelProductStore
//  many: ItemCancelProductStore
//  many: Quote
//  many: Requirement
//  many: RequirementStatus
//  many: ReturnHeader
//  many: ReturnItem
//  many: InventoryReturnItem
//  many: ReturnStatus
//  many: Shipment
//  many: CarrierServiceShipmentRouteSegment
//  many: ShipmentStatus
//  many: MainStatusValidChange
//  many: ToStatusValidChange
//  many: SurveyResponse
//  many: TestingStatus
//  many: Timesheet
//  many: UnemploymentClaim
//  many: CurrentWorkEffort
//  many: WorkEffortFixedAssetAssign
//  many: AvailabilityWorkEffortFixedAssetAssign
//  many: WorkEffortGoodStandard
//  many: WorkEffortInventoryAssign
//  many: AssignmentWorkEffortPartyAssignment
//  many: AvailabilityWorkEffortPartyAssignment
//  many: WorkEffortReview
//  many: WorkEffortStatus
//  many: Workload
//  many: WorkloadStatus
func (StatusItem) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("status_type", StatusType.Type).Ref("status_items").
			// Bind the "statusTypeId" field to this edge.
			// Field("status_type_id").
			Unique().Annotations(entproto.Field(5)),
		// m2o
		edge.To("order_headers", OrderHeader.Type).
			Annotations(entproto.Field(41)),
		// m2o
		edge.To("sync_order_headers", OrderHeader.Type).
			Annotations(entproto.Field(42)),
		// m2o
		edge.To("order_items", OrderItem.Type).
			Annotations(entproto.Field(43)),
		// m2o
		edge.To("sync_order_items", OrderItem.Type).
			Annotations(entproto.Field(44)),
		// m2o
		edge.To("order_payment_preferences", OrderPaymentPreference.Type).
			Annotations(entproto.Field(45)),
		// m2o
		edge.To("order_statuses", OrderStatus.Type).
			Annotations(entproto.Field(46)),
		// m2o
		edge.To("product_reviews", ProductReview.Type).
			Annotations(entproto.Field(62)),
		// m2o
		edge.To("header_approved_product_stores", ProductStore.Type).
			Annotations(entproto.Field(63)),
		// m2o
		edge.To("item_approved_product_stores", ProductStore.Type).
			Annotations(entproto.Field(64)),
		// m2o
		edge.To("digital_item_approved_product_stores", ProductStore.Type).
			Annotations(entproto.Field(65)),
		// m2o
		edge.To("header_declined_product_stores", ProductStore.Type).
			Annotations(entproto.Field(66)),
		// m2o
		edge.To("item_declined_product_stores", ProductStore.Type).
			Annotations(entproto.Field(67)),
		// m2o
		edge.To("header_cancel_product_stores", ProductStore.Type).
			Annotations(entproto.Field(68)),
		// m2o
		edge.To("item_cancel_product_stores", ProductStore.Type).
			Annotations(entproto.Field(69)),
		// m2o
		edge.To("main_status_valid_changes", StatusValidChange.Type).
			Annotations(entproto.Field(80)),
		// m2o
		edge.To("to_status_valid_changes", StatusValidChange.Type).
			Annotations(entproto.Field(81)),
	}
}

type ProductFeatureApplType struct {
	ent.Schema
}

// Fields of the ProductFeatureApplType.
// Unique-Indexes: productFeatureApplTypeId
func (ProductFeatureApplType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (ProductFeatureApplType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductFeatureApplType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductFeatureApplType.
//  one: ParentProductFeatureApplType
//  many: ProductFeatureAppl
//  many: ChildProductFeatureApplType
func (ProductFeatureApplType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProductFeatureApplType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("child_product_feature_appl_types", ProductFeatureApplType.Type).
			Annotations(entproto.Field(6)),
	}
}

type InventoryItemType struct {
	ent.Schema
}

// Fields of the InventoryItemType.
// Unique-Indexes: inventoryItemTypeId
func (InventoryItemType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (InventoryItemType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (InventoryItemType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the InventoryItemType.
//  one: ParentInventoryItemType
//  many: DefaultFacility
//  many: InventoryItem
//  many: ChildInventoryItemType
//  many: InventoryItemTypeAttr
//  many: Product
func (InventoryItemType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", InventoryItemType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("child_inventory_item_types", InventoryItemType.Type).
			Annotations(entproto.Field(7)),
		// m2o
		edge.To("products", Product.Type).
			Annotations(entproto.Field(9)),
	}
}

type ShipmentGatewayUps struct {
	ent.Schema
}

// Fields of the ShipmentGatewayUps.
// Unique-Indexes: shipmentGatewayConfigId
func (ShipmentGatewayUps) Fields() []ent.Field {
	return []ent.Field{
		field.String("connect_url").Optional().Annotations(entproto.Field(2)),
		field.Int("connect_timeout").Optional().Annotations(entproto.Field(3)),
		field.String("shipper_number").Optional().Annotations(entproto.Field(4)),
		field.String("bill_shipper_account_number").Optional().Annotations(entproto.Field(5)),
		field.String("access_license_number").Optional().Annotations(entproto.Field(6)),
		field.String("access_user_id").Optional().Annotations(entproto.Field(7)),
		field.String("access_password").Optional().Annotations(entproto.Field(8)),
		field.String("save_cert_info").Optional().Annotations(entproto.Field(9)),
		field.String("save_cert_path").Optional().Annotations(entproto.Field(10)),
		field.String("shipper_pickup_type").Optional().Annotations(entproto.Field(11)),
		field.String("customer_classification").Optional().Annotations(entproto.Field(12)),
		field.Float("max_estimate_weight").Optional().Annotations(entproto.Field(13)),
		field.Float("min_estimate_weight").Optional().Annotations(entproto.Field(14)),
		field.String("cod_allow_cod").Optional().Annotations(entproto.Field(15)),
		field.Float("cod_surcharge_amount").Optional().Annotations(entproto.Field(16)),
		field.String("cod_surcharge_currency_uom_id").Optional().Annotations(entproto.Field(17)),
		field.String("cod_surcharge_apply_to_package").Optional().Annotations(entproto.Field(18)),
		field.String("cod_funds_code").Optional().Annotations(entproto.Field(19)),
		field.String("default_return_label_memo").Optional().Annotations(entproto.Field(20)),
		field.String("default_return_label_subject").Optional().Annotations(entproto.Field(21)),
	}
}

//* entproto annotations ??
func (ShipmentGatewayUps) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ShipmentGatewayUps) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ShipmentGatewayUps.
//  one: ShipmentGatewayConfig
func (ShipmentGatewayUps) Edges() []ent.Edge {
	return []ent.Edge{
		// o2o (nofk)
		edge.From("shipment_gateway_config", ShipmentGatewayConfig.Type).
			Ref("shipment_gateway_ups").
			// Bind the "shipmentGatewayConfigId" field to this edge.
			// Field("shipment_gateway_config_id").
			Unique().Annotations(entproto.Field(22)),
	}
}

type ProductStore struct {
	ent.Schema
}

// Fields of the ProductStore.
// Unique-Indexes: productStoreId
func (ProductStore) Fields() []ent.Field {
	return []ent.Field{
		field.String("store_name").Optional().Annotations(entproto.Field(2)),
		field.String("company_name").Optional().Annotations(entproto.Field(3)),
		field.String("title").Optional().Annotations(entproto.Field(4)),
		field.String("subtitle").Optional().Annotations(entproto.Field(5)),
		field.Int("pay_to_party_id").Optional().Annotations(entproto.Field(6)),
		field.Int("days_to_cancel_non_pay").Optional().Annotations(entproto.Field(7)),
		field.Enum("manual_auth_is_capture").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(8)),
		field.Enum("prorate_shipping").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(9)),
		field.Enum("prorate_taxes").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(10)),
		field.Enum("view_cart_on_add").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(11)),
		field.Enum("auto_save_cart").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(12)),
		field.Enum("auto_approve_reviews").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(13)),
		field.Enum("is_demo_store").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(14)),
		field.Enum("is_immediately_fulfilled").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(15)),
		field.Int("inventory_facility_id").Optional().Annotations(entproto.Field(16)),
		field.Enum("one_inventory_facility").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(17)),
		field.Enum("check_inventory").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(18)),
		field.Enum("reserve_inventory").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(19)),
		field.Enum("require_inventory").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(20)),
		field.Enum("balance_res_on_order_creation").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(21)),
		field.String("order_number_prefix").MaxLen(32).Optional().Annotations(entproto.Field(22)),
		field.String("default_locale_string").MaxLen(10).Optional().Annotations(entproto.Field(23)),
		field.Int("default_currency_uom_id").Optional().Annotations(entproto.Field(24)),
		field.String("default_time_zone_string").MaxLen(32).Optional().Annotations(entproto.Field(25)),
		field.Enum("allow_password").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(26)),
		field.String("default_password").Optional().Annotations(entproto.Field(27)),
		field.Enum("explode_order_items").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(28)),
		field.Enum("check_gc_balance").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(29)),
		field.Enum("retry_failed_auths").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(30)),
		field.String("auth_declined_message").Optional().Annotations(entproto.Field(31)),
		field.String("auth_fraud_message").Optional().Annotations(entproto.Field(32)),
		field.String("auth_error_message").Optional().Annotations(entproto.Field(33)),
		field.Int("visual_theme_id").Optional().Annotations(entproto.Field(34)),
		field.Enum("use_primary_email_username").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(35)),
		field.Enum("require_customer_role").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(36)),
		field.Enum("auto_invoice_digital_items").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(37)),
		field.Enum("req_ship_addr_for_dig_items").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(38)),
		field.Enum("show_checkout_gift_options").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(39)),
		field.Enum("select_payment_type_per_item").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(40)),
		field.Enum("show_prices_with_vat_tax").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(41)),
		field.Enum("show_tax_is_exempt").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(42)),
		field.Int("vat_tax_auth_geo_id").Optional().Annotations(entproto.Field(43)),
		field.Int("vat_tax_auth_party_id").Optional().Annotations(entproto.Field(44)),
		field.Enum("enable_auto_suggestion_list").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(45)),
		field.Enum("enable_dig_prod_upload").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(46)),
		field.Enum("prod_search_exclude_variants").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(47)),
		field.Int("dig_prod_upload_category_id").Optional().Annotations(entproto.Field(48)),
		field.Enum("auto_order_cc_try_exp").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(49)),
		field.Enum("auto_order_cc_try_other_cards").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(50)),
		field.Enum("auto_order_cc_try_later_nsf").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(51)),
		field.Int("auto_order_cc_try_later_max").Optional().Annotations(entproto.Field(52)),
		field.Int("store_credit_valid_days").Optional().Annotations(entproto.Field(53)),
		field.Enum("auto_approve_invoice").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(54)),
		field.Enum("auto_approve_order").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(55)),
		field.Enum("ship_if_capture_fails").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(56)),
		field.Enum("set_owner_upon_issuance").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(57)),
		field.Enum("req_return_inventory_receive").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(58)),
		field.Enum("add_to_cart_remove_incompat").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(59)),
		field.Enum("add_to_cart_replace_upsell").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(60)),
		field.Enum("split_pay_pref_per_shp_grp").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(61)),
		field.Enum("managed_by_lot").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(62)),
		field.Enum("show_out_of_stock_products").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(63)),
		field.Enum("order_decimal_quantity").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(64)),
		field.Enum("allow_comment").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(65)),
		field.Enum("allocate_inventory").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(66)),
	}
}

//* entproto annotations ??
func (ProductStore) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductStore) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductStore.
//  one: PrimaryProductStoreGroup
//  one: Facility
//  one: ReserveOrderEnumeration
//  one: RequirementMethodEnumeration
//  one: Party
//  one: Uom
//  one: DefaultSalesChannelEnumeration
//  one: HeaderApprovedStatusItem
//  one: ItemApprovedStatusItem
//  one: DigitalItemApprovedStatusItem
//  one: HeaderDeclinedStatusItem
//  one: ItemDeclinedStatusItem
//  one: HeaderCancelStatusItem
//  one: ItemCancelStatusItem
//  one: VatTaxAuthority
//  one: StoreCreditAccountEnumeration
//  many: CustRequest
//  one-nofk: EbayConfig
//  many: EbayShippingMethod
//  many: GitHubUser
//  many: InventoryItemTempRes
//  many: LinkedInUser
//  many: OAuth2GitHub
//  many: OAuth2LinkedIn
//  many: OrderHeader
//  many: PartyProfileDefault
//  many: ProductReview
//  many: ProductStoreCatalog
//  many: ProductStoreEmailSetting
//  many: ProductStoreFacility
//  many: ProductStoreFinActSetting
//  many: ProductStoreGroupMember
//  many: ProductStoreKeywordOvrd
//  many: ProductStorePaymentSetting
//  many: ProductStorePromoAppl
//  many: ProductStoreRole
//  many: ProductStoreSurveyAppl
//  many: ProductStoreTelecomSetting
//  many: ProductStoreVendorPayment
//  many: ProductStoreVendorShipment
//  many: Quote
//  many: SegmentGroup
//  many: ShoppingList
//  many: TaxAuthorityRateProduct
//  many: ThirdPartyLogin
//  many: WebSite
func (ProductStore) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("primary_product_store_group", ProductStoreGroup.Type).Ref("primary_product_stores").
			// Bind the "primaryStoreGroupId" field to this edge.
			// Field("primary_store_group_id").
			Unique().Annotations(entproto.Field(67)),
		// o2m
		edge.From("reserve_order_enumeration", Enumeration.Type).Ref("reserve_order_product_stores").
			// Bind the "reserveOrderEnumId" field to this edge.
			// Field("reserve_order_enum_id").
			Unique().Annotations(entproto.Field(69)),
		// o2m
		edge.From("requirement_method_enumeration", Enumeration.Type).Ref("requirement_method_product_stores").
			// Bind the "requirementMethodEnumId" field to this edge.
			// Field("requirement_method_enum_id").
			Unique().Annotations(entproto.Field(70)),
		// o2m
		edge.From("default_sales_channel_enumeration", Enumeration.Type).Ref("default_sales_channel_product_stores").
			// Bind the "defaultSalesChannelEnumId" field to this edge.
			// Field("default_sales_channel_enum_id").
			Unique().Annotations(entproto.Field(73)),
		// o2m
		edge.From("header_approved_status_item", StatusItem.Type).Ref("header_approved_product_stores").
			// Bind the "headerApprovedStatus" field to this edge.
			// Field("header_approved_status").
			Unique().Annotations(entproto.Field(74)),
		// o2m
		edge.From("item_approved_status_item", StatusItem.Type).Ref("item_approved_product_stores").
			// Bind the "itemApprovedStatus" field to this edge.
			// Field("item_approved_status").
			Unique().Annotations(entproto.Field(75)),
		// o2m
		edge.From("digital_item_approved_status_item", StatusItem.Type).Ref("digital_item_approved_product_stores").
			// Bind the "digitalItemApprovedStatus" field to this edge.
			// Field("digital_item_approved_status").
			Unique().Annotations(entproto.Field(76)),
		// o2m
		edge.From("header_declined_status_item", StatusItem.Type).Ref("header_declined_product_stores").
			// Bind the "headerDeclinedStatus" field to this edge.
			// Field("header_declined_status").
			Unique().Annotations(entproto.Field(77)),
		// o2m
		edge.From("item_declined_status_item", StatusItem.Type).Ref("item_declined_product_stores").
			// Bind the "itemDeclinedStatus" field to this edge.
			// Field("item_declined_status").
			Unique().Annotations(entproto.Field(78)),
		// o2m
		edge.From("header_cancel_status_item", StatusItem.Type).Ref("header_cancel_product_stores").
			// Bind the "headerCancelStatus" field to this edge.
			// Field("header_cancel_status").
			Unique().Annotations(entproto.Field(79)),
		// o2m
		edge.From("item_cancel_status_item", StatusItem.Type).Ref("item_cancel_product_stores").
			// Bind the "itemCancelStatus" field to this edge.
			// Field("item_cancel_status").
			Unique().Annotations(entproto.Field(80)),
		// o2m
		edge.From("store_credit_account_enumeration", Enumeration.Type).Ref("store_credit_account_product_stores").
			// Bind the "storeCreditAccountEnumId" field to this edge.
			// Field("store_credit_account_enum_id").
			Unique().Annotations(entproto.Field(82)),
		// m2o
		edge.To("order_headers", OrderHeader.Type).
			Annotations(entproto.Field(91)),
		// m2o
		edge.To("product_reviews", ProductReview.Type).
			Annotations(entproto.Field(93)),
	}
}

type SupplierPrefOrder struct {
	ent.Schema
}

// Fields of the SupplierPrefOrder.
// Unique-Indexes: supplierPrefOrderId
func (SupplierPrefOrder) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (SupplierPrefOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (SupplierPrefOrder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the SupplierPrefOrder.
//  many: SupplierProduct
func (SupplierPrefOrder) Edges() []ent.Edge {
	return []ent.Edge{}
}

type ProductStoreGroup struct {
	ent.Schema
}

// Fields of the ProductStoreGroup.
// Unique-Indexes: productStoreGroupId
func (ProductStoreGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int("product_store_group_type_id").Optional().Annotations(entproto.Field(2)),
		field.String("product_store_group_name").Optional().Annotations(entproto.Field(3)),
		field.String("description").Optional().Annotations(entproto.Field(4)),
	}
}

//* entproto annotations ??
func (ProductStoreGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductStoreGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductStoreGroup.
//  one: ProductStoreGroupType
//  one: PrimaryParentProductStoreGroup
//  many: ProductPrice
//  many: PrimaryProductStore
//  many: ProductStoreGroupMember
//  many: ProductStoreGroupRole
//  many: CurrentProductStoreGroupRollup
//  many: ParentProductStoreGroupRollup
//  many: VendorProduct
func (ProductStoreGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProductStoreGroup.Type).
			From("parent").
			// Bind the "primaryParentGroupId" field to this edge.
			// Field("primary_parent_group_id").
			Unique().Annotations(entproto.Field(6)),
		// m2o
		edge.To("product_prices", ProductPrice.Type).
			Annotations(entproto.Field(7)),
		// m2o
		edge.To("primary_product_stores", ProductStore.Type).
			Annotations(entproto.Field(8)),
	}
}

type ProductAssocType struct {
	ent.Schema
}

// Fields of the ProductAssocType.
// Unique-Indexes: productAssocTypeId
func (ProductAssocType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (ProductAssocType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductAssocType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductAssocType.
//  one: ParentProductAssocType
//  many: ProductAssoc
//  many: ChildProductAssocType
func (ProductAssocType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProductAssocType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("product_assocs", ProductAssoc.Type).
			Annotations(entproto.Field(5)),
		// m2o
		edge.To("child_product_assoc_types", ProductAssocType.Type).
			Annotations(entproto.Field(6)),
	}
}

type OrderItemShipGroup struct {
	ent.Schema
}

// Fields of the OrderItemShipGroup.
// Unique-Indexes: orderId, shipGroupSeqId
func (OrderItemShipGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int("ship_group_seq_id").Annotations(entproto.Field(2)),
		field.Int("shipment_method_type_id").Optional().Annotations(entproto.Field(3)),
		field.Int("supplier_party_id").Optional().Annotations(entproto.Field(4)),
		field.Int("supplier_agreement_id").Optional().Annotations(entproto.Field(5)),
		field.Int("vendor_party_id").Optional().Annotations(entproto.Field(6)),
		field.Int("carrier_party_id").Optional().Annotations(entproto.Field(7)),
		field.Int("carrier_role_type_id").Optional().Annotations(entproto.Field(8)),
		field.Int("facility_id").Optional().Annotations(entproto.Field(9)),
		field.Int("contact_mech_id").Optional().Annotations(entproto.Field(10)),
		field.Int("telecom_contact_mech_id").Optional().Annotations(entproto.Field(11)),
		field.String("tracking_number").Optional().Annotations(entproto.Field(12)),
		field.String("shipping_instructions").Optional().Annotations(entproto.Field(13)),
		field.Enum("may_split").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(14)),
		field.String("gift_message").Optional().Annotations(entproto.Field(15)),
		field.Enum("is_gift").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(16)),
		field.Time("ship_after_date").
			Default(time.Now).Optional().Annotations(entproto.Field(17)),
		field.Time("ship_by_date").
			Default(time.Now).Optional().Annotations(entproto.Field(18)),
		field.Time("estimated_ship_date").
			Default(time.Now).Optional().Annotations(entproto.Field(19)),
		field.Time("estimated_delivery_date").
			Default(time.Now).Optional().Annotations(entproto.Field(20)),
	}
}

//* entproto annotations ??
func (OrderItemShipGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
func (OrderItemShipGroup) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("order_id", "ship_group_seq_id").
            Unique(),
    }
}

*/

func (OrderItemShipGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

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
		// o2m
		edge.From("order_header", OrderHeader.Type).Ref("order_item_ship_groups").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(21)),
		// m2o
		edge.To("order_adjustments", OrderAdjustment.Type).
			Annotations(entproto.Field(34)),
		// m2o
		edge.To("order_item_ship_group_assocs", OrderItemShipGroupAssoc.Type).
			Annotations(entproto.Field(37)),
		// m2o
		edge.To("order_item_ship_grp_inv_res", OrderItemShipGrpInvRes.Type).
			Annotations(entproto.Field(38)),
		// m2o
		edge.To("order_payment_preferences", OrderPaymentPreference.Type).
			Annotations(entproto.Field(39)),
	}
}

type FacilityType struct {
	ent.Schema
}

// Fields of the FacilityType.
// Unique-Indexes: facilityTypeId
func (FacilityType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (FacilityType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (FacilityType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the FacilityType.
//  one: ParentFacilityType
//  many: Facility
//  many: ChildFacilityType
//  many: FacilityTypeAttr
func (FacilityType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", FacilityType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("child_facility_types", FacilityType.Type).
			Annotations(entproto.Field(6)),
	}
}

type ProductPricePurpose struct {
	ent.Schema
}

// Fields of the ProductPricePurpose.
// Unique-Indexes: productPricePurposeId
func (ProductPricePurpose) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (ProductPricePurpose) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductPricePurpose) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductPricePurpose.
//  many: OrderPaymentPreference
//  many: ProductPaymentMethodType
//  many: ProductPrice
func (ProductPricePurpose) Edges() []ent.Edge {
	return []ent.Edge{
		// m2o
		edge.To("order_payment_preferences", OrderPaymentPreference.Type).
			Annotations(entproto.Field(3)),
		// m2o
		edge.To("product_prices", ProductPrice.Type).
			Annotations(entproto.Field(5)),
	}
}

type RejectionReason struct {
	ent.Schema
}

// Fields of the RejectionReason.
// Unique-Indexes: rejectionId
func (RejectionReason) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (RejectionReason) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (RejectionReason) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the RejectionReason.
//  many: ShipmentReceipt
func (RejectionReason) Edges() []ent.Edge {
	return []ent.Edge{}
}

type OrderHeader struct {
	ent.Schema
}

// Fields of the OrderHeader.
// Unique-Indexes: orderId
func (OrderHeader) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_type_id").Optional().Annotations(entproto.Field(2)),
		field.String("order_name").Optional().Annotations(entproto.Field(3)),
		field.Int("external_id").Optional().Annotations(entproto.Field(4)),
		field.Time("order_date").
			Default(time.Now).Optional().Annotations(entproto.Field(5)),
		field.Enum("priority").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(6)),
		field.Time("entry_date").
			Default(time.Now).Optional().Annotations(entproto.Field(7)),
		field.Time("pick_sheet_printed_date").
			Default(time.Now).Optional().Annotations(entproto.Field(8)),
		field.Int("visit_id").Optional().Annotations(entproto.Field(9)),
		field.String("created_by").Optional().Annotations(entproto.Field(10)),
		field.Int("first_attempt_order_id").Optional().Annotations(entproto.Field(11)),
		field.Int("currency_uom").Optional().Annotations(entproto.Field(12)),
		field.Int("billing_account_id").Optional().Annotations(entproto.Field(13)),
		field.Int("origin_facility_id").Optional().Annotations(entproto.Field(14)),
		field.Int("web_site_id").Optional().Annotations(entproto.Field(15)),
		field.Int("agreement_id").Optional().Annotations(entproto.Field(16)),
		field.String("terminal_id").MaxLen(32).Optional().Annotations(entproto.Field(17)),
		field.String("transaction_id").MaxLen(32).Optional().Annotations(entproto.Field(18)),
		field.Int("auto_order_shopping_list_id").Optional().Annotations(entproto.Field(19)),
		field.Enum("needs_inventory_issuance").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(20)),
		field.Enum("is_rush_order").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(21)),
		field.String("internal_code").MaxLen(32).Optional().Annotations(entproto.Field(22)),
		field.Float("remaining_sub_total").Optional().Annotations(entproto.Field(23)),
		field.Float("grand_total").Optional().Annotations(entproto.Field(24)),
		field.Enum("is_viewed").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(25)),
		field.Enum("invoice_per_shipment").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(26)),
	}
}

//* entproto annotations ??
func (OrderHeader) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (OrderHeader) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

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
		// o2m
		edge.From("sales_channel_enumeration", Enumeration.Type).Ref("sales_channel_order_headers").
			// Bind the "salesChannelEnumId" field to this edge.
			// Field("sales_channel_enum_id").
			Unique().Annotations(entproto.Field(28)),
		// o2m
		edge.From("product_store", ProductStore.Type).Ref("order_headers").
			// Bind the "productStoreId" field to this edge.
			// Field("product_store_id").
			Unique().Annotations(entproto.Field(32)),
		// o2m
		edge.From("status_item", StatusItem.Type).Ref("order_headers").
			// Bind the "statusId" field to this edge.
			// Field("status_id").
			Unique().Annotations(entproto.Field(35)),
		// o2m
		edge.From("sync_status_item", StatusItem.Type).Ref("sync_order_headers").
			// Bind the "syncStatusId" field to this edge.
			// Field("sync_status_id").
			Unique().Annotations(entproto.Field(36)),
		// m2o
		edge.To("order_adjustments", OrderAdjustment.Type).
			Annotations(entproto.Field(46)),
		// m2o
		edge.To("order_contact_meches", OrderContactMech.Type).
			Annotations(entproto.Field(48)),
		// m2o
		edge.To("order_items", OrderItem.Type).
			Annotations(entproto.Field(53)),
		// m2o
		edge.To("order_item_ship_groups", OrderItemShipGroup.Type).
			Annotations(entproto.Field(62)),
		// m2o
		edge.To("order_item_ship_group_assocs", OrderItemShipGroupAssoc.Type).
			Annotations(entproto.Field(63)),
		// m2o
		edge.To("order_item_ship_grp_inv_res", OrderItemShipGrpInvRes.Type).
			Annotations(entproto.Field(64)),
		// m2o
		edge.To("order_payment_preferences", OrderPaymentPreference.Type).
			Annotations(entproto.Field(66)),
		// m2o
		edge.To("order_roles", OrderRole.Type).
			Annotations(entproto.Field(69)),
		// m2o
		edge.To("order_statuses", OrderStatus.Type).
			Annotations(entproto.Field(71)),
	}
}

type ContentType struct {
	ent.Schema
}

// Fields of the ContentType.
// Unique-Indexes: contentTypeId
func (ContentType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (ContentType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ContentType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ContentType.
//  one: ParentContentType
//  many: Content
//  many: ChildContentType
//  many: ContentTypeAttr
func (ContentType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ContentType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("child_content_types", ContentType.Type).
			Annotations(entproto.Field(6)),
	}
}

type ShipmentGatewayConfigType struct {
	ent.Schema
}

// Fields of the ShipmentGatewayConfigType.
// Unique-Indexes: shipmentGatewayConfTypeId
func (ShipmentGatewayConfigType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (ShipmentGatewayConfigType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ShipmentGatewayConfigType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ShipmentGatewayConfigType.
//  one: ParentShipmentGatewayConfigType
//  many: SiblingShipmentGatewayConfigType
//  many: ShipmentGatewayConfig
//  many: ChildShipmentGatewayConfigType
func (ShipmentGatewayConfigType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ShipmentGatewayConfigType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("sibling_shipment_gateway_config_types", ShipmentGatewayConfigType.Type).
			Annotations(entproto.Field(5)),
		// m2o
		edge.To("shipment_gateway_configs", ShipmentGatewayConfig.Type).
			Annotations(entproto.Field(6)),
		// m2o
		edge.To("child_shipment_gateway_config_types", ShipmentGatewayConfigType.Type).
			Annotations(entproto.Field(7)),
	}
}

type ProductContentType struct {
	ent.Schema
}

// Fields of the ProductContentType.
// Unique-Indexes: productContentTypeId
func (ProductContentType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (ProductContentType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductContentType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductContentType.
//  one: ParentProductContentType
//  many: ProductContent
//  many: ChildProductContentType
//  many: ProductPromoContent
func (ProductContentType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProductContentType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("child_product_content_types", ProductContentType.Type).
			Annotations(entproto.Field(6)),
	}
}

type OrderStatus struct {
	ent.Schema
}

// Fields of the OrderStatus.
// Unique-Indexes: orderStatusId
func (OrderStatus) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_item_seq_id").Optional().Annotations(entproto.Field(2)),
		field.Time("status_datetime").
			Default(time.Now).Optional().Annotations(entproto.Field(3)),
		field.String("status_user_login").Optional().Annotations(entproto.Field(4)),
		field.String("change_reason").Optional().Annotations(entproto.Field(5)),
	}
}

//* entproto annotations ??
func (OrderStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (OrderStatus) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the OrderStatus.
//  one: StatusItem
//  one: OrderHeader
//  one-nofk: OrderItem
//  one-nofk: OrderPaymentPreference
//  one: UserLogin
func (OrderStatus) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("status_item", StatusItem.Type).Ref("order_statuses").
			// Bind the "statusId" field to this edge.
			// Field("status_id").
			Unique().Annotations(entproto.Field(6)),
		// o2m
		edge.From("order_header", OrderHeader.Type).Ref("order_statuses").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(7)),
		// o2m
		edge.From("order_item", OrderItem.Type).Ref("order_statuses").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(8)),
		// o2m
		edge.From("order_payment_preference", OrderPaymentPreference.Type).Ref("order_statuses").
			// Bind the "orderPaymentPreferenceId" field to this edge.
			// Field("order_payment_preference_id").
			Unique().Annotations(entproto.Field(9)),
	}
}

type EnumerationType struct {
	ent.Schema
}

// Fields of the EnumerationType.
// Unique-Indexes: enumTypeId
func (EnumerationType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (EnumerationType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (EnumerationType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the EnumerationType.
//  one: ParentEnumerationType
//  many: Enumeration
//  many: ChildEnumerationType
func (EnumerationType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", EnumerationType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("enumerations", Enumeration.Type).
			Annotations(entproto.Field(5)),
		// m2o
		edge.To("child_enumeration_types", EnumerationType.Type).
			Annotations(entproto.Field(6)),
	}
}

type FacilityAssocType struct {
	ent.Schema
}

// Fields of the FacilityAssocType.
// Unique-Indexes: facilityAssocTypeId
func (FacilityAssocType) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (FacilityAssocType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (FacilityAssocType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the FacilityAssocType.
//  many: ProductFacilityAssoc
func (FacilityAssocType) Edges() []ent.Edge {
	return []ent.Edge{}
}

type OrderItem struct {
	ent.Schema
}

// Fields of the OrderItem.
// Unique-Indexes: orderId, orderItemSeqId
func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_item_seq_id").Annotations(entproto.Field(2)),
		field.Int("external_id").Optional().Annotations(entproto.Field(3)),
		field.Int("order_item_type_id").Optional().Annotations(entproto.Field(4)),
		field.Int("order_item_group_seq_id").Optional().Annotations(entproto.Field(5)),
		field.Enum("is_item_group_primary").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(6)),
		field.Int("from_inventory_item_id").Optional().Annotations(entproto.Field(7)),
		field.Int("budget_id").Optional().Annotations(entproto.Field(8)),
		field.Int("budget_item_seq_id").Optional().Annotations(entproto.Field(9)),
		field.String("supplier_product_id").MaxLen(32).Optional().Annotations(entproto.Field(10)),
		field.Int("product_feature_id").Optional().Annotations(entproto.Field(11)),
		field.Int("prod_catalog_id").Optional().Annotations(entproto.Field(12)),
		field.Int("product_category_id").Optional().Annotations(entproto.Field(13)),
		field.Enum("is_promo").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(14)),
		field.Int("quote_id").Optional().Annotations(entproto.Field(15)),
		field.Int("quote_item_seq_id").Optional().Annotations(entproto.Field(16)),
		field.Int("shopping_list_id").Optional().Annotations(entproto.Field(17)),
		field.Int("shopping_list_item_seq_id").Optional().Annotations(entproto.Field(18)),
		field.Int("subscription_id").Optional().Annotations(entproto.Field(19)),
		field.Int("deployment_id").Optional().Annotations(entproto.Field(20)),
		field.Float("quantity").Optional().Annotations(entproto.Field(21)),
		field.Float("cancel_quantity").Optional().Annotations(entproto.Field(22)),
		field.Float("selected_amount").Optional().Annotations(entproto.Field(23)),
		field.Float("unit_price").Optional().Annotations(entproto.Field(24)),
		field.Float("unit_list_price").Optional().Annotations(entproto.Field(25)),
		field.Float("unit_average_cost").Optional().Annotations(entproto.Field(26)),
		field.Float("unit_recurring_price").Optional().Annotations(entproto.Field(27)),
		field.Enum("is_modified_price").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(28)),
		field.Int("recurring_freq_uom_id").Optional().Annotations(entproto.Field(29)),
		field.String("item_description").Optional().Annotations(entproto.Field(30)),
		field.String("comments").Optional().Annotations(entproto.Field(31)),
		field.Int("corresponding_po_id").Optional().Annotations(entproto.Field(32)),
		field.Time("estimated_ship_date").
			Default(time.Now).Optional().Annotations(entproto.Field(33)),
		field.Time("estimated_delivery_date").
			Default(time.Now).Optional().Annotations(entproto.Field(34)),
		field.Time("auto_cancel_date").
			Default(time.Now).Optional().Annotations(entproto.Field(35)),
		field.Time("dont_cancel_set_date").
			Default(time.Now).Optional().Annotations(entproto.Field(36)),
		field.String("dont_cancel_set_user_login").Optional().Annotations(entproto.Field(37)),
		field.Time("ship_before_date").
			Default(time.Now).Optional().Annotations(entproto.Field(38)),
		field.Time("ship_after_date").
			Default(time.Now).Optional().Annotations(entproto.Field(39)),
		field.Time("reserve_after_date").
			Default(time.Now).Optional().Annotations(entproto.Field(40)),
		field.Time("cancel_back_order_date").
			Default(time.Now).Optional().Annotations(entproto.Field(41)),
		field.Int("override_gl_account_id").Optional().Annotations(entproto.Field(42)),
		field.Int("sales_opportunity_id").Optional().Annotations(entproto.Field(43)),
		field.String("change_by_user_login_id").Optional().Annotations(entproto.Field(44)),
	}
}

//* entproto annotations ??
func (OrderItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
func (OrderItem) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("order_id", "order_item_seq_id").
            Unique(),
    }
}

*/

func (OrderItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

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
		// o2m
		edge.From("order_header", OrderHeader.Type).Ref("order_items").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(45)),
		// o2m
		edge.From("product", Product.Type).Ref("order_items").
			// Bind the "productId" field to this edge.
			// Field("product_id").
			Unique().Annotations(entproto.Field(49)),
		// o2m
		edge.From("status_item", StatusItem.Type).Ref("order_items").
			// Bind the "statusId" field to this edge.
			// Field("status_id").
			Unique().Annotations(entproto.Field(52)),
		// m2o
		edge.To("status_valid_changes", StatusValidChange.Type).
			Annotations(entproto.Field(54)),
		// o2m
		edge.From("sync_status_item", StatusItem.Type).Ref("sync_order_items").
			// Bind the "syncStatusId" field to this edge.
			// Field("sync_status_id").
			Unique().Annotations(entproto.Field(55)),
		// m2o
		edge.To("order_adjustments", OrderAdjustment.Type).
			Annotations(entproto.Field(68)),
		// m2o
		edge.To("order_item_ship_group_assocs", OrderItemShipGroupAssoc.Type).
			Annotations(entproto.Field(79)),
		// m2o
		edge.To("order_item_ship_grp_inv_res", OrderItemShipGrpInvRes.Type).
			Annotations(entproto.Field(80)),
		// m2o
		edge.To("order_payment_preferences", OrderPaymentPreference.Type).
			Annotations(entproto.Field(81)),
		// m2o
		edge.To("order_statuses", OrderStatus.Type).
			Annotations(entproto.Field(84)),
	}
}

type QuantityBreakType struct {
	ent.Schema
}

// Fields of the QuantityBreakType.
// Unique-Indexes: quantityBreakTypeId
func (QuantityBreakType) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (QuantityBreakType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (QuantityBreakType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the QuantityBreakType.
//  many: QuantityBreak
func (QuantityBreakType) Edges() []ent.Edge {
	return []ent.Edge{}
}

type OrderPaymentPreference struct {
	ent.Schema
}

// Fields of the OrderPaymentPreference.
// Unique-Indexes: orderPaymentPreferenceId
func (OrderPaymentPreference) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_item_seq_id").Optional().Annotations(entproto.Field(2)),
		field.Int("ship_group_seq_id").Optional().Annotations(entproto.Field(3)),
		field.Int("payment_method_type_id").Optional().Annotations(entproto.Field(4)),
		field.Int("payment_method_id").Optional().Annotations(entproto.Field(5)),
		field.Int("fin_account_id").Optional().Annotations(entproto.Field(6)),
		field.String("security_code").Optional().Annotations(entproto.Field(7)),
		field.String("track_2").Optional().Annotations(entproto.Field(8)),
		field.Enum("present_flag").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(9)),
		field.Enum("swiped_flag").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(10)),
		field.Enum("overflow_flag").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(11)),
		field.Float("max_amount").Optional().Annotations(entproto.Field(12)),
		field.Int("process_attempt").Optional().Annotations(entproto.Field(13)),
		field.String("billing_postal_code").Optional().Annotations(entproto.Field(14)),
		field.String("manual_auth_code").Optional().Annotations(entproto.Field(15)),
		field.String("manual_ref_num").Optional().Annotations(entproto.Field(16)),
		field.Enum("needs_nsf_retry").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(17)),
		field.Time("created_date").
			Default(time.Now).Optional().Annotations(entproto.Field(18)),
		field.String("created_by_user_login").Optional().Annotations(entproto.Field(19)),
		field.Time("last_modified_date").
			Default(time.Now).Optional().Annotations(entproto.Field(20)),
		field.String("last_modified_by_user_login").Optional().Annotations(entproto.Field(21)),
	}
}

//* entproto annotations ??
func (OrderPaymentPreference) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (OrderPaymentPreference) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the OrderPaymentPreference.
//  one: OrderHeader
//  one-nofk: OrderItem
//  one-nofk: OrderItemShipGroup
//  one: ProductPricePurpose
//  one: PaymentMethodType
//  one: PaymentMethod
//  one: FinAccount
//  one: StatusItem
//  one: UserLogin
//  one-nofk: CreditCard
//  one-nofk: EftAccount
//  one-nofk: GiftCard
//  many: OrderStatus
//  many: Payment
//  many: PaymentGatewayResponse
//  many: ReturnItemResponse
func (OrderPaymentPreference) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("order_header", OrderHeader.Type).Ref("order_payment_preferences").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(22)),
		// o2m
		edge.From("order_item", OrderItem.Type).Ref("order_payment_preferences").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(23)),
		// o2m
		edge.From("order_item_ship_group", OrderItemShipGroup.Type).Ref("order_payment_preferences").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(24)),
		// o2m
		edge.From("product_price_purpose", ProductPricePurpose.Type).Ref("order_payment_preferences").
			// Bind the "productPricePurposeId" field to this edge.
			// Field("product_price_purpose_id").
			Unique().Annotations(entproto.Field(25)),
		// o2m
		edge.From("status_item", StatusItem.Type).Ref("order_payment_preferences").
			// Bind the "statusId" field to this edge.
			// Field("status_id").
			Unique().Annotations(entproto.Field(29)),
		// m2o
		edge.To("order_statuses", OrderStatus.Type).
			Annotations(entproto.Field(34)),
	}
}

type ProductFeature struct {
	ent.Schema
}

// Fields of the ProductFeature.
// Unique-Indexes: productFeatureId
func (ProductFeature) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
		field.Int("uom_id").Optional().Annotations(entproto.Field(3)),
		field.Float("number_specified").Optional().Annotations(entproto.Field(4)),
		field.Float("default_amount").Optional().Annotations(entproto.Field(5)),
		field.Int("default_sequence_num").Optional().Annotations(entproto.Field(6)),
		field.Int("abbrev").Optional().Annotations(entproto.Field(7)),
		field.String("id_code").MaxLen(32).Optional().Annotations(entproto.Field(8)),
	}
}

//* entproto annotations ??
func (ProductFeature) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductFeature) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductFeature.
//  one: ProductFeatureCategory
//  one: ProductFeatureType
//  one: Uom
//  many: CostComponent
//  many: DesiredFeature
//  many: InvoiceItem
//  many: ProductFeatureAppl
//  many: ProductFeatureApplAttr
//  many: ProductFeatureDataResource
//  many: ProductFeatureGroupAppl
//  many: MainProductFeatureIactn
//  many: AssocProductFeatureIactn
//  many: ProductManufacturingRule
//  many: QuoteItem
//  many: ShipmentItemFeature
//  many: SupplierProductFeature
func (ProductFeature) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("product_feature_category", ProductFeatureCategory.Type).Ref("product_features").
			// Bind the "productFeatureCategoryId" field to this edge.
			// Field("product_feature_category_id").
			Unique().Annotations(entproto.Field(9)),
		// o2m
		edge.From("product_feature_type", ProductFeatureType.Type).Ref("product_features").
			// Bind the "productFeatureTypeId" field to this edge.
			// Field("product_feature_type_id").
			Unique().Annotations(entproto.Field(10)),
	}
}

type Product struct {
	ent.Schema
}

// Fields of the Product.
// Unique-Indexes: productId
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int("facility_id").Optional().Annotations(entproto.Field(2)),
		field.Time("introduction_date").
			Default(time.Now).Optional().Annotations(entproto.Field(3)),
		field.Time("release_date").
			Default(time.Now).Optional().Annotations(entproto.Field(4)),
		field.Time("support_discontinuation_date").
			Default(time.Now).Optional().Annotations(entproto.Field(5)),
		field.Time("sales_discontinuation_date").
			Default(time.Now).Optional().Annotations(entproto.Field(6)),
		field.Enum("sales_disc_when_not_avail").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(7)),
		field.String("internal_name").Optional().Annotations(entproto.Field(8)),
		field.String("brand_name").Optional().Annotations(entproto.Field(9)),
		field.String("comments").Optional().Annotations(entproto.Field(10)),
		field.String("product_name").Optional().Annotations(entproto.Field(11)),
		field.String("description").Optional().Annotations(entproto.Field(12)),
		field.String("long_description").Optional().Annotations(entproto.Field(13)),
		field.String("price_detail_text").Optional().Annotations(entproto.Field(14)),
		field.String("small_image_url").Optional().Annotations(entproto.Field(15)),
		field.String("medium_image_url").Optional().Annotations(entproto.Field(16)),
		field.String("large_image_url").Optional().Annotations(entproto.Field(17)),
		field.String("detail_image_url").Optional().Annotations(entproto.Field(18)),
		field.String("original_image_url").Optional().Annotations(entproto.Field(19)),
		field.String("detail_screen").Optional().Annotations(entproto.Field(20)),
		field.String("inventory_message").Optional().Annotations(entproto.Field(21)),
		field.Enum("require_inventory").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(22)),
		field.Int("quantity_uom_id").Optional().Annotations(entproto.Field(23)),
		field.Float("quantity_included").Optional().Annotations(entproto.Field(24)),
		field.Int("pieces_included").Optional().Annotations(entproto.Field(25)),
		field.Enum("require_amount").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(26)),
		field.Float("fixed_amount").Optional().Annotations(entproto.Field(27)),
		field.Int("amount_uom_type_id").Optional().Annotations(entproto.Field(28)),
		field.Int("weight_uom_id").Optional().Annotations(entproto.Field(29)),
		field.Float("shipping_weight").Optional().Annotations(entproto.Field(30)),
		field.Float("product_weight").Optional().Annotations(entproto.Field(31)),
		field.Int("height_uom_id").Optional().Annotations(entproto.Field(32)),
		field.Float("product_height").Optional().Annotations(entproto.Field(33)),
		field.Float("shipping_height").Optional().Annotations(entproto.Field(34)),
		field.Int("width_uom_id").Optional().Annotations(entproto.Field(35)),
		field.Float("product_width").Optional().Annotations(entproto.Field(36)),
		field.Float("shipping_width").Optional().Annotations(entproto.Field(37)),
		field.Int("depth_uom_id").Optional().Annotations(entproto.Field(38)),
		field.Float("product_depth").Optional().Annotations(entproto.Field(39)),
		field.Float("shipping_depth").Optional().Annotations(entproto.Field(40)),
		field.Int("diameter_uom_id").Optional().Annotations(entproto.Field(41)),
		field.Float("product_diameter").Optional().Annotations(entproto.Field(42)),
		field.Float("product_rating").Optional().Annotations(entproto.Field(43)),
		field.Enum("returnable").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(44)),
		field.Enum("taxable").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(45)),
		field.Enum("charge_shipping").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(46)),
		field.Enum("auto_create_keywords").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(47)),
		field.Enum("include_in_promotions").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(48)),
		field.Enum("is_virtual").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(49)),
		field.Enum("is_variant").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(50)),
		field.Int("origin_geo_id").Optional().Annotations(entproto.Field(51)),
		field.Int("bill_of_material_level").Optional().Annotations(entproto.Field(52)),
		field.Float("reserv_max_persons").Optional().Annotations(entproto.Field(53)),
		field.Float("reserv_2_nd_pp_perc").Optional().Annotations(entproto.Field(54)),
		field.Float("reserv_nth_pp_perc").Optional().Annotations(entproto.Field(55)),
		field.Int("config_id").Optional().Annotations(entproto.Field(56)),
		field.Time("created_date").
			Default(time.Now).Optional().Annotations(entproto.Field(57)),
		field.String("created_by_user_login").Optional().Annotations(entproto.Field(58)),
		field.Time("last_modified_date").
			Default(time.Now).Optional().Annotations(entproto.Field(59)),
		field.String("last_modified_by_user_login").Optional().Annotations(entproto.Field(60)),
		field.Enum("in_shipping_box").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(61)),
		field.Int("default_shipment_box_type_id").Optional().Annotations(entproto.Field(62)),
		field.String("lot_id_filled_in").Optional().Annotations(entproto.Field(63)),
		field.Enum("order_decimal_quantity").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(64)),
	}
}

//* entproto annotations ??
func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the Product.
//  one: ProductType
//  many: ProductTypeAttr
//  one: PrimaryProductCategory
//  one: Facility
//  one: QuantityUom
//  one: AmountUomType
//  one: WeightUom
//  one: HeightUom
//  one: WidthUom
//  one: DepthUom
//  one: DiameterUom
//  one: VirtualVariantMethodEnumeration
//  one: RatingEnumeration
//  one: RequirementMethodEnumeration
//  one: OriginGeo
//  one: CreatedByUserLogin
//  one: LastModifiedByUserLogin
//  many: ProductFeatureAndAppl
//  one: DefaultShipmentBoxType
//  one: InventoryItemType
//  many: Agreement
//  many: AgreementProductAppl
//  many: CartAbandonedLine
//  many: CommunicationEventProduct
//  many: CostComponent
//  many: CustRequestItem
//  many: InstanceOfFixedAsset
//  many: FixedAssetProduct
//  many: GoodIdentification
//  many: InventoryItem
//  many: InventoryItemTempRes
//  many: InvoiceItem
//  many: MrpEvent
//  many: OrderItem
//  many: OrderSummaryEntry
//  many: PartyNeed
//  many: MainProductAssoc
//  many: AssocProductAssoc
//  many: ProductAttribute
//  many: ProductAverageCost
//  one-nofk: ProductCalculatedInfo
//  many: ProductCategoryMember
//  many: ProductProductConfig
//  many: ProductProductConfigProduct
//  many: ProductProductConfigStats
//  many: ProductContent
//  many: ProductCostComponentCalc
//  many: ProductFacility
//  many: ProductFacilityAssoc
//  many: ProductFacilityLocation
//  many: ProductFeatureAppl
//  many: ProductFeatureApplAttr
//  many: ProductGeo
//  many: ProductGlAccount
//  many: ProductGroupOrder
//  many: ProductKeyword
//  many: ProductMaint
//  many: ProductManufacturingRule
//  many: ProductForProductManufacturingRule
//  many: ProductInProductManufacturingRule
//  many: ProductSubstProductManufacturingRule
//  many: ProductMeter
//  many: ProductOrderItem
//  many: ProductPaymentMethodType
//  many: ProductPrice
//  many: ProductPromoProduct
//  many: ProductReview
//  many: ProductRole
//  many: ProductStoreSurveyAppl
//  many: ProductSubscriptionResource
//  many: QuoteItem
//  many: ReorderGuideline
//  many: Requirement
//  many: ReturnItem
//  many: SalesForecastDetail
//  many: ShipmentItem
//  many: SubShipmentPackageContent
//  many: ShipmentReceipt
//  many: ShoppingListItem
//  many: Subscription
//  many: SupplierProduct
//  many: VendorProduct
//  many: WorkEffortGoodStandard
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("product_type", ProductType.Type).Ref("products").
			// Bind the "productTypeId" field to this edge.
			// Field("product_type_id").
			Unique().Annotations(entproto.Field(65)),
		// o2m
		edge.From("primary_product_category", ProductCategory.Type).Ref("primary_products").
			// Bind the "primaryProductCategoryId" field to this edge.
			// Field("primary_product_category_id").
			Unique().Annotations(entproto.Field(67)),
		// o2m
		edge.From("virtual_variant_method_enumeration", Enumeration.Type).Ref("virtual_variant_method_products").
			// Bind the "virtualVariantMethodEnum" field to this edge.
			// Field("virtual_variant_method_enum").
			Unique().Annotations(entproto.Field(76)),
		// o2m
		edge.From("rating_enumeration", Enumeration.Type).Ref("rating_products").
			// Bind the "ratingTypeEnum" field to this edge.
			// Field("rating_type_enum").
			Unique().Annotations(entproto.Field(77)),
		// o2m
		edge.From("requirement_method_enumeration", Enumeration.Type).Ref("requirement_method_products").
			// Bind the "requirementMethodEnumId" field to this edge.
			// Field("requirement_method_enum_id").
			Unique().Annotations(entproto.Field(78)),
		// o2m
		edge.From("inventory_item_type", InventoryItemType.Type).Ref("products").
			// Bind the "inventoryItemTypeId" field to this edge.
			// Field("inventory_item_type_id").
			Unique().Annotations(entproto.Field(84)),
		// m2o
		edge.To("order_items", OrderItem.Type).
			Annotations(entproto.Field(98)),
		// m2o
		edge.To("main_product_assocs", ProductAssoc.Type).
			Annotations(entproto.Field(101)),
		// m2o
		edge.To("assoc_product_assocs", ProductAssoc.Type).
			Annotations(entproto.Field(102)),
		// m2o
		edge.To("product_prices", ProductPrice.Type).
			Annotations(entproto.Field(129)),
		// m2o
		edge.To("product_reviews", ProductReview.Type).
			Annotations(entproto.Field(131)),
	}
}

type Enumeration struct {
	ent.Schema
}

// Fields of the Enumeration.
// Unique-Indexes: enumId
func (Enumeration) Fields() []ent.Field {
	return []ent.Field{
		field.String("enum_code").Optional().Annotations(entproto.Field(2)),
		field.Int("sequence_id").Optional().Annotations(entproto.Field(3)),
		field.String("description").Optional().Annotations(entproto.Field(4)),
	}
}

//* entproto annotations ??
func (Enumeration) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (Enumeration) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the Enumeration.
//  one: EnumerationType
//  many: Plan Method Enum IdAllocationPlanItem
//  many: CommunicationEvent
//  many: PrivilegeContent
//  many: ContentPurposeOperation
//  many: SalesChannelCustRequest
//  many: EbayShippingMethod
//  many: EmailTemplateSetting
//  many: ExampleFeature
//  many: ImportStatusExcelImportHistory
//  many: ThruReasonExcelImportHistory
//  many: TypeFacilityLocation
//  many: ReasonFinAccountTrans
//  many: ReplenishFinAccountType
//  many: ClassFixedAsset
//  many: GeoPointTypeGeoPoint
//  many: GiftCardFulfillment
//  many: ReasonInventoryItemDetail
//  many: JobInterview
//  many: ReasonJobManagerLock
//  many: ExamTypeJobRequisition
//  many: JobPostingTypeJobRequisition
//  many: RelationshipKeywordThesaurus
//  many: SalesChannelOrderHeader
//  many: OrderItemChange
//  many: ReasonOrderItemChange
//  many: OrderNotification
//  many: TaxFormPartyAcctgPreference
//  many: CogsPartyAcctgPreference
//  many: InvoiceSequencePartyAcctgPreference
//  many: QuoteSequencePartyAcctgPreference
//  many: OrderSequencePartyAcctgPreference
//  many: ServiceTypePaymentGatewayResponse
//  many: TranCodePaymentGatewayResponse
//  many: EmploymentStatusPerson
//  many: ResidenceStatusPerson
//  many: MaritalStatusPerson
//  many: PosTerminalInternTx
//  many: VirtualVariantMethodProduct
//  many: RatingProduct
//  many: RequirementMethodProduct
//  many: LinkTypeProductCategoryLink
//  many: RequirementMethodProductFacility
//  many: ProductFacility
//  many: ProductGeo
//  many: ProductKeyword
//  many: InputParamProductPriceCond
//  many: OperatorProductPriceCond
//  many: ActionProductPromoAction
//  many: ApplProductPromoCategory
//  many: InputParamProductPromoCond
//  many: OperatorProductPromoCond
//  many: ApplProductPromoProduct
//  many: ReserveOrderProductStore
//  many: RequirementMethodProductStore
//  many: DefaultSalesChannelProductStore
//  many: StoreCreditAccountProductStore
//  many: ProductStoreEmailSetting
//  many: ReplenishMethodProductStoreFinActSetting
//  many: ProductStoreKeywordOvrd
//  many: ProductStorePaymentSetting
//  many: ProductStoreTelecomSetting
//  many: CreditCardProductStoreVendorPayment
//  many: SalesChannelQuote
//  many: TypeSalesOpportunity
//  many: TrackingCodeVisit
//  many: PurposeUomConversionDated
//  many: VisualThemeResource
//  many: ScopeWorkEffort
//  many: ExpectationWorkEffortPartyAssignment
//  many: DelegateReasonWorkEffortPartyAssignment
//  many: WorkloadFeature
func (Enumeration) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("enumeration_type", EnumerationType.Type).Ref("enumerations").
			// Bind the "enumTypeId" field to this edge.
			// Field("enum_type_id").
			Unique().Annotations(entproto.Field(5)),
		// m2o
		edge.To("reason_inventory_item_details", InventoryItemDetail.Type).
			Annotations(entproto.Field(22)),
		// m2o
		edge.To("sales_channel_order_headers", OrderHeader.Type).
			Annotations(entproto.Field(28)),
		// m2o
		edge.To("virtual_variant_method_products", Product.Type).
			Annotations(entproto.Field(43)),
		// m2o
		edge.To("rating_products", Product.Type).
			Annotations(entproto.Field(44)),
		// m2o
		edge.To("requirement_method_products", Product.Type).
			Annotations(entproto.Field(45)),
		// m2o
		edge.To("reserve_order_product_stores", ProductStore.Type).
			Annotations(entproto.Field(58)),
		// m2o
		edge.To("requirement_method_product_stores", ProductStore.Type).
			Annotations(entproto.Field(59)),
		// m2o
		edge.To("default_sales_channel_product_stores", ProductStore.Type).
			Annotations(entproto.Field(60)),
		// m2o
		edge.To("store_credit_account_product_stores", ProductStore.Type).
			Annotations(entproto.Field(61)),
	}
}

type ProductCategoryType struct {
	ent.Schema
}

// Fields of the ProductCategoryType.
// Unique-Indexes: productCategoryTypeId
func (ProductCategoryType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (ProductCategoryType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductCategoryType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductCategoryType.
//  one: ParentProductCategoryType
//  many: ProductCategory
//  many: ChildProductCategoryType
//  many: ProductCategoryTypeAttr
func (ProductCategoryType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProductCategoryType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("product_categories", ProductCategory.Type).
			Annotations(entproto.Field(5)),
		// m2o
		edge.To("child_product_category_types", ProductCategoryType.Type).
			Annotations(entproto.Field(6)),
	}
}

type ShipmentContactMechType struct {
	ent.Schema
}

// Fields of the ShipmentContactMechType.
// Unique-Indexes: shipmentContactMechTypeId
func (ShipmentContactMechType) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (ShipmentContactMechType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ShipmentContactMechType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ShipmentContactMechType.
//  many: ShipmentContactMech
func (ShipmentContactMechType) Edges() []ent.Edge {
	return []ent.Edge{}
}

type ProductMaintType struct {
	ent.Schema
}

// Fields of the ProductMaintType.
// Unique-Indexes: productMaintTypeId
func (ProductMaintType) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (ProductMaintType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductMaintType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductMaintType.
//  one: ParentProductMaintType
//  many: FixedAssetMaint
//  many: ProductMaint
//  many: ChildProductMaintType
func (ProductMaintType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProductMaintType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(3)),
		// m2o
		edge.To("child_product_maint_types", ProductMaintType.Type).
			Annotations(entproto.Field(6)),
	}
}

type ProductCategoryContentType struct {
	ent.Schema
}

// Fields of the ProductCategoryContentType.
// Unique-Indexes: prodCatContentTypeId
func (ProductCategoryContentType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (ProductCategoryContentType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductCategoryContentType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductCategoryContentType.
//  one: ParentProductCategoryContentType
//  many: ProductCategoryContent
//  many: ChildProductCategoryContentType
func (ProductCategoryContentType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProductCategoryContentType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("child_product_category_content_types", ProductCategoryContentType.Type).
			Annotations(entproto.Field(6)),
	}
}

type OrderAdjustment struct {
	ent.Schema
}

// Fields of the OrderAdjustment.
// Unique-Indexes: orderAdjustmentId
func (OrderAdjustment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_adjustment_type_id").Optional().Annotations(entproto.Field(2)),
		field.Int("order_item_seq_id").Optional().Annotations(entproto.Field(3)),
		field.Int("ship_group_seq_id").Optional().Annotations(entproto.Field(4)),
		field.String("comments").Optional().Annotations(entproto.Field(5)),
		field.String("description").Optional().Annotations(entproto.Field(6)),
		field.Float("amount").Optional().Annotations(entproto.Field(7)),
		field.Float("recurring_amount").Optional().Annotations(entproto.Field(8)),
		field.Float("amount_already_included").Optional().Annotations(entproto.Field(9)),
		field.Int("product_promo_id").Optional().Annotations(entproto.Field(10)),
		field.Int("product_promo_rule_id").Optional().Annotations(entproto.Field(11)),
		field.Int("product_promo_action_seq_id").Optional().Annotations(entproto.Field(12)),
		field.Int("product_feature_id").Optional().Annotations(entproto.Field(13)),
		field.Int("corresponding_product_id").Optional().Annotations(entproto.Field(14)),
		field.Int("tax_authority_rate_seq_id").Optional().Annotations(entproto.Field(15)),
		field.String("source_reference_id").MaxLen(32).Optional().Annotations(entproto.Field(16)),
		field.Float("source_percentage").Optional().Annotations(entproto.Field(17)),
		field.String("customer_reference_id").MaxLen(32).Optional().Annotations(entproto.Field(18)),
		field.Int("primary_geo_id").Optional().Annotations(entproto.Field(19)),
		field.Int("secondary_geo_id").Optional().Annotations(entproto.Field(20)),
		field.Float("exempt_amount").Optional().Annotations(entproto.Field(21)),
		field.Int("tax_auth_geo_id").Optional().Annotations(entproto.Field(22)),
		field.Int("tax_auth_party_id").Optional().Annotations(entproto.Field(23)),
		field.Int("override_gl_account_id").Optional().Annotations(entproto.Field(24)),
		field.Enum("include_in_tax").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(25)),
		field.Enum("include_in_shipping").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(26)),
		field.Enum("is_manual").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(27)),
		field.Time("created_date").
			Default(time.Now).Optional().Annotations(entproto.Field(28)),
		field.String("created_by_user_login").Optional().Annotations(entproto.Field(29)),
		field.Time("last_modified_date").
			Default(time.Now).Optional().Annotations(entproto.Field(30)),
		field.String("last_modified_by_user_login").Optional().Annotations(entproto.Field(31)),
	}
}

//* entproto annotations ??
func (OrderAdjustment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (OrderAdjustment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the OrderAdjustment.
//  one: OrderAdjustmentType
//  many: OrderAdjustmentTypeAttr
//  one: OrderHeader
//  one: UserLogin
//  one-nofk: OrderItem
//  one-nofk: OrderItemShipGroup
//  one-nofk: OrderItemShipGroupAssoc
//  one: ProductPromo
//  one-nofk: ProductPromoRule
//  one-nofk: ProductPromoAction
//  one: PrimaryGeo
//  one: SecondaryGeo
//  one: TaxAuthority
//  one: OverrideGlAccount
//  one: TaxAuthorityRateProduct
//  one-nofk: OrderAdjustment
//  many: OrderAdjustmentAttribute
//  many: OrderAdjustmentBilling
//  many: ReturnAdjustment
func (OrderAdjustment) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("order_header", OrderHeader.Type).Ref("order_adjustments").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(34)),
		// o2m
		edge.From("order_item", OrderItem.Type).Ref("order_adjustments").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(36)),
		// o2m
		edge.From("order_item_ship_group", OrderItemShipGroup.Type).Ref("order_adjustments").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(37)),
		// o2m
		edge.From("order_item_ship_group_assoc", OrderItemShipGroupAssoc.Type).Ref("order_adjustments").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(38)),
		edge.To("children", OrderAdjustment.Type).
			From("parent").
			// Bind the "originalAdjustmentId" field to this edge.
			// Field("original_adjustment_id").
			Unique().Annotations(entproto.Field(47)),
	}
}

type FacilityGroup struct {
	ent.Schema
}

// Fields of the FacilityGroup.
// Unique-Indexes: facilityGroupId
func (FacilityGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("facility_group_name").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (FacilityGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (FacilityGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the FacilityGroup.
//  one: FacilityGroupType
//  one: PrimaryParentFacilityGroup
//  many: Facility
//  many: FacilityGroupMember
//  many: FacilityGroupRole
//  many: CurrentFacilityGroupRollup
//  many: ParentFacilityGroupRollup
func (FacilityGroup) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("facility_group_type", FacilityGroupType.Type).Ref("facility_groups").
			// Bind the "facilityGroupTypeId" field to this edge.
			// Field("facility_group_type_id").
			Unique().Annotations(entproto.Field(4)),
		edge.To("children", FacilityGroup.Type).
			From("parent").
			// Bind the "primaryParentGroupId" field to this edge.
			// Field("primary_parent_group_id").
			Unique().Annotations(entproto.Field(5)),
	}
}

type OrderItemShipGrpInvRes struct {
	ent.Schema
}

// Fields of the OrderItemShipGrpInvRes.
// Unique-Indexes: orderId, shipGroupSeqId, orderItemSeqId, inventoryItemId
func (OrderItemShipGrpInvRes) Fields() []ent.Field {
	return []ent.Field{
		field.Int("ship_group_seq_id").Annotations(entproto.Field(2)),
		field.Int("order_item_seq_id").Annotations(entproto.Field(3)),
		field.Int("inventory_item_id").Annotations(entproto.Field(4)),
		field.Int("reserve_order_enum_id").Optional().Annotations(entproto.Field(5)),
		field.Float("quantity").Optional().Annotations(entproto.Field(6)),
		field.Float("quantity_not_available").Optional().Annotations(entproto.Field(7)),
		field.Time("reserved_datetime").
			Default(time.Now).Optional().Annotations(entproto.Field(8)),
		field.Time("created_datetime").
			Default(time.Now).Optional().Annotations(entproto.Field(9)),
		field.Time("promised_datetime").
			Default(time.Now).Optional().Annotations(entproto.Field(10)),
		field.Time("current_promised_date").
			Default(time.Now).Optional().Annotations(entproto.Field(11)),
		field.Enum("priority").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(12)),
		field.Int("sequence_id").Optional().Annotations(entproto.Field(13)),
		field.Time("old_pick_start_date").
			Default(time.Now).Optional().Annotations(entproto.Field(14)),
	}
}

//* entproto annotations ??
func (OrderItemShipGrpInvRes) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
func (OrderItemShipGrpInvRes) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("order_id", "ship_group_seq_id", "order_item_seq_id", "inventory_item_id").
            Unique(),
    }
}

*/

func (OrderItemShipGrpInvRes) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the OrderItemShipGrpInvRes.
//  one-nofk: OrderHeader
//  one: OrderItem
//  one-nofk: OrderItemShipGroup
//  one-nofk: OrderItemShipGroupAssoc
//  one: InventoryItem
//  many: InventoryItemDetail
//  many: ItemIssuance
//  many: PicklistItem
func (OrderItemShipGrpInvRes) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("order_header", OrderHeader.Type).Ref("order_item_ship_grp_inv_res").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(15)),
		// o2m
		edge.From("order_item", OrderItem.Type).Ref("order_item_ship_grp_inv_res").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(16)),
		// o2m
		edge.From("order_item_ship_group", OrderItemShipGroup.Type).Ref("order_item_ship_grp_inv_res").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(17)),
		// o2m
		edge.From("order_item_ship_group_assoc", OrderItemShipGroupAssoc.Type).Ref("order_item_ship_grp_inv_res").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(18)),
		// m2o
		edge.To("inventory_item_details", InventoryItemDetail.Type).
			Annotations(entproto.Field(20)),
	}
}

type ShipmentGatewayFedex struct {
	ent.Schema
}

// Fields of the ShipmentGatewayFedex.
// Unique-Indexes: shipmentGatewayConfigId
func (ShipmentGatewayFedex) Fields() []ent.Field {
	return []ent.Field{
		field.String("connect_url").Optional().Annotations(entproto.Field(2)),
		field.String("connect_soap_url").Optional().Annotations(entproto.Field(3)),
		field.Int("connect_timeout").Optional().Annotations(entproto.Field(4)),
		field.String("access_account_nbr").Optional().Annotations(entproto.Field(5)),
		field.String("access_meter_number").Optional().Annotations(entproto.Field(6)),
		field.String("access_user_key").Optional().Annotations(entproto.Field(7)),
		field.String("access_user_pwd").Optional().Annotations(entproto.Field(8)),
		field.String("label_image_type").Optional().Annotations(entproto.Field(9)),
		field.String("default_dropoff_type").Optional().Annotations(entproto.Field(10)),
		field.String("default_packaging_type").Optional().Annotations(entproto.Field(11)),
		field.String("template_shipment").Optional().Annotations(entproto.Field(12)),
		field.String("template_subscription").Optional().Annotations(entproto.Field(13)),
		field.String("rate_estimate_template").Optional().Annotations(entproto.Field(14)),
	}
}

//* entproto annotations ??
func (ShipmentGatewayFedex) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ShipmentGatewayFedex) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ShipmentGatewayFedex.
//  one: ShipmentGatewayConfig
func (ShipmentGatewayFedex) Edges() []ent.Edge {
	return []ent.Edge{
		// o2o (nofk)
		edge.From("shipment_gateway_config", ShipmentGatewayConfig.Type).
			Ref("shipment_gateway_fedex").
			// Bind the "shipmentGatewayConfigId" field to this edge.
			// Field("shipment_gateway_config_id").
			Unique().Annotations(entproto.Field(15)),
	}
}

type ProductCategory struct {
	ent.Schema
}

// Fields of the ProductCategory.
// Unique-Indexes: productCategoryId
func (ProductCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("category_name").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
		field.String("long_description").Optional().Annotations(entproto.Field(4)),
		field.String("category_image_url").Optional().Annotations(entproto.Field(5)),
		field.String("link_one_image_url").Optional().Annotations(entproto.Field(6)),
		field.String("link_two_image_url").Optional().Annotations(entproto.Field(7)),
		field.String("detail_screen").Optional().Annotations(entproto.Field(8)),
		field.Enum("show_in_select").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(9)),
	}
}

//* entproto annotations ??
func (ProductCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductCategory.
//  one: ProductCategoryType
//  many: ProductCategoryTypeAttr
//  one: PrimaryParentProductCategory
//  many: PrimaryChildProductCategory
//  many: MarketInterest
//  many: PartyNeed
//  many: ProdCatalogCategory
//  many: PrimaryProduct
//  many: ProductCategoryAttribute
//  many: ProductCategoryContent
//  many: ProductCategoryGlAccount
//  many: ProductCategoryLink
//  many: ProductCategoryMember
//  many: ProductCategoryRole
//  many: CurrentProductCategoryRollup
//  many: ParentProductCategoryRollup
//  many: ProductFeatureCatGrpAppl
//  many: ProductFeatureCategoryAppl
//  many: ProductPromoCategory
//  many: ProductStoreSurveyAppl
//  many: SalesForecastDetail
//  many: Subscription
//  many: TaxAuthorityCategory
//  many: TaxAuthorityRateProduct
func (ProductCategory) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("product_category_type", ProductCategoryType.Type).Ref("product_categories").
			// Bind the "productCategoryTypeId" field to this edge.
			// Field("product_category_type_id").
			Unique().Annotations(entproto.Field(10)),
		edge.To("children", ProductCategory.Type).
			From("parent").
			// Bind the "primaryParentCategoryId" field to this edge.
			// Field("primary_parent_category_id").
			Unique().Annotations(entproto.Field(12)),
		// m2o
		edge.To("primary_child_product_categories", ProductCategory.Type).
			Annotations(entproto.Field(13)),
		// m2o
		edge.To("primary_products", Product.Type).
			Annotations(entproto.Field(17)),
	}
}

type CostComponentType struct {
	ent.Schema
}

// Fields of the CostComponentType.
// Unique-Indexes: costComponentTypeId
func (CostComponentType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (CostComponentType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (CostComponentType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the CostComponentType.
//  one: ParentCostComponentType
//  many: CostComponent
//  many: ChildCostComponentType
//  many: CostComponentTypeAttr
//  many: ProductCostComponentCalc
//  many: WorkEffortCostCalc
func (CostComponentType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", CostComponentType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("child_cost_component_types", CostComponentType.Type).
			Annotations(entproto.Field(6)),
	}
}

type SubscriptionType struct {
	ent.Schema
}

// Fields of the SubscriptionType.
// Unique-Indexes: subscriptionTypeId
func (SubscriptionType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("has_table").
			Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
		field.String("description").Optional().Annotations(entproto.Field(3)),
	}
}

//* entproto annotations ??
func (SubscriptionType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (SubscriptionType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the SubscriptionType.
//  one: ParentSubscriptionType
//  many: Subscription
//  many: ChildSubscriptionType
//  many: SubscriptionTypeAttr
func (SubscriptionType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", SubscriptionType.Type).
			From("parent").
			// Bind the "parentTypeId" field to this edge.
			// Field("parent_type_id").
			Unique().Annotations(entproto.Field(4)),
		// m2o
		edge.To("child_subscription_types", SubscriptionType.Type).
			Annotations(entproto.Field(6)),
	}
}

type InventoryItemDetail struct {
	ent.Schema
}

// Fields of the InventoryItemDetail.
// Unique-Indexes: inventoryItemId, inventoryItemDetailSeqId
func (InventoryItemDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int("inventory_item_detail_seq_id").Annotations(entproto.Field(2)),
		field.Time("effective_date").
			Default(time.Now).Optional().Annotations(entproto.Field(3)),
		field.Float("quantity_on_hand_diff").Optional().Annotations(entproto.Field(4)),
		field.Float("available_to_promise_diff").Optional().Annotations(entproto.Field(5)),
		field.Float("accounting_quantity_diff").Optional().Annotations(entproto.Field(6)),
		field.Float("unit_cost").Optional().Annotations(entproto.Field(7)),
		field.Int("order_item_seq_id").Optional().Annotations(entproto.Field(8)),
		field.Int("ship_group_seq_id").Optional().Annotations(entproto.Field(9)),
		field.Int("shipment_id").Optional().Annotations(entproto.Field(10)),
		field.Int("shipment_item_seq_id").Optional().Annotations(entproto.Field(11)),
		field.Int("return_id").Optional().Annotations(entproto.Field(12)),
		field.Int("return_item_seq_id").Optional().Annotations(entproto.Field(13)),
		field.Int("work_effort_id").Optional().Annotations(entproto.Field(14)),
		field.Int("fixed_asset_id").Optional().Annotations(entproto.Field(15)),
		field.Int("maint_hist_seq_id").Optional().Annotations(entproto.Field(16)),
		field.Int("item_issuance_id").Optional().Annotations(entproto.Field(17)),
		field.Int("receipt_id").Optional().Annotations(entproto.Field(18)),
		field.Int("physical_inventory_id").Optional().Annotations(entproto.Field(19)),
		field.String("description").Optional().Annotations(entproto.Field(20)),
	}
}

//* entproto annotations ??
func (InventoryItemDetail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
func (InventoryItemDetail) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("inventory_item_id", "inventory_item_detail_seq_id").
            Unique(),
    }
}

*/

func (InventoryItemDetail) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the InventoryItemDetail.
//  one: InventoryItem
//  one: WorkEffort
//  one-nofk: OrderItemShipGrpInvRes
//  one: FixedAssetMaint
//  one: ItemIssuance
//  one-nofk: WorkEffortInventoryAssign
//  one-nofk: WorkEffortInventoryProduced
//  one: ShipmentReceipt
//  one: PhysicalInventory
//  one: ReasonEnumeration
//  one-nofk: InventoryItemVariance
func (InventoryItemDetail) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("order_item_ship_grp_inv_res", OrderItemShipGrpInvRes.Type).Ref("inventory_item_details").
			// Bind the "orderId" field to this edge.
			// Field("order_id").
			Unique().Annotations(entproto.Field(23)),
		// o2m
		edge.From("reason_enumeration", Enumeration.Type).Ref("reason_inventory_item_details").
			// Bind the "reasonEnumId" field to this edge.
			// Field("reason_enum_id").
			Unique().Annotations(entproto.Field(30)),
	}
}

type ShipmentGatewayDhl struct {
	ent.Schema
}

// Fields of the ShipmentGatewayDhl.
// Unique-Indexes: shipmentGatewayConfigId
func (ShipmentGatewayDhl) Fields() []ent.Field {
	return []ent.Field{
		field.String("connect_url").Optional().Annotations(entproto.Field(2)),
		field.Int("connect_timeout").Optional().Annotations(entproto.Field(3)),
		field.String("head_version").Optional().Annotations(entproto.Field(4)),
		field.String("head_action").Optional().Annotations(entproto.Field(5)),
		field.String("access_user_id").Optional().Annotations(entproto.Field(6)),
		field.String("access_password").Optional().Annotations(entproto.Field(7)),
		field.String("access_account_nbr").Optional().Annotations(entproto.Field(8)),
		field.String("access_shipping_key").Optional().Annotations(entproto.Field(9)),
		field.String("label_image_format").Optional().Annotations(entproto.Field(10)),
		field.String("rate_estimate_template").Optional().Annotations(entproto.Field(11)),
	}
}

//* entproto annotations ??
func (ShipmentGatewayDhl) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ShipmentGatewayDhl) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ShipmentGatewayDhl.
//  one: ShipmentGatewayConfig
func (ShipmentGatewayDhl) Edges() []ent.Edge {
	return []ent.Edge{
		// o2o (nofk)
		edge.From("shipment_gateway_config", ShipmentGatewayConfig.Type).
			Ref("shipment_gateway_dhl").
			// Bind the "shipmentGatewayConfigId" field to this edge.
			// Field("shipment_gateway_config_id").
			Unique().Annotations(entproto.Field(12)),
	}
}

type ProductFeatureCategory struct {
	ent.Schema
}

// Fields of the ProductFeatureCategory.
// Unique-Indexes: productFeatureCategoryId
func (ProductFeatureCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional().Annotations(entproto.Field(2)),
	}
}

//* entproto annotations ??
func (ProductFeatureCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductFeatureCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductFeatureCategory.
//  one: ParentProductFeatureCategory
//  many: ProductFeature
//  many: ChildProductFeatureCategory
//  many: ProductFeatureCategoryAppl
func (ProductFeatureCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProductFeatureCategory.Type).
			From("parent").
			// Bind the "parentCategoryId" field to this edge.
			// Field("parent_category_id").
			Unique().Annotations(entproto.Field(3)),
		// m2o
		edge.To("product_features", ProductFeature.Type).
			Annotations(entproto.Field(4)),
		// m2o
		edge.To("child_product_feature_categories", ProductFeatureCategory.Type).
			Annotations(entproto.Field(5)),
	}
}

type ProductConfigItem struct {
	ent.Schema
}

// Fields of the ProductConfigItem.
// Unique-Indexes: configItemId
func (ProductConfigItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("config_item_type_id").Optional().Annotations(entproto.Field(2)),
		field.String("config_item_name").Optional().Annotations(entproto.Field(3)),
		field.String("description").Optional().Annotations(entproto.Field(4)),
		field.String("long_description").Optional().Annotations(entproto.Field(5)),
		field.String("image_url").Optional().Annotations(entproto.Field(6)),
	}
}

//* entproto annotations ??
func (ProductConfigItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
 */

func (ProductConfigItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductConfigItem.
//  many: ProdConfItemContent
//  many: ConfigItemProductConfig
//  many: ConfigItemProductConfigConfig
//  many: ConfigItemProductConfigOption
//  many: ConfigItemProductConfigOptionIactn
//  many: ConfigItemToProductConfigOptionIactn
//  many: ConfigItemProductConfigProduct
func (ProductConfigItem) Edges() []ent.Edge {
	return []ent.Edge{}
}

type ProductAssoc struct {
	ent.Schema
}

// Fields of the ProductAssoc.
// Unique-Indexes: productId, productIdTo, productAssocTypeId, fromDate
func (ProductAssoc) Fields() []ent.Field {
	return []ent.Field{
		field.Time("from_date").
			Default(time.Now).Annotations(entproto.Field(2)),
		field.Time("thru_date").
			Default(time.Now).Optional().Annotations(entproto.Field(3)),
		field.Int("sequence_num").Optional().Annotations(entproto.Field(4)),
		field.String("reason").Optional().Annotations(entproto.Field(5)),
		field.Float("quantity").Optional().Annotations(entproto.Field(6)),
		field.Float("scrap_factor").Optional().Annotations(entproto.Field(7)),
		field.String("instruction").Optional().Annotations(entproto.Field(8)),
		field.Int("routing_work_effort_id").Optional().Annotations(entproto.Field(9)),
		field.Int("recurrence_info_id").Optional().Annotations(entproto.Field(10)),
	}
}

//* entproto annotations ??
func (ProductAssoc) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

//*/

/*
func (ProductAssoc) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("product_id", "product_id_to", "product_assoc_type_id", "from_date").
            Unique(),
    }
}

*/

func (ProductAssoc) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Edges of the ProductAssoc.
//  one: ProductAssocType
//  one: MainProduct
//  one: AssocProduct
//  one: RoutingWorkEffort
//  one: CustomMethod
//  one: RecurrenceInfo
func (ProductAssoc) Edges() []ent.Edge {
	return []ent.Edge{
		// o2m
		edge.From("product_assoc_type", ProductAssocType.Type).Ref("product_assocs").
			// Bind the "productAssocTypeId" field to this edge.
			// Field("product_assoc_type_id").
			Unique().Annotations(entproto.Field(11)),
		// o2m
		edge.From("main_product", Product.Type).Ref("main_product_assocs").
			// Bind the "productId" field to this edge.
			// Field("product_id").
			Unique().Annotations(entproto.Field(12)),
		// o2m
		edge.From("assoc_product", Product.Type).Ref("assoc_product_assocs").
			// Bind the "productIdTo" field to this edge.
			// Field("product_id_to").
			Unique().Annotations(entproto.Field(13)),
		// o2m
		edge.From("custom_method", CustomMethod.Type).Ref("product_assocs").
			// Bind the "estimateCalcMethod" field to this edge.
			// Field("estimate_calc_method").
			Unique().Annotations(entproto.Field(15)),
	}
}
