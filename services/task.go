package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	config "trongpham.dev/todo/helpers"

	models "trongpham.dev/todo/models"
)

// page & limit
// func GetTasks(filter interface{}) ([]*models.Task, error) {
func GetTasks(filter interface{}) ([]*map[string]interface{}, error) {

	var tasks []*map[string]interface{}

	// id, _ := primitive.ObjectIDFromHex("63b3d337a809e2cf0efc1efb")
	// projectMatch := bson.D{
	// 	{
	// 		Key: "$match", Value: bson.D{
	// 			{Key: "projectID", Value: id},
	// 		},
	// 	}}

	userLookup := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "userID"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "user"},
		}},
	}

	projectLookup := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "projects"},
			{Key: "localField", Value: "projectID"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "project"},
		}},
	}

	projectTask := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "userID", Value: 0},
			{Key: "user._id", Value: 0},
			{Key: "project._id", Value: 0},
			{Key: "project.userID", Value: 0},
			// {Key: "projectID", Value: 0},
		}},
	}

	unwind := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$user"},
			{Key: "path", Value: "$project"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}},
	}

	sort := bson.D{
		{Key: "$sort", Value: bson.D{
			{Key: "created_at", Value: -1},
		}},
	}

	cur, err := config.TasksCollection.Aggregate(
		config.Ctx,
		mongo.Pipeline{
			// projectMatch,
			userLookup,
			projectLookup,
			projectTask,
			unwind,
			sort,
			// facet,
		},
	)

	if err != nil {
		return nil, err
	}

	for cur.Next(config.Ctx) {
		var t map[string]interface{}
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
