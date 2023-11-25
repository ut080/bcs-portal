package personnel

import (
	"fmt"
	"os"

	"github.com/ag7if/go-files"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/clients/eservices"
	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/pkg/org"
	"github.com/ut080/bcs-portal/repositories/sharepoint"
)

func loadMembershipReport(mbrReport files.File, mbrType org.MemberType) ([]org.Member, error) {
	report, err := eservices.NewMembershipReport(mbrReport, mbrType)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	members, err := report.FetchMembers()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var mbr []org.Member
	for _, member := range members {
		mbr = append(mbr, member)
	}

	return mbr, nil
}

func loadMembers(mbrReports map[org.MemberType]files.File) (map[org.MemberType][]sharepoint.Member, error) {
	mbr := make(map[org.MemberType][]sharepoint.Member)

	for mbrType, reportFile := range mbrReports {
		members, err := loadMembershipReport(reportFile, mbrType)
		if err != nil {
			return nil, err
		}

		m := make([]sharepoint.Member, 0)
		for _, v := range members {
			m = append(m, sharepoint.NewMember(v))
		}

		mbr[mbrType] = m
	}

	return mbr, nil
}

// memberTypeDirectoryName is hard-coded for now. Eventually, this will be pulled from the fileplan YAML.
// TODO: Pull this info from the file plan
func memberTypeDirectoryName(t org.MemberType) string {
	switch t {
	case org.CadetMember:
		return "3.1.3.1-Cadets"
	case org.SeniorMember:
		return "3.1.3.2-Seniors"
	case org.CadetSponsorMember:
		return "3.1.3.3-Cadet Sponsor Members"
	default:
		return "misc"
	}
}

func makeDirectories(t org.MemberType, members []sharepoint.Member, path string) error {
	dir := fmt.Sprintf("%s/%s", path, memberTypeDirectoryName(t))

	logging.Info().Str("path", dir).Msgf("Creating %ss directory", t)
	err := os.Mkdir(dir, 0755)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, m := range members {
		logging.Info().Str("path", dir).Uint("CAPID", m.SummaryInfo.CAPID).Msg("Creating personnel file directory")
		err = m.CreatePersonnelFile(dir)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func CreateDirectories(mbrReports map[org.MemberType]files.File, outputPath string) error {
	membersByType, err := loadMembers(mbrReports)
	if err != nil {
		return errors.WithStack(err)
	}

	err = makeDirectories(org.SeniorMember, membersByType[org.SeniorMember], outputPath)
	if err != nil {
		return errors.WithStack(err)
	}

	err = makeDirectories(org.CadetMember, membersByType[org.CadetMember], outputPath)
	if err != nil {
		return errors.WithStack(err)
	}

	err = makeDirectories(org.CadetSponsorMember, membersByType[org.CadetSponsorMember], outputPath)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
