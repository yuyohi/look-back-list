package issue

type IssueRepository interface {
	Store(issue Issue) error
	FindById(id IssueID) (Issue, error)
	DeleteById(id IssueID) error
}
