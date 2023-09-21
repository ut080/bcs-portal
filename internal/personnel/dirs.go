package personnel

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/clients/eservices"
	"github.com/ut080/bcs-portal/domain"
	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/repositories/sharepoint"
)

func loadMembershipReport(mbrReportPath string) ([]domain.Member, error) {
	report, err := eservices.NewMembershipReport(mbrReportPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	members, err := report.FetchMembers()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var mbr []domain.Member
	for _, member := range members {
		mbr = append(mbr, member)
	}

	return mbr, nil
}

func parseMembership(members []domain.Member) map[domain.MemberType][]sharepoint.Member {
	mbr := make(map[domain.MemberType][]sharepoint.Member)
	for _, member := range members {
		m := sharepoint.NewMember(member)
		mbr[member.MemberType] = append(mbr[member.MemberType], m)
	}

	return mbr
}

func makeDirectories(t domain.MemberType, members []sharepoint.Member, path string) error {
	mt := strings.ToLower(string(t))
	mt = fmt.Sprintf("%s%s", strings.ToUpper(mt[:1]), mt[1:])
	dir := fmt.Sprintf("%s/%ss", path, mt)

	logging.Info().Str("path", dir).Msgf("Creating %ss directory", mt)
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

func CreateDirectories(mbrReportPath, outputPath string) error {
	members, err := loadMembershipReport(mbrReportPath)
	if err != nil {
		return errors.WithStack(err)
	}

	membersByType := parseMembership(members)

	err = makeDirectories(domain.SeniorMember, membersByType[domain.SeniorMember], outputPath)
	if err != nil {
		return errors.WithStack(err)
	}

	err = makeDirectories(domain.CadetMember, membersByType[domain.CadetMember], outputPath)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
