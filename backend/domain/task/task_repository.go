package task

type TaskRepository interface {
	Store() error
	FindById(id TaskId) (error, Task)
	DeleteById(id TaskId) error
}