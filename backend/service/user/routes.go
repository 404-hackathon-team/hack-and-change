package user

import (
	"fmt"
	"net/http"

	"github.com/Jeno7u/studybud/config"
	"github.com/Jeno7u/studybud/service/auth"
	"github.com/Jeno7u/studybud/types"
	"github.com/Jeno7u/studybud/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router gin.IRouter) {
	router.POST("/login", h.handleLogin)
	router.POST("/register", h.handleRegister)
}

func (h *Handler) handleLogin(c *gin.Context) {
	// get JSON payload
	var payload types.LoginUserPayload
	if err := c.BindJSON(&payload); err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("error during JSON converting, %v", err))
		return
	}

	// validate the payload
	if err := utils.Vaildate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// check if the user exists
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("not found, invalid email or passowrd"))
		return
	}

	if !auth.ComparePassowrds(u.Password, []byte(payload.Password)) {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("not found, invalid email or passowrd"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(c.Writer, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(c *gin.Context) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Vaildate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	// if it doesnt we create the new user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}
	
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(c.Writer, http.StatusCreated, map[string]string{"token": token})

}

// func getPayload(c *gin.Context) (*types.RegisterUserPayload, error) {
// 	// get JSON payload
// 	var payload types.RegisterUserPayload
// 	if err := c.ShouldBindJSON(&payload); err != nil {
// 		utils.WriteError(c.Writer, http.StatusBadRequest, err)
// 		return nil, err

// 	return payload, nil
// }