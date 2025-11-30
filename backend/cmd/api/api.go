package api

import (
	"database/sql"
	"log"

	"github.com/Jeno7u/studybud/utils"

	"github.com/Jeno7u/studybud/service/course"
	"github.com/Jeno7u/studybud/service/notification"
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

	courseStore := course.NewStore(s.db)
	courseHandler := course.NewHandler(courseStore)
	courseHandler.CourseRoutes(router_v1)

	notificationStore := notification.NewStore(s.db)
	notificationHandler := notification.NewHandler(notificationStore)
	notificationHandler.NotificationRoutes(router_v1)
	go notificationStore.BroadcastNotifications()

	log.Println("Listening on", s.addr)
	return router.Run(s.addr)
}
