package gorm_org

import (
	"reflect"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/repositories"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) FromDomainObject(object pkg.DomainObject) (repositories.RepoObject, error) {
	t := reflect.TypeOf(object).Name()
	switch t {
	case "DutyAssignment":
		da := DutyAssignment{}
		err := da.FromDomainObject(object)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &da, nil
	case "Element":
		return nil, errors.New("Element objects cannot be created independent of Flight objects")
	case "Flight":
		f := Flight{}
		err := f.FromDomainObject(object)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &f, nil
	case "Member":
		m := Member{}
		err := m.FromDomainObject(object)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &m, nil
	case "StaffGroup":
		sg := StaffGroup{}
		err := sg.FromDomainObject(object)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &sg, nil
	case "Unit":
		u := Unit{}
		err := u.FromDomainObject(object)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &u, nil
	default:
		return nil, errors.Errorf("%s is not supported by the GORM Repo", t)
	}
}

func (r *Repository) FromDomainObjects(objects []pkg.DomainObject) ([]repositories.RepoObject, error) {
	var objs []repositories.RepoObject
	for _, v := range objects {
		obj, err := r.FromDomainObject(v)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

func (r *Repository) Create(objects []repositories.RepoObject) ([]pkg.DomainObject, error) {
	res := r.db.Create(objects)
	if res.Error != nil {
		return nil, errors.WithStack(res.Error)
	}

	var results []pkg.DomainObject
	for _, v := range objects {
		res := v.ToDomainObject()
		results = append(results, res)
	}

	return results, nil
}

func (r *Repository) CreateOne(object repositories.RepoObject) (pkg.DomainObject, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Fetch(object repositories.RepoObject) ([]pkg.DomainObject, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Update(objects []repositories.RepoObject) ([]pkg.DomainObject, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) UpdateOne(object repositories.RepoObject) (pkg.DomainObject, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) UpdateOrCreate(objects []repositories.RepoObject) ([]pkg.DomainObject, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) UpdateOrCreateOne(object repositories.RepoObject) (pkg.DomainObject, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Delete(objects []repositories.RepoObject) ([]pkg.DomainObject, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) DeleteOne(object repositories.RepoObject) (pkg.DomainObject, error) {
	//TODO implement me
	panic("implement me")
}
