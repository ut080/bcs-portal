package attendance

import (
	"fmt"
	"maps"
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
	daCfgFile, err := files.NewFile(filepath.Join(cfgDir, "cfg", "defs", "duty_assignments.yaml"), logger.DefaultLogger())
	if err != nil {
		return org.TableOfOrganization{}, errors.WithStack(err)
	}
	// TODO: Add schema validation
	err = yaml.LoadFromFile(daCfgFile, &daCfg, nil, logger)
	if err != nil {
		return org.TableOfOrganization{}, err
	}

	domainDACfg := daCfg.DomainDutyAssignments()

	yamlTO := yaml.TableOfOrganization{}
	// TODO: Add schema validation
	err = yaml.LoadFromFile(toCfg, &yamlTO, nil, logger)
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

func loadDataFromMembershipReport(reportFile files.File, memberType org.MemberType) (map[uint]org.Member, time.Time, error) {
	report, err := eservices.NewMembershipReport(reportFile, memberType)
	if err != nil {
		return nil, time.Time{}, errors.WithStack(err)
	}

	members, err := report.FetchMembers()
	if err != nil {
		return nil, time.Time{}, errors.WithStack(err)
	}

	return members, report.LastModified(), nil
}

func loadMembershipReports(to *org.TableOfOrganization, reportFiles map[org.MemberType]files.File) (time.Time, error) {
	var lastSync time.Time

	members := make(map[uint]org.Member)
	for mbrType, reportFile := range reportFiles {
		mbrs, t, err := loadDataFromMembershipReport(reportFile, mbrType)
		if err != nil {
			return time.Time{}, errors.WithStack(err)
		}

		if lastSync.IsZero() || t.Before(lastSync) {
			lastSync = t
		}

		maps.Copy(members, mbrs)
	}

	err := to.PopulateMemberData(members)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}

	return lastSync, nil
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

func BuildBarcodeLog(toCfg, outFile files.File, mbrReports map[org.MemberType]files.File, logDate time.Time, logger logging.Logger) error {
	to, err := loadTableOfOrganizationConfiguration(toCfg, logger)
	if err != nil {
		return errors.WithStack(err)
	}

	var lastSync time.Time
	/*
		// TODO: Re-enable CAPWATCH access
			if membershipReport.Empty() {
				lastSync, err = loadCapwatchData(&to, logger)
				if err != nil {
					return errors.WithStack(err)
				}

			} else {
	*/
	lastSync, err = loadMembershipReports(&to, mbrReports)
	if err != nil {
		return errors.WithStack(err)
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
