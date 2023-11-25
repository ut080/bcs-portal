package sharepoint

import (
	"fmt"
	"os"
	"time"

	"github.com/ag7if/go-files"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/pkg/org"
)

const (
	PersonnelFileSection1 = "1—Emergency Information"
	PersonnelFileSection2 = "2—Performance Documents"
	PersonnelFileSection3 = "3—Promotions"
	PersonnelFileSection4 = "4—Duty Assignments"
	PersonnelFileSection5 = "5—Miscellaneous Permanent Documents"
	PersonnelFileSection6 = "6—Miscellaneous Temporary Documents"
)

type SummaryInfo struct {
	CAPID          uint           `yaml:"capid"`
	LastName       string         `yaml:"last_name"`
	FirstName      string         `yaml:"first_name"`
	MemberType     org.MemberType `yaml:"member_type"`
	Grade          org.Grade      `yaml:"grade"`
	JoinDate       time.Time      `yaml:"join_date"`
	RankDate       time.Time      `yaml:"rank_date"`
	ExpirationDate time.Time      `yaml:"expiration_date"`
}

type Member struct {
	SummaryInfo SummaryInfo `yaml:"summary"`
}

func NewMember(member org.Member) Member {
	si := SummaryInfo{
		CAPID:          member.CAPID,
		LastName:       member.LastName,
		FirstName:      member.FirstName,
		MemberType:     member.MemberType,
		Grade:          member.Grade,
		JoinDate:       member.JoinDate,
		RankDate:       member.RankDate,
		ExpirationDate: member.ExpirationDate,
	}

	return Member{
		SummaryInfo: si,
	}
}

func (m *Member) ToDomainMember() org.Member {
	return org.Member{
		CAPID:          m.SummaryInfo.CAPID,
		LastName:       m.SummaryInfo.LastName,
		FirstName:      m.SummaryInfo.FirstName,
		MemberType:     m.SummaryInfo.MemberType,
		Grade:          m.SummaryInfo.Grade,
		JoinDate:       m.SummaryInfo.JoinDate,
		RankDate:       m.SummaryInfo.RankDate,
		ExpirationDate: m.SummaryInfo.ExpirationDate,
	}
}

func (m *Member) DirectoryName() string {
	return fmt.Sprintf("%s, %s—%d", m.SummaryInfo.LastName, m.SummaryInfo.FirstName, m.SummaryInfo.CAPID)
}

func (m *Member) CreatePersonnelFile(path string) error {
	pd := fmt.Sprintf("%s/%s", path, m.DirectoryName())

	subs := []string{
		fmt.Sprintf("%s/%s", pd, PersonnelFileSection1),
		fmt.Sprintf("%s/%s", pd, PersonnelFileSection2),
		fmt.Sprintf("%s/%s", pd, PersonnelFileSection3),
		fmt.Sprintf("%s/%s", pd, PersonnelFileSection4),
		fmt.Sprintf("%s/%s", pd, PersonnelFileSection5),
		fmt.Sprintf("%s/%s", pd, PersonnelFileSection6),
	}

	for _, d := range subs {
		// TODO: Remove direct calls to logging
		logging.Info().Str("path", d).Msg("Creating personnel file section directory")
		err := os.MkdirAll(d, 0755)
		if err != nil {
			return errors.WithMessagef(err, "failed to create subdirectory for CAPID: %d", m.SummaryInfo.CAPID)
		}
	}

	err := m.WriteSummary(fmt.Sprintf("%s/summary.yml", pd))
	if err != nil {
		return errors.WithMessagef(err, "failed to write personnel file summary for CAPID: %d", m.SummaryInfo.CAPID)
	}

	return nil
}

func (m *Member) WriteSummary(path string) error {
	f, err := files.NewFile(path, logging.DefaultLogger())
	if err != nil {
		return errors.WithMessagef(err, "failed to create a file at path: %s", path)
	}

	d, err := yaml.Marshal(m)
	if err != nil {
		return errors.WithMessagef(err, "failed to marshal CAPID: %d", m.SummaryInfo.CAPID)
	}

	err = f.WriteBytes(d)
	if err != nil {
		return errors.WithMessagef(err, "failed to write summary YAML for CAPID: %d", m.SummaryInfo.CAPID)
	}

	return nil
}
