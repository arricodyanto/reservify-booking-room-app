package config

const (
	InsertRoom            = `INSERT INTO rooms (name, room_type, capacity, status, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
	SelectRoomByID        = `SELECT id, name, room_type, capacity, status, created_at, updated_at FROM rooms WHERE id = $1`
	SelectRoomList        = `SELECT id, name, room_type, capacity, status, created_at, updated_at FROM rooms ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectRoomListStatus  = `SELECT id, name, room_type, capacity, status, created_at, updated_at FROM rooms WHERE status = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	UpdateRoomByID        = `UPDATE rooms SET name = $2, room_type = $3, capacity = $4, status = $5, updated_at = $6 WHERE id = $1 RETURNING created_at`
	UpdateRoomStatus      = `UPDATE rooms SET status = $2, updated_at = $3 WHERE id = $1 RETURNING name, room_type, capacity, created_at`
	SelectCountRoom       = `SELECT COUNT(*) FROM rooms`
	SelectCountRoomStatus = `SELECT COUNT(*) FROM rooms WHERE status = $1`

	InsertFasilities     = `INSERT INTO facilities (name, quantity) VALUES ($1, $2) RETURNING id, created_at, updated_at`
	SelectFasilitiesList = `SELECT id, name, quantity, created_at, updated_at FROM facilities ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectFasilitiesById = `SELECT id, name, quantity, created_at, updated_at FROM facilities WHERE id = $1`
	UpdateFasilities     = `UPDATE facilities SET name = $1, quantity = $2, updated_at = $3 WHERE id = $4 RETURNING id, created_at`
	TotalRowsFasilities  = `SELECT COUNT(*) FROM facilities`
)
