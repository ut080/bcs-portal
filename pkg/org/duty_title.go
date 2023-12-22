package org

import (
	"github.com/google/uuid"
)

type DutyTitle struct {
	id         uuid.UUID
	Title      string
	MemberType MemberType
	MinGrade   *Grade
	MaxGrade   *Grade
}

func NewDutyTitle(
	id uuid.UUID,
	title string,
	memberType MemberType,
	minGrade *Grade,
	maxGrade *Grade,
) DutyTitle {
	return DutyTitle{
		id:         id,
		Title:      title,
		MemberType: memberType,
		MinGrade:   minGrade,
		MaxGrade:   maxGrade,
	}
}

func (dt DutyTitle) ID() uuid.UUID {
	return dt.id
}
