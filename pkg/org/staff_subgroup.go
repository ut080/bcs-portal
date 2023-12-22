package org

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type StaffSubgroup struct {
	id            uuid.UUID
	Name          string
	Leader        DutyAssignment
	DirectReports []DutyAssignment
}

func NewStaffSubgroup(
	id uuid.UUID,
	name string,
	leader DutyAssignment,
	directReports []DutyAssignment,
) StaffSubgroup {
	return StaffSubgroup{
		id:            id,
		Name:          name,
		Leader:        leader,
		DirectReports: directReports,
	}
}

func (ssg StaffSubgroup) ID() uuid.UUID {
	return ssg.id
}

func (ssg *StaffSubgroup) PopulateMemberData(members map[uint]Member, assigned *mapset.Set[uint]) (err error) {
	if ssg.Leader.Assignee != nil {
		leader, ok := members[ssg.Leader.Assignee.CAPID]
		if !ok {
			// TODO: Instead of halting on error, continue to populate and return a slice of errors
			err = errors.Errorf("no member found with CAPID %d", ssg.Leader.Assignee.CAPID)
			return err
		}
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
			report.Assignee = &member
			(*assigned).Add(member.CAPID)

			directReports = append(directReports, report)
		}
	}

	ssg.DirectReports = directReports

	return nil
}
