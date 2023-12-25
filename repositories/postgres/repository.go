package postgres

import (
	"database/sql"
	"fmt"
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

func (r *Repository) Create(objects []repositories.RepoObject) repositories.Statement {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) CreateOne(object repositories.RepoObject) repositories.Statement {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Fetch(object repositories.RepoObject) repositories.Statement {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Update(objects []repositories.RepoObject) repositories.Statement {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) UpdateOne(object repositories.RepoObject) repositories.Statement {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) UpdateOrCreate(objects []repositories.RepoObject) repositories.Statement {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) UpdateOrCreateOne(object repositories.RepoObject) repositories.Statement {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Delete(objects []repositories.RepoObject) repositories.Statement {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) DeleteOne(object repositories.RepoObject) repositories.Statement {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) SoftDelete(objects []repositories.RepoObject) (repositories.Statement, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) SoftDeleteOne(object repositories.RepoObject) (repositories.Statement, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) HardDelete(objects []repositories.RepoObject) repositories.Statement {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) HardDeleteOne(object repositories.RepoObject) repositories.Statement {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) FromDomainObject(object pkg.DomainObject) (repositories.RepoObject, error) {
	panic(errors.New("method not implemented"))
}