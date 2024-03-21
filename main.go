package main

import (
	"restaurant-api/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/health", handlers.HealthHandler)

	e.GET("/users", handlers.GetAllUsers)
	e.GET("/users/:id", handlers.GetUserByID)
	e.POST("/users", handlers.CreateUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)

	e.Logger.Fatal(e.Start(":9000"))
}
