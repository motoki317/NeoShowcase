// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBranches(t *testing.T) {
	t.Parallel()

	query := Branches()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBranchesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
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

	count, err := Branches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBranchesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Branches().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Branches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBranchesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BranchSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Branches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBranchesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BranchExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Branch exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BranchExists to return true, but got false.")
	}
}

func testBranchesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	branchFound, err := FindBranch(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if branchFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBranchesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Branches().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBranchesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Branches().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBranchesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	branchOne := &Branch{}
	branchTwo := &Branch{}
	if err = randomize.Struct(seed, branchOne, branchDBTypes, false, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}
	if err = randomize.Struct(seed, branchTwo, branchDBTypes, false, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = branchOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = branchTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Branches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBranchesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	branchOne := &Branch{}
	branchTwo := &Branch{}
	if err = randomize.Struct(seed, branchOne, branchDBTypes, false, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}
	if err = randomize.Struct(seed, branchTwo, branchDBTypes, false, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = branchOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = branchTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Branches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func branchBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Branch) error {
	*o = Branch{}
	return nil
}

func branchAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Branch) error {
	*o = Branch{}
	return nil
}

func branchAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Branch) error {
	*o = Branch{}
	return nil
}

func branchBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Branch) error {
	*o = Branch{}
	return nil
}

func branchAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Branch) error {
	*o = Branch{}
	return nil
}

func branchBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Branch) error {
	*o = Branch{}
	return nil
}

func branchAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Branch) error {
	*o = Branch{}
	return nil
}

func branchBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Branch) error {
	*o = Branch{}
	return nil
}

func branchAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Branch) error {
	*o = Branch{}
	return nil
}

func testBranchesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Branch{}
	o := &Branch{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, branchDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Branch object: %s", err)
	}

	AddBranchHook(boil.BeforeInsertHook, branchBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	branchBeforeInsertHooks = []BranchHook{}

	AddBranchHook(boil.AfterInsertHook, branchAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	branchAfterInsertHooks = []BranchHook{}

	AddBranchHook(boil.AfterSelectHook, branchAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	branchAfterSelectHooks = []BranchHook{}

	AddBranchHook(boil.BeforeUpdateHook, branchBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	branchBeforeUpdateHooks = []BranchHook{}

	AddBranchHook(boil.AfterUpdateHook, branchAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	branchAfterUpdateHooks = []BranchHook{}

	AddBranchHook(boil.BeforeDeleteHook, branchBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	branchBeforeDeleteHooks = []BranchHook{}

	AddBranchHook(boil.AfterDeleteHook, branchAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	branchAfterDeleteHooks = []BranchHook{}

	AddBranchHook(boil.BeforeUpsertHook, branchBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	branchBeforeUpsertHooks = []BranchHook{}

	AddBranchHook(boil.AfterUpsertHook, branchAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	branchAfterUpsertHooks = []BranchHook{}
}

func testBranchesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Branches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBranchesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(branchColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Branches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBranchOneToOneWebsiteUsingWebsite(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var foreign Website
	var local Branch

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreign.BranchID = local.ID
	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Website().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.BranchID != foreign.BranchID {
		t.Errorf("want: %v, got %v", foreign.BranchID, check.BranchID)
	}

	slice := BranchSlice{&local}
	if err = local.L.LoadWebsite(ctx, tx, false, (*[]*Branch)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Website == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Website = nil
	if err = local.L.LoadWebsite(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Website == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBranchOneToOneSetOpWebsiteUsingWebsite(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Branch
	var b, c Website

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, branchDBTypes, false, strmangle.SetComplement(branchPrimaryKeyColumns, branchColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, websiteDBTypes, false, strmangle.SetComplement(websitePrimaryKeyColumns, websiteColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, websiteDBTypes, false, strmangle.SetComplement(websitePrimaryKeyColumns, websiteColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Website{&b, &c} {
		err = a.SetWebsite(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Website != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Branch != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.ID != x.BranchID {
			t.Error("foreign key was wrong value", a.ID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.BranchID))
		reflect.Indirect(reflect.ValueOf(&x.BranchID)).Set(zero)

		if err = x.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ID != x.BranchID {
			t.Error("foreign key was wrong value", a.ID, x.BranchID)
		}

		if _, err = x.Delete(ctx, tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testBranchToManyBuildLogs(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Branch
	var b, c BuildLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, buildLogDBTypes, false, buildLogColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, buildLogDBTypes, false, buildLogColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.BranchID = a.ID
	c.BranchID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.BuildLogs().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BranchID == b.BranchID {
			bFound = true
		}
		if v.BranchID == c.BranchID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BranchSlice{&a}
	if err = a.L.LoadBuildLogs(ctx, tx, false, (*[]*Branch)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BuildLogs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.BuildLogs = nil
	if err = a.L.LoadBuildLogs(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BuildLogs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBranchToManyEnvironments(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Branch
	var b, c Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, environmentDBTypes, false, environmentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, environmentDBTypes, false, environmentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.BranchID = a.ID
	c.BranchID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Environments().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BranchID == b.BranchID {
			bFound = true
		}
		if v.BranchID == c.BranchID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BranchSlice{&a}
	if err = a.L.LoadEnvironments(ctx, tx, false, (*[]*Branch)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Environments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Environments = nil
	if err = a.L.LoadEnvironments(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Environments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBranchToManyAddOpBuildLogs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Branch
	var b, c, d, e BuildLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, branchDBTypes, false, strmangle.SetComplement(branchPrimaryKeyColumns, branchColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BuildLog{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, buildLogDBTypes, false, strmangle.SetComplement(buildLogPrimaryKeyColumns, buildLogColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*BuildLog{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBuildLogs(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BranchID {
			t.Error("foreign key was wrong value", a.ID, first.BranchID)
		}
		if a.ID != second.BranchID {
			t.Error("foreign key was wrong value", a.ID, second.BranchID)
		}

		if first.R.Branch != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Branch != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.BuildLogs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.BuildLogs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.BuildLogs().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBranchToManyAddOpEnvironments(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Branch
	var b, c, d, e Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, branchDBTypes, false, strmangle.SetComplement(branchPrimaryKeyColumns, branchColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Environment{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Environment{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddEnvironments(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BranchID {
			t.Error("foreign key was wrong value", a.ID, first.BranchID)
		}
		if a.ID != second.BranchID {
			t.Error("foreign key was wrong value", a.ID, second.BranchID)
		}

		if first.R.Branch != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Branch != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Environments[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Environments[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Environments().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBranchToOneApplicationUsingApplication(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Branch
	var foreign Application

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, branchDBTypes, false, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, applicationDBTypes, false, applicationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Application struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ApplicationID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Application().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BranchSlice{&local}
	if err = local.L.LoadApplication(ctx, tx, false, (*[]*Branch)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Application == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Application = nil
	if err = local.L.LoadApplication(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Application == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBranchToOneBuildLogUsingBuild(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Branch
	var foreign BuildLog

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, buildLogDBTypes, false, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.BuildID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Build().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BranchSlice{&local}
	if err = local.L.LoadBuild(ctx, tx, false, (*[]*Branch)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Build == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Build = nil
	if err = local.L.LoadBuild(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Build == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBranchToOneSetOpApplicationUsingApplication(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Branch
	var b, c Application

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, branchDBTypes, false, strmangle.SetComplement(branchPrimaryKeyColumns, branchColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, applicationDBTypes, false, strmangle.SetComplement(applicationPrimaryKeyColumns, applicationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, applicationDBTypes, false, strmangle.SetComplement(applicationPrimaryKeyColumns, applicationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Application{&b, &c} {
		err = a.SetApplication(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Application != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Branches[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ApplicationID != x.ID {
			t.Error("foreign key was wrong value", a.ApplicationID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ApplicationID))
		reflect.Indirect(reflect.ValueOf(&a.ApplicationID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ApplicationID != x.ID {
			t.Error("foreign key was wrong value", a.ApplicationID, x.ID)
		}
	}
}
func testBranchToOneSetOpBuildLogUsingBuild(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Branch
	var b, c BuildLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, branchDBTypes, false, strmangle.SetComplement(branchPrimaryKeyColumns, branchColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, buildLogDBTypes, false, strmangle.SetComplement(buildLogPrimaryKeyColumns, buildLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, buildLogDBTypes, false, strmangle.SetComplement(buildLogPrimaryKeyColumns, buildLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*BuildLog{&b, &c} {
		err = a.SetBuild(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Build != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BuildBranches[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.BuildID, x.ID) {
			t.Error("foreign key was wrong value", a.BuildID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BuildID))
		reflect.Indirect(reflect.ValueOf(&a.BuildID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.BuildID, x.ID) {
			t.Error("foreign key was wrong value", a.BuildID, x.ID)
		}
	}
}

func testBranchToOneRemoveOpBuildLogUsingBuild(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Branch
	var b BuildLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, branchDBTypes, false, strmangle.SetComplement(branchPrimaryKeyColumns, branchColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, buildLogDBTypes, false, strmangle.SetComplement(buildLogPrimaryKeyColumns, buildLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetBuild(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveBuild(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Build().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Build != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.BuildID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.BuildBranches) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testBranchesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
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

func testBranchesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BranchSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBranchesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Branches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	branchDBTypes = map[string]string{`ID`: `varchar`, `ApplicationID`: `varchar`, `BranchName`: `varchar`, `BuildType`: `enum('image','static')`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`, `BuildID`: `varchar`}
	_             = bytes.MinRead
)

func testBranchesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(branchPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(branchAllColumns) == len(branchPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Branches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, branchDBTypes, true, branchPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBranchesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(branchAllColumns) == len(branchPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Branch{}
	if err = randomize.Struct(seed, o, branchDBTypes, true, branchColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Branches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, branchDBTypes, true, branchPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(branchAllColumns, branchPrimaryKeyColumns) {
		fields = branchAllColumns
	} else {
		fields = strmangle.SetComplement(
			branchAllColumns,
			branchPrimaryKeyColumns,
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

	slice := BranchSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBranchesUpsert(t *testing.T) {
	t.Parallel()

	if len(branchAllColumns) == len(branchPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLBranchUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Branch{}
	if err = randomize.Struct(seed, &o, branchDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Branch: %s", err)
	}

	count, err := Branches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, branchDBTypes, false, branchPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Branch struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Branch: %s", err)
	}

	count, err = Branches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
