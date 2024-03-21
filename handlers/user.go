package handlers

import (
	"database/sql"
	"net/http"
	"restaurant-api/db"
	"restaurant-api/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	db, err := db.Connect()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to connect to database")
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, username, email FROM users")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch users from database")
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email); err != nil {
			return c.String(http.StatusInternalServerError, "Failed to scan user row")

		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return c.String(http.StatusInternalServerError, "Error iterating over user rows")
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}

	db, err := db.Connect()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to connect to database")

	}
	defer db.Close()

	row := db.QueryRow("SELECT id, name, username, email FROM users WHERE id = $1", userID)

	var user models.User

	err = row.Scan(&user.ID, &user.Name, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, "User not found")

		}
		return c.String(http.StatusInternalServerError, "Failed to fetch user from database")
	}

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "Failed to decode request body")
	}

	db, err := db.Connect()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to connect to database")
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4)", user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to insert user into database")
	}

	return c.JSON(http.StatusCreated, user)
}

func UpdateUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}

	var updatedUser models.User
	if err := c.Bind(&updatedUser); err != nil {
		return c.String(http.StatusBadRequest, "Failed to decode request body")
	}

	db, err := db.Connect()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to connect to database")
	}
	defer db.Close()

	_, err = db.Exec("UPDATE users SET name = $1, username = $2, email = $3 WHERE id = $4",
		updatedUser.Name, updatedUser.Username, updatedUser.Email, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to update user in database")
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}

	db, err := db.Connect()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to connect to database")
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete user from database")
	}

	return c.String(http.StatusOK, "User deleted successfully")
}
