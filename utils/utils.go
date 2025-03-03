package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hacdan/issue-tracker-cli/types"
)

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func StringToPriority(s string) types.PriorityLevel {
	s = strings.ToLower(s)

	if s == "low" {
		return types.LOW
	}
	if s == "medium" {
		return types.MEDIUM
	}
	if s == "high" {
		return types.HIGH
	}
	if s == "critcial" {
		return types.CRITICAL
	}

	return types.LOW
}

func SStatusToStatus(s string) types.IssueStatus {
	if strings.ToLower(s) == "open" {
		return types.OPEN
	}
	if strings.ToLower(s) == "inprogress" {
		return types.INPROGRESS
	}
	if strings.ToLower(s) == "closed" {
		return types.CLOSED
	}
	return types.OPEN
}

func IssuesToJson(issues []types.Issue) string {
	sIssues, err := json.Marshal(issues)
	if err != nil {
		fmt.Println("Error marshaling json: ", err)
	}
	return string(sIssues)
}

func IssuesToCSV(issues []types.Issue) string {
	header := []string{"id, title, description, priority, status, user"}
	data := [][]string{}
	data = append(data, header)
	for _, issue := range issues {
		row := []string{fmt.Sprintf("%d", issue.Id), issue.Title, issue.Description, issue.Priority.ToString(), issue.Status.ToString(), issue.User}
		data = append(data, row)
	}

	var result []string
	for _, row := range data {
		result = append(result, strings.Join(row, ","))
	}
	return strings.Join(result, strings.Join(result, "\n"))
}
