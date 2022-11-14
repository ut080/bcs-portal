package yaml

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/pkg/errors"

	"github.com/derhabicht/herriman/pkg"
)

type TableOfOrganization struct {
	Unit        Unit         `yaml:"unit"`
	StaffGroups []StaffGroup `yaml:"staff_groups"`
	Flights     []Flight     `yaml:"flights"`
	Inactive    []uint       `yaml:"inactive"`
}

type Unit struct {
	Charter string `yaml:"charter"`
	Name    string `yaml:"name"`
}

func domainDutyAssignment(config map[string]pkg.DutyAssignment, da DutyAssignment) (pkg.DutyAssignment, error) {
	// Try to find the assignment in the configuration
	assignment, ok := config[da.OfficeSymbol]
	if !ok {
		return pkg.DutyAssignment{}, DutyAssignmentNotDefinedError{OfficeSymbol: da.OfficeSymbol}
	}

	// If a CAPID is given in the YAML duty assignment, then we need to create a pointer to an domain.Member
	var member *pkg.Member
	if da.AsigneeCAPID != nil {
		member = new(pkg.Member)
		member.CAPID = *da.AsigneeCAPID
	}

	// Should be equivalent to assignment.Assignee = nil if da.AssigneeCAPID is nil
	assignment.Assignee = member

	return assignment, nil
}

func handleDomainStaffGroupError(err error, group string, subgroup string) ([]pkg.StaffGroup, error) {
	if e, ok := err.(DutyAssignmentNotDefinedError); ok {
		e.StaffGroup = group
		e.StaffSubGroup = subgroup
		return nil, e
	}

	return nil, err
}

func domainStaffGroups(config map[string]pkg.DutyAssignment, groups []StaffGroup) ([]pkg.StaffGroup, error) {
	var domainGroupList []pkg.StaffGroup

	for _, group := range groups {
		domainGroup := pkg.StaffGroup{Name: group.Group}
		for _, subgroup := range group.Subgroups {
			leader, err := domainDutyAssignment(config, subgroup.Leader)
			if err != nil {
				return handleDomainStaffGroupError(err, group.Group, subgroup.Subgroup)
			}

			domainSubgroup := pkg.StaffSubGroup{
				Name:   subgroup.Subgroup,
				Leader: leader,
			}

			for _, report := range subgroup.DirectReports {
				dutyAssignment, err := domainDutyAssignment(config, report)
				if err != nil {
					return handleDomainStaffGroupError(err, group.Group, subgroup.Subgroup)
				}

				domainSubgroup.DirectReports = append(domainSubgroup.DirectReports, dutyAssignment)
			}

			domainGroup.SubGroups = append(domainGroup.SubGroups, domainSubgroup)
		}

		domainGroupList = append(domainGroupList, domainGroup)
	}

	return domainGroupList, nil
}

func handleDomainElementsError(err error, flight string, element string) ([]pkg.Element, error) {
	if e, ok := err.(DutyAssignmentNotDefinedError); ok {
		e.StaffGroup = flight
		e.StaffSubGroup = element
		return nil, e
	}

	return nil, err
}

func domainElements(config map[string]pkg.DutyAssignment, flight Flight) ([]pkg.Element, error) {
	var elementList []pkg.Element

	for _, element := range flight.Elements {

		domainEL, err := domainDutyAssignment(config, element.ElementLeader)
		if err != nil {
			return handleDomainElementsError(err, flight.Name, element.Name)
		}

		domainAsstEL, err := domainDutyAssignment(config, element.AsstElementLeader)
		if err != nil {
			return handleDomainElementsError(err, flight.Name, element.Name)
		}

		domainElement := pkg.Element{ElementLeader: domainEL, AsstElementLeader: domainAsstEL}

		for _, member := range element.Members {
			m := pkg.Member{CAPID: member}
			domainElement.Members = append(domainElement.Members, m)
		}

		elementList = append(elementList, domainElement)
	}

	return elementList, nil
}

func handleDomainFlightsError(err error, flight string) ([]pkg.Flight, error) {
	if e, ok := err.(DutyAssignmentNotDefinedError); ok {
		e.StaffGroup = flight
		return nil, e
	}

	return nil, err
}

func domainFlights(config map[string]pkg.DutyAssignment, flights []Flight) ([]pkg.Flight, error) {
	var flightList []pkg.Flight

	for _, flight := range flights {

		domainFltCC, err := domainDutyAssignment(config, flight.Commander)
		if err != nil {
			return handleDomainFlightsError(err, flight.Name)
		}

		domainFltCCF, err := domainDutyAssignment(config, flight.FlightSergeant)
		if err != nil {
			return handleDomainFlightsError(err, flight.Name)
		}

		orgFlight := pkg.Flight{
			Name:            flight.Name,
			FlightCommander: domainFltCC,
			FlightSergeant:  domainFltCCF,
		}

		orgElements, err := domainElements(config, flight)
		if err != nil {
			return nil, err
		}

		orgFlight.Elements = orgElements

		flightList = append(flightList, orgFlight)
	}

	return flightList, nil
}

func (to TableOfOrganization) DomainTableOfOrganization(config map[string]pkg.DutyAssignment) (pkg.TableOfOrganization, error) {
	staffGroups, err := domainStaffGroups(config, to.StaffGroups)
	if err != nil {
		return pkg.TableOfOrganization{}, errors.WithStack(err)
	}

	flights, err := domainFlights(config, to.Flights)
	if err != nil {
		return pkg.TableOfOrganization{}, errors.WithStack(err)
	}

	inactive := mapset.NewSet[uint]()
	for _, i := range to.Inactive {
		inactive.Add(i)
	}

	domainTO := pkg.TableOfOrganization{
		StaffGroups:    staffGroups,
		Flights:        flights,
		InactiveCAPIDs: inactive,
	}

	return domainTO, nil
}

type DutyAssignmentNotDefinedError struct {
	OfficeSymbol  string
	StaffGroup    string
	StaffSubGroup string
}

func (err DutyAssignmentNotDefinedError) Error() string {
	return fmt.Sprintf(
		"failed to find definiton for duty assignment %s in %s/%s",
		err.OfficeSymbol,
		err.StaffGroup,
		err.StaffSubGroup,
	)
}
