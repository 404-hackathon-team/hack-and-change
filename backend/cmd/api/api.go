package api

import (
	"database/sql"
	"github.com/Jeno7u/studybud/utils"
	"log"

	"github.com/Jeno7u/studybud/service/user"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := gin.Default()

	router_v1 := router.Group("/api/v1")

	userStore := user.NewStore(s.db)
	dataProvider := utils.NewDataProvider(s.db)
	userHandler := user.NewHandler(userStore, *dataProvider)
	userHandler.RegisterRoutes(router_v1)

	log.Println("Listening on", s.addr)
	return router.Run(s.addr)
}
