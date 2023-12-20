package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/repository"
	"booking-room-app/shared/model"
	"fmt"
)

type RoomFacilityUsecase interface {
	FindAllRoomFacility(page, size int) ([]entity.RoomFacility, model.Paging, error)
	FindRoomFacilityById(id string) (entity.RoomFacility, error)
	AddRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error)
	UpdateRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error)
}

type roomFacilityUsecase struct {
	repo repository.RoomFacilityRepository
}

// find all room-facility
func (rf *roomFacilityUsecase) FindAllRoomFacility(page, size int) ([]entity.RoomFacility, model.Paging, error) {
	if page == 0 && size == 0 {
		page = 1
		size = 5
	}
	return rf.repo.List(page, size)
}

// find room-facility by id
func (rf *roomFacilityUsecase) FindRoomFacilityById(id string) (entity.RoomFacility, error) {
	return rf.repo.GetTransactionById(id)
}

// add room-facility
func (rf *roomFacilityUsecase) AddRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error) {

	transactions, err := rf.repo.Create(payload)
	if err != nil {
		return entity.RoomFacility{}, fmt.Errorf("oppps, failed to save data transations :%v", err.Error())
	}
	return transactions, nil
}

// update room-facility
func (rf *roomFacilityUsecase) UpdateRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error) {
	// get old record
	oldRoomFacility, err := rf.repo.GetTransactionById(payload.ID)
	if err != nil {
		return entity.RoomFacility{}, fmt.Errorf("oppps, failed to get previous data :%v", err.Error())
	}

	// partial update checking
	if payload.RoomId == "" {
		payload.RoomId = oldRoomFacility.RoomId
	}
	if payload.FacilityId == "" {
		payload.FacilityId = oldRoomFacility.FacilityId
	}
	if payload.Quantity == 0 {
		payload.Quantity = oldRoomFacility.Quantity
	}

	roomFacility, err := rf.repo.UpdateRoomFacility(payload)
	if err != nil {
		return entity.RoomFacility{}, fmt.Errorf("oppps, failed to update data transations :%v", err.Error())
	}
	return roomFacility, nil
}

func NewTransactionsUsecase(repo repository.RoomFacilityRepository) RoomFacilityUsecase {
	return &roomFacilityUsecase{repo: repo}
}
