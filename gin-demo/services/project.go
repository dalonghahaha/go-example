package services

import (
	"gin-demo/models/comjia"
)

type ProjectService struct {
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) GetProject(id int64) *comjia.CjProject {
	project := new(comjia.CjProject)
	err := project.GeByID(id)
	if err != nil {
		return nil
	} else {
		return project
	}
}
