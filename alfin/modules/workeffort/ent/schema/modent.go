package schema
import (
    "entgo.io/ent"
    // "entgo.io/ent/schema/index"
	"entgo.io/ent/schema"
    "entgo.io/ent/schema/mixin"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
	"entgo.io/contrib/entproto"
	"github.com/samlet/petrel/alfin/schemamixins"
    "time"
)


type TemporalExpression struct {
    ent.Schema
}

// Fields of the TemporalExpression.
// Unique-Indexes: tempExprId
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

//* entproto annotations ??
func (TemporalExpression) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // m2o
                    edge.To("to_temporal_expression_assocs", TemporalExpressionAssoc.Type).
                        Annotations(entproto.Field(12)),
                    // m2o
                    edge.To("work_efforts", WorkEffort.Type).
                        Annotations(entproto.Field(13)),
    }
}


type PartyContentType struct {
    ent.Schema
}

// Fields of the PartyContentType.
// Unique-Indexes: partyContentTypeId
func (PartyContentType) Fields() []ent.Field {
    return []ent.Field{
        field.String("description").Optional().Annotations(entproto.Field(2)),
    }
}

//* entproto annotations ??
func (PartyContentType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (PartyContentType) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the PartyContentType.
//  one: ParentPartyContentType
//  many: PartyContent
//  many: ChildPartyContentType
func (PartyContentType) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", PartyContentType.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique().Annotations(entproto.Field(3)),
                    // m2o
                    edge.To("child_party_content_types", PartyContentType.Type).
                        Annotations(entproto.Field(5)),
    }
}


type UserPreference struct {
    ent.Schema
}

// Fields of the UserPreference.
// Unique-Indexes: userLoginId, userPrefTypeId
func (UserPreference) Fields() []ent.Field {
    return []ent.Field{
        field.String("user_pref_type_id").MaxLen(32).Annotations(entproto.Field(2)),
        field.String("user_pref_group_type_id").MaxLen(32).Optional().Annotations(entproto.Field(3)),
        field.String("user_pref_value").Optional().Annotations(entproto.Field(4)),
        field.String("user_pref_data_type").MaxLen(32).Optional().Annotations(entproto.Field(5)),
    }
}

//* entproto annotations ??
func (UserPreference) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
func (UserPreference) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("user_login_id", "user_pref_type_id").
            Unique(),
    }
}

*/

func (UserPreference) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the UserPreference.
//  one-nofk: UserLogin
//  one: UserPrefGroupType
func (UserPreference) Edges() []ent.Edge {
    return []ent.Edge{
                    // o2m
                    edge.From("user_login", UserLogin.Type).Ref("user_preferences").
                        // Bind the "userLoginId" field to this edge.
                        // Field("user_login_id").
                        Unique().Annotations(entproto.Field(6)),
    }
}


type PartyClassificationType struct {
    ent.Schema
}

// Fields of the PartyClassificationType.
// Unique-Indexes: partyClassificationTypeId
func (PartyClassificationType) Fields() []ent.Field {
    return []ent.Field{
        field.Enum("has_table").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
        field.String("description").Optional().Annotations(entproto.Field(3)),
    }
}

//* entproto annotations ??
func (PartyClassificationType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (PartyClassificationType) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the PartyClassificationType.
//  one: ParentPartyClassificationType
//  many: PartyClassificationGroup
//  many: ChildPartyClassificationType
func (PartyClassificationType) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", PartyClassificationType.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique().Annotations(entproto.Field(4)),
                    // m2o
                    edge.To("child_party_classification_types", PartyClassificationType.Type).
                        Annotations(entproto.Field(6)),
    }
}


type WorkEffortAssoc struct {
    ent.Schema
}

// Fields of the WorkEffortAssoc.
// Unique-Indexes: workEffortIdFrom, workEffortIdTo, workEffortAssocTypeId, fromDate
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

//* entproto annotations ??
func (WorkEffortAssoc) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // o2m
                    edge.From("from_work_effort", WorkEffort.Type).Ref("from_work_effort_assocs").
                        // Bind the "workEffortIdFrom" field to this edge.
                        // Field("work_effort_id_from").
                        Unique().Annotations(entproto.Field(8)),
                    // o2m
                    edge.From("to_work_effort", WorkEffort.Type).Ref("to_work_effort_assocs").
                        // Bind the "workEffortIdTo" field to this edge.
                        // Field("work_effort_id_to").
                        Unique().Annotations(entproto.Field(9)),
    }
}


type ContactMechTypePurpose struct {
    ent.Schema
}

// Fields of the ContactMechTypePurpose.
// Unique-Indexes: contactMechTypeId, contactMechPurposeTypeId
func (ContactMechTypePurpose) Fields() []ent.Field {
    return []ent.Field{
    }
}

//* entproto annotations ??
func (ContactMechTypePurpose) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
func (ContactMechTypePurpose) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("contact_mech_type_id", "contact_mech_purpose_type_id").
            Unique(),
    }
}

*/

func (ContactMechTypePurpose) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the ContactMechTypePurpose.
//  one: ContactMechType
//  one: ContactMechPurposeType
func (ContactMechTypePurpose) Edges() []ent.Edge {
    return []ent.Edge{
                    // o2m
                    edge.From("contact_mech_type", ContactMechType.Type).Ref("contact_mech_type_purposes").
                        // Bind the "contactMechTypeId" field to this edge.
                        // Field("contact_mech_type_id").
                        Unique().Annotations(entproto.Field(0)),
                    // o2m
                    edge.From("contact_mech_purpose_type", ContactMechPurposeType.Type).Ref("contact_mech_type_purposes").
                        // Bind the "contactMechPurposeTypeId" field to this edge.
                        // Field("contact_mech_purpose_type_id").
                        Unique().Annotations(entproto.Field(1)),
    }
}


type SecurityGroupPermission struct {
    ent.Schema
}

// Fields of the SecurityGroupPermission.
// Unique-Indexes: groupId, permissionId, fromDate
func (SecurityGroupPermission) Fields() []ent.Field {
    return []ent.Field{
        field.Time("from_date").
            Default(time.Now).Annotations(entproto.Field(2)),
        field.Time("thru_date").
            Default(time.Now).Optional().Annotations(entproto.Field(3)),
    }
}

