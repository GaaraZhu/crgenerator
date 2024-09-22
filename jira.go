package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type JiraIssue struct {
	Key     string
	Summary string
	Link    string
}

func (issue JiraIssue) String() string {
	return fmt.Sprintf("%s: %s %s", issue.Key, issue.Summary, issue.Link)
}

func getJiraIssues(issueNumbers []string, baseURL, userName, apiToken string) ([]string, []string, error) {
	var issues []string
	var issuesNotFound []string
	for _, issueNumber := range issueNumbers {
		issue, err := getJiraIssue(issueNumber, baseURL, userName, apiToken)
		if err != nil {
			fmt.Println(err)
			issuesNotFound = append(issuesNotFound, issueNumber)
			continue
		}
		issues = append(issues, issue.String())
	}
	return issues, issuesNotFound, nil
}

type JiraIssueDTO struct {
	Key    string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
	} `json:"fields"`
}

func getJiraIssue(issueNumber, baseURL, username, apiToken string) (*JiraIssue, error) {
	url := fmt.Sprintf("%s/rest/api/2/issue/%s", strings.TrimSuffix(baseURL, "/"), issueNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP Error: %s - %s", resp.Status, body)
	}

	var issueDto JiraIssueDTO
	err = json.NewDecoder(resp.Body).Decode(&issueDto)
	if err != nil {
		return nil, err
	}

	jiraIssue := JiraIssue{
		Key:     issueNumber,
		Summary: issueDto.Fields.Summary,
		Link:    fmt.Sprintf("%s/browse/%s", baseURL, issueNumber),
	}

	return &jiraIssue, nil
}
