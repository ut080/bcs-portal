package pkg

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type StaffGroup struct {
	Name      string
	SubGroups []StaffSubGroup
}

func (sg *StaffGroup) PopulateMemberData(members map[uint]Member, assigned *mapset.Set[uint]) error {
	var subgroups []StaffSubGroup
	for _, group := range sg.SubGroups {
		err := group.PopulateMemberData(members, assigned)
		if err != nil {
			// TODO: Instead of halting on error, continue to populate and return a slice of errors
			return errors.WithStack(err)
		}
		subgroups = append(subgroups, group)
	}

	sg.SubGroups = subgroups

	return nil
}

type StaffSubGroup struct {
	Name          string
	Leader        DutyAssignment
	DirectReports []DutyAssignment
}

func (ssg *StaffSubGroup) PopulateMemberData(members map[uint]Member, assigned *mapset.Set[uint]) error {
	if ssg.Leader.Assignee != nil {
		leader, ok := members[ssg.Leader.Assignee.CAPID]
		if !ok {
			// TODO: Instead of halting on error, continue to populate and return a slice of errors
			return errors.Errorf("no member found with CAPID %d", ssg.Leader.Assignee.CAPID)
		}
		log.Trace().Msgf("Found %s leader: %s", ssg.Name, leader)
		ssg.Leader.Assignee = &leader
		(*assigned).Add(leader.CAPID)
	}

	var directReports []DutyAssignment
	for _, report := range ssg.DirectReports {
		if report.Assignee != nil {
			member, ok := members[report.Assignee.CAPID]
			if !ok {
				// TODO: Instead of halting on error, continue to populate and return a slice of errors
				return errors.Errorf("no member found with CAPID %d", ssg.Leader.Assignee.CAPID)
			}
			log.Trace().Msgf("Found %s assignee: %s", report.DutyTitle, member)
			report.Assignee = &member
			(*assigned).Add(member.CAPID)

			directReports = append(directReports, report)
		}
	}

	ssg.DirectReports = directReports

	return nil
}
