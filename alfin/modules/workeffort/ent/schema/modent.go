package schema
import (
    "entgo.io/ent"
    // "entgo.io/ent/schema/index"
	"entgo.io/ent/schema"
    // "entgo.io/ent/schema/mixin"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
	"entgo.io/contrib/entproto"
    "time"
)


type WorkEffortPartyAssignment struct {
    ent.Schema
}

// Fields of the WorkEffortPartyAssignment.
func (WorkEffortPartyAssignment) Fields() []ent.Field {
    return []ent.Field{
        field.Int("role_type_id").Annotations(entproto.Field(2)),
        field.Time("from_date").
            Default(time.Now).Annotations(entproto.Field(3)),
        field.Time("thru_date").
            Default(time.Now).Optional().Annotations(entproto.Field(4)),
        field.Int("status_id").Optional().Annotations(entproto.Field(5)),
        field.Time("status_date_time").
            Default(time.Now).Optional().Annotations(entproto.Field(6)),
        field.Int("expectation_enum_id").Optional().Annotations(entproto.Field(7)),
        field.Int("delegate_reason_enum_id").Optional().Annotations(entproto.Field(8)),
        field.Int("facility_id").Optional().Annotations(entproto.Field(9)),
        field.String("comments").Optional().Annotations(entproto.Field(10)),
        field.Enum("must_rsvp").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(11)),
        field.Int("availability_status_id").Optional().Annotations(entproto.Field(12)),
    }
}

/* entproto annotations ??
func (WorkEffortPartyAssignment) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
func (WorkEffortPartyAssignment) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("work_effort_id", "party_id", "role_type_id", "from_date").
            Unique(),
    }
}

*/

func (WorkEffortPartyAssignment) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the WorkEffortPartyAssignment.
//  one: WorkEffort
//  one-nofk: Party
//  one: PartyRole
//  one-nofk: RoleType
//  one: AssignedByUserLogin
//  one: AssignmentStatusItem
//  one: ExpectationEnumeration
//  one: DelegateReasonEnumeration
//  one: Facility
//  one: AvailabilityStatusItem
//  many: ApplicationSandbox
func (WorkEffortPartyAssignment) Edges() []ent.Edge {
    return []ent.Edge{
                // m2o
                edge.From("work_effort", WorkEffort.Type).Ref("work_effort_party_assignments").
                    // Bind the "workEffortId" field to this edge.
                    // Field("work_effort_id").
                    Unique().Annotations(entproto.Field(13)),
                // m2o
                edge.From("party", Party.Type).Ref("work_effort_party_assignments").
                    // Bind the "partyId" field to this edge.
                    // Field("party_id").
                    Unique().Annotations(entproto.Field(14)),
                // m2o
                edge.From("party_role", PartyRole.Type).Ref("work_effort_party_assignments").
                    // Bind the "partyId" field to this edge.
                    // Field("party_id").
                    Unique().Annotations(entproto.Field(15)),
                // m2o
                edge.From("assigned_by_user_login", UserLogin.Type).Ref("assigned_by_work_effort_party_assignments").
                    // Bind the "assignedByUserLoginId" field to this edge.
                    // Field("assigned_by_user_login_id").
                    Unique().Annotations(entproto.Field(17)),
    }
}


type FixedAsset struct {
    ent.Schema
}

// Fields of the FixedAsset.
func (FixedAsset) Fields() []ent.Field {
    return []ent.Field{
        field.Int("fixed_asset_type_id").Optional().Annotations(entproto.Field(2)),
        field.Int("instance_of_product_id").Optional().Annotations(entproto.Field(3)),
        field.Int("class_enum_id").Optional().Annotations(entproto.Field(4)),
        field.Int("role_type_id").Optional().Annotations(entproto.Field(5)),
        field.String("fixed_asset_name").Optional().Annotations(entproto.Field(6)),
        field.Int("acquire_order_id").Optional().Annotations(entproto.Field(7)),
        field.Int("acquire_order_item_seq_id").Optional().Annotations(entproto.Field(8)),
        field.Time("date_acquired").
            Default(time.Now).Optional().Annotations(entproto.Field(9)),
        field.Time("date_last_serviced").
            Default(time.Now).Optional().Annotations(entproto.Field(10)),
        field.Time("date_next_service").
            Default(time.Now).Optional().Annotations(entproto.Field(11)),
        field.Time("expected_end_of_life").
            Default(time.Now).Optional().Annotations(entproto.Field(12)),
        field.Time("actual_end_of_life").
            Default(time.Now).Optional().Annotations(entproto.Field(13)),
        field.Float("production_capacity").Optional().Annotations(entproto.Field(14)),
        field.Int("uom_id").Optional().Annotations(entproto.Field(15)),
        field.Int("calendar_id").Optional().Annotations(entproto.Field(16)),
        field.String("serial_number").Optional().Annotations(entproto.Field(17)),
        field.Int("located_at_facility_id").Optional().Annotations(entproto.Field(18)),
        field.Int("located_at_location_seq_id").Optional().Annotations(entproto.Field(19)),
        field.Float("salvage_value").Optional().Annotations(entproto.Field(20)),
        field.Float("depreciation").Optional().Annotations(entproto.Field(21)),
        field.Float("purchase_cost").Optional().Annotations(entproto.Field(22)),
        field.Int("purchase_cost_uom_id").Optional().Annotations(entproto.Field(23)),
    }
}

/* entproto annotations ??
func (FixedAsset) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
*/

