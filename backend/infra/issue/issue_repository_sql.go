package issue

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
		(?, ?, ?, ?, ?, ?, ?);
	`

	issueT := fromDomain(issue)

	_, err := r.db.Exec(sqlStr, issueT.IssueId, issueT.Title, issueT.Detail, issueT.EstimatedTime, issueT.ActualTime, issueT.IsDone, issueT.CreatedAt)

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

	row := r.db.QueryRow(sqlStr, id.Value())
	if err := row.Err(); err != nil {
		return issue_domain.Issue{}, err
	}

	var issueTable issueTable
	err := row.Scan(&issueTable.IssueId, &issueTable.Title, &issueTable.Detail, &issueTable.EstimatedTime, &issueTable.ActualTime, &issueTable.IsDone, &issueTable.CreatedAt)
	if err != nil {
		return issue_domain.Issue{}, err
	}
	issue := issueTable.toDomain()

	return *issue, nil
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