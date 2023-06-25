package repository

import "taskmine/domain/entity"

type TaskRepository interface {
	Store(task *entity.Task) error
	Update(task *entity.Task) error
	Delete(task *entity.Task) error
	FindAll() ([]*entity.Task, error)
	FindByID(id uint) (*entity.Task, error)
}
