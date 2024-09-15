package main

type JiraIssue struct {
	Key    string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
	} `json:"fields"`
}

func GetJiraIssue(ticketNumber, baseURL, username, apiToken string) (*JiraIssue, error) {
	return nil, nil
}
