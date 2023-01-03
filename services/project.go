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

func GetProjects(filter interface{}) ([]*models.Project, error) {

	var projects []*models.Project

	// cur, err := config.ProjectsCollection.Find(config.Ctx, filter)
	lookupUser := bson.D{
		{"$lookup", bson.D{
			{"from", "users"},
			{"localField", "userID"},
			{"foreignField", "_id"},
			{"as", "user"},
		}},
		// {
		// 	"$unwind", bson.D{
		// 		{"path", "$user"},
		// 		{"preserveNullAndEmptyArrays", true},
		// 	},
		// },
	}
	cur, err := config.ProjectsCollection.Aggregate(config.Ctx, mongo.Pipeline{lookupUser})

	if err != nil {
		return nil, err
	}

	for cur.Next(config.Ctx) {
		var t models.Project
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
