package config

const(
	// Employee
		// done
	InsertEmployee = "INSERT INTO employees(name, username, password, role, division, position, contact, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at;"
	

	SelectAllEmployee = "SELECT id, name, username, password, role, division, position, contact, created_at, updated_at FROM employees;"
	
		// done
	SelectEmployeeByID = "SELECT id, name, username, password, role, division, position, contact, created_at, updated_at FROM employees WHERE id = $1;"
	SelectEmployeeByUsername = "SELECT id, name, username, password, role, division, position, contact, created_at, updated_at FROM employees WHERE username = $1;"
	
		// done
	UpdateEmployee = `UPDATE employees SET name = $1, username = $2, password = $3, role = $4, division = $5, position = $6, contact = $7, updated_at = $8 WHERE id = $9 RETURNING created_at`
)