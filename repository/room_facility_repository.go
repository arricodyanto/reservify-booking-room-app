package repository

// create room facilities (ADMIN) -POST
// get all room facilities (ADMIN) -GET
// get by ID room facilities (ADMIN) -GET
// update room facilites (ADMIN) -GET

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"database/sql"
	"fmt"
	"log"
	"math"
)

type RoomFacilityRepository interface {
	Create(payload entity.RoomFacility) (entity.RoomFacility, error)
	List(page, size int) ([]entity.RoomFacility, model.Paging, error)
	GetTransactionById(id string) (entity.RoomFacility, error)
	UpdateRoomFacility(payload entity.RoomFacility) (entity.RoomFacility, error)
}

type roomFacilityRepository struct {
	db *sql.DB
}

// get all room facilities (ADMIN) -GET
func (t *roomFacilityRepository) List(page, size int) ([]entity.RoomFacility, model.Paging, error) {
	var roomFacilities []entity.RoomFacility
	offset := (page - 1) * size

	rows, err := t.db.Query(config.SelectRoomFacilityList, size, offset)
	if err != nil {
		log.Println("roomFacilityRepository.Query:", err.Error())
		return nil, model.Paging{}, err
	}

	for rows.Next() {
		var roomFacility entity.RoomFacility
		err = rows.Scan(
			&roomFacility.ID,
			&roomFacility.RoomId,
			&roomFacility.FacilityId,
			&roomFacility.Quantity,
			&roomFacility.CreatedAt,
			&roomFacility.UpdatedAt)
		if err != nil {
			log.Println("roomFacilityRepository.Rows.Next():", err.Error())
			return nil, model.Paging{}, err
		}
		roomFacilities = append(roomFacilities, roomFacility)
	}

	totalRows := 0
	if err := t.db.QueryRow(config.GetCountRoomFacility).Scan(&totalRows); err != nil {
		return nil, model.Paging{}, err
	}

	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return roomFacilities, paging, nil

}

// get by ID room facilities (ADMIN) -GET
func (t *roomFacilityRepository) GetTransactionById(id string) (entity.RoomFacility, error) {
	var roomFacility entity.RoomFacility
	err := t.db.QueryRow(config.SelectRoomFacilityByID, id).Scan(
		&roomFacility.ID,
		&roomFacility.RoomId,
		&roomFacility.FacilityId,
		&roomFacility.Quantity,
		&roomFacility.CreatedAt,
		&roomFacility.UpdatedAt)
	if err != nil {
		return entity.RoomFacility{}, err
	}
	return roomFacility, nil
}

// create room facilities (ADMIN) -POST
func (t *roomFacilityRepository) Create(payload entity.RoomFacility) (entity.RoomFacility, error) {
	var roomFacilities entity.RoomFacility
	err := t.db.QueryRow(
		config.InsertRoomFacility,
		payload.RoomId,
		payload.FacilityId,
		payload.Quantity).
		Scan(
			&payload.ID,
			&payload.CreatedAt,
			&payload.UpdatedAt)
	if err != nil {
		return entity.RoomFacility{}, err
	}

	roomFacilities = payload
	return roomFacilities, err
}

// update room facilites (ADMIN) -GET
func (t *roomFacilityRepository) UpdateRoomFacility(payload entity.RoomFacility) (entity.RoomFacility, error) {
	var roomFacility entity.RoomFacility

	err := t.db.QueryRow(
		config.UpdateRoomFacility,
		payload.RoomId,
		payload.FacilityId,
		payload.Quantity,
		payload.ID).Scan(&payload.CreatedAt, &payload.UpdatedAt)
	if err != nil {
		log.Println("roomFacilityRepository.UpdateRoomFacility:", err.Error())
		return entity.RoomFacility{}, err
	}

	roomFacility = payload
	fmt.Println(payload)
	return roomFacility, err
}

func NewTransactionsRepository(db *sql.DB) RoomFacilityRepository {
	return &roomFacilityRepository{db: db}
}
