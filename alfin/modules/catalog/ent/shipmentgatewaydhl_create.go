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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewaydhl"
)

// ShipmentGatewayDhlCreate is the builder for creating a ShipmentGatewayDhl entity.
type ShipmentGatewayDhlCreate struct {
	config
	mutation *ShipmentGatewayDhlMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (sgdc *ShipmentGatewayDhlCreate) SetCreateTime(t time.Time) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetCreateTime(t)
	return sgdc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableCreateTime(t *time.Time) *ShipmentGatewayDhlCreate {
	if t != nil {
		sgdc.SetCreateTime(*t)
	}
	return sgdc
}

// SetUpdateTime sets the "update_time" field.
func (sgdc *ShipmentGatewayDhlCreate) SetUpdateTime(t time.Time) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetUpdateTime(t)
	return sgdc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableUpdateTime(t *time.Time) *ShipmentGatewayDhlCreate {
	if t != nil {
		sgdc.SetUpdateTime(*t)
	}
	return sgdc
}

// SetStringRef sets the "string_ref" field.
func (sgdc *ShipmentGatewayDhlCreate) SetStringRef(s string) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetStringRef(s)
	return sgdc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableStringRef(s *string) *ShipmentGatewayDhlCreate {
	if s != nil {
		sgdc.SetStringRef(*s)
	}
	return sgdc
}

// SetConnectURL sets the "connect_url" field.
func (sgdc *ShipmentGatewayDhlCreate) SetConnectURL(s string) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetConnectURL(s)
	return sgdc
}

// SetNillableConnectURL sets the "connect_url" field if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableConnectURL(s *string) *ShipmentGatewayDhlCreate {
	if s != nil {
		sgdc.SetConnectURL(*s)
	}
	return sgdc
}

// SetConnectTimeout sets the "connect_timeout" field.
func (sgdc *ShipmentGatewayDhlCreate) SetConnectTimeout(i int) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetConnectTimeout(i)
	return sgdc
}

// SetNillableConnectTimeout sets the "connect_timeout" field if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableConnectTimeout(i *int) *ShipmentGatewayDhlCreate {
	if i != nil {
		sgdc.SetConnectTimeout(*i)
	}
	return sgdc
}

// SetHeadVersion sets the "head_version" field.
func (sgdc *ShipmentGatewayDhlCreate) SetHeadVersion(s string) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetHeadVersion(s)
	return sgdc
}

// SetNillableHeadVersion sets the "head_version" field if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableHeadVersion(s *string) *ShipmentGatewayDhlCreate {
	if s != nil {
		sgdc.SetHeadVersion(*s)
	}
	return sgdc
}

// SetHeadAction sets the "head_action" field.
func (sgdc *ShipmentGatewayDhlCreate) SetHeadAction(s string) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetHeadAction(s)
	return sgdc
}

// SetNillableHeadAction sets the "head_action" field if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableHeadAction(s *string) *ShipmentGatewayDhlCreate {
	if s != nil {
		sgdc.SetHeadAction(*s)
	}
	return sgdc
}

// SetAccessUserID sets the "access_user_id" field.
func (sgdc *ShipmentGatewayDhlCreate) SetAccessUserID(s string) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetAccessUserID(s)
	return sgdc
}

// SetNillableAccessUserID sets the "access_user_id" field if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableAccessUserID(s *string) *ShipmentGatewayDhlCreate {
	if s != nil {
		sgdc.SetAccessUserID(*s)
	}
	return sgdc
}

// SetAccessPassword sets the "access_password" field.
func (sgdc *ShipmentGatewayDhlCreate) SetAccessPassword(s string) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetAccessPassword(s)
	return sgdc
}

// SetNillableAccessPassword sets the "access_password" field if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableAccessPassword(s *string) *ShipmentGatewayDhlCreate {
	if s != nil {
		sgdc.SetAccessPassword(*s)
	}
	return sgdc
}

// SetAccessAccountNbr sets the "access_account_nbr" field.
func (sgdc *ShipmentGatewayDhlCreate) SetAccessAccountNbr(s string) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetAccessAccountNbr(s)
	return sgdc
}