//* entproto annotations ??
func (SecurityGroupPermission) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // o2m
                    edge.From("security_group", SecurityGroup.Type).Ref("security_group_permissions").
                        // Bind the "groupId" field to this edge.
                        // Field("group_id").
                        Unique().Annotations(entproto.Field(4)),
                    // o2m
                    edge.From("security_permission", SecurityPermission.Type).Ref("security_group_permissions").
                        // Bind the "permissionId" field to this edge.
                        // Field("permission_id").
                        Unique().Annotations(entproto.Field(5)),
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
                    edge.To("parties", Party.Type).
                        Annotations(entproto.Field(47)),
                    // m2o
                    edge.To("party_statuses", PartyStatus.Type).
                        Annotations(entproto.Field(53)),
                    // m2o
                    edge.To("main_status_valid_changes", StatusValidChange.Type).
                        Annotations(entproto.Field(80)),
                    // m2o
                    edge.To("to_status_valid_changes", StatusValidChange.Type).
                        Annotations(entproto.Field(81)),
                    // m2o
                    edge.To("current_work_efforts", WorkEffort.Type).
                        Annotations(entproto.Field(86)),
                    // m2o
                    edge.To("work_effort_fixed_asset_assigns", WorkEffortFixedAssetAssign.Type).
                        Annotations(entproto.Field(87)),
                    // m2o
                    edge.To("availability_work_effort_fixed_asset_assigns", WorkEffortFixedAssetAssign.Type).
                        Annotations(entproto.Field(88)),
                    // m2o
                    edge.To("assignment_work_effort_party_assignments", WorkEffortPartyAssignment.Type).
                        Annotations(entproto.Field(91)),
                    // m2o
                    edge.To("availability_work_effort_party_assignments", WorkEffortPartyAssignment.Type).
                        Annotations(entproto.Field(92)),
    }
}


type TemporalExpressionAssoc struct {
    ent.Schema
}

// Fields of the TemporalExpressionAssoc.
// Unique-Indexes: fromTempExprId, toTempExprId
func (TemporalExpressionAssoc) Fields() []ent.Field {
    return []ent.Field{
        field.Int("expr_assoc_type").Optional().Annotations(entproto.Field(2)),
    }
}

//* entproto annotations ??
func (TemporalExpressionAssoc) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // o2m
                    edge.From("from_temporal_expression", TemporalExpression.Type).Ref("from_temporal_expression_assocs").
                        // Bind the "fromTempExprId" field to this edge.
                        // Field("from_temp_expr_id").
                        Unique().Annotations(entproto.Field(3)),
                    // o2m
                    edge.From("to_temporal_expression", TemporalExpression.Type).Ref("to_temporal_expression_assocs").
                        // Bind the "toTempExprId" field to this edge.
                        // Field("to_temp_expr_id").
                        Unique().Annotations(entproto.Field(4)),
    }
}


type PartyContactMech struct {
    ent.Schema
}

// Fields of the PartyContactMech.
// Unique-Indexes: partyId, contactMechId, fromDate
func (PartyContactMech) Fields() []ent.Field {
    return []ent.Field{
        field.Int("contact_mech_id").Annotations(entproto.Field(2)),
        field.Time("from_date").
            Default(time.Now).Annotations(entproto.Field(3)),
        field.Time("thru_date").
            Default(time.Now).Optional().Annotations(entproto.Field(4)),
        field.Enum("allow_solicitation").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(5)),
        field.String("extension").Optional().Annotations(entproto.Field(6)),
        field.Enum("verified").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(7)),
        field.String("comments").Optional().Annotations(entproto.Field(8)),
        field.Int("years_with_contact_mech").Optional().Annotations(entproto.Field(9)),
        field.Int("months_with_contact_mech").Optional().Annotations(entproto.Field(10)),
    }
}

//* entproto annotations ??
func (PartyContactMech) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
func (PartyContactMech) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("party_id", "contact_mech_id", "from_date").
            Unique(),
    }
}

*/

func (PartyContactMech) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the PartyContactMech.
//  one: Party
//  one-nofk: Person
//  one-nofk: PartyGroup
//  one: PartyRole
//  one: RoleType
//  one: ContactMech
//  one-nofk: TelecomNumber
//  one-nofk: PostalAddress
//  many: PartyContactMechPurpose
func (PartyContactMech) Edges() []ent.Edge {
    return []ent.Edge{
                    // o2m
                    edge.From("party", Party.Type).Ref("party_contact_meches").
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(11)),
                    // o2m
                    edge.From("person", Person.Type).Ref("party_contact_meches").
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(12)),
                    // o2m
                    edge.From("party_role", PartyRole.Type).Ref("party_contact_meches").
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(14)),
                    // o2m
                    edge.From("role_type", RoleType.Type).Ref("party_contact_meches").
                        // Bind the "roleTypeId" field to this edge.
                        // Field("role_type_id").
                        Unique().Annotations(entproto.Field(15)),
    }
}


type WorkEffortType struct {
    ent.Schema
}

// Fields of the WorkEffortType.
// Unique-Indexes: workEffortTypeId
func (WorkEffortType) Fields() []ent.Field {
    return []ent.Field{
        field.Enum("has_table").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
        field.String("description").Optional().Annotations(entproto.Field(3)),
    }
}

//* entproto annotations ??
func (WorkEffortType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (WorkEffortType) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the WorkEffortType.
//  one: ParentWorkEffortType
//  many: WorkEffort
//  many: ChildWorkEffortType
//  many: WorkEffortTypeAttr
func (WorkEffortType) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", WorkEffortType.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique().Annotations(entproto.Field(4)),
                    // m2o
                    edge.To("work_efforts", WorkEffort.Type).
                        Annotations(entproto.Field(5)),
                    // m2o
                    edge.To("child_work_effort_types", WorkEffortType.Type).
                        Annotations(entproto.Field(6)),
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


type CommunicationEventPrpTyp struct {
    ent.Schema
}

// Fields of the CommunicationEventPrpTyp.
// Unique-Indexes: communicationEventPrpTypId
func (CommunicationEventPrpTyp) Fields() []ent.Field {
    return []ent.Field{
        field.Enum("has_table").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
        field.String("description").Optional().Annotations(entproto.Field(3)),
    }
}

//* entproto annotations ??
func (CommunicationEventPrpTyp) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (CommunicationEventPrpTyp) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the CommunicationEventPrpTyp.
//  one: ParentCommunicationEventPrpTyp
//  many: ChildCommunicationEventPrpTyp
//  many: CommunicationEventPurpose
func (CommunicationEventPrpTyp) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", CommunicationEventPrpTyp.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique().Annotations(entproto.Field(4)),
                    // m2o
                    edge.To("child_communication_event_prp_typs", CommunicationEventPrpTyp.Type).
                        Annotations(entproto.Field(5)),
    }
}


