package capwatch

import (
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/pkg/org"
	"github.com/ut080/bcs-portal/tests"
)

type DumpTestSuite struct {
	suite.Suite
}

func (suite *DumpTestSuite) SetupTest() {
	tests.SetUpTestConfig(testDataDir)
}

func (suite *DumpTestSuite) TestFetchMembers() {
	suite.T().Skip("You need a CAPWATCH dump to run this test. Comment out the Skip() method if you want it to run.")
	orgID := viper.GetString("capwatch.orgid")
	username := viper.GetString("capwatch.username")
	password := viper.GetString("capwatch.password")
	refresh := viper.GetInt("capwatch.refresh")
	client := NewClient(orgID, username, refresh, logging.Logger{})
	client.SetCapwatchPassword([]byte(password))

	cacheFile := filepath.Join(testDataDir, "cache", "capwatch.zip")

	dump, err := client.Fetch(cacheFile, false)
	assert.NoError(suite.T(), err)

	members, err := dump.FetchMembers()
	assert.NoError(suite.T(), err)

	// NOTE: To run this test, you'll need to get your hands on a CAPWATCH download of your unit. Then, add the
	// following (test-only) keys to testdata/config/config.yaml, filling in the details with a member of your unit:
	//	test_member:
	//		capid: <CAPID>
	//		last_name: <LAST_NAME>
	//		first_name: <FIRST_NAME>
	//		grade: <GRADE>
	//		member_type: <MEMBER_TYPE>
	testMember, ok := members[viper.GetUint("test_member.capid")]
	assert.True(suite.T(), ok)

	testMemberType, err := org.ParseMemberType(viper.GetString("test_member.member_type"))
	if err != nil {
		panic(err)
	}
	testMemberGrade, err := org.ParseGrade(viper.GetString("test_member.grade"))
	if err != nil {
		panic(err)
	}

	assert.Equal(suite.T(), viper.GetUint("test_member.capid"), testMember.CAPID)
	assert.Equal(suite.T(), viper.GetString("test_member.last_name"), testMember.LastName)
	assert.Equal(suite.T(), viper.GetString("test_member.first_name"), testMember.FirstName)
	assert.Equal(suite.T(), testMemberType, testMember.MemberType)
	assert.Equal(suite.T(), testMemberGrade, testMember.Grade)
}

func TestDumpSuite(t *testing.T) {
	suite.Run(t, new(DumpTestSuite))
}
