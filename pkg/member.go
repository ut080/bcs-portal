package pkg

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

type Member struct {
	CAPID      uint
	LastName   string
	FirstName  string
	MemberType MemberType
	Grade      Grade
}

func (m Member) String() string {
	return fmt.Sprintf("%s, %s, %s", m.LastName, m.FirstName, m.Grade)
}

func (m Member) FullName() string {
	return fmt.Sprintf("%s %s %s", m.Grade, m.FirstName, m.LastName)
}

type MemberGroup struct {
	Name    string
	Seniors []Member
	Cadets  []Member
}

func capidSet(members map[uint]Member) mapset.Set[uint] {
	s := mapset.NewSet[uint]()

	for u, _ := range members {
		s.Add(u)
	}

	return s
}

func NewUnassignedMemberGroup(members map[uint]Member, assigned mapset.Set[uint], inactive mapset.Set[uint]) MemberGroup {
	mg := MemberGroup{Name: "Unassigned"}

	diff := capidSet(members).Difference(inactive).Difference(assigned)

	for capid := range diff.Iter() {
		mbr := members[capid]
		switch mbr.MemberType {
		case SeniorMember:
			mg.Seniors = append(mg.Seniors, mbr)
		case CadetMember:
			mg.Cadets = append(mg.Cadets, mbr)
		}
	}

	return mg
}
