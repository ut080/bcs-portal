package postgres

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/repositories"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func valuePlaceholders(object repositories.RepoObject, count int) (string, int) {
	values := ""
	idx := 1
	for i := 0; i < count; i++ {
		val, end := object.Parameters("$", idx)
		values += val
		idx += end + 1
	}

	return values, idx
}

func idList(start, count int) string {
	values := "("
	for i := start; i < count; i++ {
		values += fmt.Sprintf("%d,", i)
	}

	values = strings.TrimRight(values, ",") + ")"

	return values
}

func (r *Repository) prepareSoftDeleteStatement(object repositories.SoftDeletable, count int) string {
	template := object.SoftDelete()
	ids := idList(1, count)
	template = strings.Replace(template, "${IDS}", ids, 1)

	return template
}

func (r *Repository) prepareStatement(object repositories.RepoObject, statementType StatementType, count int) string {
	var template string
	switch statementType {
	case InsertStatement:
		template = object.Create()
		values, _ := valuePlaceholders(object, count)
		template = strings.Replace(template, "${VALUES}", values, 1)
	case SelectStatement:
		template = object.Fetch(true)
	case UpdateStatement:
		template = object.Update()
		values, start := valuePlaceholders(object, count)
		ids := idList(start, count)
		template = strings.Replace(template, "${VALUES}", values, 1)
		template = strings.Replace(template, "${IDS}", ids, 1)
	case UpsertStatement:
		template = object.UpdateOrCreate()
		values, start := valuePlaceholders(object, count)
		ids := idList(start, count)
		template = strings.Replace(template, "${VALUES}", values, 1)
		template = strings.Replace(template, "${IDS}", ids, 1)
	case DeleteStatement:
		template = object.Delete()
		ids := idList(1, count)
		template = strings.Replace(template, "${IDS}", ids, 1)
	}

	return template
}

func (r *Repository) Create(objects []repositories.RepoObject) (repositories.Statement, error) {
	query := r.prepareStatement(objects[0], InsertStatement, len(objects))

	var params []any
	for _, v := range objects {
		p := v.Values()
		params = append(params, p...)
	}

	stmt, err := NewStatement(r.db, InsertStatement, query, params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stmt, nil
}

func (r *Repository) CreateOne(object repositories.RepoObject) (repositories.Statement, error) {
	query := r.prepareStatement(object, InsertStatement, 1)

	stmt, err := NewStatement(r.db, InsertStatement, query, object.Values())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stmt, nil
}

func (r *Repository) Fetch(object repositories.RepoObject) (repositories.Statement, error) {
	query := r.prepareStatement(object, SelectStatement, -1)

	// TODO: Review how to make SELECT queries
	stmt, err := NewStatement(r.db, SelectStatement, query, nil)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stmt, nil
}

func (r *Repository) Update(objects []repositories.RepoObject) (repositories.Statement, error) {
	query := r.prepareStatement(objects[0], UpdateStatement, len(objects))

	var params []any
	var ids []any
	for _, v := range objects {
		p := v.Values()
		params = append(params, p...)
		ids = append(ids, v.PKValue())
	}

	params = append(params, ids...)

	stmt, err := NewStatement(r.db, UpdateStatement, query, params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stmt, nil
}

func (r *Repository) UpdateOne(object repositories.RepoObject) (repositories.Statement, error) {
	query := r.prepareStatement(object, UpdateStatement, 1)

	params := object.Values()
	params = append(params, object.PKValue())

	stmt, err := NewStatement(r.db, UpdateStatement, query, object.Values())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stmt, nil
}

func (r *Repository) UpdateOrCreate(objects []repositories.RepoObject) (repositories.Statement, error) {
	query := r.prepareStatement(objects[0], UpsertStatement, len(objects))

	var params []any
	for _, v := range objects {
		p := v.Values()
		params = append(params, p...)
	}

	stmt, err := NewStatement(r.db, UpsertStatement, query, params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stmt, nil
}

func (r *Repository) UpdateOrCreateOne(object repositories.RepoObject) (repositories.Statement, error) {
	query := r.prepareStatement(object, UpsertStatement, 1)

	stmt, err := NewStatement(r.db, UpsertStatement, query, object.Values())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stmt, nil
}

func (r *Repository) Delete(objects []repositories.RepoObject) (repositories.Statement, error) {
	var objs []repositories.SoftDeletable
	softDeletable := true
	for _, v := range objects {
		o, ok := v.(repositories.SoftDeletable)
		if !ok {
			softDeletable = false
			break
		}

		objs = append(objs, o)
	}

	if softDeletable {
		return r.SoftDelete(objs)
	}

	return r.HardDelete(objects)
}

func (r *Repository) DeleteOne(object repositories.RepoObject) (repositories.Statement, error) {
	obj, ok := object.(repositories.SoftDeletable)
	if ok {
		return r.SoftDeleteOne(obj)
	}

	return r.HardDeleteOne(object)
}

func (r *Repository) SoftDelete(objects []repositories.SoftDeletable) (repositories.Statement, error) {
	query := r.prepareSoftDeleteStatement(objects[0], len(objects))

	var ids []any
	for _, v := range objects {
		ids = append(ids, v.PKValue())
	}

	stmt, err := NewStatement(r.db, DeleteStatement, query, ids)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stmt, nil
}

func (r *Repository) SoftDeleteOne(object repositories.SoftDeletable) (repositories.Statement, error) {
	query := r.prepareSoftDeleteStatement(object, 1)

	stmt, err := NewStatement(r.db, DeleteStatement, query, []any{object.PKValue()})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stmt, nil
}

func (r *Repository) HardDelete(objects []repositories.RepoObject) (repositories.Statement, error) {
	query := r.prepareStatement(objects[0], DeleteStatement, len(objects))

	var ids []any
	for _, v := range objects {
		ids = append(ids, v.PKValue())
	}

	stmt, err := NewStatement(r.db, DeleteStatement, query, ids)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stmt, nil
}

func (r *Repository) HardDeleteOne(object repositories.RepoObject) (repositories.Statement, error) {
	query := r.prepareStatement(object, DeleteStatement, 1)

	stmt, err := NewStatement(r.db, DeleteStatement, query, []any{object.PKValue()})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stmt, nil
}

func (r *Repository) FromDomainObject(object pkg.DomainObject) (repositories.RepoObject, error) {
	switch reflect.TypeOf(object).Name() {
	default:
		return nil, errors.Errorf("%s is not supported by the Postgres repository", reflect.TypeOf(object).Name())
	}
}
