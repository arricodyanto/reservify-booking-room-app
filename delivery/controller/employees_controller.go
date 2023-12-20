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

type EmployeeController struct {
	employeeUC usecase.EmployeesUseCase
	rg         *gin.RouterGroup
}

// untuk create employee
func (e *EmployeeController) createHandler(ctx *gin.Context){
	var payload entity.Employee
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	employee, err := e.employeeUC.RegisterNewEmployee(payload)
	if err != nil{
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	common.SendCreateResponse(ctx, employee, "Created")
}

// Read
// func (e *EmployeeController) getHandler(ctx *gin.Context){
// 	employee, err := e.employeeUC.FindAllEmployees()
// 	if err != nil{
// 		common.SendErrorResponse(ctx, http.StatusNotFound, "Employee's not found")
// 		return
// 	}
// 	common.SendSingleResponse(ctx, employee, "Ok")
// }
// read by
func (e *EmployeeController) getByIdHandler(ctx *gin.Context){
	id := ctx.Param("id")
	employee, err := e.employeeUC.FindEmployeesByID(id)
	if err != nil{
		common.SendErrorResponse(ctx, http.StatusNotFound, "Employee with ID "+id+" not found")
		return
	}
	common.SendSingleResponse(ctx, employee, "Ok")
}

// update
func (e *EmployeeController) putHandler(ctx *gin.Context)  {
	var payload entity.Employee
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx,http.StatusBadRequest,"Failed to bind data")
	}
	employee, err := e.employeeUC.UpdateEmployee(payload)
	if err != nil {
		common.SendErrorResponse(ctx,http.StatusNotFound,"Employee Not Found")
	}
	common.SendSingleResponse(ctx,employee,"Updated Successfully")

}
// pagination
func (e *EmployeeController) ListHandler(ctx *gin.Context){
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	employees, paging, err := e.employeeUC.ListAll(page, size)
	if err != nil{
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}
	for _, v := range employees{
		response = append(response,v)
	}
	common.SendPagedResponse(ctx, response, paging, "Ok")
}
// route

func (e *EmployeeController) Route() {
	e.rg.GET(config.EmployeesGetById, e.getByIdHandler)
	// belum
	// e.rg.GET(config.EmployeesList, e.getHandler)
	// berhasil tapi belum ada validasi
	e.rg.POST(config.EmployeesCreate, e.createHandler)
	// put
	e.rg.PUT(config.EmployeesUpdate, e.putHandler)

	// list
	e.rg.GET(config.EmployeesList, e.ListHandler)

}


func NewEmployeeController(employeeUC usecase.EmployeesUseCase, rg *gin.RouterGroup) *EmployeeController{
	return &EmployeeController{
		employeeUC: employeeUC,
		rg:         rg,
	}
}