func (FixedAsset) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the FixedAsset.
//  one: FixedAssetType
//  many: FixedAssetTypeAttr
//  one: ParentFixedAsset
//  one: InstanceOfProduct
//  one: ClassEnumeration
//  one: Party
//  one: RoleType
//  one-nofk: PartyRole
//  one: AcquireOrderHeader
//  one: AcquireOrderItem
//  one: Uom
//  one: TechDataCalendar
//  one: LocatedAtFacility
//  one-nofk: LocatedAtFacilityLocation
//  many: AccommodationMap
//  many: AccommodationSpot
//  many: AcctgTrans
//  many: CostComponent
//  many: Delivery
//  many: ChildFixedAsset
//  many: FixedAssetAttribute
//  many: FixedAssetDepMethod
//  many: FixedAssetGeoPoint
//  many: FixedAssetIdent
//  many: FixedAssetMaint
//  many: FixedAssetMaintOrder
//  many: FixedAssetProduct
//  many: FixedAssetRegistration
//  many: FixedAssetStdCost
//  many: FixedAssetTypeGlAccount
//  many: FixedAssetInventoryItem
//  many: PartyFixedAssetAssignment
//  many: Requirement
//  many: WorkEffort
//  many: WorkEffortFixedAssetAssign
func (FixedAsset) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", FixedAsset.Type).
                    From("parent").
                    // Bind the "parentFixedAssetId" field to this edge.
                    // Field("parent_fixed_asset_id").
                    Unique().Annotations(entproto.Field(26)),
                // m2o
                edge.From("party", Party.Type).Ref("fixed_assets").
                    // Bind the "partyId" field to this edge.
                    // Field("party_id").
                    Unique().Annotations(entproto.Field(29)),
                // m2o
                edge.From("party_role", PartyRole.Type).Ref("fixed_assets").
                    // Bind the "partyId" field to this edge.
                    // Field("party_id").
                    Unique().Annotations(entproto.Field(31)),
                // m2o
                edge.To("child_fixed_assets", FixedAsset.Type).
                    Annotations(entproto.Field(43)),
                // o2o
                // m2o
                edge.To("work_efforts", WorkEffort.Type).
                    Annotations(entproto.Field(57)),
                // o2o
                // m2o
                edge.To("work_effort_fixed_asset_assigns", WorkEffortFixedAssetAssign.Type).
                    Annotations(entproto.Field(58)),
                // o2o
    }
}


type Person struct {
    ent.Schema
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
    return []ent.Field{
        field.String("salutation").Optional().Annotations(entproto.Field(2)),
        field.String("first_name").Optional().Annotations(entproto.Field(3)),
        field.String("middle_name").Optional().Annotations(entproto.Field(4)),
        field.String("last_name").Optional().Annotations(entproto.Field(5)),
        field.String("personal_title").Optional().Annotations(entproto.Field(6)),
        field.String("suffix").Optional().Annotations(entproto.Field(7)),
        field.String("nickname").Optional().Annotations(entproto.Field(8)),
        field.String("first_name_local").Optional().Annotations(entproto.Field(9)),
        field.String("middle_name_local").Optional().Annotations(entproto.Field(10)),
        field.String("last_name_local").Optional().Annotations(entproto.Field(11)),
        field.String("other_local").Optional().Annotations(entproto.Field(12)),
        field.Int("member_id").Optional().Annotations(entproto.Field(13)),
        field.Enum("gender").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(14)),
        field.Time("birth_date").
            Default(time.Now).Optional().Annotations(entproto.Field(15)),
        field.Time("deceased_date").
            Default(time.Now).Optional().Annotations(entproto.Field(16)),
        field.Float("height").Optional().Annotations(entproto.Field(17)),
        field.Float("weight").Optional().Annotations(entproto.Field(18)),
        field.String("mothers_maiden_name").Optional().Annotations(entproto.Field(19)),
        field.Enum("old_marital_status").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(20)),
        field.Int("marital_status_enum_id").Optional().Annotations(entproto.Field(21)),
        field.String("social_security_number").Optional().Annotations(entproto.Field(22)),
        field.String("passport_number").Optional().Annotations(entproto.Field(23)),
        field.Time("passport_expire_date").
            Default(time.Now).Optional().Annotations(entproto.Field(24)),
        field.Float("total_years_work_experience").Optional().Annotations(entproto.Field(25)),
        field.String("comments").Optional().Annotations(entproto.Field(26)),
        field.Int("employment_status_enum_id").Optional().Annotations(entproto.Field(27)),
        field.Int("residence_status_enum_id").Optional().Annotations(entproto.Field(28)),
        field.String("occupation").Optional().Annotations(entproto.Field(29)),
        field.Int("years_with_employer").Optional().Annotations(entproto.Field(30)),
        field.Int("months_with_employer").Optional().Annotations(entproto.Field(31)),
        field.Enum("existing_customer").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(32)),
        field.String("card_id").MaxLen(32).Optional().Annotations(entproto.Field(33)),
    }
}

/* entproto annotations ??
func (Person) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
*/

func (Person) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the Person.
//  one: Party
//  one: EmploymentStatusEnumeration
//  one: ResidenceStatusEnumeration
//  one: MaritalStatusEnumeration
//  many: PartyContactMech
//  many: PartyContactMechPurpose
//  many: ApproverPersonTraining
//  many: ProdCatalogRole
//  many: ProductStoreRole
//  many: ToShipment
//  many: FromShipment
//  many: CarrierShipmentRouteSegment
//  many: UserLogin
//  many: WebSiteRole
func (Person) Edges() []ent.Edge {
    return []ent.Edge{
                // m2o
                    edge.From("party", Party.Type).
                        Ref("person").
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(34)),
                // o2m
                // m2o
                edge.To("user_logins", UserLogin.Type).
                    Annotations(entproto.Field(46)),
                // o2o
    }
}


type PartyRole struct {
    ent.Schema
}

// Fields of the PartyRole.
func (PartyRole) Fields() []ent.Field {
    return []ent.Field{
        field.Int("role_type_id").Annotations(entproto.Field(2)),
    }
}

/* entproto annotations ??
func (PartyRole) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
func (PartyRole) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("party_id", "role_type_id").
            Unique(),
    }
}

*/

