package gorm

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/org"
)

type StaffGroup struct {
	ID        uuid.UUID
	Name      string
	Subgroups []StaffSubgroup
}

func (sg *StaffGroup) FromDomainObject(object pkg.DomainObject) error {
	staffGroup, ok := object.(org.StaffGroup)
	if !ok {
		return errors.New("not a valid domain StaffGroup object")
	}

	var subgroups []StaffSubgroup
	for _, s := range staffGroup.SubGroups {
		subgroup := StaffSubgroup{StaffGroupID: staffGroup.ID()}
		err := subgroup.FromDomainObject(s)
		if err != nil {
			return errors.WithStack(err)
		}

		subgroups = append(subgroups, subgroup)
	}

	sg.ID = staffGroup.ID()
	sg.Name = staffGroup.Name
	sg.Subgroups = subgroups

	return nil
}

func (sg *StaffGroup) ToDomainObject() pkg.DomainObject {
	var subgroups []org.StaffSubgroup
	for _, v := range sg.Subgroups {
		subgroup := v.ToDomainObject().(org.StaffSubgroup)
		subgroups = append(subgroups, subgroup)
	}

	return org.NewStaffGroup(sg.ID, sg.Name, subgroups)
}
