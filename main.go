package main

import (
	"net/http"

	"encoding/json"

	_ "log"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	config "trongpham.dev/todo/helpers"

	controllers "trongpham.dev/todo/controllers"
)

func main() {

	config.MongoConnection()

	router := mux.NewRouter()

	router.HandleFunc("/", _initServer)

	//task
	router.HandleFunc("/api/v1/task/get-all", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/api/v1/task/get-detail", controllers.GetDetail).Methods("GET")
	router.HandleFunc("/api/v1/task/create", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/v1/task/update", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/api/v1/task/delete/{id}", controllers.DeleteTask).Methods("DELETE")

	//user
	router.HandleFunc("/api/v1/user/create", controllers.CreateUser).Methods("POST")

	//project
	router.HandleFunc("/api/v1/project/create", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/api/v1/project/get-all", controllers.GetProjects).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	fmt.Println("Server listenning on port 8080 ...")
	http.ListenAndServe(":8080", handler)
}

func _initServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Welcome to the API")
}
