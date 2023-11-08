package attendance

import (
	"fmt"
	"path/filepath"
	"syscall"
	"time"

	"github.com/ag7if/go-files"
	"github.com/ag7if/go-latex"
	"github.com/pkg/errors"
	"golang.org/x/term"

	"github.com/ut080/bcs-portal/clients/capwatch"
	"github.com/ut080/bcs-portal/clients/eservices"
	"github.com/ut080/bcs-portal/clients/yaml"
	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/pkg/org"
	"github.com/ut080/bcs-portal/reports"
	"github.com/ut080/bcs-portal/reports/attendance"
)

func loadTableOfOrganizationConfiguration(toCfg files.File, logger logging.Logger) (to org.TableOfOrganization, err error) {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return org.TableOfOrganization{}, errors.WithMessage(err, "failed to access config directory")
	}

	daCfg := yaml.DutyAssignmentConfig{}
	daCfgFile, err := files.NewFile(filepath.Join(cfgDir, "cfg", "duty_assignments.yaml"), logger.DefaultLogger())
	if err != nil {
		return org.TableOfOrganization{}, errors.WithStack(err)
	}
	err = yaml.LoadFromFile(daCfgFile, &daCfg, logger)
	if err != nil {
		return org.TableOfOrganization{}, err
	}

	domainDACfg := daCfg.DomainDutyAssignments()

	yamlTO := yaml.TableOfOrganization{}
	err = yaml.LoadFromFile(toCfg, &yamlTO, logger)
	if err != nil {
		return org.TableOfOrganization{}, err
	}

	to, err = yamlTO.DomainTableOfOrganization(domainDACfg)
	if err != nil {
		return org.TableOfOrganization{}, err
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

func loadCapwatchData(to *org.TableOfOrganization, logger logging.Logger) (refreshDate time.Time, err error) {
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

func loadDataFromMembershipReport(to *org.TableOfOrganization, reportFile files.File) (lastSync time.Time, err error) {
	report, err := eservices.NewMembershipReport(reportFile)
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

func generateLaTeX(to org.TableOfOrganization, compiler *latex.Compiler, outputFile files.File, logDate, lastSync time.Time) error {
	unit := config.GetString("unit.name")
	unitPatch := config.GetString("unit.patch_image")

	bl := attendance.NewBarcodeLog(unit, "cap_command_emblem.jpg", unitPatch, logDate, lastSync)
	bl.PopulateFromTableOfOrganization(to)

	assets := []string{"cap_command_emblem.jpg", unitPatch}

	err := compiler.GenerateLaTeX(bl, outputFile, assets)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func BuildBarcodeLog(toCfg, outFile, membershipReport files.File, logDate time.Time, logger logging.Logger) error {
	to, err := loadTableOfOrganizationConfiguration(toCfg, logger)
	if err != nil {
		return errors.WithStack(err)
	}

	var lastSync time.Time
	if membershipReport.Empty() {
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

	compiler, err := reports.ConfigureLaTeXCompiler(logger)
	if err != nil {
		return errors.WithStack(err)
	}

	err = generateLaTeX(to, compiler, outFile, logDate, lastSync)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compiler.CompileLaTeX(outFile)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
