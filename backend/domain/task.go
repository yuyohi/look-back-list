package domain

import (
	"fmt"
	"time"
)

type TimeMinute int
type TaskId string

type Task struct {
	TaskId        TaskId
	Title         string
	Detail        string
	EstimatedTime TimeMinute
	ActualTime    TimeMinute
	IsDone        bool
	CreatedAt     time.Time
}

func (t *Task) Do() error {
	if t.IsDone {
		return fmt.Errorf("既に達成されたタスクを達成することはできません")
	}

	t.IsDone = true

	return nil
}