package services

import (
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	config "trongpham.dev/todo/helpers"

	models "trongpham.dev/todo/models"
)

// page & limit
func GetTasks(filter interface{}) ([]*models.Task, error) {

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

	// res := tasks []*models.Task{

	// 	"message": "Tasks fetched successfully",
	// 	"status":  200,
	// 	"data": map[string]interface{}{
	// 		"result": tasks,
	// 		"count":  len(tasks)},
	// }

	return tasks, nil
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

func GetDetailLookUp(query url.Values) (map[string]interface{}, error) {

	// objId, _ := primitive.ObjectIDFromHex(query.Get("id"))

	// filter := bson.D{{Key: "_id", Value: objId}}

	// matchStage := bson.D{{Key: "$match", Value: filter}}

	lookupStage := bson.D{{
		"$lookup",
		bson.D{
			{"from", "users"},
			{"localField", "userID"},
			{"foreignField", "_id"},
			{"as", "user"}}}}

	showLoadedCursor, err := config.TasksCollection.Aggregate(config.Ctx,
		mongo.Pipeline{lookupStage})

	if err != nil {
		return nil, err
	}

	var tasks []models.Task

	if err = showLoadedCursor.All(config.Ctx, &tasks); err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"message": "Task fetched successfully",
		"status":  200,
		"data": map[string]interface{}{
			"result": tasks,
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

func DeleteTask(id string) (map[string]interface{}, error) {

	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"_id": objID,
	}
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
