package repositories

import (
	"github.com/ut080/bcs-portal/pkg"
)

type Repository interface {
	// FromDomainObject finds and returns the correct RepoObject that corresponds with the given pkg.DomainObject. If
	// the correct pkg.DomainObject can't be found (usually because it isn't supported by this repository), then an
	// error is returned instead.
	FromDomainObject(pkg.DomainObject) (RepoObject, error)

	// Create will return a statement that attempts to create the slice of objects.
	Create([]RepoObject) ([]pkg.DomainObject, error)

	// CreateOne is the singular case of Create().
	CreateOne(RepoObject) (pkg.DomainObject, error)

	// Fetch will return a statement that queries the database. Any non-zero parameters in the passed object are used in
	// the WHERE statement.
	Fetch(RepoObject) ([]pkg.DomainObject, error)

	// Update will return a statement that attempts to update all the objects in the passed slice. Any object that does
	// not already exist in the database will return an error.
	Update([]RepoObject) ([]pkg.DomainObject, error)

	// UpdateOne is the singular form of Update().
	UpdateOne(RepoObject) (pkg.DomainObject, error)

	// UpdateOrCreate will attempt to update all the objects in the passed slice. If any object does not exist, the
	// repository will attempt to create it.
	UpdateOrCreate([]RepoObject) ([]pkg.DomainObject, error)

	// UpdateOrCreateOne is the singular form of UpdateOrCreate().
	UpdateOrCreateOne(RepoObject) (pkg.DomainObject, error)

	// Delete will attempt a soft delete on the passed objects. If the corresponding repository object is not
	// SoftDeletable, then a hard delete is executed.
	Delete([]RepoObject) ([]pkg.DomainObject, error)

	// DeleteOne is the singular form of Delete.
	DeleteOne(RepoObject) (pkg.DomainObject, error)
}

type RepoObject interface {
	// FromDomainObject populates this RepoObject from the given pkg.DomainObject.
	FromDomainObject(pkg.DomainObject) error

	// ToDomainObject converts this RepoObject into its corresponding pkg.DomainObject.
	ToDomainObject() pkg.DomainObject
}
