package controller

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"booking-room-app/shared/common"
	"booking-room-app/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomFacilityController struct {
	transactionUC usecase.RoomFacilityUsecase
	rg            *gin.RouterGroup
	// authMiddleware middleware.AuthMiddleware
}

func (t *RoomFacilityController) createHandler(ctx *gin.Context) {
	var payload entity.RoomFacility
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	transactions, err := t.transactionUC.AddRoomFacilityTransaction(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, transactions, "Created")
}

func (t *RoomFacilityController) listHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))

	transactions, paging, err := t.transactionUC.FindAllRoomFacility(page, size)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var response []interface{}
	for _, v := range transactions {
		response = append(response, v)
	}
	common.SendPagedResponse(ctx, response, paging, "Ok")
}

func (t *RoomFacilityController) getTransactionById(ctx *gin.Context) {
	id := ctx.Param("id")
	transactions, err := t.transactionUC.FindRoomFacilityById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "transaction with transaction ID "+id+" not found")
		return
	}

	common.SendSingleResponse(ctx, transactions, "Ok")
}

func (t *RoomFacilityController) updateStatusHandler(ctx *gin.Context) {
	var payload entity.RoomFacility
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	transactions, err := t.transactionUC.UpdateRoomFacilityTransaction(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, transactions, "Updated")
}

func (t *RoomFacilityController) Route() {
	t.rg.GET(config.TransactionList, t.listHandler)
	t.rg.GET(config.TransactionGetById, t.getTransactionById)
	t.rg.POST(config.TransactionCreate, t.createHandler)
	t.rg.PUT(config.TransactionUpdatePerm, t.updateStatusHandler)
}

func NewRoomFacilityController(transactionUC usecase.RoomFacilityUsecase, rg *gin.RouterGroup) *RoomFacilityController {
	return &RoomFacilityController{
		transactionUC: transactionUC,
		rg:            rg,
		// authMiddleware: authMiddleware,
	}
}
