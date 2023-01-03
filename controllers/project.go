package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	models "trongpham.dev/todo/models"
	services "trongpham.dev/todo/services"
)

func CreateProject(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	project := models.Project{}

	err := json.NewDecoder(req.Body).Decode(&project)

	if err != nil {
		fmt.Println(err.Error())
	}

	response, err = services.CreateProject(&project)

	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")

	if err := enc.Encode(response); err != nil {

		fmt.Println(err.Error())

	}
}

func GetProjects(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := []*models.Project{}

	data := map[string]interface{}{}

	err := json.NewDecoder(req.Body).Decode(&data)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer req.Body.Close()

	filter := bson.D{{}}
	response, err = services.GetProjects(filter)

	if err != nil {

		// response = []*models.Taskinterface{}{"error": err.Error()}

	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")

	if err := enc.Encode(response); err != nil {

		fmt.Println(err.Error())

	}

}
