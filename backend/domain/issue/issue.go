package issue

import (
	"fmt"
	"time"

	"github.com/yuyohi/look-back-list/util"
)

type TimeMinute int

func NewTimeMinute(value int) (TimeMinute, error) {
	if value < 0 {
		return TimeMinute(0), fmt.Errorf("TimeMinuteは0以上である必要があります")
	}

	return TimeMinute(value), nil
}

type IssueId util.Identifier

func GenerateIssueId() IssueId {
	id := IssueId(util.IDGenerator.Generate())
	return id
}

func NewIssueId(idStr string) IssueId {
	id := IssueId(util.NewIdentifier(idStr))
	return id
}

func (i IssueId) Value() string {
	return util.Identifier(i).Value()
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

func NewIssue(title string, detail string, estimatedTimeInt int, createAt time.Time) (*Issue, error) {

	estimatedTime, err := NewTimeMinute(estimatedTimeInt)
	if err != nil {
		return nil, err
	}

	issue := Issue{
		IssueId: GenerateIssueId(),
		Title: title,
		Detail: detail,
		EstimatedTime: estimatedTime,
		IsDone: false,
		CreatedAt: createAt,
	}

	return &issue, nil
}

func ReconstructIssue(idStr string, title string, detail string, estimatedTime int, actualTime int, isDone bool, createdAt time.Time) *Issue {
	issueId := NewIssueId(idStr)
	et, _ := NewTimeMinute(estimatedTime)
	at, _ := NewTimeMinute(actualTime)

	issue := Issue{
		IssueId: issueId,
		Title: title,
		Detail: detail,
		EstimatedTime: et,
		ActualTime: at,
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
