package eservices

import (
	"encoding/csv"
	"io"
	"regexp"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/internal/files"
	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/pkg/org"
)

type MembershipReport struct {
	csvReader    *csv.Reader
	lastModified time.Time
}

func openCSV(reportFile files.File) (*csv.Reader, time.Time, error) {
	info, err := reportFile.Stat()
	if err != nil {
		return nil, time.Time{}, errors.WithStack(err)
	}
	modTime := info.ModTime()

	f, err := reportFile.Open()
	if err != nil {
		return nil, time.Time{}, errors.WithStack(err)
	}

	reader := csv.NewReader(f)

	return reader, modTime, nil
}

func NewMembershipReport(reportFile files.File) (MembershipReport, error) {
	reader, modTime, err := openCSV(reportFile)
	if err != nil {
		return MembershipReport{}, errors.WithStack(err)
	}

	return MembershipReport{csvReader: reader, lastModified: modTime}, nil
}

func (mr *MembershipReport) LastModified() time.Time {
	return mr.lastModified
}

func (mr *MembershipReport) FetchMembers() (members map[uint]org.Member, err error) {
	members = make(map[uint]org.Member)

	// Naively assuming the CSV columns are always the same, the ones I'm interested in follow:
	const nameField = 2
	const capidField = 3
	const gradeField = 4
	const rankDateField = 5
	const joinDateField = 7

	// nameField is the full name. To parse it, we will need this regex:
	nameRE := regexp.MustCompile(`(\w+),\s*(\w+)`)

	// Additionally, member type will have to be assumed from the member's grade:
	memberTypeMap := org.MapGradesToMemberTypes()

	// Time layouts will be in the following format:
	const timeLayout = `02 Jan 2006`

	headerSkipped := false

	for {
		record, err := mr.csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}

		if !headerSkipped {
			headerSkipped = true
			continue
		}

		capid, err := strconv.Atoi(record[capidField])
		if err != nil {
			// TODO: Remove direct calls to logger
			logging.Warn().Err(err).Int("col", capidField).Str("capid", record[capidField]).Msg("error converting CAPID, skipping record")
			continue
		}

		name := record[nameField]
		matches := nameRE.FindStringSubmatch(name)
		if len(matches) < 3 {
			// TODO: Remove direct calls to logger
			logging.Warn().Err(err).Int("col", nameField).Int("capid", capid).Str("name", record[nameField]).Msg("error parsing name, skipping record")
			continue
		}
		lastName := matches[1]
		firstName := matches[2]

		grade, err := org.ParseGrade(record[gradeField])
		if err != nil {
			// TODO: Remove direct calls to logger
			logging.Error().Err(err).Int("capid", capid).Int("col", gradeField).Str("grade", record[gradeField]).Msg("error converting Grade, skipping record")
			continue
		}

		memberType, ok := memberTypeMap[grade]
		if !ok {
			// TODO: Remove direct calls to logger
			logging.Warn().Err(err).Int("capid", capid).Int("col", gradeField).Str("grade", record[gradeField]).Msg("error determining MemberType, skipping record")
			continue
		}

		joinDate, err := time.Parse(timeLayout, record[joinDateField])
		if err != nil {
			// TODO: Remove direct calls to logger
			logging.Error().Err(err).Int("capid", capid).Int("col", joinDateField).Str("grade", record[joinDateField]).Msg("error parsing join date, skipping record")
			continue
		}

		rankDate, err := time.Parse(timeLayout, record[rankDateField])
		if err != nil {
			// TODO: Remove direct calls to logger
			logging.Error().Err(err).Int("capid", capid).Int("col", rankDateField).Str("grade", record[rankDateField]).Msg("error parsing rank date, skipping record")
			continue
		}

		member := org.Member{
			CAPID:      uint(capid),
			LastName:   lastName,
			FirstName:  firstName,
			MemberType: memberType,
			Grade:      grade,
			JoinDate:   joinDate,
			RankDate:   rankDate,
		}

		members[uint(capid)] = member
	}

	return members, nil
}
