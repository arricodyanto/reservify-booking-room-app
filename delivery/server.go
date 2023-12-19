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
	roomUC usecase.RoomUseCase
	engine *gin.Engine
	host   string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)

	controller.NewRoomController(s.roomUC, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("Server can not run on host %s, because of error: %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	// fmt.Println(dsn)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic(err.Error())
	}

	roomRepo := repository.NewRoomRepository(db)

	roomUC := usecase.NewRoomUseCase(roomRepo)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		roomUC: roomUC,
		engine: engine,
		host:   host,
	}
}
