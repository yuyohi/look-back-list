package issue

import (
	"fmt"
	"time"

	"github.com/yuyohi/look-back-list/domain/user"
	"github.com/yuyohi/look-back-list/util"
)

type TimeMinute int

func NewTimeMinute(value int) (TimeMinute, error) {
	if value < 0 {
		return TimeMinute(0), fmt.Errorf("TimeMinuteは0以上である必要があります")
	}

	return TimeMinute(value), nil
}

type IssueID util.Identifier

func GenerateIssueID() IssueID {
	id := IssueID(util.IDGenerator.Generate())
	return id
}

func NewIssueID(idStr string) IssueID {
	id := IssueID(util.NewIdentifier(idStr))
	return id
}

func (i IssueID) Value() string {
	return util.Identifier(i).Value()
}

type Issue struct {
	IssueID       IssueID
	UserID        user.UserID
	Title         string
	Detail        string
	EstimatedTime TimeMinute
	ActualTime    TimeMinute
	IsDone        bool
	CreatedAt     time.Time
}

func NewIssue(title string, userID user.UserID, detail string, estimatedTimeInt int, createAt time.Time) (*Issue, error) {

	estimatedTime, err := NewTimeMinute(estimatedTimeInt)
	if err != nil {
		return nil, err
	}

	issue := Issue{
		IssueID:       GenerateIssueID(),
		UserID:        userID,
		Title:         title,
		Detail:        detail,
		EstimatedTime: estimatedTime,
		IsDone:        false,
		CreatedAt:     createAt,
	}

	return &issue, nil
}

func ReconstructIssue(idStr string, userIDStr string, title string, detail string, estimatedTime int, actualTime int, isDone bool, createdAt time.Time) *Issue {
	issueID := NewIssueID(idStr)
	userID := user.UserID(userIDStr)
	et, _ := NewTimeMinute(estimatedTime)
	at, _ := NewTimeMinute(actualTime)

	issue := Issue{
		IssueID:       issueID,
		UserID:        userID,
		Title:         title,
		Detail:        detail,
		EstimatedTime: et,
		ActualTime:    at,
		IsDone:        false,
		CreatedAt:     createdAt,
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
