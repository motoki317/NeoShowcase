// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// BuildStatus is an object representing the database table.
type BuildStatus struct { // ビルドの状態
	Status string `boil:"status" json:"status" toml:"status" yaml:"status"`

	R *buildStatusR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L buildStatusL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BuildStatusColumns = struct {
	Status string
}{
	Status: "status",
}

var BuildStatusTableColumns = struct {
	Status string
}{
	Status: "build_status.status",
}

// Generated where

var BuildStatusWhere = struct {
	Status whereHelperstring
}{
	Status: whereHelperstring{field: "`build_status`.`status`"},
}

// BuildStatusRels is where relationship names are stored.
var BuildStatusRels = struct {
	StatusBuilds string
}{
	StatusBuilds: "StatusBuilds",
}

// buildStatusR is where relationships are stored.
type buildStatusR struct {
	StatusBuilds BuildSlice `boil:"StatusBuilds" json:"StatusBuilds" toml:"StatusBuilds" yaml:"StatusBuilds"`
}

// NewStruct creates a new relationship struct
func (*buildStatusR) NewStruct() *buildStatusR {
	return &buildStatusR{}
}

func (r *buildStatusR) GetStatusBuilds() BuildSlice {
	if r == nil {
		return nil
	}
	return r.StatusBuilds
}

// buildStatusL is where Load methods for each relationship are stored.
type buildStatusL struct{}

var (
	buildStatusAllColumns            = []string{"status"}
	buildStatusColumnsWithoutDefault = []string{"status"}
	buildStatusColumnsWithDefault    = []string{}
	buildStatusPrimaryKeyColumns     = []string{"status"}
	buildStatusGeneratedColumns      = []string{}
)

type (
	// BuildStatusSlice is an alias for a slice of pointers to BuildStatus.
	// This should almost always be used instead of []BuildStatus.
	BuildStatusSlice []*BuildStatus
	// BuildStatusHook is the signature for custom BuildStatus hook methods
	BuildStatusHook func(context.Context, boil.ContextExecutor, *BuildStatus) error

	buildStatusQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	buildStatusType                 = reflect.TypeOf(&BuildStatus{})
	buildStatusMapping              = queries.MakeStructMapping(buildStatusType)
	buildStatusPrimaryKeyMapping, _ = queries.BindMapping(buildStatusType, buildStatusMapping, buildStatusPrimaryKeyColumns)
	buildStatusInsertCacheMut       sync.RWMutex
	buildStatusInsertCache          = make(map[string]insertCache)
	buildStatusUpdateCacheMut       sync.RWMutex
	buildStatusUpdateCache          = make(map[string]updateCache)
	buildStatusUpsertCacheMut       sync.RWMutex
	buildStatusUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var buildStatusAfterSelectHooks []BuildStatusHook

var buildStatusBeforeInsertHooks []BuildStatusHook
var buildStatusAfterInsertHooks []BuildStatusHook

var buildStatusBeforeUpdateHooks []BuildStatusHook
var buildStatusAfterUpdateHooks []BuildStatusHook

var buildStatusBeforeDeleteHooks []BuildStatusHook
var buildStatusAfterDeleteHooks []BuildStatusHook

var buildStatusBeforeUpsertHooks []BuildStatusHook
var buildStatusAfterUpsertHooks []BuildStatusHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BuildStatus) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildStatusAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BuildStatus) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildStatusBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BuildStatus) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildStatusAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BuildStatus) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildStatusBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BuildStatus) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildStatusAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BuildStatus) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildStatusBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BuildStatus) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildStatusAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BuildStatus) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildStatusBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BuildStatus) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildStatusAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBuildStatusHook registers your hook function for all future operations.
func AddBuildStatusHook(hookPoint boil.HookPoint, buildStatusHook BuildStatusHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		buildStatusAfterSelectHooks = append(buildStatusAfterSelectHooks, buildStatusHook)
	case boil.BeforeInsertHook:
		buildStatusBeforeInsertHooks = append(buildStatusBeforeInsertHooks, buildStatusHook)
	case boil.AfterInsertHook:
		buildStatusAfterInsertHooks = append(buildStatusAfterInsertHooks, buildStatusHook)
	case boil.BeforeUpdateHook:
		buildStatusBeforeUpdateHooks = append(buildStatusBeforeUpdateHooks, buildStatusHook)
	case boil.AfterUpdateHook:
		buildStatusAfterUpdateHooks = append(buildStatusAfterUpdateHooks, buildStatusHook)
	case boil.BeforeDeleteHook:
		buildStatusBeforeDeleteHooks = append(buildStatusBeforeDeleteHooks, buildStatusHook)
	case boil.AfterDeleteHook:
		buildStatusAfterDeleteHooks = append(buildStatusAfterDeleteHooks, buildStatusHook)
	case boil.BeforeUpsertHook:
		buildStatusBeforeUpsertHooks = append(buildStatusBeforeUpsertHooks, buildStatusHook)
	case boil.AfterUpsertHook:
		buildStatusAfterUpsertHooks = append(buildStatusAfterUpsertHooks, buildStatusHook)
	}
}

