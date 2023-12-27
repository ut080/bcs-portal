package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/org"
)

type Member struct {
	ID             uuid.UUID
	CAPID          uint
	LastName       string
	FirstName      string
	MemberType     org.MemberType
	Grade          org.Grade
	JoinDate       *time.Time
	RankDate       *time.Time
	ExpirationDate *time.Time
}

func (m *Member) FromDomainObject(object pkg.DomainObject) error {
	mbr, ok := object.(org.Member)
	if !ok {
		return errors.New("non-Member object passed to Member.FromDomainObject()")
	}

	m.ID = mbr.ID()
	m.CAPID = mbr.CAPID()
	m.LastName = mbr.LastName()
	m.FirstName = mbr.FirstName()
	m.MemberType = mbr.MemberType()
	m.Grade = mbr.Grade()
	m.JoinDate = mbr.JoinDate()
	m.RankDate = mbr.RankDate()
	m.ExpirationDate = mbr.ExpirationDate()

	return nil
}

func (m *Member) ToDomainObject() pkg.DomainObject {
	mbr, err := org.NewMember(
		m.ID,
		m.CAPID,
		m.LastName,
		m.FirstName,
		m.MemberType,
		m.Grade,
		m.JoinDate,
		m.RankDate,
		m.ExpirationDate,
	)
}

func (m *Member) FromValues(values map[string]any) error {
	//TODO implement me
	panic("implement me")
}

func (m *Member) Create() string {
	//TODO implement me
	panic("implement me")
}

func (m *Member) Fetch(eager bool) string {
	//TODO implement me
	panic("implement me")
}

func (m *Member) Update() string {
	//TODO implement me
	panic("implement me")
}

func (m *Member) UpdateOrCreate() string {
	//TODO implement me
	panic("implement me")
}

func (m *Member) Delete() string {
	//TODO implement me
	panic("implement me")
}

func (m *Member) ColNames() []string {
	//TODO implement me
	panic("implement me")
}

func (m *Member) Parameters(placeholder string, startIdx int) (string, int) {
	//TODO implement me
	panic("implement me")
}

func (m *Member) Values() []any {
	//TODO implement me
	panic("implement me")
}

func (m *Member) PKValue() uuid.UUID {
	return m.ID
}
