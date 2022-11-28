package issue_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	issue_domain "github.com/yuyohi/look-back-list/domain/issue"
	"github.com/yuyohi/look-back-list/domain/user"
)

func TestStore(t *testing.T) {
	jst := createJst(t)
	time := time.Date(2022, time.December, 11, 11, 1, 1, 123000, jst)
	userID := user.UserID("testUserID")
	expected, err := issue_domain.NewIssue("test", userID, "test detail", 100, time)
	if err != nil {
		t.Fatal(err)
	}

	err = tRepo.Store(*expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := tRepo.FindByID(expected.IssueID)
	if err != nil {
		t.Fatal(err)
	}

	opt := cmp.AllowUnexported(issue_domain.IssueID{})
	if diff := cmp.Diff(actual, *expected, opt); diff != "" {
		t.Errorf(diff)
	}

	t.Cleanup(func() {
		const sqlStr = `
		DELETE FROM issue
		`

		testDB.Exec(sqlStr)
	})

}

func TestDelete(t *testing.T) {
	jst := createJst(t)
	time := time.Date(2022, time.December, 11, 11, 1, 1, 123000, jst)
	userID := user.UserID("testUserID")
	newIssue, err := issue_domain.NewIssue("test", userID, "test detail", 100, time)
	if err != nil {
		t.Fatal(err)
	}

	err = tRepo.Store(*newIssue)
	if err != nil {
		t.Fatal(err)
	}

	tRepo.DeleteByID(newIssue.IssueID)

	// 直にSQL叩くtestの方が良いかも
	allIssue, err := tRepo.FindByUserID(userID)
	if err != nil {
		t.Fatal(err)
	}

	if len(allIssue) != 0 {
		t.Errorf("データが削除されていません: %v", allIssue)
	}
}

func createJst(t *testing.T) *time.Location {
	t.Helper()
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}

	return jst
}