package pkg

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/pkg/errors"
)

type TableOfOrganization struct {
	StaffGroups    []StaffGroup
	Flights        []Flight
	Unassigned     MemberGroup
	Inactive       MemberGroup
	InactiveCAPIDs mapset.Set[uint]
}

func (to *TableOfOrganization) PopulateMemberData(members map[uint]Member) error {
	assigned := mapset.NewSet[uint]()

	var staffGroups []StaffGroup
	for _, group := range to.StaffGroups {
		err := group.PopulateMemberData(members, &assigned)
		if err != nil {
			// TODO: Instead of halting on error, continue to populate and return a slice of errors
			return errors.WithStack(err)
		}

		staffGroups = append(staffGroups, group)
	}

	to.StaffGroups = staffGroups

	var flights []Flight
	for _, flight := range to.Flights {
		err := flight.PopulateMemberData(members, &assigned)
		if err != nil {
			// TODO: Instead of halting on error, continue to populate and return a slice of errors
			return errors.WithStack(err)
		}
		flights = append(flights, flight)
	}

	to.Flights = flights

	to.Unassigned = NewUnassignedMemberGroup(members, assigned, to.InactiveCAPIDs)

	return nil
}