type PartyRole struct {
    ent.Schema
}

// Fields of the PartyRole.
// Unique-Indexes: partyId, roleTypeId
func (PartyRole) Fields() []ent.Field {
    return []ent.Field{
    }
}

//* entproto annotations ??
func (PartyRole) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // o2m
                    edge.From("party", Party.Type).Ref("party_roles").
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(0)),
                    // o2m
                    edge.From("role_type", RoleType.Type).Ref("party_roles").
                        // Bind the "roleTypeId" field to this edge.
                        // Field("role_type_id").
                        Unique().Annotations(entproto.Field(1)),
                    // m2o
                    edge.To("fixed_assets", FixedAsset.Type).
                        Annotations(entproto.Field(23)),
                    // m2o
                    edge.To("party_contact_meches", PartyContactMech.Type).
                        Annotations(entproto.Field(35)),
                    // m2o
                    edge.To("work_effort_party_assignments", WorkEffortPartyAssignment.Type).
                        Annotations(entproto.Field(63)),
    }
}


type WorkEffortFixedAssetAssign struct {
    ent.Schema
}

// Fields of the WorkEffortFixedAssetAssign.
// Unique-Indexes: workEffortId, fixedAssetId, fromDate
func (WorkEffortFixedAssetAssign) Fields() []ent.Field {
    return []ent.Field{
        field.Time("from_date").
            Default(time.Now).Annotations(entproto.Field(2)),
        field.Time("thru_date").
            Default(time.Now).Optional().Annotations(entproto.Field(3)),
        field.Float("allocated_cost").Optional().Annotations(entproto.Field(4)),
        field.String("comments").Optional().Annotations(entproto.Field(5)),
    }
}

//* entproto annotations ??
func (WorkEffortFixedAssetAssign) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // o2m
                    edge.From("work_effort", WorkEffort.Type).Ref("work_effort_fixed_asset_assigns").
                        // Bind the "workEffortId" field to this edge.
                        // Field("work_effort_id").
                        Unique().Annotations(entproto.Field(6)),
                    // o2m
                    edge.From("fixed_asset", FixedAsset.Type).Ref("work_effort_fixed_asset_assigns").
                        // Bind the "fixedAssetId" field to this edge.
                        // Field("fixed_asset_id").
                        Unique().Annotations(entproto.Field(7)),
                    // o2m
                    edge.From("status_item", StatusItem.Type).Ref("work_effort_fixed_asset_assigns").
                        // Bind the "statusId" field to this edge.
                        // Field("status_id").
                        Unique().Annotations(entproto.Field(8)),
                    // o2m
                    edge.From("availability_status_item", StatusItem.Type).Ref("availability_work_effort_fixed_asset_assigns").
                        // Bind the "availabilityStatusId" field to this edge.
                        // Field("availability_status_id").
                        Unique().Annotations(entproto.Field(9)),
    }
}


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
                    edge.To("fixed_assets", FixedAsset.Type).
                        Annotations(entproto.Field(20)),
                    // m2o
                    edge.To("party_contact_meches", PartyContactMech.Type).
                        Annotations(entproto.Field(28)),
                    // m2o
                    edge.To("valid_from_party_relationship_types", PartyRelationshipType.Type).
                        Annotations(entproto.Field(35)),
                    // m2o
                    edge.To("valid_to_party_relationship_types", PartyRelationshipType.Type).
                        Annotations(entproto.Field(36)),
                    // m2o
                    edge.To("party_roles", PartyRole.Type).
                        Annotations(entproto.Field(37)),
                    // m2o
                    edge.To("child_role_types", RoleType.Type).
                        Annotations(entproto.Field(48)),
                    // m2o
                    edge.To("work_effort_party_assignments", WorkEffortPartyAssignment.Type).
                        Annotations(entproto.Field(58)),
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


type PartyType struct {
    ent.Schema
}

// Fields of the PartyType.
// Unique-Indexes: partyTypeId
func (PartyType) Fields() []ent.Field {
    return []ent.Field{
        field.Enum("has_table").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
        field.String("description").Optional().Annotations(entproto.Field(3)),
    }
}

//* entproto annotations ??
func (PartyType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (PartyType) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the PartyType.
//  one: ParentPartyType
//  many: SiblingPartyType
//  many: Party
//  many: PartyNeed
//  many: ChildPartyType
//  many: PartyTypeAttr
func (PartyType) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", PartyType.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique().Annotations(entproto.Field(4)),
                    // m2o
                    edge.To("sibling_party_types", PartyType.Type).
                        Annotations(entproto.Field(5)),
                    // m2o
                    edge.To("parties", Party.Type).
                        Annotations(entproto.Field(6)),
                    // m2o
                    edge.To("child_party_types", PartyType.Type).
                        Annotations(entproto.Field(8)),
    }
}


type TermType struct {
    ent.Schema
}

// Fields of the TermType.
// Unique-Indexes: termTypeId
func (TermType) Fields() []ent.Field {
    return []ent.Field{
        field.Enum("has_table").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
        field.String("description").Optional().Annotations(entproto.Field(3)),
    }
}

//* entproto annotations ??
func (TermType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (TermType) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the TermType.
//  one: ParentTermType
//  many: AgreementTerm
//  many: BillingAccountTerm
//  many: InvoiceTerm
//  many: OrderTerm
//  many: QuoteTerm
//  many: ChildTermType
//  many: TermTypeAttr
func (TermType) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", TermType.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique().Annotations(entproto.Field(4)),
                    // m2o
                    edge.To("child_term_types", TermType.Type).
                        Annotations(entproto.Field(10)),
    }
}


type Person struct {
    ent.Schema
}

// Fields of the Person.
// Unique-Indexes: partyId
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
        field.String("social_security_number").Optional().Annotations(entproto.Field(21)),
        field.String("passport_number").Optional().Annotations(entproto.Field(22)),
        field.Time("passport_expire_date").
            Default(time.Now).Optional().Annotations(entproto.Field(23)),
        field.Float("total_years_work_experience").Optional().Annotations(entproto.Field(24)),
        field.String("comments").Optional().Annotations(entproto.Field(25)),
        field.String("occupation").Optional().Annotations(entproto.Field(26)),
        field.Int("years_with_employer").Optional().Annotations(entproto.Field(27)),
        field.Int("months_with_employer").Optional().Annotations(entproto.Field(28)),
        field.Enum("existing_customer").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(29)),
        field.String("card_id").MaxLen(32).Optional().Annotations(entproto.Field(30)),
    }
}

