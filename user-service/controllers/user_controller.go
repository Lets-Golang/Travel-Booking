package controllers

import (
	"net/http"
	"strconv"

	"github.com/Lets-Golang/Travel-Booking/user-service/models"
	"github.com/Lets-Golang/Travel-Booking/user-service/services"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new user
// @Description Creates a user with name and email
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserCreateDTO true "User data"
// @Success 201 {object} models.UserDTO
// @Failure 400 {object} models.ErrorResponse
// @Router /users [post]
func CreateUser(service *services.UserService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var dto models.UserCreateDTO
        if err := c.BindJSON(&dto); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        user, err := service.Create(dto)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, user)
    }
}

// @Summary Get user by ID
// @Description Retrieves a user by their ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.UserDTO
// @Failure 404 {object} models.ErrorResponse
// @Router /users/{id} [get]
func GetUser(service *services.UserService) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        user, err := service.GetByID(id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusOK, user)
    }
}
