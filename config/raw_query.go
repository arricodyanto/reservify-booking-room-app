package config

const (
	// Employee
	InsertEmployee     = "INSERT INTO employees(name, username, password, role, division, position, contact, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at;"
	SelectAllEmployee  = "SELECT id, name, username, password, role, division, position, contact, created_at, updated_at FROM employees LIMIT $1 OFFSET $2;"
	SelectEmployeeByID = "SELECT id, name, username, password, role, division, position, contact, created_at, updated_at FROM employees WHERE id = $1;"
	UpdateEmployee     = `UPDATE employees SET name = $1, division = $2, position = $3, contact = $4, updated_at = $5, username = $6, password = $7, role = $8 WHERE id = $9 RETURNING created_at;`
)
