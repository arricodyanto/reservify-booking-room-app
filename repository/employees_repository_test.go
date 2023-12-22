package repository

import (
	"booking-room-app/entity"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectEmployee = entity.Employee{
	ID:       "1",
	Name: "John Doe",
	Username: "JohnDoe123",
	Password: "abc5dasar",
	Role: "GA",
	Division: "Marketing",
	Position: "Manager",
	Contact: "612906347",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type EmployeeRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    EmployeeRepository
}

func (suite *EmployeeRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewEmployeeRepository(suite.mockDb)
}

func (suite *EmployeeRepositoryTestSuite) TestGetEmployeeByID_success(id string){
	rows := sqlmock.NewRows([]string{"id, name, username, password, role, division, position, contact, created_at, updated_at"}).AddRow(expectEmployee.ID, expectEmployee.Name, expectEmployee.Username, expectEmployee.Password, expectEmployee.Role, expectEmployee.Division, expectEmployee.Position, expectEmployee.Contact, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, username, password, role, division, position, contact, created_at, updated_at FROM employees WHERE id = $1`)).WithArgs(expectEmployee.ID).WillReturnRows(rows)

	actualEmployee, actualError := suite.repo.GetEmployeesByID(expectEmployee.ID)
	assert.Nil(suite.T(), actualError)
	assert.NoError(suite.T(), actualError)
	assert.Equal(suite.T(), expectEmployee.ID, actualEmployee.ID)
}

func TestEmployeeRepositoryTestSuite(t *testing.T){
	suite.Run(t, new(EmployeeRepositoryTestSuite))
}