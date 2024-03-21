package handlers

import (
	"net/http"
	"restaurant-api/db"

	"github.com/labstack/echo/v4"
)

func HealthHandler(c echo.Context) error {
	db, err := db.Connect()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to connect to database")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to ping database")
	}

	return c.String(http.StatusOK, "Database is connected successfully!")
}
