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

func (suite *ConfigTestSuite) TestCfgDefault() {
	err := os.Unsetenv("BCSPORTAL_CONFIG")
	assert.NoError(suite.T(), err)

	usrCfgDir, err := os.UserConfigDir()
	assert.NoError(suite.T(), err)

	cfgDir, err := CfgDir()

	assert.Equal(suite.T(), fmt.Sprintf("%s/bcs-portal", usrCfgDir), cfgDir)
}

func (suite *ConfigTestSuite) TestCfgEnvVar() {
	err := os.Setenv("BCSPORTAL_CONFIG", "test/config/dir")
	assert.NoError(suite.T(), err)

	cfgDir, err := CfgDir()

	assert.Equal(suite.T(), "test/config/dir", cfgDir)
}

func (suite *ConfigTestSuite) TestCacheDefault() {
	err := os.Unsetenv("BCSPORTAL_CACHE")
	assert.NoError(suite.T(), err)

	usrCacheDir, err := os.UserCacheDir()
	assert.NoError(suite.T(), err)

	cacheDir, err := CacheDir()

	assert.Equal(suite.T(), fmt.Sprintf("%s/bcs-portal", usrCacheDir), cacheDir)
}

func (suite *ConfigTestSuite) TestCacheEnvVar() {
	err := os.Setenv("BCSPORTAL_CACHE", "test/cache/dir")
	assert.NoError(suite.T(), err)

	cacheDir, err := CacheDir()

	assert.Equal(suite.T(), "test/cache/dir", cacheDir)
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
