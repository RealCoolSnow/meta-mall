// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"user-service/internal/data/ent/predicate"
	"user-service/internal/data/ent/usergroup"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserGroupUpdate is the builder for updating UserGroup entities.
type UserGroupUpdate struct {
	config
	hooks    []Hook
	mutation *UserGroupMutation
}

// Where appends a list predicates to the UserGroupUpdate builder.
func (ugu *UserGroupUpdate) Where(ps ...predicate.UserGroup) *UserGroupUpdate {
	ugu.mutation.Where(ps...)
	return ugu
}

// SetName sets the "name" field.
func (ugu *UserGroupUpdate) SetName(s string) *UserGroupUpdate {
	ugu.mutation.SetName(s)
	return ugu
}

// SetCreateTime sets the "create_time" field.
func (ugu *UserGroupUpdate) SetCreateTime(t time.Time) *UserGroupUpdate {
	ugu.mutation.SetCreateTime(t)
	return ugu
}

// Mutation returns the UserGroupMutation object of the builder.
func (ugu *UserGroupUpdate) Mutation() *UserGroupMutation {
	return ugu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ugu *UserGroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ugu.defaults()
	if len(ugu.hooks) == 0 {
		affected, err = ugu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ugu.mutation = mutation
			affected, err = ugu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ugu.hooks) - 1; i >= 0; i-- {
			if ugu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ugu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ugu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ugu *UserGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := ugu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ugu *UserGroupUpdate) Exec(ctx context.Context) error {
	_, err := ugu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugu *UserGroupUpdate) ExecX(ctx context.Context) {
	if err := ugu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ugu *UserGroupUpdate) defaults() {
	if _, ok := ugu.mutation.CreateTime(); !ok {
		v := usergroup.UpdateDefaultCreateTime()
		ugu.mutation.SetCreateTime(v)
	}
}

func (ugu *UserGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usergroup.Table,
			Columns: usergroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: usergroup.FieldID,
			},
		},
	}
	if ps := ugu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ugu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usergroup.FieldName,
		})
	}
	if value, ok := ugu.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usergroup.FieldCreateTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ugu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usergroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserGroupUpdateOne is the builder for updating a single UserGroup entity.
type UserGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserGroupMutation
}

// SetName sets the "name" field.
func (uguo *UserGroupUpdateOne) SetName(s string) *UserGroupUpdateOne {
	uguo.mutation.SetName(s)
	return uguo
}

// SetCreateTime sets the "create_time" field.
func (uguo *UserGroupUpdateOne) SetCreateTime(t time.Time) *UserGroupUpdateOne {
	uguo.mutation.SetCreateTime(t)
	return uguo
}

// Mutation returns the UserGroupMutation object of the builder.
func (uguo *UserGroupUpdateOne) Mutation() *UserGroupMutation {
	return uguo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uguo *UserGroupUpdateOne) Select(field string, fields ...string) *UserGroupUpdateOne {
	uguo.fields = append([]string{field}, fields...)
	return uguo
}

// Save executes the query and returns the updated UserGroup entity.
func (uguo *UserGroupUpdateOne) Save(ctx context.Context) (*UserGroup, error) {
	var (
		err  error
		node *UserGroup
	)
	uguo.defaults()
	if len(uguo.hooks) == 0 {
		node, err = uguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uguo.mutation = mutation
			node, err = uguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uguo.hooks) - 1; i >= 0; i-- {
			if uguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uguo *UserGroupUpdateOne) SaveX(ctx context.Context) *UserGroup {
	node, err := uguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uguo *UserGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := uguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uguo *UserGroupUpdateOne) ExecX(ctx context.Context) {
	if err := uguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uguo *UserGroupUpdateOne) defaults() {
	if _, ok := uguo.mutation.CreateTime(); !ok {
		v := usergroup.UpdateDefaultCreateTime()
		uguo.mutation.SetCreateTime(v)
	}
}

func (uguo *UserGroupUpdateOne) sqlSave(ctx context.Context) (_node *UserGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usergroup.Table,
			Columns: usergroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: usergroup.FieldID,
			},
		},
	}
	id, ok := uguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usergroup.FieldID)
		for _, f := range fields {
			if !usergroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usergroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uguo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usergroup.FieldName,
		})
	}
	if value, ok := uguo.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usergroup.FieldCreateTime,
		})
	}
	_node = &UserGroup{config: uguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usergroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
