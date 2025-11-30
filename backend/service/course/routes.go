package course

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Jeno7u/studybud/service/auth"
	"github.com/Jeno7u/studybud/types"
	"github.com/Jeno7u/studybud/utils"
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
		// treat "no rows" as empty result and return {}
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{})
			return
		}
		log.Println(err)
		utils.WriteError(c.Writer, http.StatusBadRequest, err)
		return
	}

	if len(courses) == 0 {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	c.JSON(http.StatusOK, courses)
}
