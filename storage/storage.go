package storage

import (
	"database/sql"
	"log"
	"math"

	"github.com/hacdan/issue-tracker-cli/types"
)

type Storage struct {
	store *sql.DB
}

func NewStorage() Storage {
	db, err := sql.Open("sqlite3", ".issues.db") //TODO: Change to customizable file name
	if err != nil {
		log.Fatal(err)
	}

	return Storage{
		store: db,
	}
}

func (s *Storage) AddIssue(issue types.Issue) (types.Issue, error) {
	issue.Id = s.getNextID()

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
