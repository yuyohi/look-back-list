package issue_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	issue_domain "github.com/yuyohi/look-back-list/domain/issue"
	"github.com/yuyohi/look-back-list/domain/user"
)

func TestStore(t *testing.T) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}
	time := time.Date(2022, time.December, 11, 11, 1, 1, 123000,jst)
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
