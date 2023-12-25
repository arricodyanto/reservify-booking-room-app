package usecase_mock

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"

	"github.com/stretchr/testify/mock"
)

type EmployeesUseCaseMock struct {
	mock.Mock
}

func (e *EmployeesUseCaseMock) FindEmployeesByID(id string) (entity.Employee, error){
	args := e.Called(id)
	return args.Get(0).(entity.Employee), args.Error(1)
}

func (e *EmployeesUseCaseMock) FindEmployeesByUsername(username string) (entity.Employee, error){
	args := e.Called(username)
	return args.Get(0).(entity.Employee), args.Error(1)
}

func (e *EmployeesUseCaseMock) RegisterNewEmployee(payload string) (entity.Employee, error){
	args := e.Called(payload)
	return args.Get(0).(entity.Employee), args.Error(1)
}

func (e *EmployeesUseCaseMock) UpdateEmployee(payload string) (entity.Employee, error){
	args := e.Called(payload)
	return args.Get(0).(entity.Employee), args.Error(1)
}

func (e *EmployeesUseCaseMock) ListAll(payload string) ([]entity.Employee, model.Paging, error){
	args := e.Called(payload)
	return args.Get(0).([]entity.Employee), args.Get(1).(model.Paging), args.Error(2)
}

