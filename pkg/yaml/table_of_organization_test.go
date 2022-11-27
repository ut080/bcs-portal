package yaml

import (
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ut080/bcs-portal/app/logging"
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
	err := LoadYamlDocFromFile(filepath.Join(testDataDir, "config", "duty_assignments.yaml"), &daCfg, logging.Logger{})
	assert.NoError(suite.T(), err)

	domainDACfg := daCfg.DomainDutyAssignments()

	to := TableOfOrganization{}
	err = LoadYamlDocFromFile(filepath.Join(testDataDir, "to.yaml"), &to, logging.Logger{})
	assert.NoError(suite.T(), err)

	domainTo, err := to.DomainTableOfOrganization(domainDACfg)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), viper.GetUint("test_member.capid"), domainTo.Flights[0].FlightCommander.Assignee.CAPID)
}

func TestTableOfOrganizationSuite(t *testing.T) {
	suite.Run(t, new(TableOfOrganizationSuite))
}
