package domain

type TaskRepository interface {
	Store(task Task) error
	FindById(id TaskId) (Task, error)
	DeleteById(id TaskId) error
}
