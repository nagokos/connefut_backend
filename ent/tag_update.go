// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nagokos/connefut_backend/ent/predicate"
	"github.com/nagokos/connefut_backend/ent/recruitmenttag"
	"github.com/nagokos/connefut_backend/ent/tag"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	hooks    []Hook
	mutation *TagMutation
}

// Where appends a list predicates to the TagUpdate builder.
func (tu *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TagUpdate) SetUpdatedAt(t time.Time) *TagUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetName sets the "name" field.
func (tu *TagUpdate) SetName(s string) *TagUpdate {
	tu.mutation.SetName(s)
	return tu
}

// AddRecruitmentTagIDs adds the "recruitment_tags" edge to the RecruitmentTag entity by IDs.
func (tu *TagUpdate) AddRecruitmentTagIDs(ids ...string) *TagUpdate {
	tu.mutation.AddRecruitmentTagIDs(ids...)
	return tu
}

// AddRecruitmentTags adds the "recruitment_tags" edges to the RecruitmentTag entity.
func (tu *TagUpdate) AddRecruitmentTags(r ...*RecruitmentTag) *TagUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.AddRecruitmentTagIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tu *TagUpdate) Mutation() *TagMutation {
	return tu.mutation
}

// ClearRecruitmentTags clears all "recruitment_tags" edges to the RecruitmentTag entity.
func (tu *TagUpdate) ClearRecruitmentTags() *TagUpdate {
	tu.mutation.ClearRecruitmentTags()
	return tu
}

// RemoveRecruitmentTagIDs removes the "recruitment_tags" edge to RecruitmentTag entities by IDs.
func (tu *TagUpdate) RemoveRecruitmentTagIDs(ids ...string) *TagUpdate {
	tu.mutation.RemoveRecruitmentTagIDs(ids...)
	return tu
}

// RemoveRecruitmentTags removes "recruitment_tags" edges to RecruitmentTag entities.
func (tu *TagUpdate) RemoveRecruitmentTags(r ...*RecruitmentTag) *TagUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.RemoveRecruitmentTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TagUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TagUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := tag.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

func (tu *TagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tag.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldUpdatedAt,
		})
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldName,
		})
	}
	if tu.mutation.RecruitmentTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tag.RecruitmentTagsTable,
			Columns: []string{tag.RecruitmentTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: recruitmenttag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedRecruitmentTagsIDs(); len(nodes) > 0 && !tu.mutation.RecruitmentTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tag.RecruitmentTagsTable,
			Columns: []string{tag.RecruitmentTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: recruitmenttag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RecruitmentTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tag.RecruitmentTagsTable,
			Columns: []string{tag.RecruitmentTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: recruitmenttag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TagUpdateOne) SetUpdatedAt(t time.Time) *TagUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TagUpdateOne) SetName(s string) *TagUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// AddRecruitmentTagIDs adds the "recruitment_tags" edge to the RecruitmentTag entity by IDs.
func (tuo *TagUpdateOne) AddRecruitmentTagIDs(ids ...string) *TagUpdateOne {
	tuo.mutation.AddRecruitmentTagIDs(ids...)
	return tuo
}

// AddRecruitmentTags adds the "recruitment_tags" edges to the RecruitmentTag entity.
func (tuo *TagUpdateOne) AddRecruitmentTags(r ...*RecruitmentTag) *TagUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.AddRecruitmentTagIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tuo *TagUpdateOne) Mutation() *TagMutation {
	return tuo.mutation
}

// ClearRecruitmentTags clears all "recruitment_tags" edges to the RecruitmentTag entity.
func (tuo *TagUpdateOne) ClearRecruitmentTags() *TagUpdateOne {
	tuo.mutation.ClearRecruitmentTags()
	return tuo
}

// RemoveRecruitmentTagIDs removes the "recruitment_tags" edge to RecruitmentTag entities by IDs.
func (tuo *TagUpdateOne) RemoveRecruitmentTagIDs(ids ...string) *TagUpdateOne {
	tuo.mutation.RemoveRecruitmentTagIDs(ids...)
	return tuo
}

// RemoveRecruitmentTags removes "recruitment_tags" edges to RecruitmentTag entities.
func (tuo *TagUpdateOne) RemoveRecruitmentTags(r ...*RecruitmentTag) *TagUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.RemoveRecruitmentTagIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TagUpdateOne) Select(field string, fields ...string) *TagUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tag entity.
func (tuo *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	var (
		err  error
		node *Tag
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TagUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := tag.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

func (tuo *TagUpdateOne) sqlSave(ctx context.Context) (_node *Tag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tag.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Tag.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tag.FieldID)
		for _, f := range fields {
			if !tag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldUpdatedAt,
		})
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldName,
		})
	}
	if tuo.mutation.RecruitmentTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tag.RecruitmentTagsTable,
			Columns: []string{tag.RecruitmentTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: recruitmenttag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedRecruitmentTagsIDs(); len(nodes) > 0 && !tuo.mutation.RecruitmentTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tag.RecruitmentTagsTable,
			Columns: []string{tag.RecruitmentTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: recruitmenttag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RecruitmentTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tag.RecruitmentTagsTable,
			Columns: []string{tag.RecruitmentTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: recruitmenttag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tag{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
