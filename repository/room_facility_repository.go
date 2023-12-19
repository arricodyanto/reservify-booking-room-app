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
	"log"
	"math"
)

type RoomFacilityRepository interface {
	Create(payload entity.RoomFacility) (entity.RoomFacility, error)
	List(page, size int) ([]entity.RoomFacility, model.Paging, error)
	GetTransactionById(id string) ([]entity.RoomFacility, error)
	UpdatePemission(payload entity.RoomFacility) (entity.RoomFacility, error)
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
	if err := t.db.QueryRow(config.GetIdListRoomFacility).Scan(&totalRows); err != nil {
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
func (t *roomFacilityRepository) GetTransactionById(id string) ([]entity.RoomFacility, error) {
	var roomFacilities []entity.RoomFacility
	rows, err := t.db.Query(config.SelectRoomFacilityByID, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var roomFacility entity.RoomFacility
		err := rows.Scan(
			&roomFacility.ID,
			&roomFacility.RoomId,
			&roomFacility.FacilityId,
			&roomFacility.Quantity,
			&roomFacility.CreatedAt,
			&roomFacility.UpdatedAt)
		if err != nil {
			log.Println("roomFacilityRepository.Rows.Next():",
				err.Error())
			return nil, err
		}
		roomFacilities = append(roomFacilities, roomFacility)
	}
	return roomFacilities, nil
}

// create room facilities (ADMIN) -POST
func (t *roomFacilityRepository) Create(payload entity.RoomFacility) (entity.RoomFacility, error) {
	var roomFacilities entity.RoomFacility
	// updatedAtStr := payload.UpdatedAt.Format("2006-01-02 15:04:05")
	// createdAtStr := payload.CreatedAt.Format("2006-01-02 15:04:05")

	err := t.db.QueryRow(
		config.InsertRoomFacility,
		payload.RoomId,
		payload.FacilityId,
		payload.Quantity,
		payload.UpdatedAt).
		Scan(
			&payload.ID,
			&payload.CreatedAt)
	if err != nil {
		return entity.RoomFacility{}, err
	}
	// updatedAt, _ := time.Parse("2006-01-02 15:04:05", updatedAtStr)
	// createdAt, _ := time.Parse("2006-01-02 15:04:05", createdAtStr)

	// payload.UpdatedAt = updatedAt
	// payload.CreatedAt = createdAt

	roomFacilities = payload
	return roomFacilities, err
}

// update room facilites (ADMIN) -GET
func (t *roomFacilityRepository) UpdatePemission(payload entity.RoomFacility) (entity.RoomFacility, error) {
	var roomFacility entity.RoomFacility

	err := t.db.QueryRow(
		config.UpdateRoomFacility,
		payload.RoomId,
		payload.FacilityId,
		payload.Quantity,
		payload.UpdatedAt,
		payload.ID).Scan(&roomFacility.CreatedAt)
	if err != nil {
		log.Println("roomFacilityRepository.UpdateStatus:", err.Error())
		return entity.RoomFacility{}, err
	}

	roomFacility = payload
	return roomFacility, err
}

func NewTransactionsRepository(db *sql.DB) RoomFacilityRepository {
	return &roomFacilityRepository{db: db}
}
