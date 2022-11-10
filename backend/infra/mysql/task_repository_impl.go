package mysql

import (
	"database/sql"

	"github.com/yuyohi/look-back-list/domain"
)

type TaskRepositorySQL struct {
	db *sql.DB
}

// コンストラクタ
func NewTaskRepositoryMySQL(db *sql.DB) *TaskRepositorySQL {
	return &TaskRepositorySQL{db: db}
}

func (r *TaskRepositorySQL) Store(task domain.Task) error {
	const sqlStr = `
		INSERT INTO tasks (id, title, detail, estimated_time, actual_time, is_done, created_at) values
		(?, ?, ?, ?, ?, ?);
	`

	_, err := r.db.Exec(sqlStr, task.TaskId, task.Title, task.Detail, task.EstimatedTime, task.ActualTime, task.IsDone, task.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepositorySQL) FindById(id domain.TaskId) (domain.Task, error) {
	const sqlStr = `
		SELECT *
		FROM tasks
		WHERE id = ?;
	`

	row := r.db.QueryRow(sqlStr, id)
	if err := row.Err(); err != nil {
		return domain.Task{}, err
	}

	var task domain.Task
	err := row.Scan(&task.TaskId, &task.Title, &task.Detail, &task.EstimatedTime, &task.ActualTime, &task, task.IsDone, &task.CreatedAt)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (r *TaskRepositorySQL) DeleteById(id domain.TaskId) error {
	const sqlStr = `
	DELETE FROM tasks 
	WHERE id = ?
	`

	_, err := r.db.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}
