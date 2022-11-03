package task

import "fmt"

type TimeMinute int

type Task struct {
	Title         string
	Detail        string
	EstimatedTime TimeMinute
	ActualTime    TimeMinute
	IsDone        bool
}

func (i *Task) Do() error {
	if i.IsDone {
		return fmt.Errorf("既に達成されたタスクを達成することはできません")
	}

	i.IsDone = true

	return nil
}
