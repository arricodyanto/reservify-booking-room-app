package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/repository"
	"booking-room-app/shared/model"
	"errors"
	// "fmt"
)

type EmployeesUseCase interface {
	// FindAllEmployees() ([]entity.Employee, error)
	FindEmployeesByID(id string) (entity.Employee, error)
	RegisterNewEmployee(payload entity.Employee) (entity.Employee, error)
	UpdateEmployee(payload entity.Employee) (entity.Employee, error)
	ListAll(page, size int) ([]entity.Employee, model.Paging, error)
}

type employeesUseCase struct {
	repo repository.EmployeeRepository
}

// ListAll implements EmployeesUseCase.
func (e *employeesUseCase) ListAll(page int, size int) ([]entity.Employee, model.Paging, error) {
	return e.repo.List(page, size)
}

// FindAllEmployees implements EmployeesUseCase.
// func (e *employeesUseCase) FindAllEmployees() ([]entity.Employee, error) {
// 	return e.repo.GetAllEmployees()
// }

// FindEmployeesByID implements EmployeesUseCase.
func (e *employeesUseCase) FindEmployeesByID(id string) (entity.Employee, error) {
	if id == "" {
		return entity.Employee{}, errors.New("id harus diisi")
	}
	return e.repo.GetEmployeesByID(id)
}

// RegisterNewEmployee implements EmployeesUseCase.
func (e *employeesUseCase) RegisterNewEmployee(payload entity.Employee) (entity.Employee, error) {
	// if payload.ID == "" || payload.Name == "" || payload.Password == "" || payload.Role == "" || payload.Division == "" || payload.Position == "" || payload.Contact == "" {
	// 	return entity.Employee{}, errors.New("name harus diisi")
	// }
	return e.repo.CreateEmployee(payload)
}

// UpdateEmployee implements EmployeesUseCase.
func (e *employeesUseCase) UpdateEmployee(payload entity.Employee) (entity.Employee, error) {
	return e.repo.UpdateEmployee(payload)
}

func NewEmployeeUseCase(repo repository.EmployeeRepository) EmployeesUseCase {
	return &employeesUseCase{repo: repo}
}

