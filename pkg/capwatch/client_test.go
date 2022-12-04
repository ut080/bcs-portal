package capwatch

import (
	"path/filepath"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ut080/bcs-portal/app/logging"
	"github.com/ut080/bcs-portal/tests"
)

type ClientTestSuite struct {
	suite.Suite
}

func (suite *ClientTestSuite) SetupTest() {
	tests.SetUpTestConfig(testDataDir)
}

func (suite *ClientTestSuite) TestFetch() {
	//suite.T().Skip("This test hits CAPWATCH, comment out the Skip() method if you want it to run.")
	orgID := viper.GetString("capwatch.orgid")
	username := viper.GetString("capwatch.username")
	password := viper.GetString("capwatch.password")
	refresh := viper.GetInt("capwatch.refresh")
	client := NewClient(orgID, username, refresh, logging.Logger{})
	client.SetCapwatchPassword([]byte(password))

	cacheFile := filepath.Join(testDataDir, "cache", "capwatch.zip")

	_, err := client.Fetch(cacheFile, true)
	if err != nil {
		log.Err(err).Msg("")
	}
	assert.NoError(suite.T(), err)
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
