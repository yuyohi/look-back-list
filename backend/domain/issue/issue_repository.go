package issue

type IssueRepository interface {
	Store(issue Issue) error
	FindById(id IssueId) (Issue, error)
	DeleteById(id IssueId) error
}
