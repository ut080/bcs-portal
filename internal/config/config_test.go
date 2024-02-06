package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func (suite *ConfigTestSuite) TestDefault() {
	err := os.Unsetenv("BCSPORTAL_CONFIG")
	assert.NoError(suite.T(), err)

	usrCfgDir, err := os.UserConfigDir()
	assert.NoError(suite.T(), err)

	cfgDir, err := CfgDir()

	assert.Equal(suite.T(), fmt.Sprintf("%s/bcs-portal", usrCfgDir), cfgDir)
}

func (suite *ConfigTestSuite) TestEnvVar() {
	err := os.Setenv("BCSPORTAL_CONFIG", "test/config/dir")
	assert.NoError(suite.T(), err)

	cfgDir, err := CfgDir()

	assert.Equal(suite.T(), "test/config/dir", cfgDir)
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
