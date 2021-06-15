// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayconfig"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayups"
)

// ShipmentGatewayUpsCreate is the builder for creating a ShipmentGatewayUps entity.
type ShipmentGatewayUpsCreate struct {
	config
	mutation *ShipmentGatewayUpsMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (sguc *ShipmentGatewayUpsCreate) SetCreateTime(t time.Time) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetCreateTime(t)
	return sguc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableCreateTime(t *time.Time) *ShipmentGatewayUpsCreate {
	if t != nil {
		sguc.SetCreateTime(*t)
	}
	return sguc
}

// SetUpdateTime sets the "update_time" field.
func (sguc *ShipmentGatewayUpsCreate) SetUpdateTime(t time.Time) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetUpdateTime(t)
	return sguc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableUpdateTime(t *time.Time) *ShipmentGatewayUpsCreate {
	if t != nil {
		sguc.SetUpdateTime(*t)
	}
	return sguc
}

// SetStringRef sets the "string_ref" field.
func (sguc *ShipmentGatewayUpsCreate) SetStringRef(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetStringRef(s)
	return sguc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableStringRef(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetStringRef(*s)
	}
	return sguc
}

// SetConnectURL sets the "connect_url" field.
func (sguc *ShipmentGatewayUpsCreate) SetConnectURL(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetConnectURL(s)
	return sguc
}

// SetNillableConnectURL sets the "connect_url" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableConnectURL(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetConnectURL(*s)
	}
	return sguc
}

// SetConnectTimeout sets the "connect_timeout" field.
func (sguc *ShipmentGatewayUpsCreate) SetConnectTimeout(i int) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetConnectTimeout(i)
	return sguc
}

// SetNillableConnectTimeout sets the "connect_timeout" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableConnectTimeout(i *int) *ShipmentGatewayUpsCreate {
	if i != nil {
		sguc.SetConnectTimeout(*i)
	}
	return sguc
}

// SetShipperNumber sets the "shipper_number" field.
func (sguc *ShipmentGatewayUpsCreate) SetShipperNumber(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetShipperNumber(s)
	return sguc
}

// SetNillableShipperNumber sets the "shipper_number" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableShipperNumber(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetShipperNumber(*s)
	}
	return sguc
}

// SetBillShipperAccountNumber sets the "bill_shipper_account_number" field.
func (sguc *ShipmentGatewayUpsCreate) SetBillShipperAccountNumber(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetBillShipperAccountNumber(s)
	return sguc
}

// SetNillableBillShipperAccountNumber sets the "bill_shipper_account_number" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableBillShipperAccountNumber(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetBillShipperAccountNumber(*s)
	}
	return sguc
}

// SetAccessLicenseNumber sets the "access_license_number" field.
func (sguc *ShipmentGatewayUpsCreate) SetAccessLicenseNumber(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetAccessLicenseNumber(s)
	return sguc
}

// SetNillableAccessLicenseNumber sets the "access_license_number" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableAccessLicenseNumber(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetAccessLicenseNumber(*s)
	}
	return sguc
}

// SetAccessUserID sets the "access_user_id" field.
func (sguc *ShipmentGatewayUpsCreate) SetAccessUserID(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetAccessUserID(s)
	return sguc
}

// SetNillableAccessUserID sets the "access_user_id" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableAccessUserID(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetAccessUserID(*s)
	}
	return sguc
}

// SetAccessPassword sets the "access_password" field.
func (sguc *ShipmentGatewayUpsCreate) SetAccessPassword(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetAccessPassword(s)
	return sguc
}

// SetNillableAccessPassword sets the "access_password" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableAccessPassword(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetAccessPassword(*s)
	}
	return sguc
}

// SetSaveCertInfo sets the "save_cert_info" field.
func (sguc *ShipmentGatewayUpsCreate) SetSaveCertInfo(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetSaveCertInfo(s)
	return sguc
}

// SetNillableSaveCertInfo sets the "save_cert_info" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableSaveCertInfo(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetSaveCertInfo(*s)
	}
	return sguc
}

// SetSaveCertPath sets the "save_cert_path" field.
func (sguc *ShipmentGatewayUpsCreate) SetSaveCertPath(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetSaveCertPath(s)
	return sguc
}

