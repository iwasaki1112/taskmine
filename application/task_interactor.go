package application

import (
	"log"
	"strconv"
	"taskmine/domain/entity"
	"taskmine/domain/repository"
	"taskmine/domain/service"
)

type TaskInteractor struct {
	repository repository.TaskRepository
	notifier   service.WebhookNotifier
}

func NewTaskInteractor(repository repository.TaskRepository, notifier service.WebhookNotifier) *TaskInteractor {
	return &TaskInteractor{
		repository: repository,
		notifier:   notifier,
	}
}

func (interactor TaskInteractor) CreateTask(input CreateTaskInput) (*entity.Task, error) {
	task := entity.Task{Title: input.Title, Description: input.Description, Status: entity.TODO}
	err := interactor.repository.Store(&task)
	if err != nil {
		return nil, err
	}

	message := task.Title + " has been created"
	err = interactor.notifier.Notify(message)
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

func (interactor TaskInteractor) DeleteTask(input DeleteTaskInput) (*entity.Task, error) {
	id, err := strconv.ParseUint(input.ID, 10, 64)
	if err != nil {
		log.Fatal("faild to convert string to uint")
	}

	task, err := interactor.repository.FindByID(uint(id))
	if err != nil {
		log.Fatalf("faild to get task that id is %s", input.ID)
	}

	err = interactor.repository.Delete(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}
