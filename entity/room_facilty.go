package entity

import "time"

type RoomFacility struct {
	ID          string    `json:"id"`
<<<<<<< HEAD
	RoomId      string    `json:"roomId,omitempty"`
=======
	RoomId      string    `json:"roomId"`
>>>>>>> ca0f7afa9631e97e14f1f754df6fa18acb05d135
	FacilityId  string    `json:"facilityId"`
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
