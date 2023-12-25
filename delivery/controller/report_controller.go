package controller

import (
	"booking-room-app/shared/common"
	"booking-room-app/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	reportUC usecase.ReportUseCase
	rg       *gin.RouterGroup
}

func (r *ReportController) downloadHandler(c *gin.Context) {
	_, err := r.reportUC.PrintAllReports("", "")
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "transaction.csv"))
	c.Header("Content-Type", "text/csv")
	common.SendSingleResponse(c, http.StatusOK, "downloading file..")
	c.File("/reports/transaction.csv")
}

func (r *ReportController) Route() {
	r.rg.GET("/reports/download", r.downloadHandler)
}

func NewReportController(reportUC usecase.ReportUseCase, rg *gin.RouterGroup) *ReportController {
	return &ReportController{reportUC: reportUC, rg: rg}
}
