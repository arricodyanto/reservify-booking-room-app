package entity

import "time"

type Transaction struct {
	ID          string `json:"id"`
	EmployeeId  string `json:"employeeId"`
	RoomId      string `json:"roomId"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}