// SetNillableSaveCertPath sets the "save_cert_path" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableSaveCertPath(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetSaveCertPath(*s)
	}
	return sguc
}

// SetShipperPickupType sets the "shipper_pickup_type" field.
func (sguc *ShipmentGatewayUpsCreate) SetShipperPickupType(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetShipperPickupType(s)
	return sguc
}

// SetNillableShipperPickupType sets the "shipper_pickup_type" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableShipperPickupType(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetShipperPickupType(*s)
	}
	return sguc
}

// SetCustomerClassification sets the "customer_classification" field.
func (sguc *ShipmentGatewayUpsCreate) SetCustomerClassification(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetCustomerClassification(s)
	return sguc
}

// SetNillableCustomerClassification sets the "customer_classification" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableCustomerClassification(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetCustomerClassification(*s)
	}
	return sguc
}

// SetMaxEstimateWeight sets the "max_estimate_weight" field.
func (sguc *ShipmentGatewayUpsCreate) SetMaxEstimateWeight(f float64) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetMaxEstimateWeight(f)
	return sguc
}

// SetNillableMaxEstimateWeight sets the "max_estimate_weight" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableMaxEstimateWeight(f *float64) *ShipmentGatewayUpsCreate {
	if f != nil {
		sguc.SetMaxEstimateWeight(*f)
	}
	return sguc
}

// SetMinEstimateWeight sets the "min_estimate_weight" field.
func (sguc *ShipmentGatewayUpsCreate) SetMinEstimateWeight(f float64) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetMinEstimateWeight(f)
	return sguc
}

// SetNillableMinEstimateWeight sets the "min_estimate_weight" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableMinEstimateWeight(f *float64) *ShipmentGatewayUpsCreate {
	if f != nil {
		sguc.SetMinEstimateWeight(*f)
	}
	return sguc
}

// SetCodAllowCod sets the "cod_allow_cod" field.
func (sguc *ShipmentGatewayUpsCreate) SetCodAllowCod(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetCodAllowCod(s)
	return sguc
}

// SetNillableCodAllowCod sets the "cod_allow_cod" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableCodAllowCod(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetCodAllowCod(*s)
	}
	return sguc
}

// SetCodSurchargeAmount sets the "cod_surcharge_amount" field.
func (sguc *ShipmentGatewayUpsCreate) SetCodSurchargeAmount(f float64) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetCodSurchargeAmount(f)
	return sguc
}

// SetNillableCodSurchargeAmount sets the "cod_surcharge_amount" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableCodSurchargeAmount(f *float64) *ShipmentGatewayUpsCreate {
	if f != nil {
		sguc.SetCodSurchargeAmount(*f)
	}
	return sguc
}

// SetCodSurchargeCurrencyUomID sets the "cod_surcharge_currency_uom_id" field.
func (sguc *ShipmentGatewayUpsCreate) SetCodSurchargeCurrencyUomID(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetCodSurchargeCurrencyUomID(s)
	return sguc
}

// SetNillableCodSurchargeCurrencyUomID sets the "cod_surcharge_currency_uom_id" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableCodSurchargeCurrencyUomID(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetCodSurchargeCurrencyUomID(*s)
	}
	return sguc
}

// SetCodSurchargeApplyToPackage sets the "cod_surcharge_apply_to_package" field.
func (sguc *ShipmentGatewayUpsCreate) SetCodSurchargeApplyToPackage(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetCodSurchargeApplyToPackage(s)
	return sguc
}

// SetNillableCodSurchargeApplyToPackage sets the "cod_surcharge_apply_to_package" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableCodSurchargeApplyToPackage(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetCodSurchargeApplyToPackage(*s)
	}
	return sguc
}

// SetCodFundsCode sets the "cod_funds_code" field.
func (sguc *ShipmentGatewayUpsCreate) SetCodFundsCode(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetCodFundsCode(s)
	return sguc
}

// SetNillableCodFundsCode sets the "cod_funds_code" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableCodFundsCode(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetCodFundsCode(*s)
	}
	return sguc
}

// SetDefaultReturnLabelMemo sets the "default_return_label_memo" field.
func (sguc *ShipmentGatewayUpsCreate) SetDefaultReturnLabelMemo(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetDefaultReturnLabelMemo(s)
	return sguc
}

