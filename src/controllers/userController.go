package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/gorm"
	"kyros-api/src/models"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	conn := connectToDatabase()
	defer conn.Close()
	
	// Fetch users from database
	var users []models.User
	conn.Find(&users)
	
	// Send response
	json.NewEncoder(w).Encode(users)
}
