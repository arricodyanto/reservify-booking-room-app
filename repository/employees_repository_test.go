package repository

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"database/sql"
	"errors"
	"fmt"

	// "log"

	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	// "golang.org/x/crypto/bcrypt"
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

// var hashedPassword, err = hashPassword(expectEmployee.Password)
// func unhashPassword(hashedPassword string) (string, error) {
//     // Mengonversi hashed password menjadi byte slice
//     hashedPasswords := []byte(hashedPassword)

//     // Meng-unhash password menggunakan bcrypt
//     password, err := bcrypt.CompareHashAndPassword(hashedPasswords, []byte(password))
//     if err != nil {
//         return "", err
//     }

//     return string(password), nil
// }

func (suite *EmployeeRepositoryTestSuite) TestGetEmployeeByID_success() {
	
	rows := sqlmock.NewRows([]string{"id", "name", "username", "password", "role", "division", "position", "contact", "created_at", "updated_at"}).AddRow(expectEmployee.ID, expectEmployee.Name, expectEmployee.Username, expectEmployee.Password, expectEmployee.Role, expectEmployee.Division, expectEmployee.Position, expectEmployee.Contact, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, username, password, role, division, position, contact, created_at, updated_at FROM employees WHERE id = $1`)).WithArgs(expectEmployee.ID).WillReturnRows(rows)

	_, actualError := suite.repo.GetEmployeesByID(expectEmployee.ID)
	assert.Nil(suite.T(), actualError)
	assert.NoError(suite.T(), actualError)
	// assert.Equal(suite.T(), expectEmployee.ID, actualEmployee.ID)
}
func (suite *EmployeeRepositoryTestSuite) TestGetEmployeeByUsername_success() {
	rows := sqlmock.NewRows([]string{"id", "name", "username", "password", "role", "division", "position", "contact", "created_at", "updated_at"}).AddRow(expectEmployee.ID, expectEmployee.Name, expectEmployee.Username, expectEmployee.Password, expectEmployee.Role, expectEmployee.Division, expectEmployee.Position, expectEmployee.Contact, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, username, password, role, division, position, contact, created_at, updated_at FROM employees WHERE username = $1`)).WithArgs(expectEmployee.Username).WillReturnRows(rows)

	_, actualError := suite.repo.GetEmployeesByUsername(expectEmployee.Username)
	assert.Nil(suite.T(), actualError)
	assert.NoError(suite.T(), actualError)
	// assert.Equal(suite.T(), expectEmployee.ID, actualEmployee.ID)
}

func (suite *EmployeeRepositoryTestSuite) TestGetEmployeeById_Fail() {
	suite.mockSql.ExpectQuery(`SELECT`).WillReturnError(fmt.Errorf("error"))

	_, err := suite.repo.GetEmployeesByID("12")
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}
func (suite *EmployeeRepositoryTestSuite) TestGetEmployeeByUsername_Fail() {
	suite.mockSql.ExpectQuery(`SELECT`).WillReturnError(fmt.Errorf("error"))

	_, err := suite.repo.GetEmployeesByUsername("12")
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}


// 
// func (suite *EmployeeRepositoryTestSuite) TestCreateEmployee_success() {
// 	hashedPassword, err := hashPassword(expectEmployee.Password)
// 	if err != nil {
// 		log.Println("Gagal Hash Password : ", err.Error())
// 		return
// 	}

// 	rows := sqlmock.NewRows([]string{"id", "created_at"}).AddRow(expectEmployee.ID, expectEmployee.CreatedAt)

// 	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`INSERT INTO employees(name, username, password, role, division, position, contact, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at`)).WithArgs(expectEmployee.Name, expectEmployee.Username, hashedPassword, expectEmployee.Role, expectEmployee.Division, expectEmployee.Position, expectEmployee.Contact, expectEmployee.UpdatedAt).WillReturnRows(rows)

	
// 	actual, err := suite.repo.CreateEmployee(expectEmployee)
// 	password, err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(expectEmployee.Password))
//     suite.NoError(err)
// 	suite.Nil(suite.T(), err)
// 	// assert.NotNil(suite.T(), err)
// 	// assert.NoError(suite.T(), err)
// 	suite.Equal(suite.T(), expectEmployee.Name, actual.Name)
	
// }


func (suite *EmployeeRepositoryTestSuite) TestList_Success() {
	page := 1
	size := 10
	expectEmployee := []entity.Employee{
		{ID: "1", Name: "John Doe", Username: "johndoe123", Password: "abc5dasar", Role: "Admin", Division: "PM", Position: "Manager", Contact: "62654398564", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "2", Name: "John", Username: "johndoe", Password: "abc5dasar", Role: "Admin", Division: "PM", Position: "Manager", Contact: "62654398564", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		
	}
	expectedPaging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   2,
		TotalPages:  1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "username", "password", "role", "division", "position", "contact", "created_at", "updated_at"}).
		AddRow("1",  "John Doe", "johndoe123","abc5dasar", "Admin", "PM", "Manager", "62654398564", time.Now(), time.Now()).
		AddRow("2",  "John", "johndoe","abc5dasar", "Admin", "PM", "Manager", "62654398564", time.Now(), time.Now())
		

	suite.mockSql.ExpectQuery(`SELECT`).
		WithArgs(size, (page-1)*size).
		WillReturnRows(rows)

	suite.mockSql.ExpectQuery(`SELECT`).
		WillReturnRows(sqlmock.NewRows([]string{"total_rows"}).AddRow(2))

	employees, paging, err := suite.repo.List(page, size)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectEmployee, employees)
	assert.Equal(suite.T(), expectedPaging, paging)
}

func (suite *EmployeeRepositoryTestSuite) TestList_Fail() {
	page := 1
	size := 10

	suite.mockSql.ExpectQuery(`SELECT`).
		WithArgs(size, (page-1)*size).
		WillReturnError(errors.New("some SQL error"))

	_, _, err := suite.repo.List(page, size)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *EmployeeRepositoryTestSuite) TestList_ScanFail() {
	page := 1
	size := 10

	rows := sqlmock.NewRows([]string{"id", "name", "username", "password", "role", "division", "position", "contact", "created_at", "updated_at"}).AddRow(
		expectEmployee.ID, expectEmployee.Name, expectEmployee.Username, expectEmployee.Password, expectEmployee.Role, expectEmployee.Division, expectEmployee.Position, expectEmployee.Contact, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	// Mocking the query
	suite.mockSql.ExpectQuery(`SELECT`).
		WithArgs(size, (page-1)*size).WillReturnRows(rows)

	_, _, err := suite.repo.List(page, size)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *EmployeeRepositoryTestSuite) TestList_ScanTotalRows() {
	page := 1
	size := 10

	rows := sqlmock.NewRows([]string{"id", "name", "username", "password", "role", "division", "position", "contact", "created_at", "updated_at"}).AddRow(
		expectEmployee.ID, expectEmployee.Name, expectEmployee.Username, expectEmployee.Password, expectEmployee.Role, expectEmployee.Division, expectEmployee.Position, expectEmployee.Contact, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	// Mocking the query
	suite.mockSql.ExpectQuery(`SELECT`).
		WithArgs(size, (page-1)*size).WillReturnRows(rows)

	suite.mockSql.ExpectQuery(`SELECT`).
		WillReturnError(errors.New("error"))

	_, _, err := suite.repo.List(page, size)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}


func TestEmployeeRepositoryTestSuite(e *testing.T) {
	suite.Run(e, new(EmployeeRepositoryTestSuite))
}
