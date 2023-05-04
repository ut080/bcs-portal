package feedback

import (
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/clients/eservices"
	"github.com/ut080/bcs-portal/domain"
	"github.com/ut080/bcs-portal/internal/logging"
)

func loadDataFromMembershipReport(filepath string) (map[uint]domain.Member, error) {
	report, err := eservices.NewMembershipReport(filepath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	members, err := report.FetchMembers()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return members, nil
}

func BuildSchedule(outfile, mbrReport string) error {
	logger := logging.Logger{}

	if mbrReport == "" {
		logger.Error().Msg("Generating a schedule from CAPWATCH is not yet supported. Please supply a Membership report.")
	}

	members, err := loadDataFromMembershipReport(mbrReport)

	schedule := fetchSchedule(members)

	return nil
}
