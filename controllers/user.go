package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "trongpham.dev/todo/models"
	services "trongpham.dev/todo/services"
)

func CreateUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	user := models.User{}

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		fmt.Println(err.Error())
	}

	response, err = services.CreateUser(&user)

	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")

	if err := enc.Encode(response); err != nil {

		fmt.Println(err.Error())

	}
}
