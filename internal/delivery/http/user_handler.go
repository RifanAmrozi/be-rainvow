package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rifanamrozi/be-rainvow/internal/usecase"
)

type UserHandler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(router *gin.Engine, uc usecase.UserUsecase) {
	handler := &UserHandler{uc: uc}

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", handler.GetUsers)
		v1.GET("/users/:id", handler.GetUserByID)
		v1.POST("/users", handler.CreateUser)
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.uc.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := h.uc.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	user, err := h.uc.CreateUser(req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}
