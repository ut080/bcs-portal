package org

type DutyTitle struct {
	Title      string
	MemberType MemberType
	MinGrade   *Grade
	MaxGrade   *Grade
}

type DutyAssignment struct {
	Title        DutyTitle
	OfficeSymbol string
	Assistant    bool
	Acting       bool
	Assignee     *Member
}
