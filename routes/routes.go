package routes

import (
	"crud_api/controller"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo){
	e.GET("/users", controller.GetUsers)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
}