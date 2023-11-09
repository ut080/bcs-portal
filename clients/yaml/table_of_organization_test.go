package yaml

import (
	"path/filepath"
	"testing"

	"github.com/ag7if/go-files"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/tests"
)

type TableOfOrganizationSuite struct {
	suite.Suite
}

func (suite *TableOfOrganizationSuite) SetupTest() {
	tests.SetUpTestConfig(testDataDir)
}

func (suite *TableOfOrganizationSuite) TestLoadTableOfOrganization() {
	daCfg := DutyAssignmentConfig{}
	daCfgFile, err := files.NewFile(filepath.Join(testDataDir, "config", "defs", "duty_assignments.yaml"), logging.DefaultLogger())
	assert.NoError(suite.T(), err)

	// TODO: Include schema validation
	err = LoadFromFile(daCfgFile, &daCfg, nil, logging.Logger{})
	assert.NoError(suite.T(), err)

	domainDACfg := daCfg.DomainDutyAssignments()

	to := TableOfOrganization{}

	toFile, err := files.NewFile(filepath.Join(testDataDir, "to.yaml"), logging.DefaultLogger())
	assert.NoError(suite.T(), err)

	// TODO: Include schema validation
	err = LoadFromFile(toFile, &to, nil, logging.Logger{})
	assert.NoError(suite.T(), err)

	domainTo, err := to.DomainTableOfOrganization(domainDACfg)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), viper.GetUint("test_member.capid"), domainTo.Flights[0].FlightCommander.Assignee.CAPID)
}

func TestTableOfOrganizationSuite(t *testing.T) {
	suite.Run(t, new(TableOfOrganizationSuite))
}
