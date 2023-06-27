package database

import (
	"taskmine/domain/entity"
	"taskmine/infra/model"

	"gorm.io/gorm"
)

type MysqlTaskRepository struct {
	DB *gorm.DB
}

func NewMysqlTaskRepository(DB *gorm.DB) *MysqlTaskRepository {
	return &MysqlTaskRepository{
		DB: DB,
	}
}

func (r *MysqlTaskRepository) Store(task *entity.Task) error {
	var taskModel = toTaskModel(task)
	return r.DB.Create(&taskModel).Error
}

func (r *MysqlTaskRepository) Update(task *entity.Task) error {
	var taskModel = toTaskModel(task)
	return r.DB.Save(&taskModel).Error
}

func (r *MysqlTaskRepository) Delete(task *entity.Task) error {
	var taskModel = toTaskModel(task)
	r.DB.First(&taskModel)
	return r.DB.Delete(&taskModel).Error
}

func (r *MysqlTaskRepository) FindAll() ([]*entity.Task, error) {
	var taskModels []*model.Task
	err := r.DB.Find(&taskModels).Error

	var tasks []*entity.Task
	for _, taskModel := range taskModels {
		task := toTaskEntity(taskModel)
		tasks = append(tasks, &task)
	}
	return tasks, err
}

func (r *MysqlTaskRepository) FindByID(id uint) (*entity.Task, error) {
	var task *model.Task
	err := r.DB.Find(&task, id).Error
	taskModel := toTaskEntity(task)
	return &taskModel, err
}
