package attendance

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/app/config"
	"github.com/ut080/bcs-portal/app/logging"
	"github.com/ut080/bcs-portal/pkg"
	"github.com/ut080/bcs-portal/pkg/attendance"
	"github.com/ut080/bcs-portal/pkg/capwatch"
	"github.com/ut080/bcs-portal/pkg/files"
	"github.com/ut080/bcs-portal/pkg/yaml"
)

func loadTableOfOrganizationConfiguration(toCfg string, logger logging.Logger) (pkg.TableOfOrganization, error) {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return pkg.TableOfOrganization{}, errors.WithMessage(err, "failed to access config directory")
	}

	daCfg := yaml.DutyAssignmentConfig{}
	daCfgPath := filepath.Join(cfgDir, "cfg", "duty_assignments.yaml")
	err = yaml.LoadYamlDocFromFile(daCfgPath, &daCfg, logger)
	if err != nil {
		return pkg.TableOfOrganization{}, err
	}

	domainDACfg := daCfg.DomainDutyAssignments()

	to := yaml.TableOfOrganization{}
	err = yaml.LoadYamlDocFromFile(toCfg, &to, logger)
	if err != nil {
		return pkg.TableOfOrganization{}, err
	}

	domainTO, err := to.DomainTableOfOrganization(domainDACfg)
	if err != nil {
		return pkg.TableOfOrganization{}, err
	}

	return domainTO, nil
}

func loadCapwatchData(to *pkg.TableOfOrganization, password string, logger logging.Logger) (time.Time, error) {
	cacheDir, err := config.CacheDir()
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}

	orgID := config.GetString("capwatch.orgid")
	username := config.GetString("capwatch.username")
	refresh := config.GetInt("capwatch.refresh")
	client := capwatch.NewClient(orgID, username, password, refresh, logger)

	cacheFile := filepath.Join(cacheDir, "capwatch", fmt.Sprintf("%s.zip", orgID))

	if client.WillRefreshCache(cacheFile) && password == "" {
		logger.Error().Msg("Will need to query CAPWATCH, but no password provided. Re-execute with the --password flag.")
		return time.Time{}, errors.New("re-run with --password flag")
	}

	dump, err := client.Fetch(cacheFile, false)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}

	members, err := dump.FetchMembers()
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}

	err = to.PopulateMemberData(members)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}

	return dump.LastSync(), nil
}

func generateLaTeX(to pkg.TableOfOrganization, logDate, lastSync time.Time, logger logging.Logger) error {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return errors.WithStack(err)
	}

	cacheDir, err := config.CacheDir()
	if err != nil {
		return errors.WithStack(err)
	}

	unit := config.GetString("unit.name")
	unitPatch := config.GetString("unit.patch_image")

	bl := attendance.NewBarcodeLog(unit, "cap_command_emblem.jpg", unitPatch, logDate, lastSync)
	bl.PopulateFromTableOfOrganization(to)

	err = files.Copy(filepath.Join(cfgDir, "assets", "cap_command_emblem.jpg"), filepath.Join(cacheDir, "build", "cap_command_emblem.jpg"))
	if err != nil {
		// TODO: React to whether this build asset has already been copied
		logger.Warn().Err(err).Str("file", "cap_command_emblem.jpg").Msg("failed to copy build asset")
	}

	err = files.Copy(filepath.Join(cfgDir, "assets", unitPatch), filepath.Join(cacheDir, "build", unitPatch))
	if err != nil {
		// TODO: React to whether this build asset has already been copied
		logger.Warn().Err(err).Str("file", unitPatch).Msg("failed to copy build asset")
	}

	latexFilePath := filepath.Join(cacheDir, "build", fmt.Sprintf("%s.tex", logDate.Format("2006-01-02")))
	err = files.Write(latexFilePath, bl.LaTeX())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func compileLaTeX(logDate time.Time, outDest string) error {
	cacheDir, err := config.CacheDir()
	if err != nil {
		return errors.WithStack(err)
	}

	filename := logDate.Format("2006-01-02")

	// First run
	cmd := exec.Command("pdflatex", "-halt-on-error", fmt.Sprintf("%s.tex", filename))
	cmd.Dir = filepath.Join(cacheDir, "build")

	err = cmd.Run()
	if err != nil {
		return errors.WithStack(err)
	}

	// Second run (pdflatex usually needs two runs to get formatting right)
	cmd = exec.Command("pdflatex", "-halt-on-error", fmt.Sprintf("%s.tex", filename))
	cmd.Dir = filepath.Join(cacheDir, "build")

	err = files.Move(filepath.Join(cacheDir, "build", fmt.Sprintf("%s.pdf", filename)), outDest)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func BuildBarcodeLog(input, output, password string, logDate time.Time) error {
	logger := logging.Logger{}

	to, err := loadTableOfOrganizationConfiguration(input, logger)
	if err != nil {
		return errors.WithStack(err)
	}

	lastSync, err := loadCapwatchData(&to, password, logger)
	if err != nil {
		return errors.WithStack(err)
	}

	err = generateLaTeX(to, logDate, lastSync, logger)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compileLaTeX(logDate, output)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
