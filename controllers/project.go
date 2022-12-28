package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
