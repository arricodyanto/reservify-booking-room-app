package controller

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"booking-room-app/shared/common"
	"booking-room-app/shared/model"
	"booking-room-app/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	roomUC usecase.RoomUseCase
	rg     *gin.RouterGroup
}

func (r *RoomController) createHandler(c *gin.Context) {
	var payload entity.Room
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	room, err := r.roomUC.RegisterNewRoom(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendCreateResponse(c, room, "Created")
}

func (r *RoomController) getHandler(c *gin.Context) {
	id := c.Param("id")
	room, err := r.roomUC.FindRoomByID(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusNotFound, fmt.Sprintf("Room with ID %s not found", id))
		return
	}
	common.SendSingleResponse(c, room, "Ok")
}

func (r *RoomController) listHandler(c *gin.Context) {
	rooms, err := r.roomUC.FindAllRoom()
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var response []interface{}
	for _, v := range rooms {
		response = append(response, v)
	}
	common.SendPagedResponse(c, response, model.Paging{}, "Ok")
}

func (r *RoomController) updateDetailHandler(c *gin.Context) {
	var payload entity.Room
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	room, err := r.roomUC.UpdateRoomDetail(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendCreateResponse(c, room, "Updated")
}
func (r *RoomController) updateStatusHandler(c *gin.Context) {
	var payload entity.Room
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	room, err := r.roomUC.UpdateRoomStatus(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendSingleResponse(c, room, "Ok")
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