//* entproto annotations ??
func (Person) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // o2o (nofk)
                    edge.From("party", Party.Type).
                        Ref("person").
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(31)),
                    // o2m
                    edge.From("employment_status_enumeration", Enumeration.Type).Ref("employment_status_people").
                        // Bind the "employmentStatusEnumId" field to this edge.
                        // Field("employment_status_enum_id").
                        Unique().Annotations(entproto.Field(32)),
                    // o2m
                    edge.From("residence_status_enumeration", Enumeration.Type).Ref("residence_status_people").
                        // Bind the "residenceStatusEnumId" field to this edge.
                        // Field("residence_status_enum_id").
                        Unique().Annotations(entproto.Field(33)),
                    // o2m
                    edge.From("marital_status_enumeration", Enumeration.Type).Ref("marital_status_people").
                        // Bind the "maritalStatusEnumId" field to this edge.
                        // Field("marital_status_enum_id").
                        Unique().Annotations(entproto.Field(34)),
                    // m2o
                    edge.To("party_contact_meches", PartyContactMech.Type).
                        Annotations(entproto.Field(35)),
                    // m2o
                    edge.To("user_logins", UserLogin.Type).
                        Annotations(entproto.Field(43)),
    }
}


type PartyQualType struct {
    ent.Schema
}

// Fields of the PartyQualType.
// Unique-Indexes: partyQualTypeId
func (PartyQualType) Fields() []ent.Field {
    return []ent.Field{
        field.Enum("has_table").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
        field.String("description").Optional().Annotations(entproto.Field(3)),
    }
}

//* entproto annotations ??
func (PartyQualType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (PartyQualType) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the PartyQualType.
//  one: ParentPartyQualType
//  many: PartyQual
//  many: ChildPartyQualType
func (PartyQualType) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", PartyQualType.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique().Annotations(entproto.Field(4)),
                    // m2o
                    edge.To("child_party_qual_types", PartyQualType.Type).
                        Annotations(entproto.Field(6)),
    }
}


type FixedAsset struct {
    ent.Schema
}

// Fields of the FixedAsset.
// Unique-Indexes: fixedAssetId
func (FixedAsset) Fields() []ent.Field {
    return []ent.Field{
        field.Int("fixed_asset_type_id").Optional().Annotations(entproto.Field(2)),
        field.Int("instance_of_product_id").Optional().Annotations(entproto.Field(3)),
        field.String("fixed_asset_name").Optional().Annotations(entproto.Field(4)),
        field.Int("acquire_order_id").Optional().Annotations(entproto.Field(5)),
        field.Int("acquire_order_item_seq_id").Optional().Annotations(entproto.Field(6)),
        field.Time("date_acquired").
            Default(time.Now).Optional().Annotations(entproto.Field(7)),
        field.Time("date_last_serviced").
            Default(time.Now).Optional().Annotations(entproto.Field(8)),
        field.Time("date_next_service").
            Default(time.Now).Optional().Annotations(entproto.Field(9)),
        field.Time("expected_end_of_life").
            Default(time.Now).Optional().Annotations(entproto.Field(10)),
        field.Time("actual_end_of_life").
            Default(time.Now).Optional().Annotations(entproto.Field(11)),
        field.Float("production_capacity").Optional().Annotations(entproto.Field(12)),
        field.Int("uom_id").Optional().Annotations(entproto.Field(13)),
        field.Int("calendar_id").Optional().Annotations(entproto.Field(14)),
        field.String("serial_number").Optional().Annotations(entproto.Field(15)),
        field.Int("located_at_facility_id").Optional().Annotations(entproto.Field(16)),
        field.Int("located_at_location_seq_id").Optional().Annotations(entproto.Field(17)),
        field.Float("salvage_value").Optional().Annotations(entproto.Field(18)),
        field.Float("depreciation").Optional().Annotations(entproto.Field(19)),
        field.Float("purchase_cost").Optional().Annotations(entproto.Field(20)),
        field.Int("purchase_cost_uom_id").Optional().Annotations(entproto.Field(21)),
    }
}

//* entproto annotations ??
func (FixedAsset) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    Unique().Annotations(entproto.Field(24)),
                    // o2m
                    edge.From("class_enumeration", Enumeration.Type).Ref("class_fixed_assets").
                        // Bind the "classEnumId" field to this edge.
                        // Field("class_enum_id").
                        Unique().Annotations(entproto.Field(26)),
                    // o2m
                    edge.From("party", Party.Type).Ref("fixed_assets").
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(27)),
                    // o2m
                    edge.From("role_type", RoleType.Type).Ref("fixed_assets").
                        // Bind the "roleTypeId" field to this edge.
                        // Field("role_type_id").
                        Unique().Annotations(entproto.Field(28)),
                    // o2m
                    edge.From("party_role", PartyRole.Type).Ref("fixed_assets").
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(29)),
                    // m2o
                    edge.To("child_fixed_assets", FixedAsset.Type).
                        Annotations(entproto.Field(41)),
                    // m2o
                    edge.To("work_efforts", WorkEffort.Type).
                        Annotations(entproto.Field(55)),
                    // m2o
                    edge.To("work_effort_fixed_asset_assigns", WorkEffortFixedAssetAssign.Type).
                        Annotations(entproto.Field(56)),
    }
}


type Party struct {
    ent.Schema
}

// Fields of the Party.
// Unique-Indexes: partyId
func (Party) Fields() []ent.Field {
    return []ent.Field{
        field.Int("external_id").Optional().Annotations(entproto.Field(2)),
        field.Int("preferred_currency_uom_id").Optional().Annotations(entproto.Field(3)),
        field.String("description").Optional().Annotations(entproto.Field(4)),
        field.Time("created_date").
            Default(time.Now).Optional().Annotations(entproto.Field(5)),
        field.Time("last_modified_date").
            Default(time.Now).Optional().Annotations(entproto.Field(6)),
        field.Int("data_source_id").Optional().Annotations(entproto.Field(7)),
        field.Enum("is_unread").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(8)),
    }
}

