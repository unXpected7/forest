package routes

import (
	"go-echo-crud/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/users", controllers.GetUsers)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
}
