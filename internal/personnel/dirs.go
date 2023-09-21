package personnel

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/clients/eservices"
	"github.com/ut080/bcs-portal/domain"
	"github.com/ut080/bcs-portal/internal/logging"
)

const (
	section1 = "1—Summary and Emergency Information"
	section2 = "2—CAPVA 60-101 and BCSF 3"
	section3 = "3—Feedback Schedule and CAPF 60-9X"
	section4 = "4—Inspections, Essays, and Drill Tests"
	section5 = "5—CAPF 2A"
	section6 = "6—Miscellaneous"
)

func loadMembershipReport(mbrReportPath string) (map[domain.MemberType][]domain.Member, error) {
	report, err := eservices.NewMembershipReport(mbrReportPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	members, err := report.FetchMembers()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	mbt := make(map[domain.MemberType][]domain.Member)

	for _, member := range members {
		mbt[member.MemberType] = append(mbt[member.MemberType], member)
	}

	return mbt, nil
}

func buildDirectoryName(member domain.Member) string {
	return fmt.Sprintf("%s, %s—%d", member.LastName, member.FirstName, member.CAPID)
}

func createPersonnelDir(path string, member domain.Member) error {
	pd := fmt.Sprintf("%s/%s", path, buildDirectoryName(member))

	subs := []string{
		fmt.Sprintf("%s/%s", pd, section1),
		fmt.Sprintf("%s/%s", pd, section2),
		fmt.Sprintf("%s/%s", pd, section3),
		fmt.Sprintf("%s/%s", pd, section4),
		fmt.Sprintf("%s/%s", pd, section5),
		fmt.Sprintf("%s/%s", pd, section6),
	}

	for _, d := range subs {
		logging.Info().Str("path", d).Msg("Creating personnel file section directory")
		err := os.MkdirAll(d, 0755)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func makeDirectories(t domain.MemberType, members []domain.Member, path string) error {
	mt := strings.ToLower(string(t))
	mt = fmt.Sprintf("%s%s", strings.ToUpper(mt[:1]), mt[1:])
	dir := fmt.Sprintf("%s/%ss", path, mt)

	logging.Info().Str("path", dir).Msgf("Creating %ss directory", mt)
	err := os.Mkdir(dir, 0755)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, m := range members {
		logging.Info().Str("path", dir).Uint("CAPID", m.CAPID).Msg("Creating personnel file directory")
		err = createPersonnelDir(dir, m)
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

	err = makeDirectories(domain.SeniorMember, members[domain.SeniorMember], outputPath)
	if err != nil {
		return errors.WithStack(err)
	}

	err = makeDirectories(domain.CadetMember, members[domain.CadetMember], outputPath)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
