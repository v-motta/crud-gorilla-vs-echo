package handlers

import (
	"net/http"
	"restaurant-api/db"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		http.Error(w, "Failed to ping database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database is connected successfully!"))
}
