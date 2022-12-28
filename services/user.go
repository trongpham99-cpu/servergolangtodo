package services

import (
	config "trongpham.dev/todo/helpers"

	models "trongpham.dev/todo/models"
)

func CreateUser(user *models.User) (map[string]interface{}, error) {

	result, err := config.UsersCollection.InsertOne(config.Ctx, user)

	if err != nil {
		return nil, err
	}
	newTask := result.InsertedID

	res := map[string]interface{}{
		"message": "User created successfully",
		"status":  201,
		"data": map[string]interface{}{
			"taskID": newTask,
		},
	}

	return res, nil
}
