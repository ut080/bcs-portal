package org

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Member struct {
	id             uuid.UUID
	capid          uint
	lastName       string
	firstName      string
	memberType     MemberType
	grade          Grade
	active         bool
	joinDate       *time.Time
	rankDate       *time.Time
	expirationDate *time.Time
}

func NewMember(
	id uuid.UUID,
	capid uint,
	lastName string,
	firstName string,
	memberType MemberType,
	grade Grade,
	active bool,
	joinDate *time.Time,
	rankDate *time.Time,
	expirationDate *time.Time,
) Member {
	return Member{
		id:             id,
		capid:          capid,
		lastName:       lastName,
		firstName:      firstName,
		memberType:     memberType,
		grade:          grade,
		active:         active,
		joinDate:       joinDate,
		rankDate:       rankDate,
		expirationDate: expirationDate,
	}
}

func (m Member) ID() uuid.UUID {
	return m.id
}

func (m Member) CAPID() uint {
	return m.capid
}

func (m Member) LastName() string {
	return m.lastName
}

func (m Member) FirstName() string {
	return m.firstName
}

func (m Member) MemberType() MemberType {
	return m.memberType
}

func (m Member) Grade() Grade {
	return m.grade
}

func (m Member) Active() bool {
	return m.active
}

func (m Member) JoinDate() *time.Time {
	return m.joinDate
}

func (m Member) RankDate() *time.Time {
	return m.rankDate
}

func (m Member) ExpirationDate() *time.Time {
	return m.expirationDate
}

func (m Member) FullName() string {
	return fmt.Sprintf("%s %s %s", &m.grade, m.firstName, m.lastName)
}

func (m Member) String() string {
	return fmt.Sprintf("%s, %s, %s", m.lastName, m.firstName, &m.grade)
}
