package issue

import (
	"fmt"
	"time"

	"github.com/yuyohi/look-back-list/util"
)

type TimeMinute int
type IssueId util.Identifier

func NewIssueId() *IssueId {
	id := IssueId(util.IDGenerator.Generate())
	return &id
}

type Issue struct {
	IssueId       IssueId
	Title         string
	Detail        string
	EstimatedTime TimeMinute
	ActualTime    TimeMinute
	IsDone        bool
	CreatedAt     time.Time
}

func NewIssue(title string, detail string, estimatedTime TimeMinute, createAt time.Time) *Issue {

	issue := Issue{
		IssueId: *NewIssueId(),
		Title: title,
		Detail: detail,
		EstimatedTime: estimatedTime,
		IsDone: false,
		CreatedAt: time.Now(),
	}

	return &issue
}

func ReconstructIssue(issueId IssueId, title string, detail string, estimatedTime TimeMinute, createAt time.Time, isDone bool, createdAt time.Time) *Issue {
	issue := Issue{
		IssueId: issueId,
		Title: title,
		Detail: detail,
		EstimatedTime: estimatedTime,
		IsDone: false,
		CreatedAt: createdAt,
	}

	return &issue
}

func (t *Issue) Do() error {
	if t.IsDone {
		return fmt.Errorf("既に達成されたタスクを達成することはできません")
	}

	t.IsDone = true

	return nil
}
