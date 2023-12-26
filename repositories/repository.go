package repositories

import (
	"github.com/google/uuid"

	"github.com/ut080/bcs-portal/pkg"
)

type Statement interface {
	// Execute runs this statement against the database. No attempt is made to collect results of the statement beyond
	// error messages.
	Execute() (int64, error)

	// QueryFirst runs this statement against the database, collecting the results and returning the first result
	// as a pkg.DomainObject.
	QueryFirst() (pkg.DomainObject, error)

	// QueryAll runs this statement against the database, collecting the results and converting/returning ALL results
	// of the query.
	QueryAll() ([]pkg.DomainObject, error)
}

type Repository interface {
	// FromDomainObject finds and returns the correct RepoObject that corresponds with the given pkg.DomainObject. If
	// the correct pkg.DomainObject can't be found (usually because it isn't supported by this repository), then an
	// error is returned instead.
	FromDomainObject(pkg.DomainObject) (RepoObject, error)

	// Create will return a statement that attempts to create the slice of objects.
	Create([]RepoObject) (Statement, error)

	// CreateOne is the singular case of Create().
	CreateOne(RepoObject) (Statement, error)

	// Fetch will return a statement that queries the database. Any non-zero parameters in the passed object are used in
	// the WHERE statement.
	Fetch(RepoObject) (Statement, error)

	// Update will return a statement that attempts to update all the objects in the passed slice. Any object that does
	// not already exist in the database will return an error.
	Update([]RepoObject) (Statement, error)

	// UpdateOne is the singular form of Update().
	UpdateOne(RepoObject) (Statement, error)

	// UpdateOrCreate will attempt to update all the objects in the passed slice. If any object does not exist, the
	// repository will attempt to create it.
	UpdateOrCreate([]RepoObject) (Statement, error)

	// UpdateOrCreateOne is the singular form of UpdateOrCreate().
	UpdateOrCreateOne(RepoObject) (Statement, error)

	// Delete will attempt a soft delete on the passed objects. If the corresponding repository object is not
	// SoftDeletable, then a hard delete is executed.
	Delete([]RepoObject) (Statement, error)

	// DeleteOne is the singular form of Delete.
	DeleteOne(RepoObject) (Statement, error)

	// SoftDelete will return a statement that attempts a soft delete on the passed objects. If the RepoObject is not
	// SoftDeletable, an error is returned.
	SoftDelete([]SoftDeletable) (Statement, error)

	// SoftDeleteOne is the singular form of SoftDelete().
	SoftDeleteOne(SoftDeletable) (Statement, error)

	// HardDelete will return a statement that attempts to permanently remove the passed objects from the database.
	HardDelete([]RepoObject) (Statement, error)

	// HardDeleteOne is the singular form of HardDelete.
	HardDeleteOne(RepoObject) (Statement, error)
}

type RepoObject interface {
	// PKValue returns the value of the Primary Key ID for this RepoObject.
	PKValue() uuid.UUID

	// FromDomainObject populates this RepoObject from the given pkg.DomainObject.
	FromDomainObject(pkg.DomainObject) error

	// ToDomainObject converts this RepoObject into its corresponding pkg.DomainObject.
	ToDomainObject() pkg.DomainObject

	// FromValues populates this RepoObject from a map of values returned by the query.
	FromValues(map[string]any) error

	// Create will return a string that can be used to prepare a statement that will create this object in the database.
	Create() string

	// Fetch will return a string that can be used to prepare a statement that will query this object in the database.
	// If the eager flag is set, the returned query will include all joins necessary to populated associated objects.
	Fetch(eager bool) string

	// Update will return a string that can be used to prepare a statement that will update this object in the database.
	Update() string

	// UpdateOrCreate will return a string that can be used to prepare a statement that will update this object in the
	// database.
	UpdateOrCreate() string

	// Delete will return a string that can be used to prepare a statement that will hard delete this object from
	// the database.
	Delete() string

	// ColNames returns a list of the column names that encode this object in the data layer. This is primarily used by
	// the Statement interface to map
	ColNames() []string

	// Parameters provides a string with the correct number of placeholders that can be inserted into a parameterized
	// query. If startIdx >= 0, then a number will be appended to the placeholder starting at startIdx (this is useful
	// for, e.g. Postgres parameterized queries where the placeholders are $1, $2, ...). If indexes are appended to the
	// placeholders, this method will return the last index used (otherwise, it'll return -1).
	Parameters(placeholder string, startIdx int) (string, int)

	// Values provides a slice of values to insert into a parameterized query.
	Values() []any
}

type SoftDeletable interface {
	RepoObject

	// SoftDelete will return a string that can be used to prepare a statement that will execute a soft delete (i.e. an
	// undoable delete) on this object in the database.
	SoftDelete() string
}
