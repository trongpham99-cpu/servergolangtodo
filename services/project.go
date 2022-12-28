package services

import (
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