// SetNillableAccessAccountNbr sets the "access_account_nbr" field if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableAccessAccountNbr(s *string) *ShipmentGatewayDhlCreate {
	if s != nil {
		sgdc.SetAccessAccountNbr(*s)
	}
	return sgdc
}

// SetAccessShippingKey sets the "access_shipping_key" field.
func (sgdc *ShipmentGatewayDhlCreate) SetAccessShippingKey(s string) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetAccessShippingKey(s)
	return sgdc
}

// SetNillableAccessShippingKey sets the "access_shipping_key" field if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableAccessShippingKey(s *string) *ShipmentGatewayDhlCreate {
	if s != nil {
		sgdc.SetAccessShippingKey(*s)
	}
	return sgdc
}

// SetLabelImageFormat sets the "label_image_format" field.
func (sgdc *ShipmentGatewayDhlCreate) SetLabelImageFormat(s string) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetLabelImageFormat(s)
	return sgdc
}

// SetNillableLabelImageFormat sets the "label_image_format" field if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableLabelImageFormat(s *string) *ShipmentGatewayDhlCreate {
	if s != nil {
		sgdc.SetLabelImageFormat(*s)
	}
	return sgdc
}

// SetRateEstimateTemplate sets the "rate_estimate_template" field.
func (sgdc *ShipmentGatewayDhlCreate) SetRateEstimateTemplate(s string) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetRateEstimateTemplate(s)
	return sgdc
}

// SetNillableRateEstimateTemplate sets the "rate_estimate_template" field if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableRateEstimateTemplate(s *string) *ShipmentGatewayDhlCreate {
	if s != nil {
		sgdc.SetRateEstimateTemplate(*s)
	}
	return sgdc
}

// SetShipmentGatewayConfigID sets the "shipment_gateway_config" edge to the ShipmentGatewayConfig entity by ID.
func (sgdc *ShipmentGatewayDhlCreate) SetShipmentGatewayConfigID(id int) *ShipmentGatewayDhlCreate {
	sgdc.mutation.SetShipmentGatewayConfigID(id)
	return sgdc
}

// SetNillableShipmentGatewayConfigID sets the "shipment_gateway_config" edge to the ShipmentGatewayConfig entity by ID if the given value is not nil.
func (sgdc *ShipmentGatewayDhlCreate) SetNillableShipmentGatewayConfigID(id *int) *ShipmentGatewayDhlCreate {
	if id != nil {
		sgdc = sgdc.SetShipmentGatewayConfigID(*id)
	}
	return sgdc
}

// SetShipmentGatewayConfig sets the "shipment_gateway_config" edge to the ShipmentGatewayConfig entity.
func (sgdc *ShipmentGatewayDhlCreate) SetShipmentGatewayConfig(s *ShipmentGatewayConfig) *ShipmentGatewayDhlCreate {
	return sgdc.SetShipmentGatewayConfigID(s.ID)
}

// Mutation returns the ShipmentGatewayDhlMutation object of the builder.
func (sgdc *ShipmentGatewayDhlCreate) Mutation() *ShipmentGatewayDhlMutation {
	return sgdc.mutation
}

