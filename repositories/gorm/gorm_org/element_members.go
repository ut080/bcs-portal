package gorm_org

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
)

type ElementMember struct {
	ElementID uuid.UUID
	MemberID  uuid.UUID
	Member    Member
}

func (e *ElementMember) FromDomainObject(object pkg.DomainObject) error {
	if e.ElementID == uuid.Nil {
		return errors.New("attempt to call ElementMember.FromDomainObject without first setting ElementID")
	}

	mbr := Member{}
	err := mbr.FromDomainObject(object)
	if err != nil {
		return errors.WithStack(err)
	}

	e.MemberID = mbr.ID
	e.Member = mbr

	return nil
}

func (e *ElementMember) ToDomainObject() pkg.DomainObject {
	return e.Member.ToDomainObject()
}
