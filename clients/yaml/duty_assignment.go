package yaml

import (
	"github.com/ut080/bcs-portal/pkg/org"
)

type DutyAssignmentConfig struct {
	SquadronCommandStaff []DutyAssignment `yaml:"squadron_command_staff"`
	SeniorProgramStaff   []DutyAssignment `yaml:"senior_program_staff"`
	SeniorSupportStaff   []DutyAssignment `yaml:"senior_support_staff"`
	CadetCommandStaff    []DutyAssignment `yaml:"cadet_command_staff"`
	CadetSupportStaff    []DutyAssignment `yaml:"cadet_support_staff"`
	CadetLineStaff       []DutyAssignment `yaml:"cadet_line_staff"`
}

type DutyAssignment struct {
	OfficeSymbol  string  `yaml:"symbol"`
	Title         string  `yaml:"title"`
	CapwatchTitle *string `yaml:"capwatch_title"`
	MinGrade      *string `yaml:"min_grade"`
	MaxGrade      *string `yaml:"max_grade"`
	AsigneeCAPID  *uint   `yaml:"capid"`
}

func toDomainDutyAssignment(assignment DutyAssignment) org.DutyAssignment {
	var minGrade *org.Grade = nil
	if assignment.MinGrade != nil {
		mi, err := org.ParseGrade(*assignment.MinGrade)
		if err == nil { // TODO: Add some actual error handling
			minGrade = &mi
		}
	}

	var maxGrade *org.Grade = nil
	if assignment.MaxGrade != nil {
		mx, err := org.ParseGrade(*assignment.MaxGrade)
		if err == nil { // TODO: Add some actual error handling
			maxGrade = &mx
		}
	}

	return org.DutyAssignment{
		DutyTitle:    assignment.Title,
		OfficeSymbol: assignment.OfficeSymbol,
		MinGrade:     minGrade,
		MaxGrade:     maxGrade,
	}
}

func (dac DutyAssignmentConfig) DomainDutyAssignments() (dutyAssignments map[string]org.DutyAssignment) {
	dutyAssignments = make(map[string]org.DutyAssignment)

	for _, assignment := range dac.SquadronCommandStaff {
		dutyAssignments[assignment.OfficeSymbol] = toDomainDutyAssignment(assignment)
	}

	for _, assignment := range dac.SeniorProgramStaff {
		dutyAssignments[assignment.OfficeSymbol] = toDomainDutyAssignment(assignment)
	}

	for _, assignment := range dac.SeniorSupportStaff {
		dutyAssignments[assignment.OfficeSymbol] = toDomainDutyAssignment(assignment)
	}

	for _, assignment := range dac.CadetCommandStaff {
		dutyAssignments[assignment.OfficeSymbol] = toDomainDutyAssignment(assignment)
	}

	for _, assignment := range dac.CadetSupportStaff {
		dutyAssignments[assignment.OfficeSymbol] = toDomainDutyAssignment(assignment)
	}

	for _, assignment := range dac.CadetLineStaff {
		dutyAssignments[assignment.OfficeSymbol] = toDomainDutyAssignment(assignment)
	}

	return dutyAssignments
}