// One returns a single buildStatus record from the query.
func (q buildStatusQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BuildStatus, error) {
	o := &BuildStatus{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for build_status")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BuildStatus records from the query.
func (q buildStatusQuery) All(ctx context.Context, exec boil.ContextExecutor) (BuildStatusSlice, error) {
	var o []*BuildStatus

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BuildStatus slice")
	}

	if len(buildStatusAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BuildStatus records in the query.
func (q buildStatusQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count build_status rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q buildStatusQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if build_status exists")
	}

	return count > 0, nil
}

// StatusBuilds retrieves all the build's Builds with an executor via status column.
func (o *BuildStatus) StatusBuilds(mods ...qm.QueryMod) buildQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`builds`.`status`=?", o.Status),
	)

	return Builds(queryMods...)
}

// LoadStatusBuilds allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (buildStatusL) LoadStatusBuilds(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBuildStatus interface{}, mods queries.Applicator) error {
	var slice []*BuildStatus
	var object *BuildStatus

	if singular {
		var ok bool
		object, ok = maybeBuildStatus.(*BuildStatus)
		if !ok {
			object = new(BuildStatus)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBuildStatus)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBuildStatus))
			}
		}
	} else {
		s, ok := maybeBuildStatus.(*[]*BuildStatus)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBuildStatus)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBuildStatus))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &buildStatusR{}
		}
		args = append(args, object.Status)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &buildStatusR{}
			}

			for _, a := range args {
				if a == obj.Status {
					continue Outer
				}
			}

			args = append(args, obj.Status)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`builds`),
		qm.WhereIn(`builds.status in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load builds")
	}

	var resultSlice []*Build
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice builds")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on builds")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for builds")
	}

	if len(buildAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.StatusBuilds = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &buildR{}
			}
			foreign.R.StatusBuildStatus = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.Status == foreign.Status {
				local.R.StatusBuilds = append(local.R.StatusBuilds, foreign)
				if foreign.R == nil {
					foreign.R = &buildR{}
				}
				foreign.R.StatusBuildStatus = local
				break
			}
		}
	}

	return nil
}

// AddStatusBuilds adds the given related objects to the existing relationships
// of the build_status, optionally inserting them as new records.
// Appends related to o.R.StatusBuilds.
// Sets related.R.StatusBuildStatus appropriately.
func (o *BuildStatus) AddStatusBuilds(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Build) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.Status = o.Status
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `builds` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"status"}),
				strmangle.WhereClause("`", "`", 0, buildPrimaryKeyColumns),
			)
			values := []interface{}{o.Status, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.Status = o.Status
		}
	}

	if o.R == nil {
		o.R = &buildStatusR{
			StatusBuilds: related,
		}
	} else {
		o.R.StatusBuilds = append(o.R.StatusBuilds, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &buildR{
				StatusBuildStatus: o,
			}
		} else {
			rel.R.StatusBuildStatus = o
		}
	}
	return nil
}

// BuildStatuses retrieves all the records using an executor.
func BuildStatuses(mods ...qm.QueryMod) buildStatusQuery {
	mods = append(mods, qm.From("`build_status`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`build_status`.*"})
	}

	return buildStatusQuery{q}
}

// FindBuildStatus retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBuildStatus(ctx context.Context, exec boil.ContextExecutor, status string, selectCols ...string) (*BuildStatus, error) {
	buildStatusObj := &BuildStatus{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `build_status` where `status`=?", sel,
	)

	q := queries.Raw(query, status)

	err := q.Bind(ctx, exec, buildStatusObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from build_status")
	}

	if err = buildStatusObj.doAfterSelectHooks(ctx, exec); err != nil {
		return buildStatusObj, err
	}

	return buildStatusObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BuildStatus) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no build_status provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(buildStatusColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	buildStatusInsertCacheMut.RLock()
	cache, cached := buildStatusInsertCache[key]
	buildStatusInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			buildStatusAllColumns,
			buildStatusColumnsWithDefault,
			buildStatusColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(buildStatusType, buildStatusMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(buildStatusType, buildStatusMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `build_status` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `build_status` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `build_status` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, buildStatusPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into build_status")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Status,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for build_status")
	}

CacheNoHooks:
	if !cached {
		buildStatusInsertCacheMut.Lock()
		buildStatusInsertCache[key] = cache
		buildStatusInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BuildStatus.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BuildStatus) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	buildStatusUpdateCacheMut.RLock()
	cache, cached := buildStatusUpdateCache[key]
	buildStatusUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			buildStatusAllColumns,
			buildStatusPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update build_status, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `build_status` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, buildStatusPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(buildStatusType, buildStatusMapping, append(wl, buildStatusPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update build_status row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for build_status")
	}

	if !cached {
		buildStatusUpdateCacheMut.Lock()
		buildStatusUpdateCache[key] = cache
		buildStatusUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q buildStatusQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for build_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for build_status")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BuildStatusSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `build_status` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, buildStatusPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in buildStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all buildStatus")
	}
	return rowsAff, nil
}

var mySQLBuildStatusUniqueColumns = []string{
	"status",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BuildStatus) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no build_status provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(buildStatusColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBuildStatusUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	buildStatusUpsertCacheMut.RLock()
	cache, cached := buildStatusUpsertCache[key]
	buildStatusUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			buildStatusAllColumns,
			buildStatusColumnsWithDefault,
			buildStatusColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			buildStatusAllColumns,
			buildStatusPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert build_status, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`build_status`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `build_status` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(buildStatusType, buildStatusMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(buildStatusType, buildStatusMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for build_status")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(buildStatusType, buildStatusMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for build_status")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for build_status")
	}

CacheNoHooks:
	if !cached {
		buildStatusUpsertCacheMut.Lock()
		buildStatusUpsertCache[key] = cache
		buildStatusUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BuildStatus record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BuildStatus) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BuildStatus provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), buildStatusPrimaryKeyMapping)
	sql := "DELETE FROM `build_status` WHERE `status`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from build_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for build_status")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q buildStatusQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no buildStatusQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from build_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for build_status")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BuildStatusSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(buildStatusBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `build_status` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, buildStatusPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from buildStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for build_status")
	}

	if len(buildStatusAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BuildStatus) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBuildStatus(ctx, exec, o.Status)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BuildStatusSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BuildStatusSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `build_status`.* FROM `build_status` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, buildStatusPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BuildStatusSlice")
	}

	*o = slice

	return nil
}

// BuildStatusExists checks if the BuildStatus row exists.
func BuildStatusExists(ctx context.Context, exec boil.ContextExecutor, status string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `build_status` where `status`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, status)
	}
	row := exec.QueryRowContext(ctx, sql, status)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if build_status exists")
	}

	return exists, nil
}

// Exists checks if the BuildStatus row exists.
func (o *BuildStatus) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BuildStatusExists(ctx, exec, o.Status)
}
