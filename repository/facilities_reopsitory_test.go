package repository

import (
	"booking-room-app/entity"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectedFasilities = entity.Facilities{
	ID:        "1",
	Name:      "This is name",
	Quantity:  10,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type FasilitiesRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    FasilitiesRepository
}

func (suite *FasilitiesRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewFasilitesRepository(suite.mockDb)
}

func (suite *FasilitiesRepositoryTestSuite) TestCreate_Success() {
	suite.mockSql.ExpectQuery(`INSERT INTO fasilities`).WithArgs(
		expectedFasilities.Name,
		expectedFasilities.Quantity,
		expectedFasilities.UpdatedAt).WillReturnRows(
		sqlmock.NewRows([]string{"id", "created_at"}).AddRow(
			expectedFasilities.ID,
			expectedFasilities.CreatedAt))

	actual, err := suite.repo.Create(expectedFasilities)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedFasilities.Name, actual.Name)
}

func TestFasilitiesRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(FasilitiesRepositoryTestSuite))
}