// SetNillableDefaultReturnLabelMemo sets the "default_return_label_memo" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableDefaultReturnLabelMemo(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetDefaultReturnLabelMemo(*s)
	}
	return sguc
}

// SetDefaultReturnLabelSubject sets the "default_return_label_subject" field.
func (sguc *ShipmentGatewayUpsCreate) SetDefaultReturnLabelSubject(s string) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetDefaultReturnLabelSubject(s)
	return sguc
}

// SetNillableDefaultReturnLabelSubject sets the "default_return_label_subject" field if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableDefaultReturnLabelSubject(s *string) *ShipmentGatewayUpsCreate {
	if s != nil {
		sguc.SetDefaultReturnLabelSubject(*s)
	}
	return sguc
}

// SetShipmentGatewayConfigID sets the "shipment_gateway_config" edge to the ShipmentGatewayConfig entity by ID.
func (sguc *ShipmentGatewayUpsCreate) SetShipmentGatewayConfigID(id int) *ShipmentGatewayUpsCreate {
	sguc.mutation.SetShipmentGatewayConfigID(id)
	return sguc
}

// SetNillableShipmentGatewayConfigID sets the "shipment_gateway_config" edge to the ShipmentGatewayConfig entity by ID if the given value is not nil.
func (sguc *ShipmentGatewayUpsCreate) SetNillableShipmentGatewayConfigID(id *int) *ShipmentGatewayUpsCreate {
	if id != nil {
		sguc = sguc.SetShipmentGatewayConfigID(*id)
	}
	return sguc
}

// SetShipmentGatewayConfig sets the "shipment_gateway_config" edge to the ShipmentGatewayConfig entity.
func (sguc *ShipmentGatewayUpsCreate) SetShipmentGatewayConfig(s *ShipmentGatewayConfig) *ShipmentGatewayUpsCreate {
	return sguc.SetShipmentGatewayConfigID(s.ID)
}

// Mutation returns the ShipmentGatewayUpsMutation object of the builder.
func (sguc *ShipmentGatewayUpsCreate) Mutation() *ShipmentGatewayUpsMutation {
	return sguc.mutation
}

