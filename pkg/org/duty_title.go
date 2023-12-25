package org

import (
	"github.com/google/uuid"
)

type DutyTitle struct {
	id         uuid.UUID
	code       string
	title      string
	memberType MemberType
	minGrade   *Grade
	maxGrade   *Grade
}

func NewDutyTitle(
	id uuid.UUID,
	code string,
	title string,
	memberType MemberType,
	minGrade *Grade,
	maxGrade *Grade,
) DutyTitle {
	return DutyTitle{
		id:         id,
		code:       code,
		title:      title,
		memberType: memberType,
		minGrade:   minGrade,
		maxGrade:   maxGrade,
	}
}

func (dt DutyTitle) ID() uuid.UUID {
	return dt.id
}

func (dt DutyTitle) Code() string {
	return dt.code
}

func (dt DutyTitle) Title() string {
	return dt.title
}

func (dt DutyTitle) MemberType() MemberType {
	return dt.memberType
}

func (dt DutyTitle) MinGrade() *Grade {
	return dt.minGrade
}

func (dt DutyTitle) MaxGrade() *Grade {
	return dt.maxGrade
}
