package gorm_org

import (
	"github.com/stretchr/testify/assert"

	"github.com/ut080/bcs-portal/repositories/gorm/gorm_org"
)

func (suite *RepositorySuite) TestDutyTitle_ToDomainObject() {
	repoMbr := gorm_org.DutyTitle{}
	result := suite.db.Take(&repoMbr).Where("code = ?", "SM-CC")
	assert.NoError(suite.T(), result.Error)
}
