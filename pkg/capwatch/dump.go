package capwatch

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"io"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/app/logging"
	"github.com/ut080/bcs-portal/pkg"
)

type Dump struct {
	raw      []byte
	lastSync time.Time
}

func NewDump(dump []byte, lastSync time.Time) *Dump {
	return &Dump{
		raw:      dump,
		lastSync: lastSync,
	}
}

func (d *Dump) openCSV(filename string) (*csv.Reader, error) {
	rawArchive := bytes.NewReader(d.raw)
	archive, err := zip.NewReader(rawArchive, int64(len(d.raw)))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	f, err := archive.Open(filename)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return csv.NewReader(f), nil
}

func (d *Dump) FetchMembers() (members map[uint]pkg.Member, err error) {
	members = make(map[uint]pkg.Member)

	membersCSV, err := d.openCSV("Member.txt")

	// Naively assuming the CSV columns are always the same, the ones I'm interested (and their corresponding fields
	// in the domain.Member struct) are:
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
		memberType, err := pkg.ParseMemberType(record[21])
		if err != nil {
			// TODO: Remove direct calls to logger
			logging.Warn().Err(err).Int("capid", capid).Int("col", 21).Str("member_type", record[21]).Msg("error converting MemberType, skipping record")
			continue
		}
		grade, err := pkg.ParseGrade(record[14])
		if err != nil {
			// TODO: Remove direct calls to logger
			logging.Error().Err(err).Int("capid", capid).Int("col", 14).Str("grade", record[14]).Msg("error converting Grade, skipping record")
			continue
		}

		member := pkg.Member{
			CAPID:      uint(capid),
			LastName:   lastName,
			FirstName:  firstName,
			MemberType: memberType,
			Grade:      grade,
		}

		members[uint(capid)] = member
	}

	return members, nil
}

func (d *Dump) LastSync() time.Time {
	return d.lastSync
}
