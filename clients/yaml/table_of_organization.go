package yaml

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/domain"
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

func domainDutyAssignment(config map[string]domain.DutyAssignment, yamlDA DutyAssignment) (da domain.DutyAssignment, err error) {
	// Try to find the assignment in the configuration
	da, ok := config[yamlDA.OfficeSymbol]
	if !ok {
		err = DutyAssignmentNotDefinedError{OfficeSymbol: yamlDA.OfficeSymbol}
		return da, err
	}

	// If a CAPID is given in the YAML duty assignment, then we need to create a pointer to an domain.Member
	var member *domain.Member
	if yamlDA.AsigneeCAPID != nil {
		member = new(domain.Member)
		member.CAPID = *yamlDA.AsigneeCAPID
	}

	// Should be equivalent to assignment.Assignee = nil if yamlDA.AssigneeCAPID is nil
	da.Assignee = member

	return da, nil
}

func handleDomainStaffGroupError(dsgErr error, group string, subgroup string) (err error) {
	if err, ok := dsgErr.(DutyAssignmentNotDefinedError); ok {
		err.StaffGroup = group
		err.StaffSubGroup = subgroup
		return err
	}

	return dsgErr
}

func domainStaffGroups(config map[string]domain.DutyAssignment, groups []StaffGroup) (sg []domain.StaffGroup, err error) {
	var domainGroupList []domain.StaffGroup

	for _, group := range groups {
		domainGroup := domain.StaffGroup{Name: group.Group}
		for _, subgroup := range group.Subgroups {
			leader, err := domainDutyAssignment(config, subgroup.Leader)
			if err != nil {
				err = handleDomainStaffGroupError(err, group.Group, subgroup.Subgroup)
				return nil, err
			}

			domainSubgroup := domain.StaffSubGroup{
				Name:   subgroup.Subgroup,
				Leader: leader,
			}

			for _, report := range subgroup.DirectReports {
				dutyAssignment, err := domainDutyAssignment(config, report)
				if err != nil {
					err = handleDomainStaffGroupError(err, group.Group, subgroup.Subgroup)
					return nil, err
				}

				domainSubgroup.DirectReports = append(domainSubgroup.DirectReports, dutyAssignment)
			}

			domainGroup.SubGroups = append(domainGroup.SubGroups, domainSubgroup)
		}

		domainGroupList = append(domainGroupList, domainGroup)
	}

	return domainGroupList, nil
}

func handleDomainElementsError(elErr error, flight string, element string) (err error) {
	if err, ok := elErr.(DutyAssignmentNotDefinedError); ok {
		err.StaffGroup = flight
		err.StaffSubGroup = element
		return err
	}

	return elErr
}

func domainElements(config map[string]domain.DutyAssignment, flight Flight) (elementList []domain.Element, err error) {
	for _, element := range flight.Elements {

		domainEL, err := domainDutyAssignment(config, element.ElementLeader)
		if err != nil {
			err = handleDomainElementsError(err, flight.Name, element.Name)
			return nil, err
		}

		domainAsstEL, err := domainDutyAssignment(config, element.AsstElementLeader)
		if err != nil {
			err = handleDomainElementsError(err, flight.Name, element.Name)
			return nil, err
		}

		domainElement := domain.Element{ElementLeader: domainEL, AsstElementLeader: domainAsstEL}

		for _, member := range element.Members {
			m := domain.Member{CAPID: member}
			domainElement.Members = append(domainElement.Members, m)
		}

		elementList = append(elementList, domainElement)
	}

	return elementList, nil
}

func handleDomainFlightsError(flightErr error, flight string) (err error) {
	if err, ok := flightErr.(DutyAssignmentNotDefinedError); ok {
		err.StaffGroup = flight
		return err
	}

	return flightErr
}

func domainFlights(config map[string]domain.DutyAssignment, flights []Flight) (flightList []domain.Flight, err error) {
	for _, flight := range flights {

		domainFltCC, err := domainDutyAssignment(config, flight.Commander)
		if err != nil {
			err = handleDomainFlightsError(err, flight.Name)
			return nil, err
		}

		domainFltCCF, err := domainDutyAssignment(config, flight.FlightSergeant)
		if err != nil {
			err = handleDomainFlightsError(err, flight.Name)
			return nil, err
		}

		orgFlight := domain.Flight{
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

func (to TableOfOrganization) DomainTableOfOrganization(config map[string]domain.DutyAssignment) (domainTO domain.TableOfOrganization, err error) {
	staffGroups, err := domainStaffGroups(config, to.StaffGroups)
	if err != nil {
		return domainTO, errors.WithStack(err)
	}

	flights, err := domainFlights(config, to.Flights)
	if err != nil {
		return domainTO, errors.WithStack(err)
	}

	inactive := mapset.NewSet[uint]()
	for _, i := range to.Inactive {
		inactive.Add(i)
	}

	domainTO.StaffGroups = staffGroups
	domainTO.Flights = flights
	domainTO.InactiveCAPIDs = inactive

	return domainTO, nil
}

type DutyAssignmentNotDefinedError struct {
	OfficeSymbol  string
	StaffGroup    string
	StaffSubGroup string
}

func (err DutyAssignmentNotDefinedError) Error() (msg string) {
	msg = fmt.Sprintf(
		"failed to find definiton for duty assignment %s in %s/%s",
		err.OfficeSymbol,
		err.StaffGroup,
		err.StaffSubGroup,
	)

	return msg
}
