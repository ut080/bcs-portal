package yaml

import (
	"path/filepath"
	"testing"

	"github.com/ag7if/go-files"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/pkg/org"
)

type DutyAssignmentSuite struct {
	suite.Suite
}

func (suite *DutyAssignmentSuite) TestLoadReadDutyAssignmentConfig() {
	dac := DutyAssignmentConfig{}

	dacFile, err := files.NewFile(filepath.Join(testDataDir, "config", "duty_assignments.yaml"), logging.DefaultLogger())
	assert.NoError(suite.T(), err)

	// TODO: Include schema validation
	err = LoadFromFile(dacFile, &dac, nil, logging.Logger{})
	assert.NoError(suite.T(), err)

	// NOTE: This test assumes that C/CC is defined in duty_assignments.yaml as follows:
	//	- symbol: C/CC
	//	  title: 'Cadet Commander'
	//	  min_grade: 'C/2d Lt'
	//    max_grade: 'C/Col'
	dam := dac.DomainDutyAssignments()
	assert.Equal(suite.T(), "Cadet Commander", dam["C/CC"].DutyTitle)
	assert.Equal(suite.T(), "C/CC", dam["C/CC"].OfficeSymbol)
	assert.Equal(suite.T(), org.CdtSecondLt, *dam["C/CC"].MinGrade)
	assert.Equal(suite.T(), org.CdtCol, *dam["C/CC"].MaxGrade)
}

func TestDutyAssignmentSuite(t *testing.T) {
	suite.Run(t, new(DutyAssignmentSuite))
}
