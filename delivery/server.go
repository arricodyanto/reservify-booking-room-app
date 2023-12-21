package delivery

import (
	"booking-room-app/config"
	"booking-room-app/delivery/controller"
	"booking-room-app/repository"
	"booking-room-app/usecase"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	roomUC         usecase.RoomUseCase
	facilitiesUC   usecase.FacilitiesUseCase
	roomFacilityUc usecase.RoomFacilityUsecase
	transactionsUc usecase.TransactionsUsecase
	engine         *gin.Engine
	host           string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)

	controller.NewRoomController(s.roomUC, rg).Route()
	controller.NewFacilitiesController(s.facilitiesUC, rg).Route()
	controller.NewRoomFacilityController(s.roomFacilityUc, rg).Route()
	controller.NewTransactionsController(s.transactionsUc, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic(err.Error())
	}

	// Inject DB ke -> repository
	roomRepo := repository.NewRoomRepository(db)
	facilityRepo := repository.NewFasilitesRepository(db)
	roomFacilityRepo := repository.NewRoomFacilityRepository(db)
	transactionsRepo := repository.NewTransactionsRepository(db)

	// Inject REPO ke -> useCase
	roomUC := usecase.NewRoomUseCase(roomRepo)
	facilitiesUC := usecase.NewFacilitiesUseCase(facilityRepo)
	roomFacilityUc := usecase.NewRoomFacilityUsecase(roomFacilityRepo)
	transactionsUc := usecase.NewTransactionsUsecase(transactionsRepo)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		roomUC:         roomUC,
		facilitiesUC:   facilitiesUC,
		transactionsUc: transactionsUc,
		roomFacilityUc: roomFacilityUc,
		engine:         engine,
		host:           host,
	}
}
