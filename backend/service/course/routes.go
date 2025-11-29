package course

import (
	"log"
	"net/http"

	"github.com/Jeno7u/studybud/service/auth"
	"github.com/Jeno7u/studybud/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store types.CourseStore
}

func NewHandler(store types.CourseStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) CourseRoutes(router gin.IRouter) {
	router.GET("/courses", auth.AuthMiddleware(), h.handleCourse)
}

func (h *Handler) handleCourse(c *gin.Context) {
	userID := c.GetInt("user_id")

	courses, err := h.store.GetCoursesByUserRelatedID(userID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, courses)
}
