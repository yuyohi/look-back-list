package rdb

import (
	"database/sql"

	issue_domain "github.com/yuyohi/look-back-list/domain/issue"
)

type IssueRepositoryImpl struct {
	db *sql.DB
}

// コンストラクタ
func NewIssueRepositoryMySQL(db *sql.DB) *IssueRepositoryImpl {
	return &IssueRepositoryImpl{db: db}
}

func (r *IssueRepositoryImpl) Store(issue issue_domain.Issue) error {
	const sqlStr = `
		INSERT INTO issue (id, title, detail, estimated_time, actual_time, is_done, created_at) values
		(?, ?, ?, ?, ?, ?);
	`

	_, err := r.db.Exec(sqlStr, issue.IssueId, issue.Title, issue.Detail, issue.EstimatedTime, issue.ActualTime, issue.IsDone, issue.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *IssueRepositoryImpl) FindById(id issue_domain.IssueId) (issue_domain.Issue, error) {
	const sqlStr = `
		SELECT *
		FROM issue
		WHERE id = ?;
	`

	row := r.db.QueryRow(sqlStr, id)
	if err := row.Err(); err != nil {
		return issue_domain.Issue{}, err
	}

	var issue issue_domain.Issue
	err := row.Scan(&issue.IssueId, &issue.Title, &issue.Detail, &issue.EstimatedTime, &issue.ActualTime, &issue, issue.IsDone, &issue.CreatedAt)
	if err != nil {
		return issue_domain.Issue{}, err
	}

	return issue, nil
}

func (r *IssueRepositoryImpl) DeleteById(id issue_domain.IssueId) error {
	const sqlStr = `
	DELETE FROM issue 
	WHERE id = ?
	`

	_, err := r.db.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}
