package repository

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"database/sql"
	"fmt"
	"log"
	"math"
	"time"
)

type TransactionsRepository interface {
	Create(payload entity.Transaction) (entity.Transaction, error)
	List(page, size int, startDate, endDate time.Time) ([]entity.Transaction, model.Paging, error)
	GetTransactionById(id string) ([]entity.Transaction, error)
	GetTransactionByEmployeId(EmployeeId string) ([]entity.Transaction, error)
	UpdatePemission(payload entity.Transaction) (entity.Transaction, error)
}

type transactionsRepository struct {
	db *sql.DB
}

// list transaction (admin & GA) -GET
func (t *transactionsRepository) List(page, size int,startDate, endDate time.Time) ([]entity.Transaction, model.Paging, error) {
	var transactions []entity.Transaction
	offset := (page - 1) * size

	rows, err := t.db.Query(config.SelectTransactionList, size, offset, startDate, endDate)
	if err != nil {
		log.Println("transactionsRepository.Query:", err.Error())
		return nil, model.Paging{}, err
	}
	for rows.Next() {
		var transaction entity.Transaction
		err = rows.Scan(
			&transaction.ID,
			&transaction.EmployeeId,
			&transaction.RoomId,
			&transaction.Description,
			&transaction.Status,
			&transaction.StartTime,
			&transaction.EndTime,
			&transaction.CreatedAt,
			&transaction.UpdatedAt)
	if err != nil {
		log.Println("transactionsRepository.Rows.Next():", err.Error())
		return nil, model.Paging{}, err
	}

	RoomFacilitiesRows, err := t.db.Query(config.SelectRoomWithFacilities, transaction.RoomId)
	if err!= nil {
        log.Println("transactionsRepository.Query:", err.Error())
        return nil, model.Paging{}, err
    }
	for RoomFacilitiesRows.Next() {
		var roomFacility entity.RoomFacility
        err = RoomFacilitiesRows.Scan(
            &roomFacility.Id,
            &roomFacility.FacilityId,
            &roomFacility.Quantity,
            &roomFacility.CreatedAt,
            &roomFacility.UpdatedAt)
        if err!= nil {
            log.Println("transactionsRepository.Rows.Next():", err.Error())
            return nil, model.Paging{}, err
        }
        transaction.RoomFacilities = append(transaction.RoomFacilities, roomFacility)
    }

		transactions = append(transactions, transaction)
	}
	totalRows := 0
	if err := t.db.QueryRow(config.GetIdListTransaction).Scan(&totalRows); err != nil {
		return nil,
			model.Paging{}, err
	}

	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return transactions, paging, nil

}

