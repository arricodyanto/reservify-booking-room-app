package controller

import (
	"booking-room-app/entity"
	usecase_mock "booking-room-app/mock/usecase"
	"booking-room-app/shared/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FacilitiesControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	fum *usecase_mock.FacilitiesUseCaseMock
}

func (suite *FacilitiesControllerTestSuite) SetupTest() {
	suite.fum = new(usecase_mock.FacilitiesUseCaseMock)
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	suite.rg = router.Group("/api/v1")
}

var expectedFasilities = entity.Facilities{
	ID:        "1",
	Name:      "This is name",
	Quantity:  10,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func (suite *FacilitiesControllerTestSuite) TestListHandler_Success() {
	mockFacility := []entity.Facilities{expectedFasilities}
	mockPaging := model.Paging{
		Page:        1,
		RowsPerPage: 1,
		TotalRows:   5,
		TotalPages:  1,
	}
	suite.fum.On("FindAllFacilities").Return(mockFacility, mockPaging, nil)
	handlerFunc := NewFacilitiesController(suite.fum, suite.rg)
	request, err := http.NewRequest(http.MethodGet, "/api/v1/facilities", nil)
	assert.NoError(suite.T(), err)
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request
	ctx.Set("facilities", mockFacility)
	handlerFunc.listHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func TestFacilitiesControllerTestSuite(t *testing.T) {
	suite.Run(t, new(FacilitiesControllerTestSuite))
}
