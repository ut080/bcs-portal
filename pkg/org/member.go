package org

import (
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

type Member struct {
	CAPID          uint
	LastName       string
	FirstName      string
	MemberType     MemberType
	Grade          Grade
	JoinDate       time.Time
	RankDate       time.Time
	ExpirationDate time.Time
}

func (m Member) String() string {
	return fmt.Sprintf("%s, %s, %s", m.LastName, m.FirstName, m.Grade)
}

func (m Member) FullName() string {
	return fmt.Sprintf("%s %s %s", m.Grade, m.FirstName, m.LastName)
}

type MemberGroup struct {
	Name          string
	Cadets        []Member
	CadetSponsors []Member
	Seniors       []Member
}

func capidSet(members map[uint]Member) (s mapset.Set[uint]) {
	s = mapset.NewSet[uint]()

	for u, _ := range members {
		s.Add(u)
	}

	return s
}

func NewUnassignedMemberGroup(members map[uint]Member, assigned mapset.Set[uint], inactive mapset.Set[uint]) (mg MemberGroup) {
	mg.Name = "Unassigned"

	diff := capidSet(members).Difference(inactive).Difference(assigned)

	for capid := range diff.Iter() {
		mbr := members[capid]
		switch mbr.MemberType {
		case CadetMember:
			mg.Cadets = append(mg.Cadets, mbr)
		case CadetSponsorMember:
			mg.CadetSponsors = append(mg.CadetSponsors, mbr)
		case SeniorMember:
			mg.Seniors = append(mg.Seniors, mbr)
		}
	}

	return mg
}
