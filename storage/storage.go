package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math"

	"github.com/hacdan/issue-tracker-cli/types"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	store *sql.DB
}

func NewStorage() Storage {
	db, err := sql.Open("sqlite3", ".issues.db") //TODO: Change to customizable file name
	if err != nil {
		log.Fatal(err)
	}

	CreateDB(db)

	return Storage{
		store: db,
	}
}

func CreateDB(db *sql.DB) {
	createIssueTableStatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS issues(id INTEGER, title TEXT, description TEXT, status INTEGER, user TEXT);")
	if err != nil {
		fmt.Println(err)
	}
	_, err = createIssueTableStatement.Exec()
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Storage) AddIssue(issue types.Issue) (types.Issue, error) {
	issue.Id = s.getNextID()
	insertStatement, err := s.store.Prepare("INSERT INTO issues(id, title, description, status, user) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return types.Issue{}, err
	}
	_, err = insertStatement.Exec(issue.Id, issue.Title, issue.Description, issue.Status, issue.User)
	if err != nil {
		return types.Issue{}, err
	}

	return issue, nil
}

func (s *Storage) SearchIssuesByText(str string) ([]types.Issue, error) {
	issues := []types.Issue{}
	str = "%" + str + "%"
	getStatement, err := s.store.Prepare(`SELECT id, title, description, status, user FROM issues WHERE(title LIKE ? OR description LIKE ?)`)
	if err != nil {
		return []types.Issue{}, err
	}

	rows, err := getStatement.Query(str, str)
	if err != nil {
		return []types.Issue{}, err
	}

	for rows.Next() {
		issue := new(types.Issue)
		err = rows.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Status, &issue.User)
		if err != nil {
			return issues, err
		}
		issues = append(issues, *issue)
	}
	return issues, nil
}

func (s *Storage) SearchIssuesByTitle(str string) ([]types.Issue, error) {
	str = "%" + str + "%"
	issues := []types.Issue{}
	getStatement, err := s.store.Prepare("SELECT id, title, description, status, user FROM issues WHERE title LIKE ?")
	if err != nil {
		return []types.Issue{}, err
	}

	rows, err := getStatement.Query(str)
	if err != nil {
		return []types.Issue{}, err
	}

	for rows.Next() {
		issue := new(types.Issue)
		err = rows.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Status, &issue.User)
		if err != nil {
			return issues, err
		}
		issues = append(issues, *issue)
	}
	return issues, nil
}

func (s *Storage) SearchIssuesByDescription(str string) ([]types.Issue, error) {
	str = "%" + str + "%"
	issues := []types.Issue{}
	getStatement, err := s.store.Prepare("SELECT id, title, description, status, user FROM issues WHERE(description LIKE ?)")
	if err != nil {
		return []types.Issue{}, err
	}

	rows, err := getStatement.Query(str)
	if err != nil {
		return []types.Issue{}, err
	}

	for rows.Next() {
		issue := new(types.Issue)
		err = rows.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Status, &issue.User)
		if err != nil {
			return issues, err
		}
		issues = append(issues, *issue)
	}
	return issues, nil
}
func (s *Storage) GetIssueByStatus(status string) ([]types.Issue, error) {
	issues := []types.Issue{}
	getStatement, err := s.store.Prepare("SELECT id, title, description, status, user FROM issues WHERE status LIKE ?")
	if err != nil {
		return issues, err
	}

	rows, err := getStatement.Query(status)
	if err != nil {
		return issues, err
	}

	for rows.Next() {
		issue := new(types.Issue)
		err = rows.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Status, &issue.User)
		if err != nil {
			return issues, err
		}
		issues = append(issues, *issue)
	}
	return issues, errors.New("No issues found with that status")
}

func (s *Storage) GetIssueByUser(user string) ([]types.Issue, error) {
	issues := []types.Issue{}
	getStatement, err := s.store.Prepare("SELECT id, title, description, status, user FROM issues WHERE user LIKE ?")
	if err != nil {
		return issues, err
	}

	rows, err := getStatement.Query(user)
	if err != nil {
		return issues, err
	}

	for rows.Next() {
		issue := new(types.Issue)
		err = rows.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Status, &issue.User)
		if err != nil {
			return issues, err
		}
		issues = append(issues, *issue)
	}
	return issues, errors.New("No issues found for that user")
}

func (s *Storage) GetIssue(id int) (types.Issue, error) {
	issue := new(types.Issue)

	getStatement, err := s.store.Prepare("SELECT id, title, description, status, user FROM issues WHERE id = ?")
	if err != nil {
		return types.Issue{}, nil
	}

	rows, err := getStatement.Query(id)
	if err != nil {
		return types.Issue{}, err
	}

	if rows.Next() {
		err = rows.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Status, &issue.User)
		if err != nil {
			return *issue, err
		}
		return *issue, nil
	}
	return *issue, nil
}

func (s Storage) GetIssues() ([]types.Issue, error) {
	issues := []types.Issue{}

	getStatement, err := s.store.Prepare("SELECT id, title, description, status, user FROM issues")
	if err != nil {
		return issues, err
	}

	rows, err := getStatement.Query()
	if err != nil {
		return issues, err
	}
	for rows.Next() {
		issue := new(types.Issue)
		err = rows.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Status, &issue.User)
		if err != nil {
			return issues, err
		}
		issues = append(issues, *issue)
	}

	return issues, errors.New("No issues found")
}

func (s *Storage) UpdateIssue(issue types.Issue) error {
	updateStatement, err := s.store.Prepare("UPDATE issues SET title = ?, description = ?, status = ?, user = ? WHERE id = ? ")
	if err != nil {
		return err
	}

	_, err = updateStatement.Exec(issue.Title, issue.Description, issue.Status, issue.User, issue.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteIssue(id int) error {
	deleteStatement, err := s.store.Prepare("DELETE FROM issues WHERE ID = ?")
	if err != nil {
		return err
	}

	_, err = deleteStatement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) getNextID() int {
	rows, err := s.store.Query("SELECT id FROM issues;")
	issueIds := []int{}
	if err == sql.ErrNoRows {
		return 1
	}
	if err != nil {
		return -1
	}

	for rows.Next() {
		issue := types.Issue{}
		err = rows.Scan(&issue.Id)
		issueIds = append(issueIds, issue.Id)
	}

	if len(issueIds) == 0 {
		return 1
	}

	max := math.MinInt
	for _, v := range issueIds {
		if v > max {
			max = v
		}
	}
	return max + 1
}
