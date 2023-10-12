package yaml

import (
	"github.com/ut080/bcs-portal/pkg/org"
)

type DutyAssignmentConfig struct {
	DutyAssignments []DutyAssignment `yaml:"duty_assignments"`
}

type DutyAssignment struct {
	OfficeSymbol  string  `yaml:"symbol"`
	Title         string  `yaml:"title"`
	CapwatchTitle *string `yaml:"capwatch_title"`
	MinGrade      *string `yaml:"min_grade"`
	MaxGrade      *string `yaml:"max_grade"`
	AsigneeCAPID  *uint   `yaml:"capid"`
}

func (dac DutyAssignmentConfig) DomainDutyAssignments() (dutyAssignments map[string]org.DutyAssignment) {
	dutyAssignments = make(map[string]org.DutyAssignment)

	for _, assignment := range dac.DutyAssignments {
		var min *org.Grade = nil
		if assignment.MinGrade != nil {
			mi, err := org.ParseGrade(*assignment.MinGrade)
			if err == nil { // TODO: Add some actual error handling
				min = &mi
			}
		}

		var max *org.Grade = nil
		if assignment.MaxGrade != nil {
			mx, err := org.ParseGrade(*assignment.MaxGrade)
			if err == nil { // TODO: Add some actual error handling
				max = &mx
			}
		}

		da := org.DutyAssignment{
			DutyTitle:    assignment.Title,
			OfficeSymbol: assignment.OfficeSymbol,
			MinGrade:     min,
			MaxGrade:     max,
		}

		dutyAssignments[assignment.OfficeSymbol] = da
	}

	return dutyAssignments
}
