package repository

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"database/sql"
	"log"
	"math"
)

type TransactionsRepository interface {
	Create(payload entity.Transaction) (entity.Transaction, error)
	List(page, size int) ([]entity.Transaction, model.Paging, error)
	GetTransactionById(id string) ([]entity.Transaction, error)
	GetTransactionByEmployeId(EmployeeId string) ([]entity.Transaction, error)
	UpdatePemission(payload entity.Transaction) (entity.Transaction, error)
}

type transactionsRepository struct {
	db *sql.DB
}

// list transaction (admin & GA) -GET
func (t *transactionsRepository) List(page, size int) ([]entity.Transaction, model.Paging, error) {
	var transactions []entity.Transaction
	offset := (page - 1) * size

	rows, err := t.db.Query(config.SelectTransactionList, size, offset)
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
			&transaction.CreatedAt,
			&transaction.UpdatedAt)
		if err != nil {
			log.Println("transactionsRepository.Rows.Next():", err.Error())
			return nil, model.Paging{}, err
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
			&transaction.CreatedAt,
			&transaction.UpdatedAt)
		if err != nil {
			log.Println("transactionRepository.Rows.Next():",
				err.Error())
			return nil, err
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
			&transaction.CreatedAt,
			&transaction.UpdatedAt)
		if err != nil {
			log.Println("transactionRepository.Rows.Next():",
				err.Error())
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

// (create transaction) Request booking rooms (employee & admin) -POST
func (t *transactionsRepository) Create(payload entity.Transaction) (entity.Transaction, error) {
	var transactions entity.Transaction
	// updatedAtStr := payload.UpdatedAt.Format("2006-01-02 15:04:05")
	// createdAtStr := payload.CreatedAt.Format("2006-01-02 15:04:05")

	err := t.db.QueryRow(config.InsertTransactions,
	payload.EmployeeId,
	payload.RoomId,
	payload.Description,
	payload.UpdatedAt).Scan(&payload.ID, &payload.Status, &payload.CreatedAt)
	if err != nil {
		return entity.Transaction{}, err
	}
	// updatedAt, _ := time.Parse("2006-01-02 15:04:05", updatedAtStr)
	// createdAt, _ := time.Parse("2006-01-02 15:04:05", createdAtStr)

	// payload.UpdatedAt = updatedAt
	// payload.CreatedAt = createdAt

	transactions = payload
	return transactions, err
}
// permission list (GA) -GET (batal)
// update permission (GA) -PUT
func (t *transactionsRepository) UpdatePemission(payload entity.Transaction) (entity.Transaction, error) {
	var transactions entity.Transaction

	 err := t.db.QueryRow(config.UpdatePermission,
		payload.Status,
		payload.ID).Scan(&transactions.EmployeeId, &transactions.RoomId,&transactions.Description, &transactions.UpdatedAt)
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
