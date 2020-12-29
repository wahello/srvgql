// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/tmc/pulid"
	"github.com/tmc/srvgql/ent/campaign"
)

// CampaignCreate is the builder for creating a Campaign entity.
type CampaignCreate struct {
	config
	mutation *CampaignMutation
	hooks    []Hook
}

// SetName sets the name field.
func (cc *CampaignCreate) SetName(s string) *CampaignCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetID sets the id field.
func (cc *CampaignCreate) SetID(pu pulid.ID) *CampaignCreate {
	cc.mutation.SetID(pu)
	return cc
}

// Mutation returns the CampaignMutation object of the builder.
func (cc *CampaignCreate) Mutation() *CampaignMutation {
	return cc.mutation
}

// Save creates the Campaign in the database.
func (cc *CampaignCreate) Save(ctx context.Context) (*Campaign, error) {
	var (
		err  error
		node *Campaign
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CampaignMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CampaignCreate) SaveX(ctx context.Context) *Campaign {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (cc *CampaignCreate) defaults() {
	if _, ok := cc.mutation.ID(); !ok {
		v := campaign.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CampaignCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	return nil
}

func (cc *CampaignCreate) sqlSave(ctx context.Context) (*Campaign, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (cc *CampaignCreate) createSpec() (*Campaign, *sqlgraph.CreateSpec) {
	var (
		_node = &Campaign{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: campaign.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: campaign.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: campaign.FieldName,
		})
		_node.Name = value
	}
	return _node, _spec
}

// CampaignCreateBulk is the builder for creating a bulk of Campaign entities.
type CampaignCreateBulk struct {
	config
	builders []*CampaignCreate
}

// Save creates the Campaign entities in the database.
func (ccb *CampaignCreateBulk) Save(ctx context.Context) ([]*Campaign, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Campaign, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CampaignMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ccb *CampaignCreateBulk) SaveX(ctx context.Context) []*Campaign {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}