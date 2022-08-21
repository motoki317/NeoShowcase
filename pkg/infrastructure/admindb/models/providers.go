// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Provider is an object representing the database table.
type Provider struct { // プロバイダID
	ID string `boil:"id" json:"id" toml:"id" yaml:"id"`
	// プロバイダのドメイン
	Domain string `boil:"domain" json:"domain" toml:"domain" yaml:"domain"`
	// Webhookシークレット
	Secret string `boil:"secret" json:"secret" toml:"secret" yaml:"secret"`

	R *providerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L providerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProviderColumns = struct {
	ID     string
	Domain string
	Secret string
}{
	ID:     "id",
	Domain: "domain",
	Secret: "secret",
}

var ProviderTableColumns = struct {
	ID     string
	Domain string
	Secret string
}{
	ID:     "providers.id",
	Domain: "providers.domain",
	Secret: "providers.secret",
}

// Generated where

var ProviderWhere = struct {
	ID     whereHelperstring
	Domain whereHelperstring
	Secret whereHelperstring
}{
	ID:     whereHelperstring{field: "`providers`.`id`"},
	Domain: whereHelperstring{field: "`providers`.`domain`"},
	Secret: whereHelperstring{field: "`providers`.`secret`"},
}

// ProviderRels is where relationship names are stored.
var ProviderRels = struct {
	Repositories string
}{
	Repositories: "Repositories",
}

// providerR is where relationships are stored.
type providerR struct {
	Repositories RepositorySlice `boil:"Repositories" json:"Repositories" toml:"Repositories" yaml:"Repositories"`
}

// NewStruct creates a new relationship struct
func (*providerR) NewStruct() *providerR {
	return &providerR{}
}

func (r *providerR) GetRepositories() RepositorySlice {
	if r == nil {
		return nil
	}
	return r.Repositories
}

// providerL is where Load methods for each relationship are stored.
type providerL struct{}

var (
	providerAllColumns            = []string{"id", "domain", "secret"}
	providerColumnsWithoutDefault = []string{"id", "domain", "secret"}
	providerColumnsWithDefault    = []string{}
	providerPrimaryKeyColumns     = []string{"id"}
	providerGeneratedColumns      = []string{}
)

type (
	// ProviderSlice is an alias for a slice of pointers to Provider.
	// This should almost always be used instead of []Provider.
	ProviderSlice []*Provider
	// ProviderHook is the signature for custom Provider hook methods
	ProviderHook func(context.Context, boil.ContextExecutor, *Provider) error

	providerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	providerType                 = reflect.TypeOf(&Provider{})
	providerMapping              = queries.MakeStructMapping(providerType)
	providerPrimaryKeyMapping, _ = queries.BindMapping(providerType, providerMapping, providerPrimaryKeyColumns)
	providerInsertCacheMut       sync.RWMutex
	providerInsertCache          = make(map[string]insertCache)
	providerUpdateCacheMut       sync.RWMutex
	providerUpdateCache          = make(map[string]updateCache)
	providerUpsertCacheMut       sync.RWMutex
	providerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var providerAfterSelectHooks []ProviderHook

var providerBeforeInsertHooks []ProviderHook
var providerAfterInsertHooks []ProviderHook

var providerBeforeUpdateHooks []ProviderHook
var providerAfterUpdateHooks []ProviderHook

var providerBeforeDeleteHooks []ProviderHook
var providerAfterDeleteHooks []ProviderHook

var providerBeforeUpsertHooks []ProviderHook
var providerAfterUpsertHooks []ProviderHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Provider) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range providerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Provider) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range providerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Provider) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range providerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Provider) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range providerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Provider) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range providerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Provider) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range providerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Provider) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range providerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Provider) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range providerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Provider) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range providerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProviderHook registers your hook function for all future operations.
func AddProviderHook(hookPoint boil.HookPoint, providerHook ProviderHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		providerAfterSelectHooks = append(providerAfterSelectHooks, providerHook)
	case boil.BeforeInsertHook:
		providerBeforeInsertHooks = append(providerBeforeInsertHooks, providerHook)
	case boil.AfterInsertHook:
		providerAfterInsertHooks = append(providerAfterInsertHooks, providerHook)
	case boil.BeforeUpdateHook:
		providerBeforeUpdateHooks = append(providerBeforeUpdateHooks, providerHook)
	case boil.AfterUpdateHook:
		providerAfterUpdateHooks = append(providerAfterUpdateHooks, providerHook)
	case boil.BeforeDeleteHook:
		providerBeforeDeleteHooks = append(providerBeforeDeleteHooks, providerHook)
	case boil.AfterDeleteHook:
		providerAfterDeleteHooks = append(providerAfterDeleteHooks, providerHook)
	case boil.BeforeUpsertHook:
		providerBeforeUpsertHooks = append(providerBeforeUpsertHooks, providerHook)
	case boil.AfterUpsertHook:
		providerAfterUpsertHooks = append(providerAfterUpsertHooks, providerHook)
	}
}

