package gorm_org

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/org"
)

type Member struct {
	ID             uuid.UUID      `gorm:"column:id;type:uuid;primary_key"`
	CAPID          *uint          `gorm:"column:capid"`
	LastName       string         `gorm:"column:last_name"`
	FirstName      string         `gorm:"column:first_name"`
	MemberType     org.MemberType `gorm:"column:member_type"`
	Grade          org.Grade      `gorm:"column:grade"`
	Active         bool           `gorm:"column:active"`
	JoinDate       *time.Time     `gorm:"column:join_date"`
	RankDate       *time.Time     `gorm:"column:rank_date"`
	ExpirationDate *time.Time     `gorm:"column:expiration_date"`
}

func (m *Member) FromDomainObject(object pkg.DomainObject) error {
	mbr, ok := object.(org.Member)
	if !ok {
		return errors.New("attempt to pass non-org.Member object to Member.FromDomainObject")
	}

	capid := mbr.CAPID()
	if capid == 0 {
		m.CAPID = nil
	} else {
		m.CAPID = &capid
	}

	m.ID = mbr.ID()
	m.LastName = mbr.LastName()
	m.FirstName = mbr.FirstName()
	m.MemberType = mbr.MemberType()
	m.Grade = mbr.Grade()
	m.Active = mbr.Active()
	m.JoinDate = mbr.JoinDate()
	m.RankDate = mbr.RankDate()
	m.ExpirationDate = mbr.ExpirationDate()

	return nil
}

func (m *Member) ToDomainObject() pkg.DomainObject {
	var capid uint
	if m.CAPID != nil {
		capid = *m.CAPID
	}

	return org.NewMember(
		m.ID,
		capid,
		m.LastName,
		m.FirstName,
		m.MemberType,
		m.Grade,
		m.Active,
		m.JoinDate,
		m.RankDate,
		m.ExpirationDate,
	)
}

func (m *Member) BeforeCreate(tx *gorm.DB) error {
	if m.ID != uuid.Nil {
		return nil
	}

	var err error
	m.ID, err = uuid.NewRandom()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
