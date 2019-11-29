// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testProposals(t *testing.T) {
	t.Parallel()

	query := Proposals()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProposalsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Proposals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProposalsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Proposals().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Proposals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProposalsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProposalSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Proposals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProposalsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProposalExists(ctx, tx, o.UUID)
	if err != nil {
		t.Errorf("Unable to check if Proposal exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProposalExists to return true, but got false.")
	}
}

func testProposalsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	proposalFound, err := FindProposal(ctx, tx, o.UUID)
	if err != nil {
		t.Error(err)
	}

	if proposalFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProposalsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Proposals().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testProposalsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Proposals().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProposalsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	proposalOne := &Proposal{}
	proposalTwo := &Proposal{}
	if err = randomize.Struct(seed, proposalOne, proposalDBTypes, false, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}
	if err = randomize.Struct(seed, proposalTwo, proposalDBTypes, false, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = proposalOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = proposalTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Proposals().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProposalsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	proposalOne := &Proposal{}
	proposalTwo := &Proposal{}
	if err = randomize.Struct(seed, proposalOne, proposalDBTypes, false, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}
	if err = randomize.Struct(seed, proposalTwo, proposalDBTypes, false, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = proposalOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = proposalTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Proposals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func proposalBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Proposal) error {
	*o = Proposal{}
	return nil
}

func proposalAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Proposal) error {
	*o = Proposal{}
	return nil
}

func proposalAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Proposal) error {
	*o = Proposal{}
	return nil
}

func proposalBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Proposal) error {
	*o = Proposal{}
	return nil
}

func proposalAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Proposal) error {
	*o = Proposal{}
	return nil
}

func proposalBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Proposal) error {
	*o = Proposal{}
	return nil
}

func proposalAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Proposal) error {
	*o = Proposal{}
	return nil
}

func proposalBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Proposal) error {
	*o = Proposal{}
	return nil
}

func proposalAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Proposal) error {
	*o = Proposal{}
	return nil
}

func testProposalsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Proposal{}
	o := &Proposal{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, proposalDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Proposal object: %s", err)
	}

	AddProposalHook(boil.BeforeInsertHook, proposalBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	proposalBeforeInsertHooks = []ProposalHook{}

	AddProposalHook(boil.AfterInsertHook, proposalAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	proposalAfterInsertHooks = []ProposalHook{}

	AddProposalHook(boil.AfterSelectHook, proposalAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	proposalAfterSelectHooks = []ProposalHook{}

	AddProposalHook(boil.BeforeUpdateHook, proposalBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	proposalBeforeUpdateHooks = []ProposalHook{}

	AddProposalHook(boil.AfterUpdateHook, proposalAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	proposalAfterUpdateHooks = []ProposalHook{}

	AddProposalHook(boil.BeforeDeleteHook, proposalBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	proposalBeforeDeleteHooks = []ProposalHook{}

	AddProposalHook(boil.AfterDeleteHook, proposalAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	proposalAfterDeleteHooks = []ProposalHook{}

	AddProposalHook(boil.BeforeUpsertHook, proposalBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	proposalBeforeUpsertHooks = []ProposalHook{}

	AddProposalHook(boil.AfterUpsertHook, proposalAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	proposalAfterUpsertHooks = []ProposalHook{}
}

func testProposalsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Proposals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProposalsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(proposalColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Proposals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProposalsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProposalsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProposalSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProposalsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Proposals().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	proposalDBTypes = map[string]string{`UUID`: `text`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `Title`: `text`, `Desc`: `text`, `NBVote`: `integer`}
	_               = bytes.MinRead
)

func testProposalsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(proposalPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(proposalAllColumns) == len(proposalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Proposals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProposalsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(proposalAllColumns) == len(proposalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Proposal{}
	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Proposals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, proposalDBTypes, true, proposalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(proposalAllColumns, proposalPrimaryKeyColumns) {
		fields = proposalAllColumns
	} else {
		fields = strmangle.SetComplement(
			proposalAllColumns,
			proposalPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ProposalSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProposalsUpsert(t *testing.T) {
	t.Parallel()

	if len(proposalAllColumns) == len(proposalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Proposal{}
	if err = randomize.Struct(seed, &o, proposalDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Proposal: %s", err)
	}

	count, err := Proposals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, proposalDBTypes, false, proposalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Proposal struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Proposal: %s", err)
	}

	count, err = Proposals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
