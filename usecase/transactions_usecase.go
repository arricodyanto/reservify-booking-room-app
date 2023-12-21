package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/repository"
	"booking-room-app/shared/model"
	"fmt"
	"time"
)

type TransactionsUsecase interface {
	FindAllTransactions(page, size int,startDate, endDate time.Time) ([]entity.Transaction, model.Paging, error)
	FindTransactionsById(id string) ([]entity.Transaction, error)
	FindTransactionsByEmployeeId(employeeId string) ([]entity.Transaction, error)
	RequestNewBookingRooms(payload entity.Transaction) (entity.Transaction, error)
	AccStatusBooking(payload entity.Transaction) (entity.Transaction, error)
	
}

type transactionsUsecase struct {
	repo repository.TransactionsRepository
}

func (t *transactionsUsecase) FindAllTransactions(page, size int, startDate, endDate time.Time) ([]entity.Transaction, model.Paging, error) {
	return t.repo.List(page, size, startDate, endDate)
}

func (t *transactionsUsecase) FindTransactionsById(id string) ([]entity.Transaction, error) {
	return t.repo.GetTransactionById(id)
}

func (t *transactionsUsecase) FindTransactionsByEmployeeId(employeeId string) ([]entity.Transaction, error) {
	return t.repo.GetTransactionByEmployeId(employeeId)
}

func (t *transactionsUsecase) RequestNewBookingRooms(payload entity.Transaction) (entity.Transaction, error) {
	// updatedAtStr := payload.UpdatedAt.Format("2006-01-02 15:04:05")
	payload.UpdatedAt = time.Now()
	// updatedAt, _ := time.Parse("2006-01-02 15:04:05", updatedAtStr)
	// payload.UpdatedAt = updatedAt

	transactions, err := t.repo.Create(payload)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf("oppps, failed to save data transations :%v", err.Error())
	}
		return transactions, nil
}

func (t *transactionsUsecase) AccStatusBooking(payload entity.Transaction) (entity.Transaction, error) {
	payload.UpdatedAt = time.Now()
	transactions, err := t.repo.UpdatePemission(payload)
	if err != nil {
		// fmt.Println(payload.Status)
		return entity.Transaction{}, fmt.Errorf("oppps, failed to update data transations :%v", err.Error())
	}
		return transactions, nil
}

func NewTransactionsUsecase(repo repository.TransactionsRepository) TransactionsUsecase {
	return &transactionsUsecase{repo: repo}
}