package issue

import "github.com/yuyohi/look-back-list/domain/user"

type IssueRepository interface {
	Store(issue Issue) error
	FindById(id IssueID) (Issue, error)
	FindByUserID(userID user.UserID) ([]Issue, error)
	DeleteById(id IssueID) error
}
