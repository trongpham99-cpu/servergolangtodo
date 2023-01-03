package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	config "trongpham.dev/todo/helpers"

	models "trongpham.dev/todo/models"
)

func CreateProject(project *models.Project) (map[string]interface{}, error) {

	result, err := config.ProjectsCollection.InsertOne(config.Ctx, project)

	if err != nil {
		return nil, err
	}
	newProject := result.InsertedID

	res := map[string]interface{}{
		"message": "Project created successfully",
		"status":  201,
		"data": map[string]interface{}{
			"taskID": newProject,
		},
	}

	return res, nil
}

func GetProjects(filter interface{}) ([]*map[string]interface{}, error) {

	var projects []*map[string]interface{}

	lookupUser := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "userID"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "user"},
		}},
	}
	projectProject := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "userID", Value: 0},
			{Key: "user._id", Value: 0},
		}},
	}
	unwindUser := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$user"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}},
	}

	cur, err := config.ProjectsCollection.Aggregate(config.Ctx, mongo.Pipeline{lookupUser, projectProject, unwindUser})

	if err != nil {
		panic(err)
	}

	if err != nil {
		return nil, err
	}

	for cur.Next(config.Ctx) {
		var t map[string]interface{}
		err := cur.Decode(&t)

		if err != nil {
			return nil, err
		}

		projects = append(projects, &t)
	}

	if err := cur.Err(); err != nil {
		return nil, nil
	}

	cur.Close(config.Ctx)

	if len(projects) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return projects, nil
}
