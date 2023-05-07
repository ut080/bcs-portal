package feedback

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/clients/eservices"
	"github.com/ut080/bcs-portal/domain"
	"github.com/ut080/bcs-portal/internal/config"
	"github.com/ut080/bcs-portal/internal/files"
	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/reports"
	"github.com/ut080/bcs-portal/reports/feedback"
)

func loadDataFromMembershipReport(filepath string) (map[uint]domain.Member, time.Time, error) {
	report, err := eservices.NewMembershipReport(filepath)
	if err != nil {
		return nil, time.Time{}, errors.WithStack(err)
	}

	members, err := report.FetchMembers()
	if err != nil {
		return nil, time.Time{}, errors.WithStack(err)
	}

	lastModified := report.LastModified()

	return members, lastModified, nil
}

func fetchSchedule(members map[uint]domain.Member) (schedule map[time.Month][]domain.Member) {
	schedule = map[time.Month][]domain.Member{
		// T1
		time.October:  nil,
		time.November: nil,
		time.December: nil,
		time.January:  nil,
		// T2
		time.February: nil,
		time.March:    nil,
		time.April:    nil,
		time.May:      nil,
		// T3
		time.June:      nil,
		time.July:      nil,
		time.August:    nil,
		time.September: nil,
	}

	thing := [][]time.Month{
		{time.October, time.February, time.June},
		{time.November, time.March, time.July},
		{time.December, time.April, time.August},
		{time.January, time.May, time.September},
	}

	for _, member := range members {
		if member.MemberType != domain.CadetMember {
			continue
		}

		idx := member.CAPID % 3
		feedbackMonths := thing[idx]

		for _, month := range feedbackMonths {
			schedule[month] = append(schedule[month], member)
			domain.SortMembersByName(schedule[month])
		}
	}

	return schedule
}

func generateReport(membersByMonth map[time.Month][]domain.Member, fiscalYear uint, lastSync time.Time) (schedule *feedback.Schedule, assets []string) {
	const capCommandEmblem = "cap_command_emblem.jpg"

	unit := config.GetString("unit.name")
	unitPatch := config.GetString("unit.patch_image")

	schedule = feedback.NewSchedule(unit, capCommandEmblem, unitPatch, fiscalYear, lastSync)
	schedule.PopulateFromMap(membersByMonth)

	assets = []string{
		capCommandEmblem,
		unitPatch,
	}

	return schedule, assets
}

func BuildSchedule(output, mbrReport string) error {
	logger := logging.Logger{}

	if mbrReport == "" {
		logger.Error().Msg("Generating a schedule from CAPWATCH is not yet supported. Please supply a Membership report.")
	}

	// TODO: Allow specifying the Fiscal Year
	fiscalYear := uint(time.Now().Year())

	outputPath, filename, err := files.SplitPath(output)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to compile LaTeX.")
		return errors.WithStack(err)
	}

	if filename == "" {
		filename = fmt.Sprintf("FeedbackSchedule-FY%d", fiscalYear)
	}

	members, lastSync, err := loadDataFromMembershipReport(mbrReport)
	if err != nil {
		return errors.WithStack(err)
	}

	schedule := fetchSchedule(members)

	cfgDir, err := config.ConfigDir()
	if err != nil {
		return errors.WithStack(err)
	}

	schedReport, assets := generateReport(schedule, fiscalYear, lastSync)
	assetDir := filepath.Join(cfgDir, "assets")

	err = reports.CompileLaTeX(schedReport, assetDir, outputPath, filename, assets, logger)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to compile LaTeX.")
		return errors.WithStack(err)
	}

	return nil
}
