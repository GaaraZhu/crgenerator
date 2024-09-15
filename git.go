package main

import (
	"regexp"
)

func extractJiraTicketNumber(commitMessage string) ([]string, error) {
	re, err := regexp.Compile(`([A-Z]+-\d+)`)
	if err != nil {
		return nil, err
	}

	return re.FindAllString(commitMessage, -1), nil
}
