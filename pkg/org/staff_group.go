package org

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type StaffGroup struct {
	id        uuid.UUID
	Name      string
	SubGroups []StaffSubgroup
}

func NewStaffGroup(id uuid.UUID, name string, subgroups []StaffSubgroup) StaffGroup {
	return StaffGroup{
		id:        id,
		Name:      name,
		SubGroups: subgroups,
	}
}

func (sg StaffGroup) ID() uuid.UUID {
	return sg.id
}

func (sg *StaffGroup) PopulateMemberData(members map[uint]Member, assigned *mapset.Set[uint]) (err error) {
	var subgroups []StaffSubgroup
	for _, group := range sg.SubGroups {
		err = group.PopulateMemberData(members, assigned)
		if err != nil {
			// TODO: Instead of halting on error, continue to populate and return a slice of errors
			err = errors.WithStack(err)
			return err
		}
		subgroups = append(subgroups, group)
	}

	sg.SubGroups = subgroups

	return nil
}
