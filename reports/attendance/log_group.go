package attendance

import (
	"fmt"
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/sbani/go-humanizer/numbers"

	"github.com/ut080/bcs-portal/pkg/org"
)

const (
	logGroupBeginTemplate = `\begin{xltabular}{\textwidth}{|r|X|c|c|c|c|}
\hline
\multicolumn{6}{|c|}{\textbf{$(LOG_GROUP)}} \\
\hline
\multicolumn{1}{|c|}{\textbf{CAPID}} & 
\multicolumn{1}{|c|}{\textbf{Name}}  & 
\textbf{P}                           & 
\textbf{E}                           & 
\textbf{ID}                          & 
\textbf{U}                           \\
\hline
\endhead
`

	logGroupEndTemplate = `\end{xltabular}`

	logSubGroupTemplate = `\multicolumn{6}{|l|}{\textbf{$(LOG_SUB_GROUP)}} \\
\hline
`
)

type LogGroup struct {
	Name           string
	SubGroups      []LogSubGroup
	breakBeforeLog bool
}

func NewLogGroupFromStaffGroup(group org.StaffGroup, ignore *mapset.Set[uint]) LogGroup {
	lg := LogGroup{Name: group.Name}

	for _, subGroup := range group.SubGroups {
		lg.SubGroups = append(lg.SubGroups, NewLogSubGroupFromStaffSubGroup(subGroup, ignore))
	}

	return lg
}

func NewLogGroupFromFlight(flight org.Flight, ignore *mapset.Set[uint]) (lg LogGroup) {
	lg.Name = flight.Name

	fltStaff := LogSubGroup{Name: "Flight Staff"}
	if flight.FlightCommander.Assignee != nil && !(*ignore).Contains(flight.FlightCommander.Assignee.CAPID) {
		fltStaff.Members = append(fltStaff.Members, NewMemberFromDomainMember(*flight.FlightCommander.Assignee))
		(*ignore).Add(flight.FlightCommander.Assignee.CAPID)
	}

	if flight.FlightSergeant.Assignee != nil && !(*ignore).Contains(flight.FlightSergeant.Assignee.CAPID) {
		fltStaff.Members = append(fltStaff.Members, NewMemberFromDomainMember(*flight.FlightSergeant.Assignee))
		(*ignore).Add(flight.FlightSergeant.Assignee.CAPID)
	}

	lg.SubGroups = append(lg.SubGroups, fltStaff)

	for i, element := range flight.Elements {
		lg.SubGroups = append(lg.SubGroups, NewLogSubGroupFromElement(element, i+1, ignore))
	}

	return lg
}

func NewLogGroupFromMemberGroup(memberGroup org.MemberGroup) (lg LogGroup) {
	lg = LogGroup{Name: memberGroup.Name, breakBeforeLog: true}

	seniors := LogSubGroup{Name: "Seniors"}
	for _, senior := range memberGroup.Seniors {
		seniors.Members = append(seniors.Members, NewMemberFromDomainMember(senior))
	}
	slices.SortFunc(seniors.Members, CompareMember)

	cadetSponsors := LogSubGroup{Name: "Cadet Sponsors"}
	for _, csm := range memberGroup.CadetSponsors {
		cadetSponsors.Members = append(cadetSponsors.Members, NewMemberFromDomainMember(csm))
	}
	slices.SortFunc(cadetSponsors.Members, CompareMember)

	cadets := LogSubGroup{Name: "Cadets"}
	for _, cadet := range memberGroup.Cadets {
		cadets.Members = append(cadets.Members, NewMemberFromDomainMember(cadet))
	}
	slices.SortFunc(cadets.Members, CompareMember)

	lg.SubGroups = append(lg.SubGroups, seniors, cadetSponsors, cadets)

	return lg
}

func (lg LogGroup) BreakBeforeLog() bool {
	return lg.breakBeforeLog
}

