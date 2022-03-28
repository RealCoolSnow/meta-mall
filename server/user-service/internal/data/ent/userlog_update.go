// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"user-service/internal/data/ent/predicate"
	"user-service/internal/data/ent/userlog"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserLogUpdate is the builder for updating UserLog entities.
type UserLogUpdate struct {
	config
	hooks    []Hook
	mutation *UserLogMutation
}

// Where appends a list predicates to the UserLogUpdate builder.
func (ulu *UserLogUpdate) Where(ps ...predicate.UserLog) *UserLogUpdate {
	ulu.mutation.Where(ps...)
	return ulu
}

// SetUserID sets the "user_id" field.
func (ulu *UserLogUpdate) SetUserID(i int64) *UserLogUpdate {
	ulu.mutation.ResetUserID()
	ulu.mutation.SetUserID(i)
	return ulu
}

// AddUserID adds i to the "user_id" field.
func (ulu *UserLogUpdate) AddUserID(i int64) *UserLogUpdate {
	ulu.mutation.AddUserID(i)
	return ulu
}

// SetIP sets the "ip" field.
func (ulu *UserLogUpdate) SetIP(s string) *UserLogUpdate {
	ulu.mutation.SetIP(s)
	return ulu
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (ulu *UserLogUpdate) SetNillableIP(s *string) *UserLogUpdate {
	if s != nil {
		ulu.SetIP(*s)
	}
	return ulu
}

// SetCreateTime sets the "create_time" field.
func (ulu *UserLogUpdate) SetCreateTime(t time.Time) *UserLogUpdate {
	ulu.mutation.SetCreateTime(t)
	return ulu
}

// Mutation returns the UserLogMutation object of the builder.
func (ulu *UserLogUpdate) Mutation() *UserLogMutation {
	return ulu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ulu *UserLogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ulu.defaults()
	if len(ulu.hooks) == 0 {
		affected, err = ulu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ulu.mutation = mutation
			affected, err = ulu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ulu.hooks) - 1; i >= 0; i-- {
			if ulu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ulu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ulu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ulu *UserLogUpdate) SaveX(ctx context.Context) int {
	affected, err := ulu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ulu *UserLogUpdate) Exec(ctx context.Context) error {
	_, err := ulu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulu *UserLogUpdate) ExecX(ctx context.Context) {
	if err := ulu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ulu *UserLogUpdate) defaults() {
	if _, ok := ulu.mutation.CreateTime(); !ok {
		v := userlog.UpdateDefaultCreateTime()
		ulu.mutation.SetCreateTime(v)
	}
}

func (ulu *UserLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userlog.Table,
			Columns: userlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userlog.FieldID,
			},
		},
	}
	if ps := ulu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ulu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userlog.FieldUserID,
		})
	}
	if value, ok := ulu.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userlog.FieldUserID,
		})
	}
	if value, ok := ulu.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlog.FieldIP,
		})
	}
	if value, ok := ulu.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userlog.FieldCreateTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ulu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserLogUpdateOne is the builder for updating a single UserLog entity.
type UserLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserLogMutation
}

// SetUserID sets the "user_id" field.
func (uluo *UserLogUpdateOne) SetUserID(i int64) *UserLogUpdateOne {
	uluo.mutation.ResetUserID()
	uluo.mutation.SetUserID(i)
	return uluo
}

// AddUserID adds i to the "user_id" field.
func (uluo *UserLogUpdateOne) AddUserID(i int64) *UserLogUpdateOne {
	uluo.mutation.AddUserID(i)
	return uluo
}

// SetIP sets the "ip" field.
func (uluo *UserLogUpdateOne) SetIP(s string) *UserLogUpdateOne {
	uluo.mutation.SetIP(s)
	return uluo
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (uluo *UserLogUpdateOne) SetNillableIP(s *string) *UserLogUpdateOne {
	if s != nil {
		uluo.SetIP(*s)
	}
	return uluo
}

// SetCreateTime sets the "create_time" field.
func (uluo *UserLogUpdateOne) SetCreateTime(t time.Time) *UserLogUpdateOne {
	uluo.mutation.SetCreateTime(t)
	return uluo
}

// Mutation returns the UserLogMutation object of the builder.
func (uluo *UserLogUpdateOne) Mutation() *UserLogMutation {
	return uluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uluo *UserLogUpdateOne) Select(field string, fields ...string) *UserLogUpdateOne {
	uluo.fields = append([]string{field}, fields...)
	return uluo
}

// Save executes the query and returns the updated UserLog entity.
func (uluo *UserLogUpdateOne) Save(ctx context.Context) (*UserLog, error) {
	var (
		err  error
		node *UserLog
	)
	uluo.defaults()
	if len(uluo.hooks) == 0 {
		node, err = uluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uluo.mutation = mutation
			node, err = uluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uluo.hooks) - 1; i >= 0; i-- {
			if uluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uluo *UserLogUpdateOne) SaveX(ctx context.Context) *UserLog {
	node, err := uluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uluo *UserLogUpdateOne) Exec(ctx context.Context) error {
	_, err := uluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uluo *UserLogUpdateOne) ExecX(ctx context.Context) {
	if err := uluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uluo *UserLogUpdateOne) defaults() {
	if _, ok := uluo.mutation.CreateTime(); !ok {
		v := userlog.UpdateDefaultCreateTime()
		uluo.mutation.SetCreateTime(v)
	}
}

func (uluo *UserLogUpdateOne) sqlSave(ctx context.Context) (_node *UserLog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userlog.Table,
			Columns: userlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userlog.FieldID,
			},
		},
	}
	id, ok := uluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userlog.FieldID)
		for _, f := range fields {
			if !userlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uluo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userlog.FieldUserID,
		})
	}
	if value, ok := uluo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userlog.FieldUserID,
		})
	}
	if value, ok := uluo.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlog.FieldIP,
		})
	}
	if value, ok := uluo.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userlog.FieldCreateTime,
		})
	}
	_node = &UserLog{config: uluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}