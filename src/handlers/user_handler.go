package handlers

import (
	"net/http"
	"strconv"
	"time"

	"gin-users-api/models"
	"gin-users-api/repositories"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) GetUsers(c *gin.Context) {

	users, err := h.Repo.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.Repo.GetByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound,
			gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	user.AddedAt = time.Now()

	err := h.Repo.Create(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.Repo.GetByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound,
			gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	user.UpdatedAt = time.Now()

	h.Repo.Update(&user)

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	err := h.Repo.Delete(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
