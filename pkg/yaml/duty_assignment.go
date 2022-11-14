package yaml

import (
	"github.com/derhabicht/herriman/pkg"
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

func (dac DutyAssignmentConfig) DomainDutyAssignments() (dutyAssignments map[string]pkg.DutyAssignment) {
	dutyAssignments = make(map[string]pkg.DutyAssignment)

	for _, assignment := range dac.DutyAssignments {
		var min *pkg.Grade = nil
		if assignment.MinGrade != nil {
			mi, err := pkg.ParseGrade(*assignment.MinGrade)
			if err == nil { // TODO: Add some actual error handling
				min = &mi
			}
		}

		var max *pkg.Grade = nil
		if assignment.MaxGrade != nil {
			mx, err := pkg.ParseGrade(*assignment.MaxGrade)
			if err == nil { // TODO: Add some actual error handling
				max = &mx
			}
		}

		da := pkg.DutyAssignment{
			DutyTitle:    assignment.Title,
			OfficeSymbol: assignment.OfficeSymbol,
			MinGrade:     min,
			MaxGrade:     max,
		}

		dutyAssignments[assignment.OfficeSymbol] = da
	}

	return dutyAssignments
}
