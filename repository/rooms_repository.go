package repository

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"database/sql"
	"log"
)

type RoomRepository interface {
	Create(payload entity.Room) (entity.Room, error)
	Get(id string) (entity.Room, error)
	List() ([]entity.Room, error)
	Update(payload entity.Room) (entity.Room, error)
	UpdateStatus(payload entity.Room) (entity.Room, error)
}

type roomRepository struct {
	db *sql.DB
}

// Create implements RoomRepository.
func (r *roomRepository) Create(payload entity.Room) (entity.Room, error) {
	var room entity.Room
	err := r.db.QueryRow(config.InsertRoom, payload.Name, payload.RoomType, payload.Capacity, payload.Status, payload.UpdatedAt).Scan(&room.ID, &room.CreatedAt)
	if err != nil {
		log.Println("roomRepository.CreateQueryRow", err.Error())
		return entity.Room{}, err
	}

	room.Name = payload.Name
	room.RoomType = payload.RoomType
	room.Capacity = payload.Capacity
	room.Status = payload.Status
	room.UpdatedAt = payload.UpdatedAt

	return room, nil
}

// Get implements RoomRepository.
func (r *roomRepository) Get(id string) (entity.Room, error) {
	var room entity.Room
	err := r.db.QueryRow(config.SelectRoomByID, id).Scan(&room.ID, &room.Name, &room.RoomType, &room.Capacity, &room.Status, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		log.Println("roomRepository.GetQueryRow", err.Error())
		return entity.Room{}, err
	}

	return room, nil
}

// List implements RoomRepository.
func (r *roomRepository) List() ([]entity.Room, error) {
	var rooms []entity.Room
	rows, err := r.db.Query(config.SelectRoomList)
	if err != nil {
		log.Println("roomRepository.ListQuery", err.Error())
		return []entity.Room{}, err
	}

	for rows.Next() {
		var room entity.Room
		err := rows.Scan(&room.ID, &room.Name, &room.RoomType, &room.Capacity, &room.Status, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			log.Println("roomRepository.ListScan", err.Error())
			return []entity.Room{}, err
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}

// Update implements RoomRepository.
func (r *roomRepository) Update(payload entity.Room) (entity.Room, error) {
	var room entity.Room
	room.ID = payload.ID

	err := r.db.QueryRow(config.UpdateRoomByID, room.ID, payload.Name, payload.RoomType, payload.Capacity, payload.Capacity, payload.Status, payload.UpdatedAt).Scan(&room.CreatedAt)
	if err != nil {
		log.Println("roomRepository.UpdateQueryRow", err.Error())
		return entity.Room{}, err
	}

	room.Name = payload.Name
	room.RoomType = payload.RoomType
	room.Capacity = payload.Capacity
	room.Status = payload.Status
	room.UpdatedAt = payload.UpdatedAt

	return room, nil
}

// UpdateStatus implements RoomRepository.
func (r *roomRepository) UpdateStatus(payload entity.Room) (entity.Room, error) {
	var room entity.Room
	room.ID = payload.ID

	err := r.db.QueryRow(config.UpdateRoomStatus, room.ID, payload.Status, payload.UpdatedAt).Scan(&room.Name, &room.RoomType, &room.Capacity, &room.CreatedAt)
	if err != nil {
		log.Println("roomRepository.UpdateStatusQueryRow", err.Error())
		return entity.Room{}, err
	}

	room.Status = payload.Status
	room.UpdatedAt = payload.UpdatedAt

	return room, nil
}

// create room (ADMIN) -GET
// get all rooms (ALL ROLE) -GET
// get by room by ID (ALL ROLE) -GET
// update room status (GA & ADMIN) -PUT
// update room (ADMIN) -PUT

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepository{db: db}
}