// Height represents roughly how high this group will be when it is rendered by LaTeX in the final pdf.
// A group will be the sum of heights of its subgroups, plus 15pts for the title row and about 20pts for the space
// between tables. Hence, the math: 35 + sum(heights_of_subgroups).
func (lg LogGroup) Height() (h uint) {
	h = 35
	for _, group := range lg.SubGroups {
		h += group.Height()
	}

	return h
}

func (lg LogGroup) LaTeX() (latex string) {
	latex = strings.Replace(logGroupBeginTemplate, "$(LOG_GROUP)", lg.Name, 1)

	for _, group := range lg.SubGroups {
		latex += group.LaTeX()
	}

	latex += logGroupEndTemplate

	return latex
}

type LogSubGroup struct {
	Name    string
	Members []Member
}

func NewLogSubGroupFromStaffSubGroup(subgroup org.StaffSubGroup, ignore *mapset.Set[uint]) LogSubGroup {
	lsg := LogSubGroup{Name: subgroup.Name}

	if subgroup.Leader.Assignee != nil && !(*ignore).Contains(subgroup.Leader.Assignee.CAPID) {
		lsg.Members = append(lsg.Members, NewMemberFromDomainMember(*subgroup.Leader.Assignee))
		(*ignore).Add(subgroup.Leader.Assignee.CAPID)
	}

	// We want subordinates in the group to come after the leader, but be listed in alphabetical order.
	// So, we will load the subordinates into a temporary slice that we can sort later.
	var m []Member
	for _, report := range subgroup.DirectReports {
		if report.Assignee != nil && !(*ignore).Contains(report.Assignee.CAPID) {
			m = append(m, NewMemberFromDomainMember(*report.Assignee))
			(*ignore).Add(report.Assignee.CAPID)
		}
	}

	slices.SortFunc(m, CompareMember)
	lsg.Members = append(lsg.Members, m...)

	return lsg
}

func NewLogSubGroupFromElement(element org.Element, elementNumber int, ignore *mapset.Set[uint]) (lsg LogSubGroup) {
	lsg.Name = fmt.Sprintf("%s Element", numbers.Ordinalize(elementNumber))

	if element.ElementLeader.Assignee != nil && !(*ignore).Contains(element.ElementLeader.Assignee.CAPID) {
		lsg.Members = append(lsg.Members, NewMemberFromDomainMember(*element.ElementLeader.Assignee))
		(*ignore).Add(element.ElementLeader.Assignee.CAPID)
	}

	if element.AsstElementLeader.Assignee != nil && !(*ignore).Contains(element.AsstElementLeader.Assignee.CAPID) {
		lsg.Members = append(lsg.Members, NewMemberFromDomainMember(*element.AsstElementLeader.Assignee))
		(*ignore).Add(element.AsstElementLeader.Assignee.CAPID)
	}

	// We want element members in the group to come after the EL and asst. EL, but be listed in alphabetical order.
	// So, we will load the element members into a temporary slice that we can sort later.
	var m []Member
	for _, member := range element.Members {
		if !(*ignore).Contains(member.CAPID) {
			m = append(m, NewMemberFromDomainMember(member))
			(*ignore).Add(member.CAPID)
		}
	}

	slices.SortFunc(m, CompareMember)
	lsg.Members = append(lsg.Members, m...)

	return lsg
}

// Height represents roughly how high this subgroup will be when it is rendered by LaTeX in the final pdf.
// Each member row of the barcode log is roughly 30pts high and each title block is roughly 15 points high, hence the
// math: (number_of_members * 30) + 15
func (lsg LogSubGroup) Height() (height uint) {
	height = uint((len(lsg.Members) * 30) + 15)
	return height
}

func (lsg LogSubGroup) LaTeX() (latex string) {
	// Skip this subgroup if it doesn't contain any
	if len(lsg.Members) == 0 {
		return latex
	}

	latex = strings.Replace(logSubGroupTemplate, "$(LOG_SUB_GROUP)", lsg.Name, 1)

	for _, member := range lsg.Members {
		latex += member.LaTeX()
	}

	return latex
}
