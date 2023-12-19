package delivery

import (
	"booking-room-app/config"
	"booking-room-app/delivery/controller"
	"booking-room-app/repository"
	"booking-room-app/usecase"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	transactionsUc usecase.TransactionsUsecase
	engine *gin.Engine
	host   string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)
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
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sql.Open(cfg.Driver, dns)
	if err != nil {
		panic("connection error")
	}

	transactionsRepo := repository.NewTransactionsRepository(db)
	transactionsUc := usecase.NewTransactionsUsecase(transactionsRepo)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		transactionsUc: transactionsUc,
		engine:     engine,
		host:       host,
		// jwtService: jwtService,
	}
}