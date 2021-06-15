// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderheader"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitem"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/product"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/statusitem"
)

// OrderItem is the model entity for the OrderItem schema.
type OrderItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StringRef holds the value of the "string_ref" field.
	StringRef string `json:"string_ref,omitempty"`
	// OrderItemSeqID holds the value of the "order_item_seq_id" field.
	OrderItemSeqID int `json:"order_item_seq_id,omitempty"`
	// ExternalID holds the value of the "external_id" field.
	ExternalID int `json:"external_id,omitempty"`
	// OrderItemTypeID holds the value of the "order_item_type_id" field.
	OrderItemTypeID int `json:"order_item_type_id,omitempty"`
	// OrderItemGroupSeqID holds the value of the "order_item_group_seq_id" field.
	OrderItemGroupSeqID int `json:"order_item_group_seq_id,omitempty"`
	// IsItemGroupPrimary holds the value of the "is_item_group_primary" field.
	IsItemGroupPrimary orderitem.IsItemGroupPrimary `json:"is_item_group_primary,omitempty"`
	// FromInventoryItemID holds the value of the "from_inventory_item_id" field.
	FromInventoryItemID int `json:"from_inventory_item_id,omitempty"`
	// BudgetID holds the value of the "budget_id" field.
	BudgetID int `json:"budget_id,omitempty"`
	// BudgetItemSeqID holds the value of the "budget_item_seq_id" field.
	BudgetItemSeqID int `json:"budget_item_seq_id,omitempty"`
	// SupplierProductID holds the value of the "supplier_product_id" field.
	SupplierProductID string `json:"supplier_product_id,omitempty"`
	// ProductFeatureID holds the value of the "product_feature_id" field.
	ProductFeatureID int `json:"product_feature_id,omitempty"`
	// ProdCatalogID holds the value of the "prod_catalog_id" field.
	ProdCatalogID int `json:"prod_catalog_id,omitempty"`
	// ProductCategoryID holds the value of the "product_category_id" field.
	ProductCategoryID int `json:"product_category_id,omitempty"`
	// IsPromo holds the value of the "is_promo" field.
	IsPromo orderitem.IsPromo `json:"is_promo,omitempty"`
	// QuoteID holds the value of the "quote_id" field.
	QuoteID int `json:"quote_id,omitempty"`
	// QuoteItemSeqID holds the value of the "quote_item_seq_id" field.
	QuoteItemSeqID int `json:"quote_item_seq_id,omitempty"`
	// ShoppingListID holds the value of the "shopping_list_id" field.
	ShoppingListID int `json:"shopping_list_id,omitempty"`
	// ShoppingListItemSeqID holds the value of the "shopping_list_item_seq_id" field.
	ShoppingListItemSeqID int `json:"shopping_list_item_seq_id,omitempty"`
	// SubscriptionID holds the value of the "subscription_id" field.
	SubscriptionID int `json:"subscription_id,omitempty"`
	// DeploymentID holds the value of the "deployment_id" field.
	DeploymentID int `json:"deployment_id,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity float64 `json:"quantity,omitempty"`
	// CancelQuantity holds the value of the "cancel_quantity" field.
	CancelQuantity float64 `json:"cancel_quantity,omitempty"`
	// SelectedAmount holds the value of the "selected_amount" field.
	SelectedAmount float64 `json:"selected_amount,omitempty"`
	// UnitPrice holds the value of the "unit_price" field.
	UnitPrice float64 `json:"unit_price,omitempty"`
	// UnitListPrice holds the value of the "unit_list_price" field.
	UnitListPrice float64 `json:"unit_list_price,omitempty"`
	// UnitAverageCost holds the value of the "unit_average_cost" field.
	UnitAverageCost float64 `json:"unit_average_cost,omitempty"`
	// UnitRecurringPrice holds the value of the "unit_recurring_price" field.
	UnitRecurringPrice float64 `json:"unit_recurring_price,omitempty"`
	// IsModifiedPrice holds the value of the "is_modified_price" field.
	IsModifiedPrice orderitem.IsModifiedPrice `json:"is_modified_price,omitempty"`
	// RecurringFreqUomID holds the value of the "recurring_freq_uom_id" field.
	RecurringFreqUomID int `json:"recurring_freq_uom_id,omitempty"`
	// ItemDescription holds the value of the "item_description" field.
	ItemDescription string `json:"item_description,omitempty"`
	// Comments holds the value of the "comments" field.
	Comments string `json:"comments,omitempty"`
	// CorrespondingPoID holds the value of the "corresponding_po_id" field.
	CorrespondingPoID int `json:"corresponding_po_id,omitempty"`
	// EstimatedShipDate holds the value of the "estimated_ship_date" field.
	EstimatedShipDate time.Time `json:"estimated_ship_date,omitempty"`
	// EstimatedDeliveryDate holds the value of the "estimated_delivery_date" field.
	EstimatedDeliveryDate time.Time `json:"estimated_delivery_date,omitempty"`
	// AutoCancelDate holds the value of the "auto_cancel_date" field.
	AutoCancelDate time.Time `json:"auto_cancel_date,omitempty"`
	// DontCancelSetDate holds the value of the "dont_cancel_set_date" field.
	DontCancelSetDate time.Time `json:"dont_cancel_set_date,omitempty"`
	// DontCancelSetUserLogin holds the value of the "dont_cancel_set_user_login" field.
	DontCancelSetUserLogin string `json:"dont_cancel_set_user_login,omitempty"`
	// ShipBeforeDate holds the value of the "ship_before_date" field.
	ShipBeforeDate time.Time `json:"ship_before_date,omitempty"`
	// ShipAfterDate holds the value of the "ship_after_date" field.
	ShipAfterDate time.Time `json:"ship_after_date,omitempty"`
	// ReserveAfterDate holds the value of the "reserve_after_date" field.
	ReserveAfterDate time.Time `json:"reserve_after_date,omitempty"`
	// CancelBackOrderDate holds the value of the "cancel_back_order_date" field.
	CancelBackOrderDate time.Time `json:"cancel_back_order_date,omitempty"`
	// OverrideGlAccountID holds the value of the "override_gl_account_id" field.
	OverrideGlAccountID int `json:"override_gl_account_id,omitempty"`
	// SalesOpportunityID holds the value of the "sales_opportunity_id" field.
	SalesOpportunityID int `json:"sales_opportunity_id,omitempty"`
	// ChangeByUserLoginID holds the value of the "change_by_user_login_id" field.
	ChangeByUserLoginID string `json:"change_by_user_login_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderItemQuery when eager-loading is set.
	Edges                        OrderItemEdges `json:"edges"`
	order_header_order_items     *int
	order_role_order_items       *int
	product_order_items          *int
	status_item_order_items      *int
	status_item_sync_order_items *int
}

