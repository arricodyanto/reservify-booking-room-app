package repository

import (
	"booking-room-app/entity"
	"booking-room-app/entity/dto"
	"database/sql"
)

type RoomRepository interface {
	List() ([]entity.Rooms, error)
	Get(id string) (entity.Rooms, error)
	Create(payload entity.Rooms) (entity.Rooms, error)
	Update(payload entity.Rooms) (entity.Rooms, error)
	UpdateStatus(payload dto.RoomsDto) (dto.RoomsDto, error)
}

type roomRepository struct {
	db *sql.DB
}

// Create implements RoomRepository.
func (*roomRepository) Create(payload entity.Rooms) (entity.Rooms, error) {
	panic("unimplemented")
}

// Get implements RoomRepository.
func (*roomRepository) Get(id string) (entity.Rooms, error) {
	panic("unimplemented")
}

// List implements RoomRepository.
func (*roomRepository) List() ([]entity.Rooms, error) {
	panic("unimplemented")
}

// Update implements RoomRepository.
func (*roomRepository) Update(payload entity.Rooms) (entity.Rooms, error) {
	panic("unimplemented")
}

// UpdateStatus implements RoomRepository.
func (*roomRepository) UpdateStatus(payload dto.RoomsDto) (dto.RoomsDto, error) {
	panic("unimplemented")
}

// create room (ADMIN) -GET
// get all rooms (ALL ROLE) -GET
// get by room by ID (ALL ROLE) -GET
// update room status (GA & ADMIN) -PUT
// update room (ADMIN) -PUT

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepository{db: db}
}
