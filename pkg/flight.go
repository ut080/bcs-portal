package pkg

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/pkg/errors"
)

type Flight struct {
	Name            string
	FlightCommander DutyAssignment
	FlightSergeant  DutyAssignment
	Elements        []Element
}

func (f *Flight) PopulateMemberData(members map[uint]Member, accounted *mapset.Set[uint]) error {
	if f.FlightCommander.Assignee != nil {
		cc, ok := members[f.FlightCommander.Assignee.CAPID]
		if !ok {
			// TODO: Instead of halting on error, continue to populate and return a slice of errors
			return errors.Errorf("no member found with CAPID %d", f.FlightCommander.Assignee.CAPID)
		}
		f.FlightCommander.Assignee = &cc
		(*accounted).Add(cc.CAPID)
	}

	if f.FlightSergeant.Assignee != nil {
		ccf, ok := members[f.FlightSergeant.Assignee.CAPID]
		if !ok {
			// TODO: Instead of halting on error, continue to populate and return a slice of errors
			return errors.Errorf("no member found with CAPID %d", f.FlightSergeant.Assignee.CAPID)
		}
		f.FlightSergeant.Assignee = &ccf
		(*accounted).Add(ccf.CAPID)
	}

	var elements []Element
	for _, element := range f.Elements {
		err := element.PopulateMemberData(members, accounted)
		if err != nil {
			// TODO: Instead of halting on error, continue to populate and return a slice of errors
			return errors.WithStack(err)
		}

		elements = append(elements, element)
	}

	f.Elements = elements

	return nil
}

type Element struct {
	ElementLeader     DutyAssignment
	AsstElementLeader DutyAssignment
	Members           []Member
}

func (e *Element) PopulateMemberData(members map[uint]Member, accounted *mapset.Set[uint]) error {
	if e.ElementLeader.Assignee != nil {
		el, ok := members[e.ElementLeader.Assignee.CAPID]
		if !ok {
			// TODO: Instead of halting on error, continue to populate and return a slice of errors
			return errors.Errorf("no member found with CAPID %d", e.ElementLeader.Assignee.CAPID)
		}
		e.ElementLeader.Assignee = &el
		(*accounted).Add(el.CAPID)
	}

	if e.AsstElementLeader.Assignee != nil {
		ael, ok := members[e.AsstElementLeader.Assignee.CAPID]
		if !ok {
			// TODO: Instead of halting on error, continue to populate and return a slice of errors
			return errors.Errorf("no member found with CAPID %d", e.AsstElementLeader.Assignee.CAPID)
		}
		e.AsstElementLeader.Assignee = &ael
		(*accounted).Add(ael.CAPID)
	}

	var elementMembers []Member
	for _, member := range e.Members {
		// TODO: Instead of halting on error, continue to populate and return a slice of errors
		mbr, ok := members[member.CAPID]
		if !ok {
			return errors.Errorf("no member found with CAPID %d", member.CAPID)
		}
		elementMembers = append(elementMembers, mbr)
		(*accounted).Add(mbr.CAPID)
	}

	e.Members = elementMembers

	return nil
}
