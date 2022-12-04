package attendance

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ut080/bcs-portal/app/logging"
	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/capwatch"
	"github.com/ut080/bcs-portal/pkg/yaml"
	"github.com/ut080/bcs-portal/tests"
)

type BarcodeLogSuite struct {
	suite.Suite
	to         pkg.TableOfOrganization
	lastCWSync time.Time
}

func loadTOConfig() (pkg.TableOfOrganization, error) {
	daCfg := yaml.DutyAssignmentConfig{}
	err := yaml.LoadYamlDocFromFile(filepath.Join(testDataDir, "config", "duty_assignments.yaml"), &daCfg, logging.Logger{})
	if err != nil {
		return pkg.TableOfOrganization{}, err
	}

	domainDACfg := daCfg.DomainDutyAssignments()

	to := yaml.TableOfOrganization{}
	err = yaml.LoadYamlDocFromFile(filepath.Join(testDataDir, "to.yaml"), &to, logging.Logger{})
	if err != nil {
		return pkg.TableOfOrganization{}, err
	}

	domainTO, err := to.DomainTableOfOrganization(domainDACfg)
	if err != nil {
		return pkg.TableOfOrganization{}, err
	}

	return domainTO, nil
}

func loadCAPWATCHData() (map[uint]pkg.Member, time.Time, error) {
	orgID := viper.GetString("capwatch.orgid")
	username := viper.GetString("capwatch.username")
	password := viper.GetString("capwatch.password")
	refresh := viper.GetInt("capwatch.refresh")
	client := capwatch.NewClient(orgID, username, refresh, logging.Logger{})
	client.SetCapwatchPassword([]byte(password))

	cacheFile := filepath.Join(testDataDir, "cache", "capwatch.zip")

	dump, err := client.Fetch(cacheFile, false)
	if err != nil {
		return nil, time.Time{}, err
	}

	members, err := dump.FetchMembers()
	if err != nil {
		return nil, time.Time{}, err
	}

	return members, dump.LastSync(), nil
}

func (suite *BarcodeLogSuite) SetupTest() {
	tests.SetUpTestConfig(testDataDir)

	to, err := loadTOConfig()
	if err != nil {
		panic(err)
	}

	members, lastSync, err := loadCAPWATCHData()
	if err != nil {
		panic(err)
	}
	suite.lastCWSync = lastSync

	err = to.PopulateMemberData(members)
	if err != nil {
		panic(err)
	}

	suite.to = to
}

func (suite *BarcodeLogSuite) TestLaTeX() {
	assetPath := "../../assets"
	unit := "Blachawk Cadet Squadron"
	commandEmblemPath := filepath.Join(assetPath, "img", "cap_command_emblem.jpg")
	unitPatchPath := filepath.Join(assetPath, "img", "ut080_color.png")
	logDate := time.Now()

	bl := NewBarcodeLog(unit, commandEmblemPath, unitPatchPath, logDate, suite.lastCWSync)
	bl.PopulateFromTableOfOrganization(suite.to)

	latex := bl.LaTeX()

	file, err := os.Create(filepath.Join(testDataDir, "cache", "test_barcode_log.tex"))
	assert.NoError(suite.T(), err)
	defer file.Close()

	_, err = file.WriteString(latex)
	assert.NoError(suite.T(), err)

	err = file.Sync()
	assert.NoError(suite.T(), err)
}

func TestBarcodeLogSuite(t *testing.T) {
	suite.Run(t, new(BarcodeLogSuite))
}