// One returns a single provider record from the query.
func (q providerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Provider, error) {
	o := &Provider{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for providers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Provider records from the query.
func (q providerQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProviderSlice, error) {
	var o []*Provider

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Provider slice")
	}

	if len(providerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Provider records in the query.
func (q providerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count providers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q providerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if providers exists")
	}

	return count > 0, nil
}

// Repositories retrieves all the repository's Repositories with an executor.
func (o *Provider) Repositories(mods ...qm.QueryMod) repositoryQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`repositories`.`provider_id`=?", o.ID),
	)

	return Repositories(queryMods...)
}

// LoadRepositories allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (providerL) LoadRepositories(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProvider interface{}, mods queries.Applicator) error {
	var slice []*Provider
	var object *Provider

	if singular {
		object = maybeProvider.(*Provider)
	} else {
		slice = *maybeProvider.(*[]*Provider)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &providerR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &providerR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`repositories`),
		qm.WhereIn(`repositories.provider_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load repositories")
	}

	var resultSlice []*Repository
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice repositories")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on repositories")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for repositories")
	}

	if len(repositoryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Repositories = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &repositoryR{}
			}
			foreign.R.Provider = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ProviderID {
				local.R.Repositories = append(local.R.Repositories, foreign)
				if foreign.R == nil {
					foreign.R = &repositoryR{}
				}
				foreign.R.Provider = local
				break
			}
		}
	}

	return nil
}

// AddRepositories adds the given related objects to the existing relationships
// of the provider, optionally inserting them as new records.
// Appends related to o.R.Repositories.
// Sets related.R.Provider appropriately.
func (o *Provider) AddRepositories(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Repository) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ProviderID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `repositories` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"provider_id"}),
				strmangle.WhereClause("`", "`", 0, repositoryPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ProviderID = o.ID
		}
	}

	if o.R == nil {
		o.R = &providerR{
			Repositories: related,
		}
	} else {
		o.R.Repositories = append(o.R.Repositories, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &repositoryR{
				Provider: o,
			}
		} else {
			rel.R.Provider = o
		}
	}
	return nil
}

// Providers retrieves all the records using an executor.
func Providers(mods ...qm.QueryMod) providerQuery {
	mods = append(mods, qm.From("`providers`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`providers`.*"})
	}

	return providerQuery{q}
}

// FindProvider retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProvider(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Provider, error) {
	providerObj := &Provider{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `providers` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, providerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from providers")
	}

	if err = providerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return providerObj, err
	}

	return providerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Provider) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no providers provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(providerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	providerInsertCacheMut.RLock()
	cache, cached := providerInsertCache[key]
	providerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			providerAllColumns,
			providerColumnsWithDefault,
			providerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(providerType, providerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(providerType, providerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `providers` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `providers` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `providers` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, providerPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into providers")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for providers")
	}

CacheNoHooks:
	if !cached {
		providerInsertCacheMut.Lock()
		providerInsertCache[key] = cache
		providerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Provider.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Provider) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	providerUpdateCacheMut.RLock()
	cache, cached := providerUpdateCache[key]
	providerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			providerAllColumns,
			providerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update providers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `providers` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, providerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(providerType, providerMapping, append(wl, providerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update providers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for providers")
	}

	if !cached {
		providerUpdateCacheMut.Lock()
		providerUpdateCache[key] = cache
		providerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q providerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for providers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for providers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProviderSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), providerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `providers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, providerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in provider slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all provider")
	}
	return rowsAff, nil
}

var mySQLProviderUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Provider) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no providers provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(providerColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLProviderUniqueColumns, o)

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

	providerUpsertCacheMut.RLock()
	cache, cached := providerUpsertCache[key]
	providerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			providerAllColumns,
			providerColumnsWithDefault,
			providerColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			providerAllColumns,
			providerPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert providers, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`providers`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `providers` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(providerType, providerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(providerType, providerMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for providers")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(providerType, providerMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for providers")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for providers")
	}

CacheNoHooks:
	if !cached {
		providerUpsertCacheMut.Lock()
		providerUpsertCache[key] = cache
		providerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Provider record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Provider) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Provider provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), providerPrimaryKeyMapping)
	sql := "DELETE FROM `providers` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from providers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for providers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q providerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no providerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from providers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for providers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProviderSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(providerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), providerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `providers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, providerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from provider slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for providers")
	}

	if len(providerAfterDeleteHooks) != 0 {
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
func (o *Provider) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProvider(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProviderSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProviderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), providerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `providers`.* FROM `providers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, providerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProviderSlice")
	}

	*o = slice

	return nil
}

// ProviderExists checks if the Provider row exists.
func ProviderExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `providers` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if providers exists")
	}

	return exists, nil
}