// Save creates the ShipmentGatewayUps in the database.
func (sguc *ShipmentGatewayUpsCreate) Save(ctx context.Context) (*ShipmentGatewayUps, error) {
	var (
		err  error
		node *ShipmentGatewayUps
	)
	sguc.defaults()
	if len(sguc.hooks) == 0 {
		if err = sguc.check(); err != nil {
			return nil, err
		}
		node, err = sguc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShipmentGatewayUpsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sguc.check(); err != nil {
				return nil, err
			}
			sguc.mutation = mutation
			if node, err = sguc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sguc.hooks) - 1; i >= 0; i-- {
			mut = sguc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sguc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sguc *ShipmentGatewayUpsCreate) SaveX(ctx context.Context) *ShipmentGatewayUps {
	v, err := sguc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (sguc *ShipmentGatewayUpsCreate) defaults() {
	if _, ok := sguc.mutation.CreateTime(); !ok {
		v := shipmentgatewayups.DefaultCreateTime()
		sguc.mutation.SetCreateTime(v)
	}
	if _, ok := sguc.mutation.UpdateTime(); !ok {
		v := shipmentgatewayups.DefaultUpdateTime()
		sguc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sguc *ShipmentGatewayUpsCreate) check() error {
	if _, ok := sguc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := sguc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	return nil
}

func (sguc *ShipmentGatewayUpsCreate) sqlSave(ctx context.Context) (*ShipmentGatewayUps, error) {
	_node, _spec := sguc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sguc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sguc *ShipmentGatewayUpsCreate) createSpec() (*ShipmentGatewayUps, *sqlgraph.CreateSpec) {
	var (
		_node = &ShipmentGatewayUps{config: sguc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: shipmentgatewayups.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shipmentgatewayups.FieldID,
			},
		}
	)
	if value, ok := sguc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipmentgatewayups.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := sguc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipmentgatewayups.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := sguc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := sguc.mutation.ConnectURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldConnectURL,
		})
		_node.ConnectURL = value
	}
	if value, ok := sguc.mutation.ConnectTimeout(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipmentgatewayups.FieldConnectTimeout,
		})
		_node.ConnectTimeout = value
	}
	if value, ok := sguc.mutation.ShipperNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldShipperNumber,
		})
		_node.ShipperNumber = value
	}
	if value, ok := sguc.mutation.BillShipperAccountNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldBillShipperAccountNumber,
		})
		_node.BillShipperAccountNumber = value
	}
	if value, ok := sguc.mutation.AccessLicenseNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldAccessLicenseNumber,
		})
		_node.AccessLicenseNumber = value
	}
	if value, ok := sguc.mutation.AccessUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldAccessUserID,
		})
		_node.AccessUserID = value
	}
	if value, ok := sguc.mutation.AccessPassword(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldAccessPassword,
		})
		_node.AccessPassword = value
	}
	if value, ok := sguc.mutation.SaveCertInfo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldSaveCertInfo,
		})
		_node.SaveCertInfo = value
	}
	if value, ok := sguc.mutation.SaveCertPath(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldSaveCertPath,
		})
		_node.SaveCertPath = value
	}
	if value, ok := sguc.mutation.ShipperPickupType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldShipperPickupType,
		})
		_node.ShipperPickupType = value
	}
	if value, ok := sguc.mutation.CustomerClassification(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldCustomerClassification,
		})
		_node.CustomerClassification = value
	}
	if value, ok := sguc.mutation.MaxEstimateWeight(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: shipmentgatewayups.FieldMaxEstimateWeight,
		})
		_node.MaxEstimateWeight = value
	}
	if value, ok := sguc.mutation.MinEstimateWeight(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: shipmentgatewayups.FieldMinEstimateWeight,
		})
		_node.MinEstimateWeight = value
	}
	if value, ok := sguc.mutation.CodAllowCod(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldCodAllowCod,
		})
		_node.CodAllowCod = value
	}
	if value, ok := sguc.mutation.CodSurchargeAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: shipmentgatewayups.FieldCodSurchargeAmount,
		})
		_node.CodSurchargeAmount = value
	}
	if value, ok := sguc.mutation.CodSurchargeCurrencyUomID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldCodSurchargeCurrencyUomID,
		})
		_node.CodSurchargeCurrencyUomID = value
	}
	if value, ok := sguc.mutation.CodSurchargeApplyToPackage(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldCodSurchargeApplyToPackage,
		})
		_node.CodSurchargeApplyToPackage = value
	}
	if value, ok := sguc.mutation.CodFundsCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldCodFundsCode,
		})
		_node.CodFundsCode = value
	}
	if value, ok := sguc.mutation.DefaultReturnLabelMemo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldDefaultReturnLabelMemo,
		})
		_node.DefaultReturnLabelMemo = value
	}
	if value, ok := sguc.mutation.DefaultReturnLabelSubject(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewayups.FieldDefaultReturnLabelSubject,
		})
		_node.DefaultReturnLabelSubject = value
	}
	if nodes := sguc.mutation.ShipmentGatewayConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   shipmentgatewayups.ShipmentGatewayConfigTable,
			Columns: []string{shipmentgatewayups.ShipmentGatewayConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shipmentgatewayconfig.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.shipment_gateway_config_shipment_gateway_ups = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ShipmentGatewayUpsCreateBulk is the builder for creating many ShipmentGatewayUps entities in bulk.
type ShipmentGatewayUpsCreateBulk struct {
	config
	builders []*ShipmentGatewayUpsCreate
}

// Save creates the ShipmentGatewayUps entities in the database.
func (sgucb *ShipmentGatewayUpsCreateBulk) Save(ctx context.Context) ([]*ShipmentGatewayUps, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sgucb.builders))
	nodes := make([]*ShipmentGatewayUps, len(sgucb.builders))
	mutators := make([]Mutator, len(sgucb.builders))
	for i := range sgucb.builders {
		func(i int, root context.Context) {
			builder := sgucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShipmentGatewayUpsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sgucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sgucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sgucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sgucb *ShipmentGatewayUpsCreateBulk) SaveX(ctx context.Context) []*ShipmentGatewayUps {
	v, err := sgucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}