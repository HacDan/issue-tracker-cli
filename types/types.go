package types

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
