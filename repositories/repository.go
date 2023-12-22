package repositories

import (
	"github.com/ut080/bcs-portal/pkg"
)

type RepoObject interface {
	FromDomainObject(pkg.DomainObject) error
	ToDomainObject() pkg.DomainObject
}