//* entproto annotations ??
func (Party) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // o2m
                    edge.From("party_type", PartyType.Type).Ref("parties").
                        // Bind the "partyTypeId" field to this edge.
                        // Field("party_type_id").
                        Unique().Annotations(entproto.Field(9)),
                    // o2m
                    edge.From("created_by_user_login", UserLogin.Type).Ref("created_by_parties").
                        // Bind the "createdByUserLogin" field to this edge.
                        // Field("created_by_user_login").
                        Unique().Annotations(entproto.Field(10)),
                    // o2m
                    edge.From("last_modified_by_user_login", UserLogin.Type).Ref("last_modified_by_parties").
                        // Bind the "lastModifiedByUserLogin" field to this edge.
                        // Field("last_modified_by_user_login").
                        Unique().Annotations(entproto.Field(11)),
                    // o2m
                    edge.From("status_item", StatusItem.Type).Ref("parties").
                        // Bind the "statusId" field to this edge.
                        // Field("status_id").
                        Unique().Annotations(entproto.Field(13)),
                    // m2o
                    edge.To("fixed_assets", FixedAsset.Type).
                        Annotations(entproto.Field(63)),
                    // m2o
                    edge.To("party_contact_meches", PartyContactMech.Type).
                        Annotations(entproto.Field(100)),
                    // m2o
                    edge.To("party_roles", PartyRole.Type).
                        Annotations(entproto.Field(123)),
                    // m2o
                    edge.To("party_statuses", PartyStatus.Type).
                        Annotations(entproto.Field(125)),
                    // o2o
                    edge.To("person", Person.Type).
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(137)),
                    // m2o
                    edge.To("user_logins", UserLogin.Type).
                        Annotations(entproto.Field(187)),
                    // m2o
                    edge.To("work_effort_party_assignments", WorkEffortPartyAssignment.Type).
                        Annotations(entproto.Field(195)),
    }
}


type UserLoginSecurityGroup struct {
    ent.Schema
}

// Fields of the UserLoginSecurityGroup.
// Unique-Indexes: userLoginId, groupId, fromDate
func (UserLoginSecurityGroup) Fields() []ent.Field {
    return []ent.Field{
        field.Time("from_date").
            Default(time.Now).Annotations(entproto.Field(2)),
        field.Time("thru_date").
            Default(time.Now).Optional().Annotations(entproto.Field(3)),
    }
}

//* entproto annotations ??
func (UserLoginSecurityGroup) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // o2m
                    edge.From("user_login", UserLogin.Type).Ref("user_login_security_groups").
                        // Bind the "userLoginId" field to this edge.
                        // Field("user_login_id").
                        Unique().Annotations(entproto.Field(4)),
                    // o2m
                    edge.From("security_group", SecurityGroup.Type).Ref("user_login_security_groups").
                        // Bind the "groupId" field to this edge.
                        // Field("group_id").
                        Unique().Annotations(entproto.Field(5)),
                    // m2o
                    edge.To("security_group_permissions", SecurityGroupPermission.Type).
                        Annotations(entproto.Field(6)),
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
                    edge.To("class_fixed_assets", FixedAsset.Type).
                        Annotations(entproto.Field(19)),
                    // m2o
                    edge.To("employment_status_people", Person.Type).
                        Annotations(entproto.Field(39)),
                    // m2o
                    edge.To("residence_status_people", Person.Type).
                        Annotations(entproto.Field(40)),
                    // m2o
                    edge.To("marital_status_people", Person.Type).
                        Annotations(entproto.Field(41)),
                    // m2o
                    edge.To("scope_work_efforts", WorkEffort.Type).
                        Annotations(entproto.Field(73)),
                    // m2o
                    edge.To("expectation_work_effort_party_assignments", WorkEffortPartyAssignment.Type).
                        Annotations(entproto.Field(74)),
                    // m2o
                    edge.To("delegate_reason_work_effort_party_assignments", WorkEffortPartyAssignment.Type).
                        Annotations(entproto.Field(75)),
    }
}


type WorkEffortPartyAssignment struct {
    ent.Schema
}

// Fields of the WorkEffortPartyAssignment.
// Unique-Indexes: workEffortId, partyId, roleTypeId, fromDate
func (WorkEffortPartyAssignment) Fields() []ent.Field {
    return []ent.Field{
        field.Time("from_date").
            Default(time.Now).Annotations(entproto.Field(2)),
        field.Time("thru_date").
            Default(time.Now).Optional().Annotations(entproto.Field(3)),
        field.Time("status_date_time").
            Default(time.Now).Optional().Annotations(entproto.Field(4)),
        field.Int("facility_id").Optional().Annotations(entproto.Field(5)),
        field.String("comments").Optional().Annotations(entproto.Field(6)),
        field.Enum("must_rsvp").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(7)),
    }
}

//* entproto annotations ??
func (WorkEffortPartyAssignment) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // o2m
                    edge.From("work_effort", WorkEffort.Type).Ref("work_effort_party_assignments").
                        // Bind the "workEffortId" field to this edge.
                        // Field("work_effort_id").
                        Unique().Annotations(entproto.Field(8)),
                    // o2m
                    edge.From("party", Party.Type).Ref("work_effort_party_assignments").
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(9)),
                    // o2m
                    edge.From("party_role", PartyRole.Type).Ref("work_effort_party_assignments").
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(10)),
                    // o2m
                    edge.From("role_type", RoleType.Type).Ref("work_effort_party_assignments").
                        // Bind the "roleTypeId" field to this edge.
                        // Field("role_type_id").
                        Unique().Annotations(entproto.Field(11)),
                    // o2m
                    edge.From("assigned_by_user_login", UserLogin.Type).Ref("assigned_by_work_effort_party_assignments").
                        // Bind the "assignedByUserLoginId" field to this edge.
                        // Field("assigned_by_user_login_id").
                        Unique().Annotations(entproto.Field(12)),
                    // o2m
                    edge.From("assignment_status_item", StatusItem.Type).Ref("assignment_work_effort_party_assignments").
                        // Bind the "statusId" field to this edge.
                        // Field("status_id").
                        Unique().Annotations(entproto.Field(13)),
                    // o2m
                    edge.From("expectation_enumeration", Enumeration.Type).Ref("expectation_work_effort_party_assignments").
                        // Bind the "expectationEnumId" field to this edge.
                        // Field("expectation_enum_id").
                        Unique().Annotations(entproto.Field(14)),
                    // o2m
                    edge.From("delegate_reason_enumeration", Enumeration.Type).Ref("delegate_reason_work_effort_party_assignments").
                        // Bind the "delegateReasonEnumId" field to this edge.
                        // Field("delegate_reason_enum_id").
                        Unique().Annotations(entproto.Field(15)),
                    // o2m
                    edge.From("availability_status_item", StatusItem.Type).Ref("availability_work_effort_party_assignments").
                        // Bind the "availabilityStatusId" field to this edge.
                        // Field("availability_status_id").
                        Unique().Annotations(entproto.Field(17)),
    }
}


