package org

import (
	"github.com/google/uuid"
)

type StaffSubgroup struct {
	id            uuid.UUID
	name          string
	leader        DutyAssignment
	directReports []DutyAssignment
}

func NewStaffSubgroup(
	id uuid.UUID,
	name string,
	leader DutyAssignment,
	directReports []DutyAssignment,
) StaffSubgroup {
	return StaffSubgroup{
		id:            id,
		name:          name,
		leader:        leader,
		directReports: directReports,
	}
}

func (ssg StaffSubgroup) ID() uuid.UUID {
	return ssg.id
}

func (ssg StaffSubgroup) Name() string {
	return ssg.name
}

func (ssg StaffSubgroup) Leader() DutyAssignment {
	return ssg.leader
}

func (ssg StaffSubgroup) DirectReports() []DutyAssignment {
	return ssg.directReports
}
