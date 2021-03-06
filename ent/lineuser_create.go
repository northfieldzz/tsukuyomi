// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tsukuyomi/ent/lineuser"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LineUserCreate is the builder for creating a LineUser entity.
type LineUserCreate struct {
	config
	mutation *LineUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIsActive sets the "is_active" field.
func (luc *LineUserCreate) SetIsActive(b bool) *LineUserCreate {
	luc.mutation.SetIsActive(b)
	return luc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (luc *LineUserCreate) SetNillableIsActive(b *bool) *LineUserCreate {
	if b != nil {
		luc.SetIsActive(*b)
	}
	return luc
}

// SetCreateAt sets the "create_at" field.
func (luc *LineUserCreate) SetCreateAt(t time.Time) *LineUserCreate {
	luc.mutation.SetCreateAt(t)
	return luc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (luc *LineUserCreate) SetNillableCreateAt(t *time.Time) *LineUserCreate {
	if t != nil {
		luc.SetCreateAt(*t)
	}
	return luc
}

// SetUpdateAt sets the "update_at" field.
func (luc *LineUserCreate) SetUpdateAt(t time.Time) *LineUserCreate {
	luc.mutation.SetUpdateAt(t)
	return luc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (luc *LineUserCreate) SetNillableUpdateAt(t *time.Time) *LineUserCreate {
	if t != nil {
		luc.SetUpdateAt(*t)
	}
	return luc
}

// SetID sets the "id" field.
func (luc *LineUserCreate) SetID(s string) *LineUserCreate {
	luc.mutation.SetID(s)
	return luc
}

// Mutation returns the LineUserMutation object of the builder.
func (luc *LineUserCreate) Mutation() *LineUserMutation {
	return luc.mutation
}

// Save creates the LineUser in the database.
func (luc *LineUserCreate) Save(ctx context.Context) (*LineUser, error) {
	var (
		err  error
		node *LineUser
	)
	luc.defaults()
	if len(luc.hooks) == 0 {
		if err = luc.check(); err != nil {
			return nil, err
		}
		node, err = luc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LineUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = luc.check(); err != nil {
				return nil, err
			}
			luc.mutation = mutation
			if node, err = luc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(luc.hooks) - 1; i >= 0; i-- {
			if luc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (luc *LineUserCreate) SaveX(ctx context.Context) *LineUser {
	v, err := luc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (luc *LineUserCreate) Exec(ctx context.Context) error {
	_, err := luc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luc *LineUserCreate) ExecX(ctx context.Context) {
	if err := luc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luc *LineUserCreate) defaults() {
	if _, ok := luc.mutation.IsActive(); !ok {
		v := lineuser.DefaultIsActive
		luc.mutation.SetIsActive(v)
	}
	if _, ok := luc.mutation.CreateAt(); !ok {
		v := lineuser.DefaultCreateAt()
		luc.mutation.SetCreateAt(v)
	}
	if _, ok := luc.mutation.UpdateAt(); !ok {
		v := lineuser.DefaultUpdateAt()
		luc.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luc *LineUserCreate) check() error {
	if _, ok := luc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "LineUser.is_active"`)}
	}
	if _, ok := luc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "LineUser.create_at"`)}
	}
	if _, ok := luc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "LineUser.update_at"`)}
	}
	if v, ok := luc.mutation.ID(); ok {
		if err := lineuser.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "LineUser.id": %w`, err)}
		}
	}
	return nil
}

func (luc *LineUserCreate) sqlSave(ctx context.Context) (*LineUser, error) {
	_node, _spec := luc.createSpec()
	if err := sqlgraph.CreateNode(ctx, luc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected LineUser.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (luc *LineUserCreate) createSpec() (*LineUser, *sqlgraph.CreateSpec) {
	var (
		_node = &LineUser{config: luc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lineuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: lineuser.FieldID,
			},
		}
	)
	_spec.OnConflict = luc.conflict
	if id, ok := luc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := luc.mutation.IsActive(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: lineuser.FieldIsActive,
		})
		_node.IsActive = value
	}
	if value, ok := luc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lineuser.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := luc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lineuser.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LineUser.Create().
//		SetIsActive(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LineUserUpsert) {
//			SetIsActive(v+v).
//		}).
//		Exec(ctx)
//
func (luc *LineUserCreate) OnConflict(opts ...sql.ConflictOption) *LineUserUpsertOne {
	luc.conflict = opts
	return &LineUserUpsertOne{
		create: luc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LineUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (luc *LineUserCreate) OnConflictColumns(columns ...string) *LineUserUpsertOne {
	luc.conflict = append(luc.conflict, sql.ConflictColumns(columns...))
	return &LineUserUpsertOne{
		create: luc,
	}
}

type (
	// LineUserUpsertOne is the builder for "upsert"-ing
	//  one LineUser node.
	LineUserUpsertOne struct {
		create *LineUserCreate
	}

	// LineUserUpsert is the "OnConflict" setter.
	LineUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetIsActive sets the "is_active" field.
func (u *LineUserUpsert) SetIsActive(v bool) *LineUserUpsert {
	u.Set(lineuser.FieldIsActive, v)
	return u
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *LineUserUpsert) UpdateIsActive() *LineUserUpsert {
	u.SetExcluded(lineuser.FieldIsActive)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *LineUserUpsert) SetCreateAt(v time.Time) *LineUserUpsert {
	u.Set(lineuser.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *LineUserUpsert) UpdateCreateAt() *LineUserUpsert {
	u.SetExcluded(lineuser.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *LineUserUpsert) SetUpdateAt(v time.Time) *LineUserUpsert {
	u.Set(lineuser.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *LineUserUpsert) UpdateUpdateAt() *LineUserUpsert {
	u.SetExcluded(lineuser.FieldUpdateAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.LineUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(lineuser.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *LineUserUpsertOne) UpdateNewValues() *LineUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(lineuser.FieldID)
		}
		if _, exists := u.create.mutation.CreateAt(); exists {
			s.SetIgnore(lineuser.FieldCreateAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.LineUser.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *LineUserUpsertOne) Ignore() *LineUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LineUserUpsertOne) DoNothing() *LineUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LineUserCreate.OnConflict
// documentation for more info.
func (u *LineUserUpsertOne) Update(set func(*LineUserUpsert)) *LineUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LineUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsActive sets the "is_active" field.
func (u *LineUserUpsertOne) SetIsActive(v bool) *LineUserUpsertOne {
	return u.Update(func(s *LineUserUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *LineUserUpsertOne) UpdateIsActive() *LineUserUpsertOne {
	return u.Update(func(s *LineUserUpsert) {
		s.UpdateIsActive()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *LineUserUpsertOne) SetCreateAt(v time.Time) *LineUserUpsertOne {
	return u.Update(func(s *LineUserUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *LineUserUpsertOne) UpdateCreateAt() *LineUserUpsertOne {
	return u.Update(func(s *LineUserUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *LineUserUpsertOne) SetUpdateAt(v time.Time) *LineUserUpsertOne {
	return u.Update(func(s *LineUserUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *LineUserUpsertOne) UpdateUpdateAt() *LineUserUpsertOne {
	return u.Update(func(s *LineUserUpsert) {
		s.UpdateUpdateAt()
	})
}

// Exec executes the query.
func (u *LineUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LineUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LineUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LineUserUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: LineUserUpsertOne.ID is not supported by MySQL driver. Use LineUserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LineUserUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LineUserCreateBulk is the builder for creating many LineUser entities in bulk.
type LineUserCreateBulk struct {
	config
	builders []*LineUserCreate
	conflict []sql.ConflictOption
}

// Save creates the LineUser entities in the database.
func (lucb *LineUserCreateBulk) Save(ctx context.Context) ([]*LineUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lucb.builders))
	nodes := make([]*LineUser, len(lucb.builders))
	mutators := make([]Mutator, len(lucb.builders))
	for i := range lucb.builders {
		func(i int, root context.Context) {
			builder := lucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LineUserMutation)
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
					_, err = mutators[i+1].Mutate(root, lucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, lucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lucb *LineUserCreateBulk) SaveX(ctx context.Context) []*LineUser {
	v, err := lucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lucb *LineUserCreateBulk) Exec(ctx context.Context) error {
	_, err := lucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lucb *LineUserCreateBulk) ExecX(ctx context.Context) {
	if err := lucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LineUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LineUserUpsert) {
//			SetIsActive(v+v).
//		}).
//		Exec(ctx)
//
func (lucb *LineUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *LineUserUpsertBulk {
	lucb.conflict = opts
	return &LineUserUpsertBulk{
		create: lucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LineUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lucb *LineUserCreateBulk) OnConflictColumns(columns ...string) *LineUserUpsertBulk {
	lucb.conflict = append(lucb.conflict, sql.ConflictColumns(columns...))
	return &LineUserUpsertBulk{
		create: lucb,
	}
}

// LineUserUpsertBulk is the builder for "upsert"-ing
// a bulk of LineUser nodes.
type LineUserUpsertBulk struct {
	create *LineUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.LineUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(lineuser.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *LineUserUpsertBulk) UpdateNewValues() *LineUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(lineuser.FieldID)
				return
			}
			if _, exists := b.mutation.CreateAt(); exists {
				s.SetIgnore(lineuser.FieldCreateAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LineUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *LineUserUpsertBulk) Ignore() *LineUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LineUserUpsertBulk) DoNothing() *LineUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LineUserCreateBulk.OnConflict
// documentation for more info.
func (u *LineUserUpsertBulk) Update(set func(*LineUserUpsert)) *LineUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LineUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsActive sets the "is_active" field.
func (u *LineUserUpsertBulk) SetIsActive(v bool) *LineUserUpsertBulk {
	return u.Update(func(s *LineUserUpsert) {
		s.SetIsActive(v)
	})
}

// UpdateIsActive sets the "is_active" field to the value that was provided on create.
func (u *LineUserUpsertBulk) UpdateIsActive() *LineUserUpsertBulk {
	return u.Update(func(s *LineUserUpsert) {
		s.UpdateIsActive()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *LineUserUpsertBulk) SetCreateAt(v time.Time) *LineUserUpsertBulk {
	return u.Update(func(s *LineUserUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *LineUserUpsertBulk) UpdateCreateAt() *LineUserUpsertBulk {
	return u.Update(func(s *LineUserUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *LineUserUpsertBulk) SetUpdateAt(v time.Time) *LineUserUpsertBulk {
	return u.Update(func(s *LineUserUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *LineUserUpsertBulk) UpdateUpdateAt() *LineUserUpsertBulk {
	return u.Update(func(s *LineUserUpsert) {
		s.UpdateUpdateAt()
	})
}

// Exec executes the query.
func (u *LineUserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LineUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LineUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LineUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
