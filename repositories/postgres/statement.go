package postgres

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/repositories"
)

type StatementType int

const (
	InsertStatement StatementType = iota
	SelectStatement
	UpdateStatement
	UpsertStatement
	DeleteStatement
)

type Statement struct {
	repoObject    repositories.RepoObject
	statementType StatementType
	statement     *sql.Stmt
	parameters    []any
}

func NewStatement(db *sql.DB, st StatementType, query string, params []any) (*Statement, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Statement{
		statementType: st,
		statement:     stmt,
		parameters:    params,
	}, nil
}

func (s *Statement) mapRow(rows *sql.Rows) (map[string]any, error) {
	cols := s.repoObject.ColNames()
	vals := make([]any, len(cols))

	err := rows.Scan(&vals)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	mapping := make(map[string]any)
	for i, v := range cols {
		mapping[v] = vals[i]
	}

	return mapping, nil
}

func (s *Statement) runQuery() (*sql.Rows, error) {
	if s.statementType != SelectStatement {
		return nil, errors.New("improper use of Statement.QueryFirst() for a non-SELECT query")
	}

	rows, err := s.statement.Query(s.parameters...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return rows, nil
}

func (s *Statement) Execute() (int64, error) {
	if s.statementType == SelectStatement {
		return -1, errors.New("improper use of Statement.Execute() for a SELECT query")
	}

	res, err := s.statement.Exec(s.parameters...)
	if err != nil {
		return -1, errors.WithStack(err)
	}

	return res.RowsAffected()
}

func (s *Statement) QueryFirst() (pkg.DomainObject, error) {
	rows, err := s.runQuery()
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logging.Error().Err(errors.WithStack(err)).Msg("Failed to close rows from query.")
		}
	}(rows)

	ok := rows.Next()
	if !ok {
		return nil, errors.New("no results returned from query")
	}

	mapping, err := s.mapRow(rows)
	err = s.repoObject.FromValues(mapping)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return s.repoObject.ToDomainObject(), nil
}

func (s *Statement) QueryAll() ([]pkg.DomainObject, error) {
	rows, err := s.runQuery()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logging.Error().Err(errors.WithStack(err)).Msg("Failed to close rows from query.")
		}
	}(rows)

	var objs []pkg.DomainObject
	for rows.Next() {
		mapping, err := s.mapRow(rows)
		err = s.repoObject.FromValues(mapping)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		objs = append(objs, s.repoObject.ToDomainObject())
	}

	return objs, nil
}
