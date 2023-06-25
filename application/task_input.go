package application

type CreateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskInput struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DeleteTaskInput struct {
	ID string `json:"id"`
}
