package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/mock/repo_mock"
	// "booking-room-app/shared/model"
	// "fmt"
	// "testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectEmployee = entity.Employee{
	ID:        "1",
	Name:      "John Doe",
	Username:  "JohnDoe123",
	Password:  "abc5dasar", // password telah ditentukan
	Role:      "GA",
	Division:  "Marketing",
	Position:  "Manager",
	Contact:   "612906347",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type EmployeeUseCaseTestSuite struct {
	suite.Suite
	erm *repo_mock.EmployeeRepoMock
	euc EmployeesUseCase
}

func (suite *EmployeeUseCaseTestSuite) SetupTest() {
	suite.erm = new(repo_mock.EmployeeRepoMock)
	suite.euc = NewEmployeeUseCase(suite.erm)
}

func (suite *EmployeeUseCaseTestSuite) TestFindEmployeesByUsername_EmptyUsername() {
    useCase := employeesUseCase{repo: suite.erm}

    _, err := useCase.FindEmployeesByUsername("")

    assert.EqualError(suite.T(), err, "id harus diisi")
}