type ContactMechType struct {
    ent.Schema
}

// Fields of the ContactMechType.
// Unique-Indexes: contactMechTypeId
func (ContactMechType) Fields() []ent.Field {
    return []ent.Field{
        field.Enum("has_table").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
        field.String("description").Optional().Annotations(entproto.Field(3)),
    }
}

//* entproto annotations ??
func (ContactMechType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (ContactMechType) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the ContactMechType.
//  one: ParentContactMechType
//  many: CommunicationEvent
//  many: ContacMechTypeCommunicationEventType
//  many: ContactList
//  many: ContactMech
//  many: ChildContactMechType
//  many: ContactMechTypeAttr
//  many: ContactMechTypePurpose
//  many: ValidContactMechRole
func (ContactMechType) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", ContactMechType.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique().Annotations(entproto.Field(4)),
                    // m2o
                    edge.To("contac_mech_type_communication_event_types", CommunicationEventType.Type).
                        Annotations(entproto.Field(6)),
                    // m2o
                    edge.To("child_contact_mech_types", ContactMechType.Type).
                        Annotations(entproto.Field(9)),
                    // m2o
                    edge.To("contact_mech_type_purposes", ContactMechTypePurpose.Type).
                        Annotations(entproto.Field(11)),
    }
}


type SkillType struct {
    ent.Schema
}

// Fields of the SkillType.
// Unique-Indexes: skillTypeId
func (SkillType) Fields() []ent.Field {
    return []ent.Field{
        field.Enum("has_table").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
        field.String("description").Optional().Annotations(entproto.Field(3)),
    }
}

//* entproto annotations ??
func (SkillType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (SkillType) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the SkillType.
//  one: ParentSkillType
//  many: JobRequisition
//  many: PartySkill
//  many: QuoteItem
//  many: ChildSkillType
//  many: WorkEffortSkillStandard
func (SkillType) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", SkillType.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique().Annotations(entproto.Field(4)),
                    // m2o
                    edge.To("child_skill_types", SkillType.Type).
                        Annotations(entproto.Field(8)),
                    // m2o
                    edge.To("work_effort_skill_standards", WorkEffortSkillStandard.Type).
                        Annotations(entproto.Field(9)),
    }
}


type WorkEffort struct {
    ent.Schema
}

// Fields of the WorkEffort.
// Unique-Indexes: workEffortId
func (WorkEffort) Fields() []ent.Field {
    return []ent.Field{
        field.Time("last_status_update").
            Default(time.Now).Optional().Annotations(entproto.Field(2)),
        field.Int("work_effort_purpose_type_id").Optional().Annotations(entproto.Field(3)),
        field.Int("priority").Optional().Annotations(entproto.Field(4)),
        field.Int("percent_complete").Optional().Annotations(entproto.Field(5)),
        field.String("work_effort_name").Optional().Annotations(entproto.Field(6)),
        field.Int("show_as_enum_id").Optional().Annotations(entproto.Field(7)),
        field.Enum("send_notification_email").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(8)),
        field.String("description").Optional().Annotations(entproto.Field(9)),
        field.String("location_desc").Optional().Annotations(entproto.Field(10)),
        field.Time("estimated_start_date").
            Default(time.Now).Optional().Annotations(entproto.Field(11)),
        field.Time("estimated_completion_date").
            Default(time.Now).Optional().Annotations(entproto.Field(12)),
        field.Time("actual_start_date").
            Default(time.Now).Optional().Annotations(entproto.Field(13)),
        field.Time("actual_completion_date").
            Default(time.Now).Optional().Annotations(entproto.Field(14)),
        field.Float("estimated_milli_seconds").Optional().Annotations(entproto.Field(15)),
        field.Float("estimated_setup_millis").Optional().Annotations(entproto.Field(16)),
        field.Int("estimate_calc_method").Optional().Annotations(entproto.Field(17)),
        field.Float("actual_milli_seconds").Optional().Annotations(entproto.Field(18)),
        field.Float("actual_setup_millis").Optional().Annotations(entproto.Field(19)),
        field.Float("total_milli_seconds_allowed").Optional().Annotations(entproto.Field(20)),
        field.Float("total_money_allowed").Optional().Annotations(entproto.Field(21)),
        field.Int("money_uom_id").Optional().Annotations(entproto.Field(22)),
        field.String("special_terms").Optional().Annotations(entproto.Field(23)),
        field.Int("time_transparency").Optional().Annotations(entproto.Field(24)),
        field.String("universal_id").Optional().Annotations(entproto.Field(25)),
        field.String("source_reference_id").MaxLen(32).Optional().Annotations(entproto.Field(26)),
        field.Int("facility_id").Optional().Annotations(entproto.Field(27)),
        field.String("info_url").Optional().Annotations(entproto.Field(28)),
        field.Int("recurrence_info_id").Optional().Annotations(entproto.Field(29)),
        field.Int("runtime_data_id").Optional().Annotations(entproto.Field(30)),
        field.Int("note_id").Optional().Annotations(entproto.Field(31)),
        field.String("service_loader_name").Optional().Annotations(entproto.Field(32)),
        field.Float("quantity_to_produce").Optional().Annotations(entproto.Field(33)),
        field.Float("quantity_produced").Optional().Annotations(entproto.Field(34)),
        field.Float("quantity_rejected").Optional().Annotations(entproto.Field(35)),
        field.Float("reserv_persons").Optional().Annotations(entproto.Field(36)),
        field.Float("reserv_2_nd_pp_perc").Optional().Annotations(entproto.Field(37)),
        field.Float("reserv_nth_pp_perc").Optional().Annotations(entproto.Field(38)),
        field.Int("accommodation_map_id").Optional().Annotations(entproto.Field(39)),
        field.Int("accommodation_spot_id").Optional().Annotations(entproto.Field(40)),
        field.Int("revision_number").Optional().Annotations(entproto.Field(41)),
        field.Time("created_date").
            Default(time.Now).Optional().Annotations(entproto.Field(42)),
        field.String("created_by_user_login").Optional().Annotations(entproto.Field(43)),
        field.Time("last_modified_date").
            Default(time.Now).Optional().Annotations(entproto.Field(44)),
        field.String("last_modified_by_user_login").Optional().Annotations(entproto.Field(45)),
        field.Int("sequence_num").Optional().Annotations(entproto.Field(46)),
    }
}

