package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	models "trongpham.dev/todo/models"
	services "trongpham.dev/todo/services"
)

func GetTasks(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := []*models.Task{}

	data := map[string]interface{}{}

	err := json.NewDecoder(req.Body).Decode(&data)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer req.Body.Close()

	filter := bson.D{{}}
	response, err = services.GetTasks(filter)

	if err != nil {

		// response = []*models.Taskinterface{}{"error": err.Error()}

	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")

	if err := enc.Encode(response); err != nil {

		fmt.Println(err.Error())

	}
}

func GetDetail(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	query := req.URL.Query()

	response, err := services.GetDetailLookUp(query)

	if err != nil {

		// response = []*models.Taskinterface{}{"error": err.Error()}

	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")

	if err := enc.Encode(response); err != nil {

		fmt.Println(err.Error())

	}
}

func CreateTask(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	task := models.Task{}

	err := json.NewDecoder(req.Body).Decode(&task)

	if err != nil {
		fmt.Println(err.Error())
	}

	response, err = services.CreateTask(&task)

	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")

	if err := enc.Encode(response); err != nil {

		fmt.Println(err.Error())

	}
}

func UpdateTask(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	task := models.Task{}

	err := json.NewDecoder(req.Body).Decode(&task)

	if err != nil {
		fmt.Println(err.Error())
	}

	filter := bson.D{{"_id", task.ID}}
	update := bson.D{{"$set", task}}

	response, err = services.UpdateTask(filter, update)

	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")

	if err := enc.Encode(response); err != nil {

		fmt.Println(err.Error())

	}
}

func DeleteTask(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	vars := mux.Vars(req)

	response, err := services.DeleteTask(vars["id"])

	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")

	if err := enc.Encode(response); err != nil {

		fmt.Println(err.Error())

	}
}
