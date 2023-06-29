package mysql

import (
	"taskmine/domain/entity"
	"taskmine/infra/database"
	"taskmine/infra/model"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(DB *gorm.DB) *TaskRepository {
	return &TaskRepository{
		DB: DB,
	}
}

func (r *TaskRepository) Store(task *entity.Task) error {
	var taskModel = database.ToTaskModel(task)
	return r.DB.Create(&taskModel).Error
}

func (r *TaskRepository) Update(task *entity.Task) error {
	var taskModel = database.ToTaskUpdateModel(task)
	return r.DB.Model(&taskModel).Where("id = ?", task.ID).Updates(taskModel).Error
}

func (r *TaskRepository) Delete(task *entity.Task) error {
	var taskModel = database.ToTaskModel(task)
	r.DB.First(&taskModel)
	return r.DB.Delete(&taskModel).Error
}

func (r *TaskRepository) FindAll() ([]*entity.Task, error) {
	var taskModels []*model.Task
	err := r.DB.Find(&taskModels).Error

	var tasks []*entity.Task
	for _, taskModel := range taskModels {
		task := database.ToTaskEntity(taskModel)
		tasks = append(tasks, &task)
	}
	return tasks, err
}

func (r *TaskRepository) FindByID(id uint) (*entity.Task, error) {
	var task *model.Task
	err := r.DB.Find(&task, id).Error
	taskModel := database.ToTaskEntity(task)
	return &taskModel, err
}
