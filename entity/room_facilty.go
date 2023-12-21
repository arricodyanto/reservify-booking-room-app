package entity

import "time"

type RoomFacility struct {
	ID         string    `json:"id"`
	RoomId     string    `json:"roomId"`
	FacilityId string    `json:"facilityId"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
