package issue

import (
	"database/sql"

	issue_domain "github.com/yuyohi/look-back-list/domain/issue"
	user_domain "github.com/yuyohi/look-back-list/domain/user"
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
		INSERT INTO issue (id, user_id, title, detail, estimated_time, actual_time, is_done, created_at) values
		(?, ?, ?, ?, ?, ?, ?, ?);
	`

	issueT := fromDomain(issue)

	_, err := r.db.Exec(sqlStr, issueT.IssueID, issue.UserID, issueT.Title, issueT.Detail, issueT.EstimatedTime, issueT.ActualTime, issueT.IsDone, issueT.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *IssueRepositoryImpl) FindByID(id issue_domain.IssueID) (issue_domain.Issue, error) {
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
	err := row.Scan(&issueTable.IssueID, &issueTable.userID, &issueTable.Title, &issueTable.Detail, &issueTable.EstimatedTime, &issueTable.ActualTime, &issueTable.IsDone, &issueTable.CreatedAt)
	if err != nil {
		return issue_domain.Issue{}, err
	}
	issue := issueTable.toDomain()

	return *issue, nil
}

func (r *IssueRepositoryImpl) FindByUserID(userID user_domain.UserID) ([]issue_domain.Issue, error) {
	const sqlStr = `
	SELECT *
	FROM issue
	WHERE user_id = ?;
	`

	rows, err := r.db.Query(sqlStr, string(userID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	issues := make([]issue_domain.Issue, 0)
	for rows.Next() {
		var issueTable issueTable
		err := rows.Scan(&issueTable.IssueID, &issueTable.userID, &issueTable.Title, &issueTable.Detail, &issueTable.EstimatedTime, &issueTable.ActualTime, &issueTable.IsDone, &issueTable.CreatedAt)

		if err != nil {
			return nil, err
		}

		issues = append(issues, *issueTable.toDomain())
	}

	return issues, nil
}

func (r *IssueRepositoryImpl) DeleteByID(id issue_domain.IssueID) error {
	const sqlStr = `
	DELETE FROM issue 
	WHERE id = ?
	`

	_, err := r.db.Exec(sqlStr, id.Value())
	if err != nil {
		return err
	}

	return nil
}