func (PartyRole) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the PartyRole.
//  one: Party
//  one: RoleType
//  many: RoleTypeAttr
//  many: PartyAttribute
//  many: AcctgTrans
//  many: AcctgTransEntry
//  many: FromAgreement
//  many: ToAgreement
//  many: AgreementRole
//  many: BillingAccountRole
//  many: BudgetRole
//  many: CarrierShipmentMethod
//  many: ToCommunicationEvent
//  many: FromCommunicationEvent
//  many: CommunicationEventRole
//  many: ContentRole
//  many: CustRequestParty
//  many: DataResourceRole
//  many: ToEmployment
//  many: FromEmployment
//  many: FacilityGroupRole
//  many: FacilityParty
//  many: FinAccountRole
//  many: FixedAsset
//  many: GlAccountOrganization
//  many: GlAccountRole
//  many: Invoice
//  many: InvoiceRole
//  many: ItemIssuanceRole
//  many: MarketingCampaignRole
//  many: OrderItemRole
//  many: CarrierOrderItemShipGroup
//  many: OrderRole
//  many: ToPartyBenefit
//  many: FromPartyBenefit
//  many: PartyContactMech
//  many: PartyFixedAssetAssignment
//  many: PartyGlAccount
//  many: PartyNeed
//  many: FromPartyRelationship
//  many: ToPartyRelationship
//  many: ToPayment
//  many: PayrollPreference
//  many: EmployeePerfReview
//  many: ManagerPerfReview
//  many: EmployeePerfReviewItem
//  many: PerformanceNote
//  many: PicklistRole
//  many: ProdCatalogRole
//  many: ProductCategoryRole
//  many: ProductRole
//  many: ProductStoreGroupRole
//  many: ProductStoreRole
//  many: QuoteRole
//  many: RequirementRole
//  many: SalesOpportunityRole
//  many: SegmentGroupRole
//  many: ShipmentCostEstimate
//  many: ShipmentReceiptRole
//  many: Subscription
//  many: OriginatedFromSubscription
//  many: TimesheetRole
//  many: WebSiteRole
//  many: WorkEffortPartyAssignment
func (PartyRole) Edges() []ent.Edge {
    return []ent.Edge{
                // m2o
                edge.From("party", Party.Type).Ref("party_roles").
                    // Bind the "partyId" field to this edge.
                    // Field("party_id").
                    Unique().Annotations(entproto.Field(3)),
                // m2o
                edge.To("fixed_assets", FixedAsset.Type).
                    Annotations(entproto.Field(26)),
                // o2o
                // m2o
                edge.To("work_effort_party_assignments", WorkEffortPartyAssignment.Type).
                    Annotations(entproto.Field(66)),
                // o2o
    }
}


type UserLogin struct {
    ent.Schema
}

// Fields of the UserLogin.
func (UserLogin) Fields() []ent.Field {
    return []ent.Field{
        field.String("current_password").Optional().Annotations(entproto.Field(2)),
        field.String("password_hint").Optional().Annotations(entproto.Field(3)),
        field.Enum("is_system").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(4)),
        field.Enum("enabled").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(5)),
        field.Enum("has_logged_out").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(6)),
        field.Enum("require_password_change").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(7)),
        field.Int("last_currency_uom").Optional().Annotations(entproto.Field(8)),
        field.String("last_locale").MaxLen(10).Optional().Annotations(entproto.Field(9)),
        field.String("last_time_zone").MaxLen(32).Optional().Annotations(entproto.Field(10)),
        field.Time("disabled_date_time").
            Default(time.Now).Optional().Annotations(entproto.Field(11)),
        field.Int("successive_failed_logins").Optional().Annotations(entproto.Field(12)),
        field.String("external_auth_id").Optional().Annotations(entproto.Field(13)),
        field.String("user_ldap_dn").Optional().Annotations(entproto.Field(14)),
        field.String("disabled_by").Optional().Annotations(entproto.Field(15)),
    }
}

/* entproto annotations ??
func (UserLogin) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
*/

func (UserLogin) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the UserLogin.
//  one: Party
//  one-nofk: Person
//  one-nofk: PartyGroup
//  many: Created ByAllocationPlanHeader
//  many: Last Modified ByAllocationPlanHeader
//  many: Created By User LoginAllocationPlanItem
//  many: Last Modified By User LoginAllocationPlanItem
//  many: ChangeByBudgetStatus
//  many: CreatedByContactList
//  many: LastModifiedByContactList
//  many: ChangeByContactListCommStatus
//  many: CreatedByContent
//  many: LastModifiedByContent
//  many: CreatedByContentAssoc
//  many: LastModifiedByContentAssoc
//  many: ChangeByCustRequestStatus
//  many: CreatedByDataResource
//  many: LastModifiedByDataResource
//  many: ExampleStatus
//  many: UserLoginExcelImportHistory
//  many: FinAccountStatus
//  many: InventoryItemStatus
//  many: ChangeByInvoiceStatus
//  many: IssuedByItemIssuance
//  many: AuthJobSandbox
//  many: RunAsJobSandbox
//  many: ChangeOldPicklistStatusHistory
//  many: OrderAdjustment
//  many: CreatedByOrderHeader
//  many: DontCancelSetOrderItem
//  many: ChangeByOrderItem
//  many: OrderItemChange
//  many: OrderPaymentPreference
//  many: OrderStatus
//  many: CreatedByParty
//  many: LastModifiedByParty
//  many: ChangeByPartyStatus
//  many: CreatedByPicklistRole
//  many: LastModifiedByPicklistRole
//  many: ChangePicklistStatus
//  many: CreatedByProduct
//  many: LastModifiedByProduct
//  many: CreatedByProductFeaturePrice
//  many: LastModifiedByProductFeaturePrice
//  many: CreatedByProductPrice
//  many: LastModifiedByProductPrice
//  many: ChangedByProductPriceChange
//  many: CreatedByProductPromo
//  many: LastModifiedByProductPromo
//  many: CreatedByProductPromoCode
//  many: LastModifiedByProductPromoCode
//  many: ProductReview
//  many: QuoteAdjustment
//  many: ChangeByRequirementStatus
//  many: ReturnAdjustment
//  many: ReturnHeader
//  many: ChangeByReturnStatus
//  many: CreatedBySalesForecast
//  many: ModifiedBySalesForecast
//  many: ModifiedBySalesForecastHistory
//  many: SalesOpportunity
//  many: SalesOpportunityHistory
//  many: ShipmentReceipt
//  many: ChangeByShipmentStatus
//  many: ChangeByTestingStatus
//  many: ApprovedByTimesheet
//  many: UserLoginHistory
//  many: OriginUserLoginHistory
//  many: UserLoginPasswordHistory
//  many: UserLoginSecurityGroup
//  one-nofk: UserLoginSession
//  many: UserPreference
//  many: WebUserPreference
//  many: AssignedByWorkEffortPartyAssignment
//  many: WorkEffortReview
//  many: SetByWorkEffortStatus
//  many: WorkloadStatus
func (UserLogin) Edges() []ent.Edge {
    return []ent.Edge{
                // m2o
                edge.From("party", Party.Type).Ref("user_logins").
                    // Bind the "partyId" field to this edge.
                    // Field("party_id").
                    Unique().Annotations(entproto.Field(16)),
                // m2o
                edge.From("person", Person.Type).Ref("user_logins").
                    // Bind the "partyId" field to this edge.
                    // Field("party_id").
                    Unique().Annotations(entproto.Field(17)),
                // m2o
                edge.To("created_by_parties", Party.Type).
                    Annotations(entproto.Field(50)),
                // o2o
                // m2o
                edge.To("last_modified_by_parties", Party.Type).
                    Annotations(entproto.Field(51)),
                // o2o
                // m2o
                edge.To("change_by_party_statuses", PartyStatus.Type).
                    Annotations(entproto.Field(52)),
                // o2o
                // m2o
                edge.To("user_login_security_groups", UserLoginSecurityGroup.Type).
                    Annotations(entproto.Field(85)),
                // o2o
                // m2o
                edge.To("assigned_by_work_effort_party_assignments", WorkEffortPartyAssignment.Type).
                    Annotations(entproto.Field(89)),
                // o2o
    }
}


