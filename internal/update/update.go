package update

import (
	"time"

	"github.com/ag7if/go-files"
	"github.com/pkg/errors"

	"github.com/ut080/bcs-portal/clients/eservices"
	"github.com/ut080/bcs-portal/internal/database"
	"github.com/ut080/bcs-portal/pkg/org"
	"github.com/ut080/bcs-portal/repositories/gorm/gorm_org"
)

func loadMembershipReport(reportFile files.File, memberType org.MemberType) (map[uint]org.Member, time.Time, error) {
	report, err := eservices.NewMembershipReport(reportFile, memberType)
	if err != nil {
		return nil, time.Time{}, errors.WithStack(err)
	}

	members, err := report.FetchMembers()
	if err != nil {
		return nil, time.Time{}, errors.WithStack(err)
	}

	return members, report.LastModified(), nil
}

func loadAllMembershipReports(reportFiles map[org.MemberType]files.File) ([]org.Member, error) {
	var members []org.Member

	for mbrType, reportFile := range reportFiles {
		mbrs, _, err := loadMembershipReport(reportFile, mbrType)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		for _, v := range mbrs {
			members = append(members, v)
		}
	}

	return members, nil
}

func updateMembers(repo gorm_org.Repository, members []org.Member) error {
	repoMbrs, err := repo.FromDomainObjects(members)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func Update(capwatch bool, mbrReports map[org.MemberType]files.File) error {
	db, err := database.GetGormDB()
	if err != nil {
		return errors.WithStack(err)
	}

	repo := gorm_org.NewRepository(db)

	members, err := loadAllMembershipReports(mbrReports)
	if err != nil {
		return errors.WithStack(err)
	}

	err = updateMembers(repo, members)

	return nil
}
