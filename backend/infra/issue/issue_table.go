package issue

import (
	"time"

	issue_domain "github.com/yuyohi/look-back-list/domain/issue"
)

type issueTable struct {
	IssueId       string
	Title         string
	Detail        string
	EstimatedTime int
	ActualTime    int
	IsDone        bool
	CreatedAt     time.Time
}

func (i issueTable) toDomain() *issue_domain.Issue {
	issue := issue_domain.ReconstructIssue(i.IssueId, i.Title, i.Detail, i.EstimatedTime, i.ActualTime, i.IsDone, i.CreatedAt)
	return issue
}

func fromDomain(i issue_domain.Issue) *issueTable {
	return &issueTable{
		IssueId: i.IssueId.Value(),
		Title: i.Title,
		Detail: i.Detail,
		EstimatedTime: int(i.EstimatedTime),
		ActualTime: int(i.ActualTime),
		IsDone: i.IsDone,
		CreatedAt: i.CreatedAt,
	}
}