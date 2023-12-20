package entity

import "time"

type RoomFacility struct {
	Id string `json:"id"`
	RoomId     string  `json:"roomId,omitempty"`
	FacilityId string  `json:"facilityId"`
	Quantity   int `json:"quantity"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