type WorkEffort struct {
    ent.Schema
}

// Fields of the WorkEffort.
func (WorkEffort) Fields() []ent.Field {
    return []ent.Field{
        field.Int("work_effort_type_id").Optional().Annotations(entproto.Field(2)),
        field.Int("current_status_id").Optional().Annotations(entproto.Field(3)),
        field.Time("last_status_update").
            Default(time.Now).Optional().Annotations(entproto.Field(4)),
        field.Int("work_effort_purpose_type_id").Optional().Annotations(entproto.Field(5)),
        field.Int("scope_enum_id").Optional().Annotations(entproto.Field(6)),
        field.Int("priority").Optional().Annotations(entproto.Field(7)),
        field.Int("percent_complete").Optional().Annotations(entproto.Field(8)),
        field.String("work_effort_name").Optional().Annotations(entproto.Field(9)),
        field.Int("show_as_enum_id").Optional().Annotations(entproto.Field(10)),
        field.Enum("send_notification_email").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(11)),
        field.String("description").Optional().Annotations(entproto.Field(12)),
        field.String("location_desc").Optional().Annotations(entproto.Field(13)),
        field.Time("estimated_start_date").
            Default(time.Now).Optional().Annotations(entproto.Field(14)),
        field.Time("estimated_completion_date").
            Default(time.Now).Optional().Annotations(entproto.Field(15)),
        field.Time("actual_start_date").
            Default(time.Now).Optional().Annotations(entproto.Field(16)),
        field.Time("actual_completion_date").
            Default(time.Now).Optional().Annotations(entproto.Field(17)),
        field.Float("estimated_milli_seconds").Optional().Annotations(entproto.Field(18)),
        field.Float("estimated_setup_millis").Optional().Annotations(entproto.Field(19)),
        field.Int("estimate_calc_method").Optional().Annotations(entproto.Field(20)),
        field.Float("actual_milli_seconds").Optional().Annotations(entproto.Field(21)),
        field.Float("actual_setup_millis").Optional().Annotations(entproto.Field(22)),
        field.Float("total_milli_seconds_allowed").Optional().Annotations(entproto.Field(23)),
        field.Float("total_money_allowed").Optional().Annotations(entproto.Field(24)),
        field.Int("money_uom_id").Optional().Annotations(entproto.Field(25)),
        field.String("special_terms").Optional().Annotations(entproto.Field(26)),
        field.Int("time_transparency").Optional().Annotations(entproto.Field(27)),
        field.String("universal_id").Optional().Annotations(entproto.Field(28)),
        field.String("source_reference_id").MaxLen(32).Optional().Annotations(entproto.Field(29)),
        field.Int("facility_id").Optional().Annotations(entproto.Field(30)),
        field.String("info_url").Optional().Annotations(entproto.Field(31)),
        field.Int("recurrence_info_id").Optional().Annotations(entproto.Field(32)),
        field.Int("runtime_data_id").Optional().Annotations(entproto.Field(33)),
        field.Int("note_id").Optional().Annotations(entproto.Field(34)),
        field.String("service_loader_name").Optional().Annotations(entproto.Field(35)),
        field.Float("quantity_to_produce").Optional().Annotations(entproto.Field(36)),
        field.Float("quantity_produced").Optional().Annotations(entproto.Field(37)),
        field.Float("quantity_rejected").Optional().Annotations(entproto.Field(38)),
        field.Float("reserv_persons").Optional().Annotations(entproto.Field(39)),
        field.Float("reserv_2_nd_pp_perc").Optional().Annotations(entproto.Field(40)),
        field.Float("reserv_nth_pp_perc").Optional().Annotations(entproto.Field(41)),
        field.Int("accommodation_map_id").Optional().Annotations(entproto.Field(42)),
        field.Int("accommodation_spot_id").Optional().Annotations(entproto.Field(43)),
        field.Int("revision_number").Optional().Annotations(entproto.Field(44)),
        field.Time("created_date").
            Default(time.Now).Optional().Annotations(entproto.Field(45)),
        field.String("created_by_user_login").Optional().Annotations(entproto.Field(46)),
        field.Time("last_modified_date").
            Default(time.Now).Optional().Annotations(entproto.Field(47)),
        field.String("last_modified_by_user_login").Optional().Annotations(entproto.Field(48)),
        field.Int("sequence_num").Optional().Annotations(entproto.Field(49)),
    }
}

