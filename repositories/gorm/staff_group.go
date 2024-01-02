package gorm

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/org"
)

type StaffGroup struct {
	ID                uuid.UUID
	Name              string
	StaffSupergroupID uuid.UUID
	LeaderID          uuid.UUID
	Leader            DutyAssignment
	Subgroups         []StaffGroup
}

func (sg *StaffGroup) FromDomainObject(object pkg.DomainObject) error {
	group, ok := object.(org.StaffGroup)
	if !ok {
		return errors.New("attempt to pass non-org.StaffGroup object to StaffGroup.FromDomainObject")
	}

	leader := DutyAssignment{}
	err := leader.FromDomainObject(group.Leader())

	var subgroups []StaffGroup
	for _, v := range group.Subgroups() {
		gp := StaffGroup{StaffSupergroupID: group.ID()}
		err = gp.FromDomainObject(v)
		if err != nil {
			return errors.WithStack(err)
		}

		subgroups = append(subgroups, gp)
	}

	sg.ID = group.ID()
	sg.Name = group.Name()
	sg.LeaderID = group.Leader().ID()
	sg.Leader = leader
	sg.Subgroups = subgroups

	return nil
}

func (sg *StaffGroup) ToDomainObject() pkg.DomainObject {
	obj := sg.Leader.ToDomainObject()
	leader := obj.(org.DutyAssignment)

	var subgroups []org.StaffGroup
	for _, v := range sg.Subgroups {
		obj = v.ToDomainObject()
		group := obj.(org.StaffGroup)
		subgroups = append(subgroups, group)
	}

	return org.NewStaffGroup(
		sg.ID,
		sg.Name,
		subgroups,
		leader,
	)
}

func (sg *StaffGroup) BeforeCreate(tx *gorm.DB) error {
	if sg.ID != uuid.Nil {
		return nil
	}

	var err error
	sg.ID, err = uuid.NewRandom()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
