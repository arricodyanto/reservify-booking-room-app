package repository

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"database/sql"
	"log"
)

type EmployeeRepository interface {
	GetAllEmployees() ([]entity.Employee, error)
	GetEmployeesByID(id string) (entity.Employee, error)
	CreateEmployee(payload entity.Employee) (entity.Employee, error)
	UpdateEmployee(id string, payload entity.Employee) (entity.Employee, error)
}

type employeeRepository struct {
	db *sql.DB
}

// CreateEmployee implements EmployeeRepository.
func (e *employeeRepository) CreateEmployee(payload entity.Employee)(entity.Employee, error) {
	var employee entity.Employee
	err := e.db.QueryRow(config.InsertEmployee,
		payload.Name,
		payload.Division,
		payload.Position,
		payload.Contact).Scan(&employee.ID, &employee.CreatedAt)

	if err != nil {
		log.Println("employeeRepository.QueryRow: ", err.Error())
		return entity.Employee{}, err
	}
	employee.Name = payload.Name
	employee.Division = payload.Division
	employee.Position = payload.Position
	employee.Contact = payload.Contact

	return employee, nil

}

// GetAllEmployees implements EmployeeRepository.
func (e *employeeRepository) GetAllEmployees() ([]entity.Employee, error) {
	var employee []entity.Employee
	rows, err := e.db.Query(config.SelectAllEmployee)
	if err != nil{
		log.Println("employeeRepository.GetAllEmployees.Query: ", err.Error())
		return []entity.Employee{}, err
	}
	for rows.Next(){
		var emp entity.Employee
		err := rows.Scan(&emp.ID, &emp.Name, &emp.Division, &emp.Position, &emp.Contact, &emp.CreatedAt, &emp.UpdatedAt)
		if err != nil{
			log.Println("employeeRepository.rows.Next(): ", err.Error())
			return []entity.Employee{}, err
		}
		employee = append(employee, emp)
	}
	return employee, nil
}

// GetEmployeesByID implements EmployeeRepository.
func (e *employeeRepository) GetEmployeesByID(id string) (entity.Employee, error) {
	var employee entity.Employee
	err := e.db.QueryRow(config.SelectEmployeeByID, id).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Division,
		&employee.Position,
		&employee.Contact,
		&employee.CreatedAt,
		&employee.UpdatedAt)
	if err != nil{
		log.Println("employeeRepository.GetEmployeeByID.QueryRow: ", err.Error())
		return entity.Employee{}, err
	}
	return employee, nil
}

// UpdateEmployee implements EmployeeRepository.
func (e *employeeRepository) UpdateEmployee(id string, payload entity.Employee) (entity.Employee, error) {
	var employee entity.Employee
	err := e.db.QueryRow(config.UpdateEmployee,
		payload.Name,
		payload.Division,
		payload.Position,
		payload.Contact,
		id).Scan(&employee.ID, &employee.UpdatedAt)

	if err != nil {
		log.Println("employeeRepository.QueryRow: ", err.Error())
		return entity.Employee{}, err
	}
	employee.Name = payload.Name
	employee.Division = payload.Division
	employee.Position = payload.Position
	employee.Contact = payload.Contact

	return employee, nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

// get all employees (GA, ADMIN) -GET

// get employees by id (GA, ADMIN) -GET

// create employee (ADMIN) -POST
// update employee (ADMIN) -POS

// delete an employee (ADMIN) -DELETE (nanti dulu deh!!!)