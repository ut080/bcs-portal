package gorm_org

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ut080/bcs-portal/pkg/org"
	"github.com/ut080/bcs-portal/repositories/gorm/gorm_org"
)

func (suite *RepositorySuite) TestDutyAssignment_ToDomainObject() {
	ccDArepo := gorm_org.DutyAssignment{OfficeSymbol: "CC"}
	result := suite.db.Preload("DutyTitle").Preload("Assignee").Where("office_symbol = ?", "CC").Take(&ccDArepo)
	assert.NoError(suite.T(), result.Error)

	dm := ccDArepo.ToDomainObject()
	ccDAdomain, ok := dm.(org.DutyAssignment)
	assert.True(suite.T(), ok)

	assert.Equal(suite.T(), uuid.MustParse("e81f0253-5268-4b1f-a3f9-9c33b6df6819"), ccDAdomain.ID())
	assert.Equal(suite.T(), "Commander", ccDAdomain.Title())
	assert.Equal(suite.T(), ccDAdomain.DutyTitle().Title(), ccDAdomain.Title())
	assert.Equal(suite.T(), "CC", ccDAdomain.OfficeSymbol())
	assert.Equal(suite.T(), false, ccDAdomain.Assistant())

	assert.Equal(suite.T(), uuid.MustParse("defb6475-a2d2-4961-a063-a360c2a59aec"), ccDAdomain.DutyTitle().ID())
	assert.Equal(suite.T(), "SM-CC", ccDAdomain.DutyTitle().Code())
	assert.Equal(suite.T(), "Commander", ccDAdomain.DutyTitle().Title())
	assert.Equal(suite.T(), org.SeniorMember, ccDAdomain.DutyTitle().MemberType())
	assert.Nil(suite.T(), ccDAdomain.DutyTitle().MinGrade())
	assert.Nil(suite.T(), ccDAdomain.DutyTitle().MaxGrade())

	assert.Equal(suite.T(), uuid.MustParse("aab94819-ee34-47ad-9a44-31e627a97ed5"), ccDAdomain.Assignee().ID())
	assert.Equal(suite.T(), uint(688662), ccDAdomain.Assignee().CAPID())
	assert.Equal(suite.T(), "Morrison", ccDAdomain.Assignee().LastName())
	assert.Equal(suite.T(), "Shelly", ccDAdomain.Assignee().FirstName())
	assert.Equal(suite.T(), org.SeniorMember, ccDAdomain.Assignee().MemberType())
	assert.Equal(suite.T(), org.Maj, ccDAdomain.Assignee().Grade())

	fltCCdaRepo := gorm_org.DutyAssignment{}
	res2 := suite.db.Preload("DutyTitle").Preload("Assignee").Where("office_symbol = ?", "AFlt/CC").Take(&fltCCdaRepo)
	assert.NoError(suite.T(), res2.Error)

	dm2 := fltCCdaRepo.ToDomainObject()
	fltCCdaDomain, ok := dm2.(org.DutyAssignment)
	assert.True(suite.T(), ok)

	assert.Equal(suite.T(), uuid.MustParse("2bd2e4e2-57bd-4867-be4e-f46203d69ed8"), fltCCdaDomain.ID())
	assert.Equal(suite.T(), "Alpha Flight Commander", fltCCdaDomain.Title())
	assert.NotEqual(suite.T(), fltCCdaDomain.DutyTitle().Title(), fltCCdaDomain.Title())
	assert.Equal(suite.T(), "AFlt/CC", fltCCdaDomain.OfficeSymbol())
	assert.Equal(suite.T(), false, fltCCdaDomain.Assistant())

	assert.Equal(suite.T(), uuid.MustParse("f1095ef3-6851-4497-bb74-b704c467e089"), fltCCdaDomain.DutyTitle().ID())
	assert.Equal(suite.T(), "CDT-FLT-CC", fltCCdaDomain.DutyTitle().Code())
	assert.Equal(suite.T(), "Cadet Flight Commander", fltCCdaDomain.DutyTitle().Title())
	assert.Equal(suite.T(), org.CadetMember, fltCCdaDomain.DutyTitle().MemberType())
	minGrade := fltCCdaDomain.DutyTitle().MinGrade()
	assert.Equal(suite.T(), org.CdtMSgt, *minGrade)
	maxGrade := fltCCdaDomain.DutyTitle().MaxGrade()
	assert.Equal(suite.T(), org.CdtCapt, *maxGrade)

	assert.Equal(suite.T(), uuid.MustParse("77c8547a-4c2e-4a2f-8ade-8c2ec0201c0f"), fltCCdaDomain.Assignee().ID())
	assert.Equal(suite.T(), uint(381729), fltCCdaDomain.Assignee().CAPID())
	assert.Equal(suite.T(), "Ross", fltCCdaDomain.Assignee().LastName())
	assert.Equal(suite.T(), "Gregg", fltCCdaDomain.Assignee().FirstName())
	assert.Equal(suite.T(), org.CadetMember, fltCCdaDomain.Assignee().MemberType())
	assert.Equal(suite.T(), org.CdtFirstLt, fltCCdaDomain.Assignee().Grade())
}
