package config

const (
	SelectTransactionList         = `SELECT id, employe_id, room_id, description, status, created_at, updated_at FROM transactions ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	GetIdListTransaction          = `SELECT COUNT(*) FROM transactions`
	SelectTransactionByID         = `SELECT id, employee_id, room_id, description, status, created_at, updated_at FROM transactions WHERE id = $1`
	SelectTransactionByEmployeeID = `SELECT id, employee_id, room_id, description, status, created_at, updated_at FROM transactions WHERE employee_id = $1`

	InsertTransactions = `INSERT INTO transactions (employee_id, room_id, description, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, status, created_at`

	UpdatePermission = `UPDATE transactions SET status = $1 WHERE id = $2 RETURNING employee_id, room_id, description,  updated_at`
	// `SELECT id, date, amount, transaction_type, balance, description, created_at, updated_at FROM expenses WHERE LOWER(transaction_type::text) = LOWER($1)`
)
