package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/om1cael/scooter-api/internal/services"
)

type UserHandler interface {
	Register(c *gin.Context, name, email, password string)
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{service: service}
}

func (h *userHandler) Register(c *gin.Context, name, email, password string) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required, email"`
		Password string `json:"password" binding:"required, min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(c.Request.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, user)
}
