package config

const (
	InsertRoom       = `INSERT INTO rooms (name, room_type, capacity, status, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
	SelectRoomByID   = `SELECT id, name, room_type, capacity, status, created_at, updated_at FROM rooms WHERE id = $1`
	SelectRoomList   = `SELECT id, name, room_type, capacity, status, created_at, updated_at FROM rooms`
	UpdateRoomByID   = `UPDATE rooms SET name = $2, room_type = $3, capacity = $4, status = $5, updated_at = $6 WHERE id = $1 RETURNING created_at`
	UpdateRoomStatus = `UPDATE room SET status = $2, updated_at = $3 WHERE id = $1 RETURNING name, room_type, capacity, created_at`
)
