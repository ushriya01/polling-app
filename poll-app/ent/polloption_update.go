// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"poll-app/ent/poll"
	"poll-app/ent/polloption"
	"poll-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PollOptionUpdate is the builder for updating PollOption entities.
type PollOptionUpdate struct {
	config
	hooks    []Hook
	mutation *PollOptionMutation
}

// Where appends a list predicates to the PollOptionUpdate builder.
func (pou *PollOptionUpdate) Where(ps ...predicate.PollOption) *PollOptionUpdate {
	pou.mutation.Where(ps...)
	return pou
}

// SetText sets the "text" field.
func (pou *PollOptionUpdate) SetText(s string) *PollOptionUpdate {
	pou.mutation.SetText(s)
	return pou
}

// SetNillableText sets the "text" field if the given value is not nil.
func (pou *PollOptionUpdate) SetNillableText(s *string) *PollOptionUpdate {
	if s != nil {
		pou.SetText(*s)
	}
	return pou
}

// SetVotes sets the "votes" field.
func (pou *PollOptionUpdate) SetVotes(i int) *PollOptionUpdate {
	pou.mutation.ResetVotes()
	pou.mutation.SetVotes(i)
	return pou
}

// SetNillableVotes sets the "votes" field if the given value is not nil.
func (pou *PollOptionUpdate) SetNillableVotes(i *int) *PollOptionUpdate {
	if i != nil {
		pou.SetVotes(*i)
	}
	return pou
}

// AddVotes adds i to the "votes" field.
func (pou *PollOptionUpdate) AddVotes(i int) *PollOptionUpdate {
	pou.mutation.AddVotes(i)
	return pou
}

