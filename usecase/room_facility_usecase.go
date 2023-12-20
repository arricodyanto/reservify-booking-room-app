package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/repository"
	"booking-room-app/shared/model"
	"fmt"
	"time"
)

type RoomFacilityUsecase interface {
	FindAllRoomFacility(page, size int) ([]entity.RoomFacility, model.Paging, error)
	FindRoomFacilityById(id string) ([]entity.RoomFacility, error)
	AddRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error)
	UpdateRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error)
}

type roomFacilityUsecase struct {
	repo repository.RoomFacilityRepository
}

func (rf *roomFacilityUsecase) FindAllRoomFacility(page, size int) ([]entity.RoomFacility, model.Paging, error) {
	if page == 0 && size == 0 {
		page = 1
		size = 5
	}
	return rf.repo.List(page, size)
}

func (rf *roomFacilityUsecase) FindRoomFacilityById(id string) ([]entity.RoomFacility, error) {
	return rf.repo.GetTransactionById(id)
}

func (rf *roomFacilityUsecase) AddRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error) {
	// updatedAtStr := payload.UpdatedAt.Format("2006-01-02 15:04:05")
	payload.UpdatedAt = time.Now()
	// updatedAt, _ := time.Parse("2006-01-02 15:04:05", updatedAtStr)
	// payload.UpdatedAt = updatedAt

	transactions, err := rf.repo.Create(payload)
	if err != nil {
		return entity.RoomFacility{}, fmt.Errorf("oppps, failed to save data transations :%v", err.Error())
	}
	return transactions, nil
}

func (rf *roomFacilityUsecase) UpdateRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error) {
	transactions, err := rf.repo.UpdatePemission(payload)
	if err != nil {
		// fmt.Println(payload.Status)
		return entity.RoomFacility{}, fmt.Errorf("oppps, failed to update data transations :%v", err.Error())
	}
	return transactions, nil
}

func NewTransactionsUsecase(repo repository.RoomFacilityRepository) RoomFacilityUsecase {
	return &roomFacilityUsecase{repo: repo}
}
