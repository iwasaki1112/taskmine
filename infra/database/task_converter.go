package database

import (
	"taskmine/domain/entity"
	"taskmine/infra/model"

	"gorm.io/gorm"
)

func ToTaskModel(task *entity.Task) model.Task {
	return model.Task{
		Model: gorm.Model{
			ID:        task.ID,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
			DeletedAt: gorm.DeletedAt{},
		},
		Title:       "",
		Description: "",
		Status:      0,
	}
}

func ToTaskEntity(task *model.Task) entity.Task {
	return entity.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      entity.TaskStatus(task.Status),
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
