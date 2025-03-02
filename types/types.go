package types

import (
	"fmt"
)

type Issue struct {
	Id          int
	Title       string
	Description string
	Status      IssueStatus
	User        string //TODO: Change to user struct later on
}

type IssueStatus string

const (
	OPEN       IssueStatus = "OPEN"
	INPROGRESS IssueStatus = "INPROGRESS"
	CLOSED     IssueStatus = "CLOSED"
)

func (i *Issue) Print() {
	fmt.Println(i.Id, "\t", i.Title, "\t", i.Description, "\t", i.Status, "\t", i.User)
}
