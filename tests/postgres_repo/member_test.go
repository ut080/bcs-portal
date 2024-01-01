package postgres_repo

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ut080/bcs-portal/pkg/org"
	"github.com/ut080/bcs-portal/repositories/gorm"
)

func (suite *RepositorySuite) TestMember_BeforeCreate() {
	mbr := org.NewMember(
		uuid.Nil,
		303191,
		"Hawk",
		"Robert",
		org.SeniorMember,
		org.Capt,
		nil,
		nil,
		nil,
	)

	repoMbr := gorm.Member{}
	err := repoMbr.FromDomainObject(mbr)
	assert.NoError(suite.T(), err)

	result := suite.db.Create(&repoMbr)
	assert.NoError(suite.T(), result.Error)
	assert.NotEqual(suite.T(), uuid.Nil, repoMbr.ID)
}

func (suite *RepositorySuite) TestMember_ToDomainObject() {
	/*
		- id: aab94819-ee34-47ad-9a44-31e627a97ed5
		capid: 688662
		last_name: Morrison
		first_name: Shelly
		member_type: SENIOR MEMBER
		grade: Maj
	*/

	repoMbr := gorm.Member{}
	result := suite.db.Take(&repoMbr).Where("capid = ?", 688662)
	assert.NoError(suite.T(), result.Error)

	dm := repoMbr.ToDomainObject()
	domainMbr, ok := dm.(org.Member)
	assert.True(suite.T(), ok)

	assert.Equal(suite.T(), uuid.MustParse("aab94819-ee34-47ad-9a44-31e627a97ed5"), domainMbr.ID())
	assert.Equal(suite.T(), uint(688662), domainMbr.CAPID())
	assert.Equal(suite.T(), "Morrison", domainMbr.LastName())
	assert.Equal(suite.T(), "Shelly", domainMbr.FirstName())
	assert.Equal(suite.T(), org.SeniorMember, domainMbr.MemberType())
	assert.Equal(suite.T(), org.Maj, domainMbr.Grade())
}
