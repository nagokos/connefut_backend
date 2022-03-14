// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nagokos/connefut_backend/ent/predicate"
	"github.com/nagokos/connefut_backend/ent/recruitment"
	"github.com/nagokos/connefut_backend/ent/recruitmenttag"
	"github.com/nagokos/connefut_backend/ent/tag"
)

// RecruitmentTagUpdate is the builder for updating RecruitmentTag entities.
type RecruitmentTagUpdate struct {
	config
	hooks    []Hook
	mutation *RecruitmentTagMutation
}

// Where appends a list predicates to the RecruitmentTagUpdate builder.
func (rtu *RecruitmentTagUpdate) Where(ps ...predicate.RecruitmentTag) *RecruitmentTagUpdate {
	rtu.mutation.Where(ps...)
	return rtu
}

// SetUpdatedAt sets the "updated_at" field.
func (rtu *RecruitmentTagUpdate) SetUpdatedAt(t time.Time) *RecruitmentTagUpdate {
	rtu.mutation.SetUpdatedAt(t)
	return rtu
}

// SetRecruitmentID sets the "recruitment_id" field.
func (rtu *RecruitmentTagUpdate) SetRecruitmentID(s string) *RecruitmentTagUpdate {
	rtu.mutation.SetRecruitmentID(s)
	return rtu
}

// SetTagID sets the "tag_id" field.
func (rtu *RecruitmentTagUpdate) SetTagID(s string) *RecruitmentTagUpdate {
	rtu.mutation.SetTagID(s)
	return rtu
}

// SetRecruitment sets the "recruitment" edge to the Recruitment entity.
func (rtu *RecruitmentTagUpdate) SetRecruitment(r *Recruitment) *RecruitmentTagUpdate {
	return rtu.SetRecruitmentID(r.ID)
}

// SetTag sets the "tag" edge to the Tag entity.
func (rtu *RecruitmentTagUpdate) SetTag(t *Tag) *RecruitmentTagUpdate {
	return rtu.SetTagID(t.ID)
}

// Mutation returns the RecruitmentTagMutation object of the builder.
func (rtu *RecruitmentTagUpdate) Mutation() *RecruitmentTagMutation {
	return rtu.mutation
}

// ClearRecruitment clears the "recruitment" edge to the Recruitment entity.
func (rtu *RecruitmentTagUpdate) ClearRecruitment() *RecruitmentTagUpdate {
	rtu.mutation.ClearRecruitment()
	return rtu
}

// ClearTag clears the "tag" edge to the Tag entity.
func (rtu *RecruitmentTagUpdate) ClearTag() *RecruitmentTagUpdate {
	rtu.mutation.ClearTag()
	return rtu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rtu *RecruitmentTagUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	rtu.defaults()
	if len(rtu.hooks) == 0 {
		if err = rtu.check(); err != nil {
			return 0, err
		}
		affected, err = rtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecruitmentTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rtu.check(); err != nil {
				return 0, err
			}
			rtu.mutation = mutation
			affected, err = rtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rtu.hooks) - 1; i >= 0; i-- {
			if rtu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtu *RecruitmentTagUpdate) SaveX(ctx context.Context) int {
	affected, err := rtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rtu *RecruitmentTagUpdate) Exec(ctx context.Context) error {
	_, err := rtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtu *RecruitmentTagUpdate) ExecX(ctx context.Context) {
	if err := rtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rtu *RecruitmentTagUpdate) defaults() {
	if _, ok := rtu.mutation.UpdatedAt(); !ok {
		v := recruitmenttag.UpdateDefaultUpdatedAt()
		rtu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtu *RecruitmentTagUpdate) check() error {
	if _, ok := rtu.mutation.RecruitmentID(); rtu.mutation.RecruitmentCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"recruitment\"")
	}
	if _, ok := rtu.mutation.TagID(); rtu.mutation.TagCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"tag\"")
	}
	return nil
}

