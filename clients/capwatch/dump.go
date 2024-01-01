package capwatch

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/internal/logging"
	"github.com/ut080/bcs-portal/pkg/org"
)

type Dump struct {
	raw      []byte
	lastSync time.Time
}

func NewDump(dump []byte, lastSync time.Time) (d *Dump) {
	nd := Dump{
		raw:      dump,
		lastSync: lastSync,
	}

	d = &nd
	return d
}

func (d *Dump) openCSV(filename string) (reader *csv.Reader, err error) {
	rawArchive := bytes.NewReader(d.raw)
	archive, err := zip.NewReader(rawArchive, int64(len(d.raw)))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	f, err := archive.Open(filename)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	reader = csv.NewReader(f)
	return reader, nil
}

func (d *Dump) FetchMembers() (members map[uint]org.Member, err error) {
	members = make(map[uint]org.Member)

	membersCSV, err := d.openCSV("Member.txt")

	// Naively assuming the CSV columns are always the same, the ones I'm interested (and their corresponding fields
	// in the pkg.Member struct) are:
	// 		 0: CAPID
	//		 2: LastName
	//		 3: FirstName
	//		21: MemberType
	//		14: Grade

	for {
		record, err := membersCSV.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}

		// Skip the header row
		if record[0] == "CAPID" {
			continue
		}

		capid, err := strconv.Atoi(record[0])
		if err != nil {
			// TODO: Remove direct calls to logger
			logging.Warn().Err(err).Int("col", 0).Str("capid", record[0]).Msg("error converting CAPID, skipping record")
			continue
		}
		lastName := record[2]
		firstName := record[3]
		memberType, err := org.ParseMemberType(record[21])
		if err != nil {
			// TODO: Remove direct calls to logger
			logging.Warn().Err(err).Int("capid", capid).Int("col", 21).Str("member_type", record[21]).Msg("error converting MemberType, skipping record")
			continue
		}
		grade, err := org.ParseGrade(record[14])
		if err != nil {
			// TODO: Remove direct calls to logger
			logging.Error().Err(err).Int("capid", capid).Int("col", 14).Str("grade", record[14]).Msg("error converting Grade, skipping record")
			continue
		}

		// TODO: Parse JoinDate
		// TODO: Parse RankDate

		member := org.NewMember(
			uuid.Nil,
			uint(capid),
			lastName,
			firstName,
			memberType,
			grade,
			nil,
			nil,
			nil,
		)

		members[uint(capid)] = member
	}

	return members, nil
}

func (d *Dump) LastSync() time.Time {
	return d.lastSync
}
