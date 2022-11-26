package yaml

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ut080/bcs-portal/pkg"
)

type DutyAssignmentSuite struct {
	suite.Suite
}

func (suite *DutyAssignmentSuite) TestLoadReadDutyAssignmentConfig() {
	dac := DutyAssignmentConfig{}

	err := LoadYamlDocFromFile(filepath.Join(testDataDir, "config", "duty_assignments.yaml"), &dac)
	assert.NoError(suite.T(), err)

	// NOTE: This test assumes that C/CC is defined in duty_assignments.yaml as follows:
	//	- symbol: C/CC
	//	  title: 'Cadet Commander'
	//	  min_grade: 'C/2d Lt'
	//    max_grade: 'C/Col'
	dam := dac.DomainDutyAssignments()
	assert.Equal(suite.T(), "Cadet Commander", dam["C/CC"].DutyTitle)
	assert.Equal(suite.T(), "C/CC", dam["C/CC"].OfficeSymbol)
	assert.Equal(suite.T(), pkg.CdtSecondLt, *dam["C/CC"].MinGrade)
	assert.Equal(suite.T(), pkg.CdtCol, *dam["C/CC"].MaxGrade)
}

func TestDutyAssignmentSuite(t *testing.T) {
	suite.Run(t, new(DutyAssignmentSuite))
}