func (rtu *RecruitmentTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recruitmenttag.Table,
			Columns: recruitmenttag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: recruitmenttag.FieldID,
			},
		},
	}
	if ps := rtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: recruitmenttag.FieldUpdatedAt,
		})
	}
	if rtu.mutation.RecruitmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recruitmenttag.RecruitmentTable,
			Columns: []string{recruitmenttag.RecruitmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: recruitment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtu.mutation.RecruitmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recruitmenttag.RecruitmentTable,
			Columns: []string{recruitmenttag.RecruitmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: recruitment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rtu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recruitmenttag.TagTable,
			Columns: []string{recruitmenttag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recruitmenttag.TagTable,
			Columns: []string{recruitmenttag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recruitmenttag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RecruitmentTagUpdateOne is the builder for updating a single RecruitmentTag entity.
type RecruitmentTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecruitmentTagMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (rtuo *RecruitmentTagUpdateOne) SetUpdatedAt(t time.Time) *RecruitmentTagUpdateOne {
	rtuo.mutation.SetUpdatedAt(t)
	return rtuo
}

// SetRecruitmentID sets the "recruitment_id" field.
func (rtuo *RecruitmentTagUpdateOne) SetRecruitmentID(s string) *RecruitmentTagUpdateOne {
	rtuo.mutation.SetRecruitmentID(s)
	return rtuo
}

// SetTagID sets the "tag_id" field.
func (rtuo *RecruitmentTagUpdateOne) SetTagID(s string) *RecruitmentTagUpdateOne {
	rtuo.mutation.SetTagID(s)
	return rtuo
}

// SetRecruitment sets the "recruitment" edge to the Recruitment entity.
func (rtuo *RecruitmentTagUpdateOne) SetRecruitment(r *Recruitment) *RecruitmentTagUpdateOne {
	return rtuo.SetRecruitmentID(r.ID)
}

// SetTag sets the "tag" edge to the Tag entity.
func (rtuo *RecruitmentTagUpdateOne) SetTag(t *Tag) *RecruitmentTagUpdateOne {
	return rtuo.SetTagID(t.ID)
}

// Mutation returns the RecruitmentTagMutation object of the builder.
func (rtuo *RecruitmentTagUpdateOne) Mutation() *RecruitmentTagMutation {
	return rtuo.mutation
}

// ClearRecruitment clears the "recruitment" edge to the Recruitment entity.
func (rtuo *RecruitmentTagUpdateOne) ClearRecruitment() *RecruitmentTagUpdateOne {
	rtuo.mutation.ClearRecruitment()
	return rtuo
}

// ClearTag clears the "tag" edge to the Tag entity.
func (rtuo *RecruitmentTagUpdateOne) ClearTag() *RecruitmentTagUpdateOne {
	rtuo.mutation.ClearTag()
	return rtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rtuo *RecruitmentTagUpdateOne) Select(field string, fields ...string) *RecruitmentTagUpdateOne {
	rtuo.fields = append([]string{field}, fields...)
	return rtuo
}

// Save executes the query and returns the updated RecruitmentTag entity.
func (rtuo *RecruitmentTagUpdateOne) Save(ctx context.Context) (*RecruitmentTag, error) {
	var (
		err  error
		node *RecruitmentTag
	)
	rtuo.defaults()
	if len(rtuo.hooks) == 0 {
		if err = rtuo.check(); err != nil {
			return nil, err
		}
		node, err = rtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecruitmentTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rtuo.check(); err != nil {
				return nil, err
			}
			rtuo.mutation = mutation
			node, err = rtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rtuo.hooks) - 1; i >= 0; i-- {
			if rtuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rtuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtuo *RecruitmentTagUpdateOne) SaveX(ctx context.Context) *RecruitmentTag {
	node, err := rtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rtuo *RecruitmentTagUpdateOne) Exec(ctx context.Context) error {
	_, err := rtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtuo *RecruitmentTagUpdateOne) ExecX(ctx context.Context) {
	if err := rtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rtuo *RecruitmentTagUpdateOne) defaults() {
	if _, ok := rtuo.mutation.UpdatedAt(); !ok {
		v := recruitmenttag.UpdateDefaultUpdatedAt()
		rtuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtuo *RecruitmentTagUpdateOne) check() error {
	if _, ok := rtuo.mutation.RecruitmentID(); rtuo.mutation.RecruitmentCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"recruitment\"")
	}
	if _, ok := rtuo.mutation.TagID(); rtuo.mutation.TagCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"tag\"")
	}
	return nil
}

func (rtuo *RecruitmentTagUpdateOne) sqlSave(ctx context.Context) (_node *RecruitmentTag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recruitmenttag.Table,
			Columns: recruitmenttag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: recruitmenttag.FieldID,
			},
		},
	}
	id, ok := rtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing RecruitmentTag.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := rtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recruitmenttag.FieldID)
		for _, f := range fields {
			if !recruitmenttag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != recruitmenttag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: recruitmenttag.FieldUpdatedAt,
		})
	}
	if rtuo.mutation.RecruitmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recruitmenttag.RecruitmentTable,
			Columns: []string{recruitmenttag.RecruitmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: recruitment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtuo.mutation.RecruitmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recruitmenttag.RecruitmentTable,
			Columns: []string{recruitmenttag.RecruitmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: recruitment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rtuo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recruitmenttag.TagTable,
			Columns: []string{recruitmenttag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtuo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recruitmenttag.TagTable,
			Columns: []string{recruitmenttag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RecruitmentTag{config: rtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recruitmenttag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}