// Save creates the ShipmentGatewayDhl in the database.
func (sgdc *ShipmentGatewayDhlCreate) Save(ctx context.Context) (*ShipmentGatewayDhl, error) {
	var (
		err  error
		node *ShipmentGatewayDhl
	)
	sgdc.defaults()
	if len(sgdc.hooks) == 0 {
		if err = sgdc.check(); err != nil {
			return nil, err
		}
		node, err = sgdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShipmentGatewayDhlMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sgdc.check(); err != nil {
				return nil, err
			}
			sgdc.mutation = mutation
			if node, err = sgdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sgdc.hooks) - 1; i >= 0; i-- {
			mut = sgdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sgdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sgdc *ShipmentGatewayDhlCreate) SaveX(ctx context.Context) *ShipmentGatewayDhl {
	v, err := sgdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (sgdc *ShipmentGatewayDhlCreate) defaults() {
	if _, ok := sgdc.mutation.CreateTime(); !ok {
		v := shipmentgatewaydhl.DefaultCreateTime()
		sgdc.mutation.SetCreateTime(v)
	}
	if _, ok := sgdc.mutation.UpdateTime(); !ok {
		v := shipmentgatewaydhl.DefaultUpdateTime()
		sgdc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sgdc *ShipmentGatewayDhlCreate) check() error {
	if _, ok := sgdc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := sgdc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	return nil
}

func (sgdc *ShipmentGatewayDhlCreate) sqlSave(ctx context.Context) (*ShipmentGatewayDhl, error) {
	_node, _spec := sgdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sgdc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sgdc *ShipmentGatewayDhlCreate) createSpec() (*ShipmentGatewayDhl, *sqlgraph.CreateSpec) {
	var (
		_node = &ShipmentGatewayDhl{config: sgdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: shipmentgatewaydhl.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shipmentgatewaydhl.FieldID,
			},
		}
	)
	if value, ok := sgdc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipmentgatewaydhl.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := sgdc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipmentgatewaydhl.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := sgdc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewaydhl.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := sgdc.mutation.ConnectURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewaydhl.FieldConnectURL,
		})
		_node.ConnectURL = value
	}
	if value, ok := sgdc.mutation.ConnectTimeout(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipmentgatewaydhl.FieldConnectTimeout,
		})
		_node.ConnectTimeout = value
	}
	if value, ok := sgdc.mutation.HeadVersion(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewaydhl.FieldHeadVersion,
		})
		_node.HeadVersion = value
	}
	if value, ok := sgdc.mutation.HeadAction(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewaydhl.FieldHeadAction,
		})
		_node.HeadAction = value
	}
	if value, ok := sgdc.mutation.AccessUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewaydhl.FieldAccessUserID,
		})
		_node.AccessUserID = value
	}
	if value, ok := sgdc.mutation.AccessPassword(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewaydhl.FieldAccessPassword,
		})
		_node.AccessPassword = value
	}
	if value, ok := sgdc.mutation.AccessAccountNbr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewaydhl.FieldAccessAccountNbr,
		})
		_node.AccessAccountNbr = value
	}
	if value, ok := sgdc.mutation.AccessShippingKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewaydhl.FieldAccessShippingKey,
		})
		_node.AccessShippingKey = value
	}
	if value, ok := sgdc.mutation.LabelImageFormat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewaydhl.FieldLabelImageFormat,
		})
		_node.LabelImageFormat = value
	}
	if value, ok := sgdc.mutation.RateEstimateTemplate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentgatewaydhl.FieldRateEstimateTemplate,
		})
		_node.RateEstimateTemplate = value
	}
	if nodes := sgdc.mutation.ShipmentGatewayConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   shipmentgatewaydhl.ShipmentGatewayConfigTable,
			Columns: []string{shipmentgatewaydhl.ShipmentGatewayConfigColumn},
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
		_node.shipment_gateway_config_shipment_gateway_dhl = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ShipmentGatewayDhlCreateBulk is the builder for creating many ShipmentGatewayDhl entities in bulk.
type ShipmentGatewayDhlCreateBulk struct {
	config
	builders []*ShipmentGatewayDhlCreate
}

// Save creates the ShipmentGatewayDhl entities in the database.
func (sgdcb *ShipmentGatewayDhlCreateBulk) Save(ctx context.Context) ([]*ShipmentGatewayDhl, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sgdcb.builders))
	nodes := make([]*ShipmentGatewayDhl, len(sgdcb.builders))
	mutators := make([]Mutator, len(sgdcb.builders))
	for i := range sgdcb.builders {
		func(i int, root context.Context) {
			builder := sgdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShipmentGatewayDhlMutation)
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
					_, err = mutators[i+1].Mutate(root, sgdcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sgdcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sgdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sgdcb *ShipmentGatewayDhlCreateBulk) SaveX(ctx context.Context) []*ShipmentGatewayDhl {
	v, err := sgdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
