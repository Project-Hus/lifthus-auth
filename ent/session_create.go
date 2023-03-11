// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lifthus-auth/ent/session"
	"lifthus-auth/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SessionCreate is the builder for creating a Session entity.
type SessionCreate struct {
	config
	mutation *SessionMutation
	hooks    []Hook
}

// SetUID sets the "uid" field.
func (sc *SessionCreate) SetUID(u uuid.UUID) *SessionCreate {
	sc.mutation.SetUID(u)
	return sc
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (sc *SessionCreate) SetNillableUID(u *uuid.UUID) *SessionCreate {
	if u != nil {
		sc.SetUID(*u)
	}
	return sc
}

// SetConnectedAt sets the "connected_at" field.
func (sc *SessionCreate) SetConnectedAt(t time.Time) *SessionCreate {
	sc.mutation.SetConnectedAt(t)
	return sc
}

// SetNillableConnectedAt sets the "connected_at" field if the given value is not nil.
func (sc *SessionCreate) SetNillableConnectedAt(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetConnectedAt(*t)
	}
	return sc
}

// SetSignedAt sets the "signed_at" field.
func (sc *SessionCreate) SetSignedAt(t time.Time) *SessionCreate {
	sc.mutation.SetSignedAt(t)
	return sc
}

// SetNillableSignedAt sets the "signed_at" field if the given value is not nil.
func (sc *SessionCreate) SetNillableSignedAt(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetSignedAt(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SessionCreate) SetID(u uuid.UUID) *SessionCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SessionCreate) SetNillableID(u *uuid.UUID) *SessionCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (sc *SessionCreate) SetUserID(id uuid.UUID) *SessionCreate {
	sc.mutation.SetUserID(id)
	return sc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (sc *SessionCreate) SetNillableUserID(id *uuid.UUID) *SessionCreate {
	if id != nil {
		sc = sc.SetUserID(*id)
	}
	return sc
}

// SetUser sets the "user" edge to the User entity.
func (sc *SessionCreate) SetUser(u *User) *SessionCreate {
	return sc.SetUserID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (sc *SessionCreate) Mutation() *SessionMutation {
	return sc.mutation
}

// Save creates the Session in the database.
func (sc *SessionCreate) Save(ctx context.Context) (*Session, error) {
	sc.defaults()
	return withHooks[*Session, SessionMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SessionCreate) SaveX(ctx context.Context) *Session {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SessionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SessionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SessionCreate) defaults() {
	if _, ok := sc.mutation.ConnectedAt(); !ok {
		v := session.DefaultConnectedAt()
		sc.mutation.SetConnectedAt(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := session.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SessionCreate) check() error {
	if _, ok := sc.mutation.ConnectedAt(); !ok {
		return &ValidationError{Name: "connected_at", err: errors.New(`ent: missing required field "Session.connected_at"`)}
	}
	return nil
}

func (sc *SessionCreate) sqlSave(ctx context.Context) (*Session, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SessionCreate) createSpec() (*Session, *sqlgraph.CreateSpec) {
	var (
		_node = &Session{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(session.Table, sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.ConnectedAt(); ok {
		_spec.SetField(session.FieldConnectedAt, field.TypeTime, value)
		_node.ConnectedAt = value
	}
	if value, ok := sc.mutation.SignedAt(); ok {
		_spec.SetField(session.FieldSignedAt, field.TypeTime, value)
		_node.SignedAt = &value
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SessionCreateBulk is the builder for creating many Session entities in bulk.
type SessionCreateBulk struct {
	config
	builders []*SessionCreate
}

// Save creates the Session entities in the database.
func (scb *SessionCreateBulk) Save(ctx context.Context) ([]*Session, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Session, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SessionMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SessionCreateBulk) SaveX(ctx context.Context) []*Session {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SessionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SessionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