//* entproto annotations ??
func (WorkEffort) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // o2m
                    edge.From("work_effort_type", WorkEffortType.Type).Ref("work_efforts").
                        // Bind the "workEffortTypeId" field to this edge.
                        // Field("work_effort_type_id").
                        Unique().Annotations(entproto.Field(47)),
                edge.To("children", WorkEffort.Type).
                    From("parent").
                    // Bind the "workEffortParentId" field to this edge.
                    // Field("work_effort_parent_id").
                    Unique().Annotations(entproto.Field(49)),
                    // o2m
                    edge.From("current_status_item", StatusItem.Type).Ref("current_work_efforts").
                        // Bind the "currentStatusId" field to this edge.
                        // Field("current_status_id").
                        Unique().Annotations(entproto.Field(51)),
                    // o2m
                    edge.From("scope_enumeration", Enumeration.Type).Ref("scope_work_efforts").
                        // Bind the "scopeEnumId" field to this edge.
                        // Field("scope_enum_id").
                        Unique().Annotations(entproto.Field(52)),
                    // o2m
                    edge.From("fixed_asset", FixedAsset.Type).Ref("work_efforts").
                        // Bind the "fixedAssetId" field to this edge.
                        // Field("fixed_asset_id").
                        Unique().Annotations(entproto.Field(53)),
                    // o2m
                    edge.From("temporal_expression", TemporalExpression.Type).Ref("work_efforts").
                        // Bind the "tempExprId" field to this edge.
                        // Field("temp_expr_id").
                        Unique().Annotations(entproto.Field(57)),
                    // m2o
                    edge.To("child_work_efforts", WorkEffort.Type).
                        Annotations(entproto.Field(83)),
                    // m2o
                    edge.To("from_work_effort_assocs", WorkEffortAssoc.Type).
                        Annotations(entproto.Field(84)),
                    // m2o
                    edge.To("to_work_effort_assocs", WorkEffortAssoc.Type).
                        Annotations(entproto.Field(85)),
                    // m2o
                    edge.To("work_effort_fixed_asset_assigns", WorkEffortFixedAssetAssign.Type).
                        Annotations(entproto.Field(93)),
                    // m2o
                    edge.To("work_effort_party_assignments", WorkEffortPartyAssignment.Type).
                        Annotations(entproto.Field(101)),
                    // m2o
                    edge.To("work_effort_skill_standards", WorkEffortSkillStandard.Type).
                        Annotations(entproto.Field(103)),
    }
}


type PartyRelationshipType struct {
    ent.Schema
}

// Fields of the PartyRelationshipType.
// Unique-Indexes: partyRelationshipTypeId
func (PartyRelationshipType) Fields() []ent.Field {
    return []ent.Field{
        field.Enum("has_table").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
        field.String("party_relationship_name").Optional().Annotations(entproto.Field(3)),
        field.String("description").Optional().Annotations(entproto.Field(4)),
    }
}

//* entproto annotations ??
func (PartyRelationshipType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (PartyRelationshipType) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the PartyRelationshipType.
//  one: ParentPartyRelationshipType
//  one: ValidFromRoleType
//  one: ValidToRoleType
//  many: PartyRelationship
//  many: ChildPartyRelationshipType
func (PartyRelationshipType) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", PartyRelationshipType.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique().Annotations(entproto.Field(5)),
                    // o2m
                    edge.From("valid_from_role_type", RoleType.Type).Ref("valid_from_party_relationship_types").
                        // Bind the "roleTypeIdValidFrom" field to this edge.
                        // Field("role_type_id_valid_from").
                        Unique().Annotations(entproto.Field(6)),
                    // o2m
                    edge.From("valid_to_role_type", RoleType.Type).Ref("valid_to_party_relationship_types").
                        // Bind the "roleTypeIdValidTo" field to this edge.
                        // Field("role_type_id_valid_to").
                        Unique().Annotations(entproto.Field(7)),
                    // m2o
                    edge.To("child_party_relationship_types", PartyRelationshipType.Type).
                        Annotations(entproto.Field(9)),
    }
}


type ContactMechPurposeType struct {
    ent.Schema
}

// Fields of the ContactMechPurposeType.
// Unique-Indexes: contactMechPurposeTypeId
func (ContactMechPurposeType) Fields() []ent.Field {
    return []ent.Field{
        field.Int("parent_type_id").Optional().Annotations(entproto.Field(2)),
        field.Enum("has_table").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(3)),
        field.String("description").Optional().Annotations(entproto.Field(4)),
    }
}

//* entproto annotations ??
func (ContactMechPurposeType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (ContactMechPurposeType) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the ContactMechPurposeType.
//  many: ContactMechTypePurpose
//  many: FacilityContactMechPurpose
//  many: InvoiceContactMech
//  many: OrderContactMech
//  many: OrderItemContactMech
//  many: PartyContactMechPurpose
//  many: ReturnContactMech
func (ContactMechPurposeType) Edges() []ent.Edge {
    return []ent.Edge{
                    // m2o
                    edge.To("contact_mech_type_purposes", ContactMechTypePurpose.Type).
                        Annotations(entproto.Field(5)),
    }
}


type PartyStatus struct {
    ent.Schema
}

// Fields of the PartyStatus.
// Unique-Indexes: statusId, partyId, statusDate
func (PartyStatus) Fields() []ent.Field {
    return []ent.Field{
        field.Time("status_date").
            Default(time.Now).Annotations(entproto.Field(2)),
    }
}

//* entproto annotations ??
func (PartyStatus) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // o2m
                    edge.From("status_item", StatusItem.Type).Ref("party_statuses").
                        // Bind the "statusId" field to this edge.
                        // Field("status_id").
                        Unique().Annotations(entproto.Field(3)),
                    // o2m
                    edge.From("party", Party.Type).Ref("party_statuses").
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(4)),
                    // o2m
                    edge.From("change_by_user_login", UserLogin.Type).Ref("change_by_party_statuses").
                        // Bind the "changeByUserLoginId" field to this edge.
                        // Field("change_by_user_login_id").
                        Unique().Annotations(entproto.Field(5)),
    }
}