/* entproto annotations ??
func (WorkEffort) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
*/

func (WorkEffort) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the WorkEffort.
//  one: WorkEffortType
//  one: WorkEffortPurposeType
//  one: ParentWorkEffort
//  many: WorkEffortTypeAttr
//  one: CurrentStatusItem
//  one: ScopeEnumeration
//  one: FixedAsset
//  one: Facility
//  one: MoneyUom
//  one: RecurrenceInfo
//  one: TemporalExpression
//  one: RuntimeData
//  one: NoteData
//  one: CustomMethod
//  one: AccommodationMap
//  one: AccommodationSpot
//  many: AcctgTrans
//  many: AgreementWorkEffortApplic
//  many: CommunicationEventWorkEff
//  many: CostComponent
//  many: CustRequestItemWorkEffort
//  many: CustRequestWorkEffort
//  many: ScheduleFixedAssetMaint
//  many: InventoryItemDetail
//  many: OrderHeaderWorkEffort
//  many: PersonTraining
//  many: RoutingProductAssoc
//  many: MaintTemplateProductMaint
//  many: QuoteItem
//  many: QuoteWorkEffort
//  many: RateAmount
//  many: SalesOpportunityWorkEffort
//  many: EstimatedShipShipment
//  many: EstimatedArrivalShipment
//  many: ShoppingListWorkEffort
//  many: TimeEntry
//  many: ChildWorkEffort
//  many: FromWorkEffortAssoc
//  many: ToWorkEffortAssoc
//  many: WorkEffortAttribute
//  many: WorkEffortBilling
//  many: WorkEffortContactMech
//  many: WorkEffortContent
//  many: WorkEffortCostCalc
//  many: WorkEffortDeliverableProd
//  many: WorkEffortEventReminder
//  many: WorkEffortFixedAssetAssign
//  many: WorkEffortFixedAssetStd
//  many: WorkEffortGoodStandard
//  one-nofk: WorkEffortIcalData
//  many: WorkEffortInventoryAssign
//  many: WorkEffortInventoryProduced
//  many: WorkEffortKeyword
//  many: WorkEffortNote
//  many: WorkEffortPartyAssignment
//  many: WorkEffortReview
//  many: WorkEffortSkillStandard
//  many: WorkEffortStatus
//  many: WorkEffortSurveyAppl
//  many: WorkEffortTransBox
//  many: WorkOrderItemFulfillment
//  many: WorkRequirementFulfillment
func (WorkEffort) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", WorkEffort.Type).
                    From("parent").
                    // Bind the "workEffortParentId" field to this edge.
                    // Field("work_effort_parent_id").
                    Unique().Annotations(entproto.Field(52)),
                // m2o
                edge.From("fixed_asset", FixedAsset.Type).Ref("work_efforts").
                    // Bind the "fixedAssetId" field to this edge.
                    // Field("fixed_asset_id").
                    Unique().Annotations(entproto.Field(56)),
                // m2o
                edge.From("temporal_expression", TemporalExpression.Type).Ref("work_efforts").
                    // Bind the "tempExprId" field to this edge.
                    // Field("temp_expr_id").
                    Unique().Annotations(entproto.Field(60)),
                // m2o
                edge.To("child_work_efforts", WorkEffort.Type).
                    Annotations(entproto.Field(86)),
                // o2o
                // m2o
                edge.To("from_work_effort_assocs", WorkEffortAssoc.Type).
                    Annotations(entproto.Field(87)),
                // o2o
                // m2o
                edge.To("to_work_effort_assocs", WorkEffortAssoc.Type).
                    Annotations(entproto.Field(88)),
                // o2o
                // m2o
                edge.To("work_effort_fixed_asset_assigns", WorkEffortFixedAssetAssign.Type).
                    Annotations(entproto.Field(96)),
                // o2o
                // m2o
                edge.To("work_effort_party_assignments", WorkEffortPartyAssignment.Type).
                    Annotations(entproto.Field(104)),
                // o2o
    }
}


type WorkEffortFixedAssetAssign struct {
    ent.Schema
}

// Fields of the WorkEffortFixedAssetAssign.
func (WorkEffortFixedAssetAssign) Fields() []ent.Field {
    return []ent.Field{
        field.Int("status_id").Optional().Annotations(entproto.Field(2)),
        field.Time("from_date").
            Default(time.Now).Annotations(entproto.Field(3)),
        field.Time("thru_date").
            Default(time.Now).Optional().Annotations(entproto.Field(4)),
        field.Int("availability_status_id").Optional().Annotations(entproto.Field(5)),
        field.Float("allocated_cost").Optional().Annotations(entproto.Field(6)),
        field.String("comments").Optional().Annotations(entproto.Field(7)),
    }
}

/* entproto annotations ??
func (WorkEffortFixedAssetAssign) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
func (WorkEffortFixedAssetAssign) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("work_effort_id", "fixed_asset_id", "from_date").
            Unique(),
    }
}

*/

func (WorkEffortFixedAssetAssign) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the WorkEffortFixedAssetAssign.
//  one: WorkEffort
//  one: FixedAsset
//  one: StatusItem
//  one: AvailabilityStatusItem
func (WorkEffortFixedAssetAssign) Edges() []ent.Edge {
    return []ent.Edge{
                // m2o
                edge.From("work_effort", WorkEffort.Type).Ref("work_effort_fixed_asset_assigns").
                    // Bind the "workEffortId" field to this edge.
                    // Field("work_effort_id").
                    Unique().Annotations(entproto.Field(8)),
                // m2o
                edge.From("fixed_asset", FixedAsset.Type).Ref("work_effort_fixed_asset_assigns").
                    // Bind the "fixedAssetId" field to this edge.
                    // Field("fixed_asset_id").
                    Unique().Annotations(entproto.Field(9)),
    }
}


type SecurityGroupPermission struct {
    ent.Schema
}

