package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/mock/repo_mock"
	"fmt"
	"testing"
	"time"

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


type TransactionUseCaseTestSuite struct {
	suite.Suite
	trm *repo_mock.TransactionsRepoMock
	tuc TransactionsUsecase
}

func (suite *TransactionUseCaseTestSuite) SetupTest() {
	suite.trm = new(repo_mock.TransactionsRepoMock)
	suite.tuc = NewTransactionsUsecase(suite.trm)
}

func (suite *TransactionUseCaseTestSuite) TestRequestNewBookingRooms_Success() {
	var expectedTransactions = entity.Transaction{
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
	suite.trm.On("Create", expectedTransactions).Return(expectedTransactions, nil)
	_, err := suite.tuc.RequestNewBookingRooms(expectedTransactions)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *TransactionUseCaseTestSuite) TestRequestNewBookingRooms_Fail() {
	var expectedTransactions = entity.Transaction{
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
	suite.trm.On("Create", expectedTransactions).Return(entity.Transaction{} ,fmt.Errorf("error"))
	_, err := suite.tuc.RequestNewBookingRooms(expectedTransactions)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *TransactionUseCaseTestSuite) TestAccStatusBooking_Success() {
	var expectedTransactions = entity.Transaction{
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
	suite.trm.On("UpdatePemission", expectedTransactions).Return(expectedTransactions, nil)
	_, err := suite.tuc.AccStatusBooking(expectedTransactions)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *TransactionUseCaseTestSuite) TestAccStatusBooking_Fail() {
	var expectedTransactions = entity.Transaction{
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
	suite.trm.On("UpdatePemission", expectedTransactions).Return(entity.Transaction{} ,fmt.Errorf("error"))
	_, err := suite.tuc.AccStatusBooking(expectedTransactions)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *TransactionUseCaseTestSuite) TestGetTransactionById_Success() {
	suite.trm.On("GetTransactionById", expectedTransactions.ID).Return(expectedTransactions, nil)
	actual, err := suite.tuc.FindTransactionsById(expectedTransactions.ID)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedTransactions.Description, actual.Description)
}


func TestTransactionUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionUseCaseTestSuite))
}
