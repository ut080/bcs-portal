package attendance

import (
	"fmt"
	"path/filepath"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/term"

	"github.com/ut080/bcs-portal/clients/capwatch"
	"github.com/ut080/bcs-portal/clients/eservices"
	"github.com/ut080/bcs-portal/clients/yaml"
	"github.com/ut080/bcs-portal/domain"
	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/files"
	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/reports"
	"github.com/ut080/bcs-portal/reports/attendance"
)

func loadTableOfOrganizationConfiguration(toCfg string, logger logging.Logger) (to domain.TableOfOrganization, err error) {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return domain.TableOfOrganization{}, errors.WithMessage(err, "failed to access config directory")
	}

	daCfg := yaml.DutyAssignmentConfig{}
	daCfgPath := filepath.Join(cfgDir, "cfg", "duty_assignments.yaml")
	err = yaml.LoadYamlDocFromFile(daCfgPath, &daCfg, logger)
	if err != nil {
		return domain.TableOfOrganization{}, err
	}

	domainDACfg := daCfg.DomainDutyAssignments()

	yamlTo := yaml.TableOfOrganization{}
	err = yaml.LoadYamlDocFromFile(toCfg, &yamlTo, logger)
	if err != nil {
		return domain.TableOfOrganization{}, err
	}

	to, err = yamlTo.DomainTableOfOrganization(domainDACfg)
	if err != nil {
		return domain.TableOfOrganization{}, err
	}

	return to, nil
}

func getCapwatchPassword(capwatchUsername string) (password []byte, err error) {
	fmt.Printf("Enter password for %s: ", capwatchUsername)
	password, err = term.ReadPassword(syscall.Stdin)
	if err != nil {
		return nil, err
	}

	return password, nil
}

func loadCapwatchData(to *domain.TableOfOrganization, logger logging.Logger) (refreshDate time.Time, err error) {
	cacheDir, err := config.CacheDir()
	if err != nil {
		return refreshDate, errors.WithStack(err)
	}

	orgID := config.GetString("capwatch.orgid")
	username := config.GetString("capwatch.username")
	refresh := config.GetInt("capwatch.refresh")
	client := capwatch.NewClient(orgID, username, refresh, logger)

	cacheFile := filepath.Join(cacheDir, "capwatch", fmt.Sprintf("%s.zip", orgID))

	if client.WillRefreshCache(cacheFile) {
		password, err := getCapwatchPassword(username)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to read CAPWATCH password")
			return refreshDate, err
		}

		client.SetCapwatchPassword(password)
	}

	dump, err := client.Fetch(cacheFile, false)
	if err != nil {
		return refreshDate, errors.WithStack(err)
	}

	members, err := dump.FetchMembers()
	if err != nil {
		return refreshDate, errors.WithStack(err)
	}

	err = to.PopulateMemberData(members)
	if err != nil {
		return refreshDate, errors.WithStack(err)
	}

	refreshDate = dump.LastSync()

	return refreshDate, nil
}

func loadDataFromMembershipReport(to *domain.TableOfOrganization, filepath string) (lastSync time.Time, err error) {
	report, err := eservices.NewMembershipReport(filepath)
	if err != nil {
		return lastSync, errors.WithStack(err)
	}

	members, err := report.FetchMembers()
	if err != nil {
		return lastSync, errors.WithStack(err)
	}

	err = to.PopulateMemberData(members)
	if err != nil {
		return lastSync, errors.WithStack(err)
	}

	return report.LastModified(), nil
}

func generateReport(to domain.TableOfOrganization, logDate, lastSync time.Time) (bl *attendance.BarcodeLog, assets []string) {
	const capCommandEmblem = "cap_command_emblem.jpg"

	unit := config.GetString("unit.name")
	unitPatch := config.GetString("unit.patch_image")

	bl = attendance.NewBarcodeLog(unit, capCommandEmblem, unitPatch, logDate, lastSync)
	bl.PopulateFromTableOfOrganization(to)

	assets = []string{
		capCommandEmblem,
		unitPatch,
	}

	return bl, assets
}

func BuildBarcodeLog(input, output, membershipReport string, logDate time.Time) (err error) {
	logger := logging.Logger{}

	outputPath, filename, err := files.SplitPath(output)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to compile LaTeX.")
		return errors.WithStack(err)
	}

	if filename == "" {
		filename = logDate.Format("2006-01-02")
	}

	to, err := loadTableOfOrganizationConfiguration(input, logger)
	if err != nil {
		return errors.WithStack(err)
	}

	var lastSync time.Time
	if membershipReport == "" {
		lastSync, err = loadCapwatchData(&to, logger)
		if err != nil {
			return errors.WithStack(err)
		}
	} else {
		lastSync, err = loadDataFromMembershipReport(&to, membershipReport)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	cfgDir, err := config.ConfigDir()
	if err != nil {
		return errors.WithStack(err)
	}

	bl, assets := generateReport(to, logDate, lastSync)
	assetDir := filepath.Join(cfgDir, "assets")

	err = reports.CompileLaTeX(bl, assetDir, outputPath, filename, assets, logger)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to compile LaTeX.")
		return errors.WithStack(err)
	}

	return nil
}