// Fields of the SecurityGroupPermission.
func (SecurityGroupPermission) Fields() []ent.Field {
    return []ent.Field{
        field.String("permission_id").MaxLen(32).Annotations(entproto.Field(2)),
        field.Time("from_date").
            Default(time.Now).Annotations(entproto.Field(3)),
        field.Time("thru_date").
            Default(time.Now).Optional().Annotations(entproto.Field(4)),
    }
}

/* entproto annotations ??
func (SecurityGroupPermission) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
func (SecurityGroupPermission) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("group_id", "permission_id", "from_date").
            Unique(),
    }
}

*/

func (SecurityGroupPermission) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the SecurityGroupPermission.
//  one: SecurityGroup
//  one-nofk: SecurityPermission
func (SecurityGroupPermission) Edges() []ent.Edge {
    return []ent.Edge{
                // m2o
                edge.From("security_group", SecurityGroup.Type).Ref("security_group_permissions").
                    // Bind the "groupId" field to this edge.
                    // Field("group_id").
                    Unique().Annotations(entproto.Field(5)),
    }
}


type TemporalExpressionAssoc struct {
    ent.Schema
}

// Fields of the TemporalExpressionAssoc.
func (TemporalExpressionAssoc) Fields() []ent.Field {
    return []ent.Field{
        field.Int("expr_assoc_type").Optional().Annotations(entproto.Field(2)),
    }
}

/* entproto annotations ??
func (TemporalExpressionAssoc) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
func (TemporalExpressionAssoc) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("from_temp_expr_id", "to_temp_expr_id").
            Unique(),
    }
}

*/

func (TemporalExpressionAssoc) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the TemporalExpressionAssoc.
//  one: FromTemporalExpression
//  one: ToTemporalExpression
func (TemporalExpressionAssoc) Edges() []ent.Edge {
    return []ent.Edge{
                // m2o
                edge.From("from_temporal_expression", TemporalExpression.Type).Ref("from_temporal_expression_assocs").
                    // Bind the "fromTempExprId" field to this edge.
                    // Field("from_temp_expr_id").
                    Unique().Annotations(entproto.Field(3)),
                // m2o
                edge.From("to_temporal_expression", TemporalExpression.Type).Ref("to_temporal_expression_assocs").
                    // Bind the "toTempExprId" field to this edge.
                    // Field("to_temp_expr_id").
                    Unique().Annotations(entproto.Field(4)),
    }
}


type UserLoginSecurityGroup struct {
    ent.Schema
}

// Fields of the UserLoginSecurityGroup.
func (UserLoginSecurityGroup) Fields() []ent.Field {
    return []ent.Field{
        field.Time("from_date").
            Default(time.Now).Annotations(entproto.Field(2)),
        field.Time("thru_date").
            Default(time.Now).Optional().Annotations(entproto.Field(3)),
    }
}

/* entproto annotations ??
func (UserLoginSecurityGroup) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
func (UserLoginSecurityGroup) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("user_login_id", "group_id", "from_date").
            Unique(),
    }
}

*/

func (UserLoginSecurityGroup) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the UserLoginSecurityGroup.
//  one: UserLogin
//  one: SecurityGroup
//  many: SecurityGroupPermission
func (UserLoginSecurityGroup) Edges() []ent.Edge {
    return []ent.Edge{
                // m2o
                edge.From("user_login", UserLogin.Type).Ref("user_login_security_groups").
                    // Bind the "userLoginId" field to this edge.
                    // Field("user_login_id").
                    Unique().Annotations(entproto.Field(4)),
                // m2o
                edge.From("security_group", SecurityGroup.Type).Ref("user_login_security_groups").
                    // Bind the "groupId" field to this edge.
                    // Field("group_id").
                    Unique().Annotations(entproto.Field(5)),
                // m2o
                edge.To("security_group_permissions", SecurityGroupPermission.Type).
                    Annotations(entproto.Field(6)),
                // o2o
    }
}


type SecurityGroup struct {
    ent.Schema
}

// Fields of the SecurityGroup.
func (SecurityGroup) Fields() []ent.Field {
    return []ent.Field{
        field.String("group_name").Optional().Annotations(entproto.Field(2)),
        field.String("description").Optional().Annotations(entproto.Field(3)),
    }
}

/* entproto annotations ??
func (SecurityGroup) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
*/

func (SecurityGroup) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the SecurityGroup.
//  many: PartyRelationship
//  many: PortalPage
//  many: ProtectedView
//  many: SecurityGroupPermission
//  many: UserLoginSecurityGroup
func (SecurityGroup) Edges() []ent.Edge {
    return []ent.Edge{
                // m2o
                edge.To("security_group_permissions", SecurityGroupPermission.Type).
                    Annotations(entproto.Field(7)),
                // o2o
                // m2o
                edge.To("user_login_security_groups", UserLoginSecurityGroup.Type).
                    Annotations(entproto.Field(8)),
                // o2o
    }
}


type WorkEffortAssoc struct {
    ent.Schema
}

// Fields of the WorkEffortAssoc.
func (WorkEffortAssoc) Fields() []ent.Field {
    return []ent.Field{
        field.Int("work_effort_assoc_type_id").Annotations(entproto.Field(2)),
        field.Int("sequence_num").Optional().Annotations(entproto.Field(3)),
        field.Time("from_date").
            Default(time.Now).Annotations(entproto.Field(4)),
        field.Time("thru_date").
            Default(time.Now).Optional().Annotations(entproto.Field(5)),
    }
}

/* entproto annotations ??
func (WorkEffortAssoc) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
func (WorkEffortAssoc) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("work_effort_id_from", "work_effort_id_to", "work_effort_assoc_type_id", "from_date").
            Unique(),
    }
}

*/

func (WorkEffortAssoc) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the WorkEffortAssoc.
//  one: WorkEffortAssocType
//  many: WorkEffortAssocTypeAttr
//  one: FromWorkEffort
//  one: ToWorkEffort
//  many: WorkEffortAssocAttribute
func (WorkEffortAssoc) Edges() []ent.Edge {
    return []ent.Edge{
                // m2o
                edge.From("from_work_effort", WorkEffort.Type).Ref("from_work_effort_assocs").
                    // Bind the "workEffortIdFrom" field to this edge.
                    // Field("work_effort_id_from").
                    Unique().Annotations(entproto.Field(8)),
                // m2o
                edge.From("to_work_effort", WorkEffort.Type).Ref("to_work_effort_assocs").
                    // Bind the "workEffortIdTo" field to this edge.
                    // Field("work_effort_id_to").
                    Unique().Annotations(entproto.Field(9)),
    }
}


