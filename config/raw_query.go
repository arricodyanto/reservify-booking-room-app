package config

const (
	SelectTransactionList = `SELECT id, employee_id, room_id, description, status, start_time, end_time, created_at, updated_at FROM transactions WHERE created_at BETWEEN $3 AND ($4::date + 1) - interval '1 second' ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	SelectRoomWithFacilities = `SELECT r.id, r.facility_id, r.quantity, r.created_at, r.updated_at FROM rooms t JOIN trx_room_facility r on t.id = r.room_id WHERE t.id = $1;`

	GetIdListTransaction = `SELECT COUNT(*) FROM transactions`

	SelectTransactionByID = `SELECT id, employee_id, room_id, description, status, start_time, end_time, created_at, updated_at FROM transactions WHERE id = $1`

	SelectTransactionByEmployeeID = `SELECT id, employee_id, room_id, description, status, start_time, end_time, created_at, updated_at FROM transactions WHERE employee_id = $1`

	InsertTransactions = `INSERT INTO transactions (employee_id, room_id, description, start_time, end_time, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, status, created_at`

	UpdatePermission = `UPDATE transactions SET status = $1, updated_at = $3 WHERE id = $2 RETURNING employee_id, room_id, description, start_time, end_time, created_at`

	InsertRoomFacility     = `INSERT INTO trx_room_facility (room_id, facility_id, quantity) VALUES ($1, $2, $3)`
	UpdateFacilityQuantity = `UPDATE facilities SET quantity - $1 WHERE id = $2`
	// `SELECT id, date, amount, transaction_type, balance, description, created_at, updated_at FROM expenses WHERE LOWER(transaction_type::text) = LOWER($1)`
)