type UserLogin struct {
    ent.Schema
}

// Fields of the UserLogin.
// Unique-Indexes: userLoginId
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

//* entproto annotations ??
func (UserLogin) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // o2m
                    edge.From("party", Party.Type).Ref("user_logins").
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(16)),
                    // o2m
                    edge.From("person", Person.Type).Ref("user_logins").
                        // Bind the "partyId" field to this edge.
                        // Field("party_id").
                        Unique().Annotations(entproto.Field(17)),
                    // m2o
                    edge.To("created_by_parties", Party.Type).
                        Annotations(entproto.Field(50)),
                    // m2o
                    edge.To("last_modified_by_parties", Party.Type).
                        Annotations(entproto.Field(51)),
                    // m2o
                    edge.To("change_by_party_statuses", PartyStatus.Type).
                        Annotations(entproto.Field(52)),
                    // m2o
                    edge.To("user_login_security_groups", UserLoginSecurityGroup.Type).
                        Annotations(entproto.Field(85)),
                    // m2o
                    edge.To("user_preferences", UserPreference.Type).
                        Annotations(entproto.Field(87)),
                    // m2o
                    edge.To("assigned_by_work_effort_party_assignments", WorkEffortPartyAssignment.Type).
                        Annotations(entproto.Field(89)),
    }
}


type SecurityGroup struct {
    ent.Schema
}

// Fields of the SecurityGroup.
// Unique-Indexes: groupId
func (SecurityGroup) Fields() []ent.Field {
    return []ent.Field{
        field.String("group_name").Optional().Annotations(entproto.Field(2)),
        field.String("description").Optional().Annotations(entproto.Field(3)),
    }
}

//* entproto annotations ??
func (SecurityGroup) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

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
                    // m2o
                    edge.To("user_login_security_groups", UserLoginSecurityGroup.Type).
                        Annotations(entproto.Field(8)),
    }
}


type WorkEffortSkillStandard struct {
    ent.Schema
}

// Fields of the WorkEffortSkillStandard.
// Unique-Indexes: workEffortId, skillTypeId
func (WorkEffortSkillStandard) Fields() []ent.Field {
    return []ent.Field{
        field.Float("estimated_num_people").Optional().Annotations(entproto.Field(2)),
        field.Float("estimated_duration").Optional().Annotations(entproto.Field(3)),
        field.Float("estimated_cost").Optional().Annotations(entproto.Field(4)),
    }
}

//* entproto annotations ??
func (WorkEffortSkillStandard) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
func (WorkEffortSkillStandard) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("work_effort_id", "skill_type_id").
            Unique(),
    }
}

*/

func (WorkEffortSkillStandard) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the WorkEffortSkillStandard.
//  one: WorkEffort
//  one: SkillType
func (WorkEffortSkillStandard) Edges() []ent.Edge {
    return []ent.Edge{
                    // o2m
                    edge.From("work_effort", WorkEffort.Type).Ref("work_effort_skill_standards").
                        // Bind the "workEffortId" field to this edge.
                        // Field("work_effort_id").
                        Unique().Annotations(entproto.Field(5)),
                    // o2m
                    edge.From("skill_type", SkillType.Type).Ref("work_effort_skill_standards").
                        // Bind the "skillTypeId" field to this edge.
                        // Field("skill_type_id").
                        Unique().Annotations(entproto.Field(6)),
    }
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


type SecurityPermission struct {
    ent.Schema
}

// Fields of the SecurityPermission.
// Unique-Indexes: permissionId
func (SecurityPermission) Fields() []ent.Field {
    return []ent.Field{
        field.String("description").Optional().Annotations(entproto.Field(2)),
    }
}

//* entproto annotations ??
func (SecurityPermission) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (SecurityPermission) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the SecurityPermission.
//  many: SecurityGroupPermission
func (SecurityPermission) Edges() []ent.Edge {
    return []ent.Edge{
                    // m2o
                    edge.To("security_group_permissions", SecurityGroupPermission.Type).
                        Annotations(entproto.Field(3)),
    }
}


type CommunicationEventType struct {
    ent.Schema
}

// Fields of the CommunicationEventType.
// Unique-Indexes: communicationEventTypeId
func (CommunicationEventType) Fields() []ent.Field {
    return []ent.Field{
        field.Enum("has_table").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
        field.String("description").Optional().Annotations(entproto.Field(3)),
    }
}

//* entproto annotations ??
func (CommunicationEventType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (CommunicationEventType) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the CommunicationEventType.
//  one: ParentCommunicationEventType
//  one: ContacMechTypeContactMechType
//  many: CommunicationEvent
//  many: ChildCommunicationEventType
func (CommunicationEventType) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", CommunicationEventType.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique().Annotations(entproto.Field(4)),
                    // o2m
                    edge.From("contac_mech_type_contact_mech_type", ContactMechType.Type).Ref("contac_mech_type_communication_event_types").
                        // Bind the "contactMechTypeId" field to this edge.
                        // Field("contact_mech_type_id").
                        Unique().Annotations(entproto.Field(5)),
                    // m2o
                    edge.To("child_communication_event_types", CommunicationEventType.Type).
                        Annotations(entproto.Field(7)),
    }
}


type PartyIdentificationType struct {
    ent.Schema
}

// Fields of the PartyIdentificationType.
// Unique-Indexes: partyIdentificationTypeId
func (PartyIdentificationType) Fields() []ent.Field {
    return []ent.Field{
        field.Enum("has_table").
            Values("Yes", "No", "Unknown").Optional().Annotations(entproto.Field(2)),
        field.String("description").Optional().Annotations(entproto.Field(3)),
    }
}

//* entproto annotations ??
func (PartyIdentificationType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
        schemamixins.StringRefMixin{},
    }
}
//*/

/*
*/

func (PartyIdentificationType) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entproto.Message(),
        entproto.Service(), // also generate a gRPC service definition
    }
}


// Edges of the PartyIdentificationType.
//  one: ParentPartyIdentificationType
//  many: PartyIdentification
//  many: ChildPartyIdentificationType
func (PartyIdentificationType) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", PartyIdentificationType.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique().Annotations(entproto.Field(4)),
                    // m2o
                    edge.To("child_party_identification_types", PartyIdentificationType.Type).
                        Annotations(entproto.Field(6)),
    }
}
