package org

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Member struct {
	id             uuid.UUID
	capid          uint
	lastName       string
	firstName      string
	memberType     MemberType
	grade          Grade
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
	joinDate *time.Time,
	rankDate *time.Time,
	expirationDate *time.Time,
) (Member, error) {
	if (capid != 0) && (capid < 100000) {
		return Member{}, errors.Errorf("invalid CAPID for normal member: %d", capid)
	}

	return Member{
		id:             id,
		capid:          capid,
		lastName:       lastName,
		firstName:      firstName,
		memberType:     memberType,
		grade:          grade,
		joinDate:       joinDate,
		rankDate:       rankDate,
		expirationDate: expirationDate,
	}, nil
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
