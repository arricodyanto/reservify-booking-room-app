package config

const(
	// Employee
	InsertEmployee = "INSERT INTO employees(name, division, position, contact) VALUES($1, $2, $3, $4) RETURNING id, created_at;"
	SelectAllEmployee = "SELECT id, name, division, position, contact, created_at, updated_at FROM employees;"
	SelectEmployeeByID = "SELECT id, name, division, position, contact, created_at, updated_at FROM employees WHERE id = $1;"
	UpdateEmployee = `UPDATE employees SET name = $1, division = $2, position = $3, contact = $4 WHERE id = $5`
)