// SetIsActive sets the "is_active" field.
func (pou *PollOptionUpdate) SetIsActive(b bool) *PollOptionUpdate {
	pou.mutation.SetIsActive(b)
	return pou
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (pou *PollOptionUpdate) SetNillableIsActive(b *bool) *PollOptionUpdate {
	if b != nil {
		pou.SetIsActive(*b)
	}
	return pou
}

// SetPollID sets the "poll" edge to the Poll entity by ID.
func (pou *PollOptionUpdate) SetPollID(id int) *PollOptionUpdate {
	pou.mutation.SetPollID(id)
	return pou
}

// SetNillablePollID sets the "poll" edge to the Poll entity by ID if the given value is not nil.
func (pou *PollOptionUpdate) SetNillablePollID(id *int) *PollOptionUpdate {
	if id != nil {
		pou = pou.SetPollID(*id)
	}
	return pou
}

// SetPoll sets the "poll" edge to the Poll entity.
func (pou *PollOptionUpdate) SetPoll(p *Poll) *PollOptionUpdate {
	return pou.SetPollID(p.ID)
}

// Mutation returns the PollOptionMutation object of the builder.
func (pou *PollOptionUpdate) Mutation() *PollOptionMutation {
	return pou.mutation
}

// ClearPoll clears the "poll" edge to the Poll entity.
func (pou *PollOptionUpdate) ClearPoll() *PollOptionUpdate {
	pou.mutation.ClearPoll()
	return pou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pou *PollOptionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pou.sqlSave, pou.mutation, pou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pou *PollOptionUpdate) SaveX(ctx context.Context) int {
	affected, err := pou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pou *PollOptionUpdate) Exec(ctx context.Context) error {
	_, err := pou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pou *PollOptionUpdate) ExecX(ctx context.Context) {
	if err := pou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pou *PollOptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(polloption.Table, polloption.Columns, sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt))
	if ps := pou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pou.mutation.Text(); ok {
		_spec.SetField(polloption.FieldText, field.TypeString, value)
	}
	if value, ok := pou.mutation.Votes(); ok {
		_spec.SetField(polloption.FieldVotes, field.TypeInt, value)
	}
	if value, ok := pou.mutation.AddedVotes(); ok {
		_spec.AddField(polloption.FieldVotes, field.TypeInt, value)
	}
	if value, ok := pou.mutation.IsActive(); ok {
		_spec.SetField(polloption.FieldIsActive, field.TypeBool, value)
	}
	if pou.mutation.PollCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   polloption.PollTable,
			Columns: []string{polloption.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pou.mutation.PollIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   polloption.PollTable,
			Columns: []string{polloption.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{polloption.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pou.mutation.done = true
	return n, nil
}

// PollOptionUpdateOne is the builder for updating a single PollOption entity.
type PollOptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PollOptionMutation
}

// SetText sets the "text" field.
func (pouo *PollOptionUpdateOne) SetText(s string) *PollOptionUpdateOne {
	pouo.mutation.SetText(s)
	return pouo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (pouo *PollOptionUpdateOne) SetNillableText(s *string) *PollOptionUpdateOne {
	if s != nil {
		pouo.SetText(*s)
	}
	return pouo
}

// SetVotes sets the "votes" field.
func (pouo *PollOptionUpdateOne) SetVotes(i int) *PollOptionUpdateOne {
	pouo.mutation.ResetVotes()
	pouo.mutation.SetVotes(i)
	return pouo
}

// SetNillableVotes sets the "votes" field if the given value is not nil.
func (pouo *PollOptionUpdateOne) SetNillableVotes(i *int) *PollOptionUpdateOne {
	if i != nil {
		pouo.SetVotes(*i)
	}
	return pouo
}

// AddVotes adds i to the "votes" field.
func (pouo *PollOptionUpdateOne) AddVotes(i int) *PollOptionUpdateOne {
	pouo.mutation.AddVotes(i)
	return pouo
}

// SetIsActive sets the "is_active" field.
func (pouo *PollOptionUpdateOne) SetIsActive(b bool) *PollOptionUpdateOne {
	pouo.mutation.SetIsActive(b)
	return pouo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (pouo *PollOptionUpdateOne) SetNillableIsActive(b *bool) *PollOptionUpdateOne {
	if b != nil {
		pouo.SetIsActive(*b)
	}
	return pouo
}

// SetPollID sets the "poll" edge to the Poll entity by ID.
func (pouo *PollOptionUpdateOne) SetPollID(id int) *PollOptionUpdateOne {
	pouo.mutation.SetPollID(id)
	return pouo
}

// SetNillablePollID sets the "poll" edge to the Poll entity by ID if the given value is not nil.
func (pouo *PollOptionUpdateOne) SetNillablePollID(id *int) *PollOptionUpdateOne {
	if id != nil {
		pouo = pouo.SetPollID(*id)
	}
	return pouo
}

// SetPoll sets the "poll" edge to the Poll entity.
func (pouo *PollOptionUpdateOne) SetPoll(p *Poll) *PollOptionUpdateOne {
	return pouo.SetPollID(p.ID)
}

// Mutation returns the PollOptionMutation object of the builder.
func (pouo *PollOptionUpdateOne) Mutation() *PollOptionMutation {
	return pouo.mutation
}

// ClearPoll clears the "poll" edge to the Poll entity.
func (pouo *PollOptionUpdateOne) ClearPoll() *PollOptionUpdateOne {
	pouo.mutation.ClearPoll()
	return pouo
}

// Where appends a list predicates to the PollOptionUpdate builder.
func (pouo *PollOptionUpdateOne) Where(ps ...predicate.PollOption) *PollOptionUpdateOne {
	pouo.mutation.Where(ps...)
	return pouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pouo *PollOptionUpdateOne) Select(field string, fields ...string) *PollOptionUpdateOne {
	pouo.fields = append([]string{field}, fields...)
	return pouo
}

// Save executes the query and returns the updated PollOption entity.
func (pouo *PollOptionUpdateOne) Save(ctx context.Context) (*PollOption, error) {
	return withHooks(ctx, pouo.sqlSave, pouo.mutation, pouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pouo *PollOptionUpdateOne) SaveX(ctx context.Context) *PollOption {
	node, err := pouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pouo *PollOptionUpdateOne) Exec(ctx context.Context) error {
	_, err := pouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pouo *PollOptionUpdateOne) ExecX(ctx context.Context) {
	if err := pouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pouo *PollOptionUpdateOne) sqlSave(ctx context.Context) (_node *PollOption, err error) {
	_spec := sqlgraph.NewUpdateSpec(polloption.Table, polloption.Columns, sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt))
	id, ok := pouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PollOption.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, polloption.FieldID)
		for _, f := range fields {
			if !polloption.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != polloption.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pouo.mutation.Text(); ok {
		_spec.SetField(polloption.FieldText, field.TypeString, value)
	}
	if value, ok := pouo.mutation.Votes(); ok {
		_spec.SetField(polloption.FieldVotes, field.TypeInt, value)
	}
	if value, ok := pouo.mutation.AddedVotes(); ok {
		_spec.AddField(polloption.FieldVotes, field.TypeInt, value)
	}
	if value, ok := pouo.mutation.IsActive(); ok {
		_spec.SetField(polloption.FieldIsActive, field.TypeBool, value)
	}
	if pouo.mutation.PollCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   polloption.PollTable,
			Columns: []string{polloption.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pouo.mutation.PollIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   polloption.PollTable,
			Columns: []string{polloption.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PollOption{config: pouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{polloption.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pouo.mutation.done = true
	return _node, nil
}