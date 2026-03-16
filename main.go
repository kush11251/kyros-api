package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	loadEnvVariables()
	conn := connectToDatabase()
	defer conn.Close()
	
	router := mux.NewRouter()
	
	// Define routes
	router.HandleFunc("/users", getUsers).Methods("GET")
	
	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", router)
}
