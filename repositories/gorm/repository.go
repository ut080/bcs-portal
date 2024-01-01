package gorm

import (
	"reflect"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/repositories"
)

type Repository struct {
	db gorm.DB
}

func NewRepository(db gorm.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) FromDomainObject(object pkg.DomainObject) (repositories.RepoObject, error) {
	switch reflect.ValueOf(object).Type().Name() {
	case "Member":
		mbr := Member{}
		err := mbr.FromDomainObject(object)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &mbr, err
	default:
		return nil, errors.Errorf("%s is not supported by the GORM repository", reflect.ValueOf(object).Type().Name())
	}
}

func (r *Repository) Create(objects []repositories.RepoObject) ([]pkg.DomainObject, error) {
	result := r.db.Create(objects)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	var objs []pkg.DomainObject
	for _, v := range objects {
		objs = append(objs, v.ToDomainObject())
	}

	return objs, nil
}

func (r *Repository) Fetch(object repositories.RepoObject) (pkg.DomainObject, error) {
	result := r.db.Create(object)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	return object.ToDomainObject(), nil
}

func (r *Repository) Update(objects []repositories.RepoObject) ([]pkg.DomainObject, error) {
	result := r.db.Updates(objects)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	var objs []pkg.DomainObject
	for _, v := range objects {
		objs = append(objs, v.ToDomainObject())
	}

	return objs, nil
}

func (r *Repository) UpdateOrCreate(objects []repositories.RepoObject) ([]pkg.DomainObject, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) UpdateOrCreateOne(object repositories.RepoObject) (pkg.DomainObject, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Delete(objects []repositories.RepoObject) (pkg.DomainObject, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) DeleteOne(object repositories.RepoObject) (pkg.DomainObject, error) {
	//TODO implement me
	panic("implement me")
}
