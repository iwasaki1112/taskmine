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

func (repository *MysqlTaskRepository) Store(task *entity.Task) error {
	return repository.DB.Create(task).Error
}

func (repository *MysqlTaskRepository) Update(task *entity.Task) error {
	return repository.DB.Save(task).Error
}

func (repository *MysqlTaskRepository) Delete(task *entity.Task) error {
	repository.DB.First(task)
	return repository.DB.Delete(task).Error
}

func (repository *MysqlTaskRepository) FindAll() ([]*entity.Task, error) {
	var tasks []*entity.Task
	err := repository.DB.Find(&tasks).Error
	return tasks, err
}

func (repository *MysqlTaskRepository) FindByID(id uint) (*entity.Task, error) {
	var task *entity.Task
	err := repository.DB.Find(&task, id).Error
	return task, err
}
