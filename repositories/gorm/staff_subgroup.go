package gorm

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/org"
)

type StaffSubgroup struct {
	ID            uuid.UUID
	Name          string
	StaffGroupID  uuid.UUID
	LeaderID      uuid.UUID
	Leader        DutyAssignment
	DirectReports []StaffSubgroupDirectReport
}

func (ssg *StaffSubgroup) FromDomainObject(object pkg.DomainObject) error {
	if ssg.StaffGroupID == uuid.Nil {
		return errors.New("attempt to convert StaffSubgroup domain object without setting StaffGroupID")
	}

	staffSubGroup, ok := object.(org.StaffSubgroup)
	if !ok {
		return errors.New("not a valid domain StaffSubgroup object")
	}

	leader := DutyAssignment{}
	err := leader.FromDomainObject(staffSubGroup.Leader)
	if err != nil {
		return errors.WithStack(err)
	}

	var directReports []StaffSubgroupDirectReport
	for _, directReport := range staffSubGroup.DirectReports {
		dr := StaffSubgroupDirectReport{StaffSubgroupID: staffSubGroup.ID()}
		err = dr.FromDomainObject(directReport)
		if err != nil {
			return errors.WithStack(err)
		}

		directReports = append(directReports, dr)
	}

	ssg.ID = staffSubGroup.ID()
	ssg.Name = staffSubGroup.Name
	ssg.LeaderID = staffSubGroup.Leader.ID()
	ssg.Leader = leader
	ssg.DirectReports = directReports

	return nil
}

func (ssg *StaffSubgroup) ToDomainObject() pkg.DomainObject {
	leader := ssg.Leader.ToDomainObject().(org.DutyAssignment)

	var directReports []org.DutyAssignment
	for _, directReport := range ssg.DirectReports {
		dr := directReport.ToDomainObject().(org.DutyAssignment)
		directReports = append(directReports, dr)
	}

	return org.NewStaffSubgroup(ssg.ID, ssg.Name, leader, directReports)
}

type StaffSubgroupDirectReport struct {
	StaffSubgroupID  uuid.UUID
	DutyAssignmentID uuid.UUID
	DutyAssignment   DutyAssignment
}

func (ssgdr *StaffSubgroupDirectReport) FromDomainObject(object pkg.DomainObject) error {
	if ssgdr.StaffSubgroupID == uuid.Nil {
		return errors.New("attempt to convert StaffSubgroup domain object without setting StaffGroupID")
	}

	dutyAssignment, ok := object.(org.DutyAssignment)
	if !ok {
		return errors.New("not a valid domain DutyAssignment object")
	}

	da := DutyAssignment{}
	err := da.FromDomainObject(dutyAssignment)
	if err != nil {
		return errors.WithStack(err)
	}

	ssgdr.DutyAssignmentID = dutyAssignment.ID()
	ssgdr.DutyAssignment = da

	return nil
}

func (ssgdr *StaffSubgroupDirectReport) ToDomainObject() pkg.DomainObject {
	return ssgdr.DutyAssignment.ToDomainObject()
}
