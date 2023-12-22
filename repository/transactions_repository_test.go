package repository

import (
	"booking-room-app/entity"
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectedTransactions = entity.Transaction{
	ID:        "1",
    EmployeeId: "1",
    RoomId:    "1",
	RoomFacilities: []entity.RoomFacility{expectedRoomFacilities},
	Status: "pending",
	StartTime:  time.Now(),
	EndTime:  time.Now(),
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
}

var expectedRoomFacilities = entity.RoomFacility {
	ID:        "1",
    RoomId:    "1",
    FacilityId: "1",
    Quantity: 1,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
}

type TransactionsRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    TransactionsRepository
}

func (suite *TransactionsRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewTransactionsRepository(suite.mockDb)
}

func (suite *TransactionsRepositoryTestSuite) TestCreate_Success() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`INSERT INTO transactions (employee_id, room_id, description, start_time, end_time, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, status, created_at`)).WithArgs(
        expectedTransactions.EmployeeId,
        expectedTransactions.RoomId,
        expectedTransactions.Description,
        expectedTransactions.StartTime,
        expectedTransactions.EndTime,
		expectedTransactions.UpdatedAt).WillReturnRows(
        sqlmock.NewRows([]string{"id", "status", "created_at"}).AddRow(
		expectedTransactions.ID, 
		expectedTransactions.Status,
		expectedTransactions.CreatedAt))
		
	suite.mockSql.ExpectQuery(regexp.QuoteMeta( `INSERT INTO trx_room_facility (room_id, facility_id, quantity, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`)).WithArgs(
		expectedRoomFacilities.RoomId, 
		expectedRoomFacilities.FacilityId, 
		expectedRoomFacilities.Quantity, 
		expectedRoomFacilities.UpdatedAt).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(
		expectedRoomFacilities.ID, 
		expectedRoomFacilities.CreatedAt, 
		expectedRoomFacilities.UpdatedAt))

	rows := sqlmock.NewRows([]string{"quantity"}).AddRow(expectedFasilities.Quantity)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT quantity FROM facilities WHERE id = $1`)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)
		
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`UPDATE facilities SET quantity = quantity - $1 WHERE id = $2 RETURNING id, created_at, updated_at`)).WithArgs(expectedRoomFacilities.Quantity, expectedFasilities.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(expectedFasilities.ID, expectedFasilities.CreatedAt, expectedFasilities.UpdatedAt))


    actual, err := suite.repo.Create(expectedTransactions)
    assert.Nil(suite.T(), err)
    assert.Equal(suite.T(), expectedTransactions.Description, actual.Description)
	}

func (suite *TransactionsRepositoryTestSuite) TestCreate_Fail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`INSERT INTO transactions (employee_id, room_id, description, start_time, end_time, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, status, created_at`)).WithArgs(
        expectedTransactions.EmployeeId,
        expectedTransactions.RoomId,
        expectedTransactions.Description,
        expectedTransactions.StartTime,
        expectedTransactions.EndTime,
		expectedTransactions.UpdatedAt).WillReturnError(fmt.Errorf("error"))

    _, err := suite.repo.Create(expectedTransactions)
    assert.NotNil(suite.T(), err)
    assert.Error(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestCreate_RoomFacilitiesNil() {
	var expected = entity.Transaction{
		ID:        "1",
		EmployeeId: "1",
		RoomId:    "1",
		RoomFacilities: nil,
		Status: "pending",
		StartTime:  time.Now(),
		EndTime:  time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`INSERT INTO transactions (employee_id, room_id, description, start_time, end_time, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, status, created_at`)).WithArgs(
        expected.EmployeeId,
        expected.RoomId,
        expected.Description,
        expected.StartTime,
        expected.EndTime,
		expected.UpdatedAt).WillReturnRows(
        sqlmock.NewRows([]string{"id", "status", "created_at"}).AddRow(
			expected.ID, 
			expected.Status,
			expected.CreatedAt))

	actual, err := suite.repo.Create(expected)
    assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.RoomFacilities, actual.RoomFacilities)
} 

func (suite *TransactionsRepositoryTestSuite) TestCreate_RoomFacilitiesScanFaill() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`INSERT INTO transactions (employee_id, room_id, description, start_time, end_time, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, status, created_at`)).WithArgs(
        expectedTransactions.EmployeeId,
        expectedTransactions.RoomId,
        expectedTransactions.Description,
        expectedTransactions.StartTime,
        expectedTransactions.EndTime,
		expectedTransactions.UpdatedAt).WillReturnRows(
        sqlmock.NewRows([]string{"id", "status", "created_at"}).AddRow(
		expectedTransactions.ID, 
		expectedTransactions.Status,
		expectedTransactions.CreatedAt))
		
	suite.mockSql.ExpectQuery(regexp.QuoteMeta( `INSERT INTO trx_room_facility (room_id, facility_id, quantity, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`)).WithArgs(
		expectedRoomFacilities.RoomId, 
		expectedRoomFacilities.FacilityId, 
		expectedRoomFacilities.Quantity, 
		expectedRoomFacilities.UpdatedAt).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(
		expectedRoomFacilities.ID, 
		expectedRoomFacilities.CreatedAt))
		
	_, err := suite.repo.Create(expectedTransactions)
    assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	
}

func (suite *TransactionsRepositoryTestSuite) TestCreate_RoomFacilitiesScanQuantityFaill() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`INSERT INTO transactions (employee_id, room_id, description, start_time, end_time, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, status, created_at`)).WithArgs(
        expectedTransactions.EmployeeId,
        expectedTransactions.RoomId,
        expectedTransactions.Description,
        expectedTransactions.StartTime,
        expectedTransactions.EndTime,
		expectedTransactions.UpdatedAt).WillReturnRows(
        sqlmock.NewRows([]string{"id", "status", "created_at"}).AddRow(
		expectedTransactions.ID, 
		expectedTransactions.Status,
		expectedTransactions.CreatedAt))
		
	suite.mockSql.ExpectQuery(regexp.QuoteMeta( `INSERT INTO trx_room_facility (room_id, facility_id, quantity, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`)).WithArgs(
		expectedRoomFacilities.RoomId, 
		expectedRoomFacilities.FacilityId, 
		expectedRoomFacilities.Quantity, 
		expectedRoomFacilities.UpdatedAt).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(
		expectedRoomFacilities.ID, 
		expectedRoomFacilities.CreatedAt,
		expectedRoomFacilities.UpdatedAt))

		// rows := sqlmock.NewRows([]string{"quantity"}).AddRow(expectedFasilities.Quantity)
		suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT quantity FROM facilities WHERE id = $1`)).WillReturnError(fmt.Errorf("error"))
		
	_, err := suite.repo.Create(expectedTransactions)
    assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	
}

func (suite *TransactionsRepositoryTestSuite) TestCreate_RoomFacilitiesQuantityFaill() {

	var expectedF = entity.Facilities{
		ID:        "1",
		Name:      "This is name",
		Quantity:  0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`INSERT INTO transactions (employee_id, room_id, description, start_time, end_time, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, status, created_at`)).WithArgs(
        expectedTransactions.EmployeeId,
        expectedTransactions.RoomId,
        expectedTransactions.Description,
        expectedTransactions.StartTime,
        expectedTransactions.EndTime,
		expectedTransactions.UpdatedAt).WillReturnRows(
        sqlmock.NewRows([]string{"id", "status", "created_at"}).AddRow(
		expectedTransactions.ID, 
		expectedTransactions.Status,
		expectedTransactions.CreatedAt))
		
	suite.mockSql.ExpectQuery(regexp.QuoteMeta( `INSERT INTO trx_room_facility (room_id, facility_id, quantity, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`)).WithArgs(
		expectedRoomFacilities.RoomId, 
		expectedRoomFacilities.FacilityId, 
		expectedRoomFacilities.Quantity, 
		expectedRoomFacilities.UpdatedAt).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(
		expectedRoomFacilities.ID, 
		expectedRoomFacilities.CreatedAt, 
		expectedRoomFacilities.UpdatedAt))

	rows := sqlmock.NewRows([]string{"quantity"}).AddRow(expectedF.Quantity)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT quantity FROM facilities WHERE id = $1`)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)

	// expectedF.Quantity < expectedRoomFacilities.Quantity
	_, err := suite.repo.Create(expectedTransactions)
    assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestUpdate_Fail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`INSERT INTO transactions (employee_id, room_id, description, start_time, end_time, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, status, created_at`)).WithArgs(
        expectedTransactions.EmployeeId,
        expectedTransactions.RoomId,
        expectedTransactions.Description,
        expectedTransactions.StartTime,
        expectedTransactions.EndTime,
		expectedTransactions.UpdatedAt).WillReturnRows(
        sqlmock.NewRows([]string{"id", "status", "created_at"}).AddRow(
		expectedTransactions.ID, 
		expectedTransactions.Status,
		expectedTransactions.CreatedAt))
		
	suite.mockSql.ExpectQuery(regexp.QuoteMeta( `INSERT INTO trx_room_facility (room_id, facility_id, quantity, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`)).WithArgs(
		expectedRoomFacilities.RoomId, 
		expectedRoomFacilities.FacilityId, 
		expectedRoomFacilities.Quantity, 
		expectedRoomFacilities.UpdatedAt).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(
		expectedRoomFacilities.ID, 
		expectedRoomFacilities.CreatedAt, 
		expectedRoomFacilities.UpdatedAt))

	rows := sqlmock.NewRows([]string{"quantity"}).AddRow(expectedFasilities.Quantity)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT quantity FROM facilities WHERE id = $1`)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)
		
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`UPDATE facilities SET quantity = quantity - $1 WHERE id = $2 RETURNING id, created_at, updated_at`)).WithArgs(expectedRoomFacilities.Quantity, expectedFasilities.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(expectedFasilities.ID, expectedFasilities.CreatedAt))


    _, err := suite.repo.Create(expectedTransactions)
    assert.NotNil(suite.T(), err)
    assert.Error(suite.T(), err)
    
	}




func TestTransactionsRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionsRepositoryTestSuite))
}
	
