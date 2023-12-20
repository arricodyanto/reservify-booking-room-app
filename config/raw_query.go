package config

const (
	SelectRoomFacilityList = `SELECT id, room_id, facility_id, quantity, created_at, updated_at FROM trx_room_facility ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectRoomFacilityByID = `SELECT id, room_id, facility_id, quantity, created_at, updated_at FROM trx_room_facility WHERE id = $1`
	InsertRoomFacility     = `INSERT INTO trx_room_facility (room_id, facility_id, quantity) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	UpdateRoomFacility     = ``
	GetCountRoomFacility   = `SELECT COUNT(*) FROM trx_room_facility`
)
