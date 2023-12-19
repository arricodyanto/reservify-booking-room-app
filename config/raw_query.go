package config

const (
	SelectRoomFacilityList = `SELECT id, room_id, facility_id, quantity, created_at, updated_at FROM trx_room_facility ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectRoomFacilityByID = ``
	InsertRoomFacility     = ``
	UpdateRoomFacility     = ``
	GetCountRoomFacility   = `SELECT COUNT(*) FROM trx_room_facility`
)
