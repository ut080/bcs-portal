package postgres_repo

import (
	"github.com/stretchr/testify/assert"

	"github.com/ut080/bcs-portal/repositories/gorm"
)

func (suite *RepositorySuite) TestDutyTitle_ToDomainObject() {
	repoMbr := gorm.DutyTitle{}
	result := suite.db.Take(&repoMbr).Where("code = ?", "SM-CC")
	assert.NoError(suite.T(), result.Error)
}
