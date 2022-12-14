package issue_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	issue_repository "github.com/yuyohi/look-back-list/infra/issue"
)

var testDB *sql.DB

var tRepo *issue_repository.IssueRepositoryImpl

var dbConn = mysql.Config{
	DBName:               "sampledb",
	User:                 "docker",
	Passwd:               "docker",
	Addr:                 "127.0.0.1:3306",
	Net:                  "tcp",
	ParseTime:            true,
	AllowNativePasswords: true,
}

func connectDB() error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}
	dbConn.Loc = jst

	testDB, err = sql.Open("mysql", dbConn.FormatDSN())
	if err != nil {
		return err
	}

	return nil
}

func setupDB() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/createTable.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func cleanupDB() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/cleanupDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func setup() error {
	if err := connectDB(); err != nil {
		fmt.Println("connect", err)
		return err
	}

	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup", err)
		return err
	}
	if err := setupDB(); err != nil {
		fmt.Println("setup", err)
		return err
	}

	return nil
}

func teardown() {
	cleanupDB()
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	tRepo = issue_repository.NewIssueRepositoryMySQL(testDB)

	m.Run()

	teardown()
}
