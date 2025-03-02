package types

import (
	"fmt"
)

type Issue struct {
	Id          int           `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Priority    PriorityLevel `json:"priority"`
	Status      IssueStatus   `json:"status"`
	User        string        `json:"user"` //TODO: Change to user struct later on
}

type IssueStatus string

const (
	OPEN       IssueStatus = "OPEN"
	INPROGRESS IssueStatus = "INPROGRESS"
	CLOSED     IssueStatus = "CLOSED"
)

type PriorityLevel int

const (
	LOW      PriorityLevel = 4
	MEDIUM   PriorityLevel = 3
	HIGH     PriorityLevel = 2
	CRITICAL PriorityLevel = 1
)

func (i *Issue) Print() {
	fmt.Println(i.Id, "\t", i.Title, "\t", i.Description, "\t", i.Status, "\t", i.User)
}

func (i IssueStatus) ToString() string {
	switch i {
	case OPEN:
		return "OPEN"
	case INPROGRESS:
		return "INPROGRESS"
	case CLOSED:
		return "CLOSED"
	default:
		return "OPEN"
	}
}

func (p PriorityLevel) ToString() string {
	switch p {
	case LOW:
		return "LOW"
	case MEDIUM:
		return "MEDIUM"
	case HIGH:
		return "HIGH"
	case CRITICAL:
		return "CRITICAL"
	default:
		return "LOW"
	}
}
