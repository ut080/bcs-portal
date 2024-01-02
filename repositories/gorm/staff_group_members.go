package gorm

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/org"
)

type StaffGroupMember struct {
	StaffGroupID     uuid.UUID
	DutyAssignmentID uuid.UUID
	DutyAssignment   DutyAssignment
}

func (s *StaffGroupMember) FromDomainObject(object pkg.DomainObject) error {
	if s.StaffGroupID == uuid.Nil {
		return errors.New("attempt to call StaffGroup.FromDomainObject without first setting StaffGroupID")
	}

	group, ok := object.(org.DutyAssignment)
	if !ok {
		return errors.New("attempt to pass non-org.DutyAssignment object to StaffGroupMember.FromDomainObject")
	}

	assignment := DutyAssignment{}
	err := assignment.FromDomainObject(group)
	if err != nil {
		return errors.WithStack(err)
	}

	s.DutyAssignmentID = group.ID()
	s.DutyAssignment = assignment

	return nil
}

func (s *StaffGroupMember) ToDomainObject() pkg.DomainObject {
	return s.DutyAssignment.ToDomainObject()
}
