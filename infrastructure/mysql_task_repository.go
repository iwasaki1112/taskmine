package infrastructure

import (
	"taskmine/domain/entity"

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
	return r.DB.Create(task).Error
}

func (r *MysqlTaskRepository) Update(task *entity.Task) error {
	return r.DB.Save(task).Error
}

func (r *MysqlTaskRepository) Delete(task *entity.Task) error {
	r.DB.First(task)
	return r.DB.Delete(task).Error
}

func (r *MysqlTaskRepository) FindAll() ([]*entity.Task, error) {
	var tasks []*entity.Task
	err := r.DB.Find(&tasks).Error
	return tasks, err
}

func (r *MysqlTaskRepository) FindByID(id uint) (*entity.Task, error) {
	var task *entity.Task
	err := r.DB.Find(&task, id).Error
	return task, err
}
