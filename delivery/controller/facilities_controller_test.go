package controller

import (
	"booking-room-app/entity"
	usecase_mock "booking-room-app/mock/usecase"
	"booking-room-app/shared/model"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
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
	suite.fum.On("FindAllFacilities", mockFacility).Return(mockFacility, mockPaging, nil)
	// request, err := http.NewRequest(http.MethodGet, "/api/v1/facilities", nil)
	// assert.NoError(suite.T(), err)
	// responseRecorder := httptest.NewRecorder()
	// handlerFunc := NewFacilitiesController(suite.fum, suite.rg).listHandler()

}

func TestFacilitiesControllerTestSuite(t *testing.T) {
	suite.Run(t, new(FacilitiesControllerTestSuite))
}
