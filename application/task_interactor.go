package application

import (
	"log"
	"strconv"
	"taskmine/domain/entity"
	"taskmine/domain/repository"
)

type TaskInteractor struct {
	repository repository.TaskRepository
}

func NewTaskInteractor(repository repository.TaskRepository) *TaskInteractor {
	return &TaskInteractor{
		repository: repository,
	}
}

func (interactor TaskInteractor) CreateTask(input CreateTaskInput) (*entity.Task, error) {
	task := entity.Task{Title: input.Title, Description: input.Description, Status: entity.TODO}
	err := interactor.repository.Store(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (interactor TaskInteractor) UpdateTask(input UpdateTaskInput) (*entity.Task, error) {
	id, err := strconv.ParseUint(input.ID, 10, 64)
	if err != nil {
		log.Fatal("faild to convert string to uint")
	}

	task, err := interactor.repository.FindByID(uint(id))
	if err != nil {
		log.Fatalf("faild to get task that id is %s", input.ID)
	}

	task.Title = input.Title
	task.Description = input.Description
	task.Status = entity.TODO

	err = interactor.repository.Update(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}