// OrderItemEdges holds the relations/edges for other nodes in the graph.
type OrderItemEdges struct {
	// OrderHeader holds the value of the order_header edge.
	OrderHeader *OrderHeader `json:"order_header,omitempty"`
	// Product holds the value of the product edge.
	Product *Product `json:"product,omitempty"`
	// StatusItem holds the value of the status_item edge.
	StatusItem *StatusItem `json:"status_item,omitempty"`
	// StatusValidChanges holds the value of the status_valid_changes edge.
	StatusValidChanges []*StatusValidChange `json:"status_valid_changes,omitempty"`
	// SyncStatusItem holds the value of the sync_status_item edge.
	SyncStatusItem *StatusItem `json:"sync_status_item,omitempty"`
	// OrderAdjustments holds the value of the order_adjustments edge.
	OrderAdjustments []*OrderAdjustment `json:"order_adjustments,omitempty"`
	// OrderItemShipGroupAssocs holds the value of the order_item_ship_group_assocs edge.
	OrderItemShipGroupAssocs []*OrderItemShipGroupAssoc `json:"order_item_ship_group_assocs,omitempty"`
	// OrderItemShipGrpInvRes holds the value of the order_item_ship_grp_inv_res edge.
	OrderItemShipGrpInvRes []*OrderItemShipGrpInvRes `json:"order_item_ship_grp_inv_res,omitempty"`
	// OrderPaymentPreferences holds the value of the order_payment_preferences edge.
	OrderPaymentPreferences []*OrderPaymentPreference `json:"order_payment_preferences,omitempty"`
	// OrderStatuses holds the value of the order_statuses edge.
	OrderStatuses []*OrderStatus `json:"order_statuses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [10]bool
}

// OrderHeaderOrErr returns the OrderHeader value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderItemEdges) OrderHeaderOrErr() (*OrderHeader, error) {
	if e.loadedTypes[0] {
		if e.OrderHeader == nil {
			// The edge order_header was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: orderheader.Label}
		}
		return e.OrderHeader, nil
	}
	return nil, &NotLoadedError{edge: "order_header"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderItemEdges) ProductOrErr() (*Product, error) {
	if e.loadedTypes[1] {
		if e.Product == nil {
			// The edge product was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// StatusItemOrErr returns the StatusItem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderItemEdges) StatusItemOrErr() (*StatusItem, error) {
	if e.loadedTypes[2] {
		if e.StatusItem == nil {
			// The edge status_item was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: statusitem.Label}
		}
		return e.StatusItem, nil
	}
	return nil, &NotLoadedError{edge: "status_item"}
}

// StatusValidChangesOrErr returns the StatusValidChanges value or an error if the edge
// was not loaded in eager-loading.
func (e OrderItemEdges) StatusValidChangesOrErr() ([]*StatusValidChange, error) {
	if e.loadedTypes[3] {
		return e.StatusValidChanges, nil
	}
	return nil, &NotLoadedError{edge: "status_valid_changes"}
}

// SyncStatusItemOrErr returns the SyncStatusItem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderItemEdges) SyncStatusItemOrErr() (*StatusItem, error) {
	if e.loadedTypes[4] {
		if e.SyncStatusItem == nil {
			// The edge sync_status_item was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: statusitem.Label}
		}
		return e.SyncStatusItem, nil
	}
	return nil, &NotLoadedError{edge: "sync_status_item"}
}

// OrderAdjustmentsOrErr returns the OrderAdjustments value or an error if the edge
// was not loaded in eager-loading.
func (e OrderItemEdges) OrderAdjustmentsOrErr() ([]*OrderAdjustment, error) {
	if e.loadedTypes[5] {
		return e.OrderAdjustments, nil
	}
	return nil, &NotLoadedError{edge: "order_adjustments"}
}

// OrderItemShipGroupAssocsOrErr returns the OrderItemShipGroupAssocs value or an error if the edge
// was not loaded in eager-loading.
func (e OrderItemEdges) OrderItemShipGroupAssocsOrErr() ([]*OrderItemShipGroupAssoc, error) {
	if e.loadedTypes[6] {
		return e.OrderItemShipGroupAssocs, nil
	}
	return nil, &NotLoadedError{edge: "order_item_ship_group_assocs"}
}

// OrderItemShipGrpInvResOrErr returns the OrderItemShipGrpInvRes value or an error if the edge
// was not loaded in eager-loading.
func (e OrderItemEdges) OrderItemShipGrpInvResOrErr() ([]*OrderItemShipGrpInvRes, error) {
	if e.loadedTypes[7] {
		return e.OrderItemShipGrpInvRes, nil
	}
	return nil, &NotLoadedError{edge: "order_item_ship_grp_inv_res"}
}

// OrderPaymentPreferencesOrErr returns the OrderPaymentPreferences value or an error if the edge
// was not loaded in eager-loading.
func (e OrderItemEdges) OrderPaymentPreferencesOrErr() ([]*OrderPaymentPreference, error) {
	if e.loadedTypes[8] {
		return e.OrderPaymentPreferences, nil
	}
	return nil, &NotLoadedError{edge: "order_payment_preferences"}
}

// OrderStatusesOrErr returns the OrderStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e OrderItemEdges) OrderStatusesOrErr() ([]*OrderStatus, error) {
	if e.loadedTypes[9] {
		return e.OrderStatuses, nil
	}
	return nil, &NotLoadedError{edge: "order_statuses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderItem) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldQuantity, orderitem.FieldCancelQuantity, orderitem.FieldSelectedAmount, orderitem.FieldUnitPrice, orderitem.FieldUnitListPrice, orderitem.FieldUnitAverageCost, orderitem.FieldUnitRecurringPrice:
			values[i] = new(sql.NullFloat64)
		case orderitem.FieldID, orderitem.FieldOrderItemSeqID, orderitem.FieldExternalID, orderitem.FieldOrderItemTypeID, orderitem.FieldOrderItemGroupSeqID, orderitem.FieldFromInventoryItemID, orderitem.FieldBudgetID, orderitem.FieldBudgetItemSeqID, orderitem.FieldProductFeatureID, orderitem.FieldProdCatalogID, orderitem.FieldProductCategoryID, orderitem.FieldQuoteID, orderitem.FieldQuoteItemSeqID, orderitem.FieldShoppingListID, orderitem.FieldShoppingListItemSeqID, orderitem.FieldSubscriptionID, orderitem.FieldDeploymentID, orderitem.FieldRecurringFreqUomID, orderitem.FieldCorrespondingPoID, orderitem.FieldOverrideGlAccountID, orderitem.FieldSalesOpportunityID:
			values[i] = new(sql.NullInt64)
		case orderitem.FieldStringRef, orderitem.FieldIsItemGroupPrimary, orderitem.FieldSupplierProductID, orderitem.FieldIsPromo, orderitem.FieldIsModifiedPrice, orderitem.FieldItemDescription, orderitem.FieldComments, orderitem.FieldDontCancelSetUserLogin, orderitem.FieldChangeByUserLoginID:
			values[i] = new(sql.NullString)
		case orderitem.FieldCreateTime, orderitem.FieldUpdateTime, orderitem.FieldEstimatedShipDate, orderitem.FieldEstimatedDeliveryDate, orderitem.FieldAutoCancelDate, orderitem.FieldDontCancelSetDate, orderitem.FieldShipBeforeDate, orderitem.FieldShipAfterDate, orderitem.FieldReserveAfterDate, orderitem.FieldCancelBackOrderDate:
			values[i] = new(sql.NullTime)
		case orderitem.ForeignKeys[0]: // order_header_order_items
			values[i] = new(sql.NullInt64)
		case orderitem.ForeignKeys[1]: // order_role_order_items
			values[i] = new(sql.NullInt64)
		case orderitem.ForeignKeys[2]: // product_order_items
			values[i] = new(sql.NullInt64)
		case orderitem.ForeignKeys[3]: // status_item_order_items
			values[i] = new(sql.NullInt64)
		case orderitem.ForeignKeys[4]: // status_item_sync_order_items
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderItem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderItem fields.
func (oi *OrderItem) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oi.ID = int(value.Int64)
		case orderitem.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				oi.CreateTime = value.Time
			}
		case orderitem.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				oi.UpdateTime = value.Time
			}
		case orderitem.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				oi.StringRef = value.String
			}
		case orderitem.FieldOrderItemSeqID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_item_seq_id", values[i])
			} else if value.Valid {
				oi.OrderItemSeqID = int(value.Int64)
			}
		case orderitem.FieldExternalID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field external_id", values[i])
			} else if value.Valid {
				oi.ExternalID = int(value.Int64)
			}
		case orderitem.FieldOrderItemTypeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_item_type_id", values[i])
			} else if value.Valid {
				oi.OrderItemTypeID = int(value.Int64)
			}
		case orderitem.FieldOrderItemGroupSeqID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_item_group_seq_id", values[i])
			} else if value.Valid {
				oi.OrderItemGroupSeqID = int(value.Int64)
			}
		case orderitem.FieldIsItemGroupPrimary:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field is_item_group_primary", values[i])
			} else if value.Valid {
				oi.IsItemGroupPrimary = orderitem.IsItemGroupPrimary(value.String)
			}
		case orderitem.FieldFromInventoryItemID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field from_inventory_item_id", values[i])
			} else if value.Valid {
				oi.FromInventoryItemID = int(value.Int64)
			}
		case orderitem.FieldBudgetID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field budget_id", values[i])
			} else if value.Valid {
				oi.BudgetID = int(value.Int64)
			}
		case orderitem.FieldBudgetItemSeqID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field budget_item_seq_id", values[i])
			} else if value.Valid {
				oi.BudgetItemSeqID = int(value.Int64)
			}
		case orderitem.FieldSupplierProductID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field supplier_product_id", values[i])
			} else if value.Valid {
				oi.SupplierProductID = value.String
			}
		case orderitem.FieldProductFeatureID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_feature_id", values[i])
			} else if value.Valid {
				oi.ProductFeatureID = int(value.Int64)
			}
		case orderitem.FieldProdCatalogID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field prod_catalog_id", values[i])
			} else if value.Valid {
				oi.ProdCatalogID = int(value.Int64)
			}
		case orderitem.FieldProductCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_category_id", values[i])
			} else if value.Valid {
				oi.ProductCategoryID = int(value.Int64)
			}
		case orderitem.FieldIsPromo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field is_promo", values[i])
			} else if value.Valid {
				oi.IsPromo = orderitem.IsPromo(value.String)
			}
		case orderitem.FieldQuoteID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quote_id", values[i])
			} else if value.Valid {
				oi.QuoteID = int(value.Int64)
			}
		case orderitem.FieldQuoteItemSeqID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quote_item_seq_id", values[i])
			} else if value.Valid {
				oi.QuoteItemSeqID = int(value.Int64)
			}
		case orderitem.FieldShoppingListID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field shopping_list_id", values[i])
			} else if value.Valid {
				oi.ShoppingListID = int(value.Int64)
			}
		case orderitem.FieldShoppingListItemSeqID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field shopping_list_item_seq_id", values[i])
			} else if value.Valid {
				oi.ShoppingListItemSeqID = int(value.Int64)
			}
		case orderitem.FieldSubscriptionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field subscription_id", values[i])
			} else if value.Valid {
				oi.SubscriptionID = int(value.Int64)
			}
		case orderitem.FieldDeploymentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_id", values[i])
			} else if value.Valid {
				oi.DeploymentID = int(value.Int64)
			}
		case orderitem.FieldQuantity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				oi.Quantity = value.Float64
			}
		case orderitem.FieldCancelQuantity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field cancel_quantity", values[i])
			} else if value.Valid {
				oi.CancelQuantity = value.Float64
			}
		case orderitem.FieldSelectedAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field selected_amount", values[i])
			} else if value.Valid {
				oi.SelectedAmount = value.Float64
			}
		case orderitem.FieldUnitPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field unit_price", values[i])
			} else if value.Valid {
				oi.UnitPrice = value.Float64
			}
		case orderitem.FieldUnitListPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field unit_list_price", values[i])
			} else if value.Valid {
				oi.UnitListPrice = value.Float64
			}
		case orderitem.FieldUnitAverageCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field unit_average_cost", values[i])
			} else if value.Valid {
				oi.UnitAverageCost = value.Float64
			}
		case orderitem.FieldUnitRecurringPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field unit_recurring_price", values[i])
			} else if value.Valid {
				oi.UnitRecurringPrice = value.Float64
			}
		case orderitem.FieldIsModifiedPrice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field is_modified_price", values[i])
			} else if value.Valid {
				oi.IsModifiedPrice = orderitem.IsModifiedPrice(value.String)
			}
		case orderitem.FieldRecurringFreqUomID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field recurring_freq_uom_id", values[i])
			} else if value.Valid {
				oi.RecurringFreqUomID = int(value.Int64)
			}
		case orderitem.FieldItemDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item_description", values[i])
			} else if value.Valid {
				oi.ItemDescription = value.String
			}
		case orderitem.FieldComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				oi.Comments = value.String
			}
		case orderitem.FieldCorrespondingPoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field corresponding_po_id", values[i])
			} else if value.Valid {
				oi.CorrespondingPoID = int(value.Int64)
			}
		case orderitem.FieldEstimatedShipDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field estimated_ship_date", values[i])
			} else if value.Valid {
				oi.EstimatedShipDate = value.Time
			}
		case orderitem.FieldEstimatedDeliveryDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field estimated_delivery_date", values[i])
			} else if value.Valid {
				oi.EstimatedDeliveryDate = value.Time
			}
		case orderitem.FieldAutoCancelDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field auto_cancel_date", values[i])
			} else if value.Valid {
				oi.AutoCancelDate = value.Time
			}
		case orderitem.FieldDontCancelSetDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dont_cancel_set_date", values[i])
			} else if value.Valid {
				oi.DontCancelSetDate = value.Time
			}
		case orderitem.FieldDontCancelSetUserLogin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dont_cancel_set_user_login", values[i])
			} else if value.Valid {
				oi.DontCancelSetUserLogin = value.String
			}
		case orderitem.FieldShipBeforeDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ship_before_date", values[i])
			} else if value.Valid {
				oi.ShipBeforeDate = value.Time
			}
		case orderitem.FieldShipAfterDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ship_after_date", values[i])
			} else if value.Valid {
				oi.ShipAfterDate = value.Time
			}
		case orderitem.FieldReserveAfterDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field reserve_after_date", values[i])
			} else if value.Valid {
				oi.ReserveAfterDate = value.Time
			}
		case orderitem.FieldCancelBackOrderDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field cancel_back_order_date", values[i])
			} else if value.Valid {
				oi.CancelBackOrderDate = value.Time
			}
		case orderitem.FieldOverrideGlAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field override_gl_account_id", values[i])
			} else if value.Valid {
				oi.OverrideGlAccountID = int(value.Int64)
			}
		case orderitem.FieldSalesOpportunityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sales_opportunity_id", values[i])
			} else if value.Valid {
				oi.SalesOpportunityID = int(value.Int64)
			}
		case orderitem.FieldChangeByUserLoginID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field change_by_user_login_id", values[i])
			} else if value.Valid {
				oi.ChangeByUserLoginID = value.String
			}
		case orderitem.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_header_order_items", value)
			} else if value.Valid {
				oi.order_header_order_items = new(int)
				*oi.order_header_order_items = int(value.Int64)
			}
		case orderitem.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_role_order_items", value)
			} else if value.Valid {
				oi.order_role_order_items = new(int)
				*oi.order_role_order_items = int(value.Int64)
			}
		case orderitem.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field product_order_items", value)
			} else if value.Valid {
				oi.product_order_items = new(int)
				*oi.product_order_items = int(value.Int64)
			}
		case orderitem.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field status_item_order_items", value)
			} else if value.Valid {
				oi.status_item_order_items = new(int)
				*oi.status_item_order_items = int(value.Int64)
			}
		case orderitem.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field status_item_sync_order_items", value)
			} else if value.Valid {
				oi.status_item_sync_order_items = new(int)
				*oi.status_item_sync_order_items = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOrderHeader queries the "order_header" edge of the OrderItem entity.
func (oi *OrderItem) QueryOrderHeader() *OrderHeaderQuery {
	return (&OrderItemClient{config: oi.config}).QueryOrderHeader(oi)
}

// QueryProduct queries the "product" edge of the OrderItem entity.
func (oi *OrderItem) QueryProduct() *ProductQuery {
	return (&OrderItemClient{config: oi.config}).QueryProduct(oi)
}

// QueryStatusItem queries the "status_item" edge of the OrderItem entity.
func (oi *OrderItem) QueryStatusItem() *StatusItemQuery {
	return (&OrderItemClient{config: oi.config}).QueryStatusItem(oi)
}

// QueryStatusValidChanges queries the "status_valid_changes" edge of the OrderItem entity.
func (oi *OrderItem) QueryStatusValidChanges() *StatusValidChangeQuery {
	return (&OrderItemClient{config: oi.config}).QueryStatusValidChanges(oi)
}

// QuerySyncStatusItem queries the "sync_status_item" edge of the OrderItem entity.
func (oi *OrderItem) QuerySyncStatusItem() *StatusItemQuery {
	return (&OrderItemClient{config: oi.config}).QuerySyncStatusItem(oi)
}

// QueryOrderAdjustments queries the "order_adjustments" edge of the OrderItem entity.
func (oi *OrderItem) QueryOrderAdjustments() *OrderAdjustmentQuery {
	return (&OrderItemClient{config: oi.config}).QueryOrderAdjustments(oi)
}

// QueryOrderItemShipGroupAssocs queries the "order_item_ship_group_assocs" edge of the OrderItem entity.
func (oi *OrderItem) QueryOrderItemShipGroupAssocs() *OrderItemShipGroupAssocQuery {
	return (&OrderItemClient{config: oi.config}).QueryOrderItemShipGroupAssocs(oi)
}

// QueryOrderItemShipGrpInvRes queries the "order_item_ship_grp_inv_res" edge of the OrderItem entity.
func (oi *OrderItem) QueryOrderItemShipGrpInvRes() *OrderItemShipGrpInvResQuery {
	return (&OrderItemClient{config: oi.config}).QueryOrderItemShipGrpInvRes(oi)
}

// QueryOrderPaymentPreferences queries the "order_payment_preferences" edge of the OrderItem entity.
func (oi *OrderItem) QueryOrderPaymentPreferences() *OrderPaymentPreferenceQuery {
	return (&OrderItemClient{config: oi.config}).QueryOrderPaymentPreferences(oi)
}

// QueryOrderStatuses queries the "order_statuses" edge of the OrderItem entity.
func (oi *OrderItem) QueryOrderStatuses() *OrderStatusQuery {
	return (&OrderItemClient{config: oi.config}).QueryOrderStatuses(oi)
}

// Update returns a builder for updating this OrderItem.
// Note that you need to call OrderItem.Unwrap() before calling this method if this OrderItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (oi *OrderItem) Update() *OrderItemUpdateOne {
	return (&OrderItemClient{config: oi.config}).UpdateOne(oi)
}

// Unwrap unwraps the OrderItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oi *OrderItem) Unwrap() *OrderItem {
	tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderItem is not a transactional entity")
	}
	oi.config.driver = tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *OrderItem) String() string {
	var builder strings.Builder
	builder.WriteString("OrderItem(")
	builder.WriteString(fmt.Sprintf("id=%v", oi.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(oi.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(oi.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(oi.StringRef)
	builder.WriteString(", order_item_seq_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.OrderItemSeqID))
	builder.WriteString(", external_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.ExternalID))
	builder.WriteString(", order_item_type_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.OrderItemTypeID))
	builder.WriteString(", order_item_group_seq_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.OrderItemGroupSeqID))
	builder.WriteString(", is_item_group_primary=")
	builder.WriteString(fmt.Sprintf("%v", oi.IsItemGroupPrimary))
	builder.WriteString(", from_inventory_item_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.FromInventoryItemID))
	builder.WriteString(", budget_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.BudgetID))
	builder.WriteString(", budget_item_seq_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.BudgetItemSeqID))
	builder.WriteString(", supplier_product_id=")
	builder.WriteString(oi.SupplierProductID)
	builder.WriteString(", product_feature_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.ProductFeatureID))
	builder.WriteString(", prod_catalog_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.ProdCatalogID))
	builder.WriteString(", product_category_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.ProductCategoryID))
	builder.WriteString(", is_promo=")
	builder.WriteString(fmt.Sprintf("%v", oi.IsPromo))
	builder.WriteString(", quote_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.QuoteID))
	builder.WriteString(", quote_item_seq_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.QuoteItemSeqID))
	builder.WriteString(", shopping_list_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.ShoppingListID))
	builder.WriteString(", shopping_list_item_seq_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.ShoppingListItemSeqID))
	builder.WriteString(", subscription_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.SubscriptionID))
	builder.WriteString(", deployment_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.DeploymentID))
	builder.WriteString(", quantity=")
	builder.WriteString(fmt.Sprintf("%v", oi.Quantity))
	builder.WriteString(", cancel_quantity=")
	builder.WriteString(fmt.Sprintf("%v", oi.CancelQuantity))
	builder.WriteString(", selected_amount=")
	builder.WriteString(fmt.Sprintf("%v", oi.SelectedAmount))
	builder.WriteString(", unit_price=")
	builder.WriteString(fmt.Sprintf("%v", oi.UnitPrice))
	builder.WriteString(", unit_list_price=")
	builder.WriteString(fmt.Sprintf("%v", oi.UnitListPrice))
	builder.WriteString(", unit_average_cost=")
	builder.WriteString(fmt.Sprintf("%v", oi.UnitAverageCost))
	builder.WriteString(", unit_recurring_price=")
	builder.WriteString(fmt.Sprintf("%v", oi.UnitRecurringPrice))
	builder.WriteString(", is_modified_price=")
	builder.WriteString(fmt.Sprintf("%v", oi.IsModifiedPrice))
	builder.WriteString(", recurring_freq_uom_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.RecurringFreqUomID))
	builder.WriteString(", item_description=")
	builder.WriteString(oi.ItemDescription)
	builder.WriteString(", comments=")
	builder.WriteString(oi.Comments)
	builder.WriteString(", corresponding_po_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.CorrespondingPoID))
	builder.WriteString(", estimated_ship_date=")
	builder.WriteString(oi.EstimatedShipDate.Format(time.ANSIC))
	builder.WriteString(", estimated_delivery_date=")
	builder.WriteString(oi.EstimatedDeliveryDate.Format(time.ANSIC))
	builder.WriteString(", auto_cancel_date=")
	builder.WriteString(oi.AutoCancelDate.Format(time.ANSIC))
	builder.WriteString(", dont_cancel_set_date=")
	builder.WriteString(oi.DontCancelSetDate.Format(time.ANSIC))
	builder.WriteString(", dont_cancel_set_user_login=")
	builder.WriteString(oi.DontCancelSetUserLogin)
	builder.WriteString(", ship_before_date=")
	builder.WriteString(oi.ShipBeforeDate.Format(time.ANSIC))
	builder.WriteString(", ship_after_date=")
	builder.WriteString(oi.ShipAfterDate.Format(time.ANSIC))
	builder.WriteString(", reserve_after_date=")
	builder.WriteString(oi.ReserveAfterDate.Format(time.ANSIC))
	builder.WriteString(", cancel_back_order_date=")
	builder.WriteString(oi.CancelBackOrderDate.Format(time.ANSIC))
	builder.WriteString(", override_gl_account_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.OverrideGlAccountID))
	builder.WriteString(", sales_opportunity_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.SalesOpportunityID))
	builder.WriteString(", change_by_user_login_id=")
	builder.WriteString(oi.ChangeByUserLoginID)
	builder.WriteByte(')')
	return builder.String()
}

// OrderItems is a parsable slice of OrderItem.
type OrderItems []*OrderItem

func (oi OrderItems) config(cfg config) {
	for _i := range oi {
		oi[_i].config = cfg
	}
}