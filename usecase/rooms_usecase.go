package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/repository"
	"fmt"
	"time"
)

type RoomUseCase interface {
	RegisterNewRoom(payload entity.Room) (entity.Room, error)
	FindRoomByID(id string) (entity.Room, error)
	FindAllRoom() ([]entity.Room, error)
	UpdateRoomDetail(payload entity.Room) (entity.Room, error)
	UpdateRoomStatus(payload entity.Room) (entity.Room, error)
}

type roomUseCase struct {
	repo repository.RoomRepository
}

// FindAllRoom implements RoomUseCase.
func (r *roomUseCase) FindAllRoom() ([]entity.Room, error) {
	return r.repo.List()
}

// FindRoomByID implements RoomUseCase.
func (r *roomUseCase) FindRoomByID(id string) (entity.Room, error) {
	return r.repo.Get(id)
}

// RegisterNewRoom implements RoomUseCase.
func (r *roomUseCase) RegisterNewRoom(payload entity.Room) (entity.Room, error) {
	if payload.Name == "" || payload.RoomType == "" || payload.Capacity == 0 {
		return entity.Room{}, fmt.Errorf("oops, field required")
	}

	payload.UpdatedAt = time.Now()
	room, err := r.repo.Create(payload)
	if err != nil {
		return entity.Room{}, fmt.Errorf("failed to create a new room list: %v", err.Error())
	}
	return room, nil
}

// UpdateRoomDetail implements RoomUseCase.
func (r *roomUseCase) UpdateRoomDetail(payload entity.Room) (entity.Room, error) {
	if payload.ID == "" || payload.Name == "" || payload.RoomType == "" || payload.Capacity == 0 || payload.Status == "" {
		return entity.Room{}, fmt.Errorf("oops, field required")
	}

	payload.UpdatedAt = time.Now()
	room, err := r.repo.Update(payload)
	if err != nil {
		return entity.Room{}, fmt.Errorf("failed to update room with ID %s: %v", payload.ID, err.Error())
	}
	return room, nil
}

// UpdateRoomStatus implements RoomUseCase.
func (r *roomUseCase) UpdateRoomStatus(payload entity.Room) (entity.Room, error) {
	if payload.ID == "" || payload.Status == "" {
		return entity.Room{}, fmt.Errorf("oops, field required")
	}

	payload.UpdatedAt = time.Now()
	room, err := r.repo.UpdateStatus(payload)
	if err != nil {
		return entity.Room{}, fmt.Errorf("failed to update room with ID %s: %v", payload.ID, err.Error())
	}
	return room, nil
}

func NewRoomUseCase(repo repository.RoomRepository) RoomUseCase {
	return &roomUseCase{repo: repo}
}
