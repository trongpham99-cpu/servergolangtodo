package services

import (
	"go.mongodb.org/mongo-driver/mongo"

	config "trongpham.dev/todo/helpers"

	models "trongpham.dev/todo/models"
)

func GetTasks(filter interface{}) (map[string]interface{}, error) {

	var tasks []*models.Task

	cur, err := config.TasksCollection.Find(config.Ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(config.Ctx) {
		var t models.Task
		err := cur.Decode(&t)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &t)
	}

	if err := cur.Err(); err != nil {
		return nil, nil
	}

	cur.Close(config.Ctx)

	if len(tasks) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	res := map[string]interface{}{

		"message": "Tasks fetched successfully",
		"status":  200,
		"data": map[string]interface{}{
			"result": tasks,
			"count":  len(tasks)},
	}

	return res, nil
}

func GetDetail(filter interface{}) (map[string]interface{}, error) {

	var task models.Task

	err := config.TasksCollection.FindOne(config.Ctx, filter).Decode(&task)

	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"message": "Task fetched successfully",
		"status":  200,
		"data": map[string]interface{}{
			"result": task,
		},
	}

	return res, nil
}

func CreateTask(task *models.Task) (map[string]interface{}, error) {

	result, err := config.TasksCollection.InsertOne(config.Ctx, task)

	if err != nil {
		return nil, err
	}
	newTask := result.InsertedID

	res := map[string]interface{}{
		"message": "Task created successfully",
		"status":  201,
		"data": map[string]interface{}{
			"taskID": newTask,
		},
	}

	return res, nil
}

func UpdateTask(filter interface{}, update interface{}) (map[string]interface{}, error) {

	result, err := config.TasksCollection.UpdateOne(config.Ctx, filter, update)

	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"message": "Task updated successfully",
		"status":  200,
		"data": map[string]interface{}{
			"result": result,
		},
	}

	return res, nil
}

func DeleteTask(filter interface{}) (map[string]interface{}, error) {

	result, err := config.TasksCollection.DeleteOne(config.Ctx, filter)

	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"message": "Task deleted successfully",
		"status":  200,
		"data": map[string]interface{}{
			"result": result,
		},
	}

	return res, nil
}
