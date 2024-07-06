package controllers

import (
	"go-echo-crud/database"
	"go-echo-crud/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c echo.Context) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User"
// @Success 201 {object} models.User
// @Router /users [post]
func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	database.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update an existing user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	database.DB.First(&user, id)
	if err := c.Bind(&user); err != nil {
		return err
	}
	database.DB.Save(&user)
	return c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete an existing user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	database.DB.First(&user, id)
	database.DB.Delete(&user)
	return c.NoContent(http.StatusNoContent)
}