// get transaction by id (GA) - GET
func (t *transactionsRepository) GetTransactionById(id string) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	rows, err := t.db.Query(config.SelectTransactionByID, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var transaction entity.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.EmployeeId,
			&transaction.RoomId,
			&transaction.Description,
			&transaction.Status,
			&transaction.StartTime, 
			&transaction.EndTime,
			&transaction.CreatedAt,
			&transaction.UpdatedAt)
		if err != nil {
			log.Println("transactionRepository.Rows.Next():",
				err.Error())
			return nil, err
		}

		RoomFacilitiesRows, err := t.db.Query(config.SelectRoomWithFacilities, transaction.RoomId)
		if err!= nil {
			log.Println("transactionsRepository.Query:", err.Error())
			return nil, err
		}
		for RoomFacilitiesRows.Next() {
			var roomFacility entity.RoomFacility
			err = RoomFacilitiesRows.Scan(
				&roomFacility.Id,
				&roomFacility.FacilityId,
				&roomFacility.Quantity,
				&roomFacility.CreatedAt,
				&roomFacility.UpdatedAt)
			if err!= nil {
				log.Println("transactionsRepository.Rows.Next():", err.Error())
				return nil, err
			}
			transaction.RoomFacilities = append(transaction.RoomFacilities, roomFacility)
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

// list transaction by employee ID (employee) -GET
func (t *transactionsRepository) GetTransactionByEmployeId(employeeId string) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	rows, err := t.db.Query(config.SelectTransactionByEmployeeID, employeeId)
	
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var transaction entity.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.EmployeeId,
			&transaction.RoomId,
			&transaction.Description,
			&transaction.Status,
			&transaction.StartTime, 
			&transaction.EndTime,
			&transaction.CreatedAt,
			&transaction.UpdatedAt)
		if err != nil {
			log.Println("transactionRepository.Rows.Next():",
				err.Error())
			return nil, err
		}
		RoomFacilitiesRows, err := t.db.Query(config.SelectRoomWithFacilities, transaction.RoomId)
		if err!= nil {
			log.Println("transactionsRepository.Query:", err.Error())
			return nil, err
		}
		for RoomFacilitiesRows.Next() {
			var roomFacility entity.RoomFacility
			err = RoomFacilitiesRows.Scan(
				&roomFacility.Id,
				&roomFacility.FacilityId,
				&roomFacility.Quantity,
				&roomFacility.CreatedAt,
				&roomFacility.UpdatedAt)
			if err!= nil {
				log.Println("transactionsRepository.Rows.Next():", err.Error())
				return nil, err
			}
			transaction.RoomFacilities = append(transaction.RoomFacilities, roomFacility)
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

// (create transaction) Request booking rooms (employee & admin) -POST
func (t *transactionsRepository) Create(payload entity.Transaction) (entity.Transaction, error) {
	var transactions entity.Transaction	
	err := t.db.QueryRow(config.InsertTransactions,
		payload.EmployeeId,
		payload.RoomId,
		payload.Description,
		payload.StartTime,
		payload.EndTime,
		payload.UpdatedAt).Scan(&payload.ID, &payload.Status, &payload.CreatedAt)
		if err != nil {
			return entity.Transaction{}, err
		}
		  // Insert ke tabel roomFacilities dan kurangi quantity di facilities
		  var roomFacilities[] entity.RoomFacility
		  for _, roomFacility := range payload.RoomFacilities {
			err = t.db.QueryRow(config.InsertRoomFacility,
				payload.RoomId,
				roomFacility.FacilityId,
				roomFacility.Quantity,
				payload.UpdatedAt).Scan(&roomFacility.Id, &roomFacility.CreatedAt, &roomFacility.UpdatedAt)
	
			if err != nil {
				return entity.Transaction{}, err
			}
			var quantity int
			err = t.db.QueryRow(config.SelectQuantityFacility,
				roomFacility.FacilityId).Scan(&quantity)
			if err != nil {
				return entity.Transaction{}, err
			}
			if roomFacility.Quantity > quantity {
				return entity.Transaction{}, fmt.Errorf("quantity more than stock")
			}
	
			// Kurangi quantity di tabel facilities
			err = t.db.QueryRow(config.UpdateFacilityQuantity,
				roomFacility.Quantity,
				roomFacility.FacilityId).Scan(&roomFacility.Id, &roomFacility.CreatedAt, &roomFacility.UpdatedAt)
			if err != nil {
				return entity.Transaction{}, err
			}
			roomFacilities = append(roomFacilities, roomFacility)
		}
		payload.RoomFacilities = roomFacilities


	transactions = payload
	return transactions, err
}
// permission list (GA) -GET (batal)
// update permission (GA) -PUT
func (t *transactionsRepository) UpdatePemission(payload entity.Transaction) (entity.Transaction, error) {
	var transactions entity.Transaction
	
	err := t.db.QueryRow(config.UpdatePermission,
		payload.Status,
		payload.ID,
		payload.UpdatedAt).Scan(&payload.EmployeeId, &payload.RoomId,&payload.Description,&payload.StartTime, &payload.EndTime, &payload.CreatedAt)
		if err != nil {
			log.Println("transactionsRepository.UpdateStatus:", err.Error())
			return entity.Transaction{}, err
		}

	transactions = payload
	return transactions, err
}

func NewTransactionsRepository(db *sql.DB) TransactionsRepository {
	return &transactionsRepository{db: db}
}
