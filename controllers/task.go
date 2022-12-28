package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	models "trongpham.dev/todo/models"
	services "trongpham.dev/todo/services"
)

func GetTasks(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	data := map[string]interface{}{}

	err := json.NewDecoder(req.Body).Decode(&data)

	if err != nil {
		fmt.Println(err.Error())
	}

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
	id := query.Get("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objId}}

	response, err := services.GetDetail(filter)

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

	response, err = services.UpdateTask(filter, &task)

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

	task := models.Task{}

	err := json.NewDecoder(req.Body).Decode(&task)

	if err != nil {
		fmt.Println(err.Error())
	}

	response, err = services.DeleteTask(&task)

	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")

	if err := enc.Encode(response); err != nil {

		fmt.Println(err.Error())

	}
}
