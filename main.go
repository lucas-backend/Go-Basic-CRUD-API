package main

import (
	"log"
	"my-api/config"
	"my-api/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Coba load file .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect ke database
	config.ConnectDB()
	defer config.DB.Close()

	// Buat router baru dari mux router
	r := mux.NewRouter()

	// Membuat subrouter khusus path /user
	userRoute := r.PathPrefix("/user").Subrouter()

	// Get all users
	userRoute.HandleFunc("/", controllers.GetAllUsers).Methods("GET")

	// Get user by id
	userRoute.HandleFunc("/{id}", controllers.GetUser).Methods("GET")

	// Add new user
	userRoute.HandleFunc("/", controllers.AddUser).Methods("POST")

	// Edit user by id
	userRoute.HandleFunc("/{id}", controllers.EditUsername).Methods("PUT")

	// Delete user by id
	userRoute.HandleFunc("/{id}", controllers.DeleteUser).Methods("DELETE")

	port := os.Getenv("SERVER_PORT")
	log.Println("Server berjalan di port " + port)
	http.ListenAndServe(":"+port, r)
}
