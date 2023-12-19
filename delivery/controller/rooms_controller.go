package controller

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"booking-room-app/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	roomUC usecase.RoomUseCase
	rg     *gin.RouterGroup
}

func (r *RoomController) createHandler(c *gin.Context) {
	var payload entity.Room
	if err := c.ShouldBindJSON(&payload); err != nil {
		// send error
		return
	}

	room, err := r.roomUC.RegisterNewRoom(payload)
	if err != nil {
		// senderror
		return
	}
	// sendsingle
	fmt.Println(room)
}

func (r *RoomController) getHandler(c *gin.Context) {
	id := c.Param("id")
	room, err := r.roomUC.FindRoomByID(id)
	if err != nil {
		// senderror
		return
	}
	// sendsingle
	fmt.Println(room)
}

func (r *RoomController) listHandler(c *gin.Context) {
	room, err := r.roomUC.FindAllRoom()
	if err != nil {
		// senderror
		return
	}
	// sendpaged
	fmt.Println(room)
}

func (r *RoomController) updateDetailHandler(c *gin.Context) {
	var payload entity.Room
	if err := c.ShouldBindJSON(&payload); err != nil {
		// send error
		return
	}

	room, err := r.roomUC.UpdateRoomDetail(payload)
	if err != nil {
		// senderror
		return
	}
	// sendsingle
	fmt.Println(room)
}
func (r *RoomController) updateStatusHandler(c *gin.Context) {
	var payload entity.Room
	if err := c.ShouldBindJSON(&payload); err != nil {
		// send error
		return
	}

	room, err := r.roomUC.UpdateRoomStatus(payload)
	if err != nil {
		// senderror
		return
	}
	// sendsingle
	fmt.Println(room)
}

func (r *RoomController) Route() {
	r.rg.POST(config.RoomCreate, r.createHandler)
	r.rg.GET(config.RoomList, r.listHandler)
	r.rg.GET(config.RoomGetById, r.getHandler)
	r.rg.PUT(config.RoomUpdate, r.updateDetailHandler)
	r.rg.PUT(config.RoomUpdateStatus, r.updateStatusHandler)
}

func NewRoomController(roomUC usecase.RoomUseCase, rg *gin.RouterGroup) *RoomController {
	return &RoomController{roomUC: roomUC, rg: rg}
}