type TemporalExpression struct {
    ent.Schema
}

// Fields of the TemporalExpression.
func (TemporalExpression) Fields() []ent.Field {
    return []ent.Field{
        field.Int("temp_expr_type_id").Optional().Annotations(entproto.Field(2)),
        field.String("description").Optional().Annotations(entproto.Field(3)),
        field.Time("date_1").
            Default(time.Now).Optional().Annotations(entproto.Field(4)),
        field.Time("date_2").
            Default(time.Now).Optional().Annotations(entproto.Field(5)),
        field.Int("integer_1").Optional().Annotations(entproto.Field(6)),
        field.Int("integer_2").Optional().Annotations(entproto.Field(7)),
        field.Int("string_1").Optional().Annotations(entproto.Field(8)),
        field.Int("string_2").Optional().Annotations(entproto.Field(9)),
    }
}

/* entproto annotations ??
func (TemporalExpression) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
*/

func (TemporalExpression) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the TemporalExpression.
//  many: JobSandbox
//  many: FromTemporalExpressionAssoc
//  many: ToTemporalExpressionAssoc
//  many: WorkEffort
func (TemporalExpression) Edges() []ent.Edge {
    return []ent.Edge{
                // m2o
                edge.To("from_temporal_expression_assocs", TemporalExpressionAssoc.Type).
                    Annotations(entproto.Field(11)),
                // o2o
                // m2o
                edge.To("to_temporal_expression_assocs", TemporalExpressionAssoc.Type).
                    Annotations(entproto.Field(12)),
                // o2o
                // m2o
                edge.To("work_efforts", WorkEffort.Type).
                    Annotations(entproto.Field(13)),
                // o2o
    }
}


type PartyStatus struct {
    ent.Schema
}

// Fields of the PartyStatus.
func (PartyStatus) Fields() []ent.Field {
    return []ent.Field{
        field.Int("status_id").Annotations(entproto.Field(2)),
        field.Time("status_date").
            Default(time.Now).Annotations(entproto.Field(3)),
    }
}

/* entproto annotations ??
func (PartyStatus) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
func (PartyStatus) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("status_id", "party_id", "status_date").
            Unique(),
    }
}

*/

func (PartyStatus) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the PartyStatus.
//  one: StatusItem
//  one: Party
//  one: ChangeByUserLogin
func (PartyStatus) Edges() []ent.Edge {
    return []ent.Edge{
                // m2o
                edge.From("party", Party.Type).Ref("party_statuses").
                    // Bind the "partyId" field to this edge.
                    // Field("party_id").
                    Unique().Annotations(entproto.Field(5)),
                // m2o
                edge.From("change_by_user_login", UserLogin.Type).Ref("change_by_party_statuses").
                    // Bind the "changeByUserLoginId" field to this edge.
                    // Field("change_by_user_login_id").
                    Unique().Annotations(entproto.Field(6)),
    }
}


type Party struct {
    ent.Schema
}

// Fields of the Party.
func (Party) Fields() []ent.Field {
    return []ent.Field{
        field.Int("party_type_id").Optional().Annotations(entproto.Field(2)),
        field.Int("external_id").Optional().Annotations(entproto.Field(3)),
        field.Int("preferred_currency_uom_id").Optional().Annotations(entproto.Field(4)),
        field.String("description").Optional().Annotations(entproto.Field(5)),
        field.Int("status_id").Optional().Annotations(entproto.Field(6)),
        field.Time("created_date").
            Default(time.Now).Optional().Annotations(entproto.Field(7)),
        field.Time("last_modified_date").
            Default(time.Now).Optional().Annotations(entproto.Field(8)),
        field.Int("data_source_id").Optional().Annotations(entproto.Field(9)),
        field.Enum("is_unread").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(10)),
    }
}

/* entproto annotations ??
func (Party) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
*/

/*
*/

