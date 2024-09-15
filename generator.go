package main

import (
	"fmt"
	"os"
)

type JiraIssueInfo struct {
	Key    string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
	} `json:"fields"`
}

func pullDetails(startCommit, endCommit, baseURL, userName, apiToken string) ([]string, error) {
	ticketNumbers, err := extractJiraTicketNumbers(startCommit, endCommit)
	if err != nil {
		return nil, err
	}

	changes, err := getJiraIssues(ticketNumbers, baseURL, userName, apiToken)
	if err != nil {
		return nil, err
	}

	return changes, nil
}

func main() {
	baseUrl := os.Getenv("JIRA_BASE_URL")
	userName := os.Getenv("JIRA_USER_NAME")
	apiToken := os.Getenv("JIRA_API_TOKEN")
	if len(baseUrl) == 0 || len(userName) == 0 || len(apiToken) == 0 {
		fmt.Println("Required environment variables: $JIRA_BASE_URL, $JIRA_USER_NAME and $JIRA_API_TOKEN.")
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Required command arguments in order: start commit hash, end commit hash")
		return
	}

	changes, _ := pullDetails(os.Args[0], os.Args[1], baseUrl, userName, apiToken)
	for _, change := range changes {
		fmt.Println(change)
	}

}
