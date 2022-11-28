package issue

import (
	"time"

	issue_domain "github.com/yuyohi/look-back-list/domain/issue"
)

type issueTable struct {
	IssueID       string
	userID        string
	Title         string
	Detail        string
	EstimatedTime int
	ActualTime    int
	IsDone        bool
	CreatedAt     time.Time
}

func (i issueTable) toDomain() *issue_domain.Issue {
	issue := issue_domain.ReconstructIssue(i.IssueID, i.userID, i.Title, i.Detail, i.EstimatedTime, i.ActualTime, i.IsDone, i.CreatedAt)
	return issue
}

func fromDomain(i issue_domain.Issue) *issueTable {
	return &issueTable{
		IssueID: i.IssueID.Value(),
		userID: string(i.UserID),
		Title: i.Title,
		Detail: i.Detail,
		EstimatedTime: int(i.EstimatedTime),
		ActualTime: int(i.ActualTime),
		IsDone: i.IsDone,
		CreatedAt: i.CreatedAt,
	}
}