func (Party) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the Party.
//  one: PartyType
//  one: CreatedByUserLogin
//  one: LastModifiedByUserLogin
//  one: Uom
//  one: StatusItem
//  many: PartyTypeAttr
//  one: DataSource
//  many: AcctgTrans
//  many: AcctgTransEntry
//  one-nofk: Affiliate
//  many: FromAgreement
//  many: ToAgreement
//  many: AgreementPartyApplic
//  many: AgreementRole
//  many: BillingAccountRole
//  many: BudgetReview
//  many: BudgetRole
//  many: CarrierShipmentBoxType
//  many: CarrierShipmentMethod
//  many: ToCommunicationEvent
//  many: FromCommunicationEvent
//  many: CommunicationEventRole
//  many: OwnerContactList
//  many: ContactListCommStatus
//  many: ContactListParty
//  many: ContentApproval
//  many: CommittedByContentRevision
//  many: ContentRole
//  many: CostComponent
//  many: FromCustRequest
//  many: CustRequestParty
//  many: CustRequestType
//  many: OrganizationCustomTimePeriod
//  many: DataResourceRole
//  many: EbayShippingMethod
//  many: EmplLeave
//  many: ApproverEmplLeave
//  many: EmplPosition
//  many: EmplPositionFulfillment
//  many: ToEmployment
//  many: FromEmployment
//  many: ApplyingEmploymentApp
//  many: ReferredByEmploymentApp
//  many: ApproverEmploymentApp
//  many: OwnerFacility
//  many: FacilityCarrierShipment
//  many: FacilityGroupRole
//  many: FacilityParty
//  many: OrganizationFinAccount
//  many: OwnerFinAccount
//  many: FinAccountRole
//  many: FinAccountTrans
//  many: PerformedByFinAccountTrans
//  many: OrganizationFinAccountTypeGlAccount
//  many: FixedAsset
//  many: GovAgencyFixedAssetRegistration
//  many: FixedAssetTypeGlAccount
//  many: GiftCardFulfillment
//  many: GlAccountHistory
//  many: GlAccountOrganization
//  many: GlAccountRole
//  many: OrganizationGlAccountTypeDefault
//  many: GlJournal
//  many: GlReconciliation
//  many: InventoryItem
//  many: OwnerInventoryItem
//  many: FromInvoice
//  many: Invoice
//  many: TaxAuthorityInvoiceItem
//  many: OverrideOrgInvoiceItem
//  many: FromInvoiceItemAssoc
//  many: ToInvoiceItemAssoc
//  many: OrganizationInvoiceItemTypeGlAccount
//  many: InvoiceRole
//  many: ItemIssuanceRole
//  many: IntervieweeJobInterview
//  many: InterviewerJobInterview
//  many: MarketingCampaignRole
//  many: NoteNoteData
//  many: OrderItemRole
//  many: SupplierOrderItemShipGroup
//  many: VendorOrderItemShipGroup
//  many: CarrierOrderItemShipGroup
//  many: OrderRole
//  one-nofk: PartyAcctgPreference
//  many: PartyAttribute
//  many: ToPartyBenefit
//  many: FromPartyBenefit
//  many: PartyCarrierAccount
//  many: CarrierPartyCarrierAccount
//  many: PartyClassification
//  many: PartyContactMech
//  many: PartyContactMechPurpose
//  many: PartyContent
//  many: PartyDataSource
//  many: PartyFixedAssetAssignment
//  many: PartyGeoPoint
//  many: OrganizationPartyGlAccount
//  many: PartyGlAccount
//  one-nofk: PartyGroup
//  one-nofk: PartyIcsAvsOverride
//  many: PartyIdentification
//  many: PartyInvitation
//  many: ToPartyInvitationGroupAssoc
//  many: PartyNameHistory
//  many: PartyNeed
//  many: PartyNote
//  many: PartyPrefDocTypeTpl
//  many: PartyProfileDefault
//  many: PartyQual
//  many: PartyRate
//  many: FromPartyRelationship
//  many: ToPartyRelationship
//  many: PartyResume
//  many: PartyRole
//  many: PartySkill
//  many: PartyStatus
//  many: PartyTaxAuthInfo
//  many: FromPayment
//  many: ToPayment
//  many: PaymentGlAccountTypeMap
//  many: PaymentMethod
//  many: OrganizationPaymentMethodTypeGlAccount
//  many: PayrollPreference
//  many: EmployeePerfReview
//  many: ManagerPerfReview
//  many: EmployeePerfReviewItem
//  many: PerformanceNote
//  one-nofk: Person
//  many: PersonTraining
//  many: ProdCatalogRole
//  many: ProductAverageCost
//  many: ProductCategoryGlAccount
//  many: ProductCategoryRole
//  many: ProductGlAccount
//  many: TaxAuthorityProductPrice
//  many: ProductPromo
//  many: ProductPromoCodeParty
//  many: ProductPromoUse
//  many: ProductRole
//  many: ProductStore
//  many: ProductStoreGroupRole
//  many: ProductStoreRole
//  many: ProductStoreShipmentMeth
//  many: VendorProductStoreVendorPayment
//  many: VendorProductStoreVendorShipment
//  many: CarrierProductStoreVendorShipment
//  many: Quote
//  many: QuoteRole
//  many: RateAmount
//  many: ReorderGuideline
//  many: RequirementRole
//  many: RespondingParty
//  many: ReturnHeader
//  many: ToReturnHeader
//  many: OrganizationSalesForecast
//  many: InternalSalesForecast
//  many: OrganizationSalesForecastHistory
//  many: InternalSalesForecastHistory
//  many: SalesOpportunityRole
//  many: SegmentGroupRole
//  many: ToShipment
//  many: FromShipment
//  many: ShipmentCostEstimate
//  many: ShipmentReceiptRole
//  many: CarrierShipmentRouteSegment
//  many: ShoppingList
//  many: Subscription
//  many: OriginatedFromSubscription
//  many: SupplierProduct
//  many: SupplierProductFeature
//  many: SurveyResponse
//  many: TaxAuthTaxAuthority
//  many: OrganizationTaxAuthorityGlAccount
//  many: TimeEntry
//  many: Timesheet
//  many: ClientTimesheet
//  many: TimesheetRole
//  many: UserLogin
//  many: UserLoginHistory
//  many: OrganizationVarianceReasonGlAccount
//  one-nofk: Vendor
//  many: VendorVendorProduct
//  many: WebSiteRole
//  many: WebUserPreference
//  many: WorkEffortEventReminder
//  many: WorkEffortPartyAssignment
func (Party) Edges() []ent.Edge {
    return []ent.Edge{
                // m2o
                edge.From("created_by_user_login", UserLogin.Type).Ref("created_by_parties").
                    // Bind the "createdByUserLogin" field to this edge.
                    // Field("created_by_user_login").
                    Unique().Annotations(entproto.Field(12)),
                // m2o
                edge.From("last_modified_by_user_login", UserLogin.Type).Ref("last_modified_by_parties").
                    // Bind the "lastModifiedByUserLogin" field to this edge.
                    // Field("last_modified_by_user_login").
                    Unique().Annotations(entproto.Field(13)),
                // m2o
                edge.To("fixed_assets", FixedAsset.Type).
                    Annotations(entproto.Field(65)),
                // o2o
                // m2o
                edge.To("party_roles", PartyRole.Type).
                    Annotations(entproto.Field(125)),
                // o2o
                // m2o
                edge.To("party_statuses", PartyStatus.Type).
                    Annotations(entproto.Field(127)),
                // o2o
                // m2o
                    edge.To("person", Person.Type).
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(139)),
                // m2o
                edge.To("user_logins", UserLogin.Type).
                    Annotations(entproto.Field(189)),
                // o2o
                // m2o
                edge.To("work_effort_party_assignments", WorkEffortPartyAssignment.Type).
                    Annotations(entproto.Field(197)),
                // o2o
    }
}
