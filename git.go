package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func extractJiraTicketNumbers(startCommit, endCommit string) ([]string, error) {
	commits, err := getCommitMessages(startCommit, endCommit)
	if err != nil {
		return nil, err
	}

	ticketNumbers := mapset.NewSet[string]()
	for _, commit := range commits {
		ticketNumbersInCommit, err := extractJiraTicketNumber(commit)
		if err != nil {
			return nil, err
		}

		for _, ticketNumber := range ticketNumbersInCommit {
			ticketNumbers.Add(ticketNumber)
		}
	}
	ticketNumbersSlice := ticketNumbers.ToSlice()
	slices.Sort(ticketNumbersSlice)
	return ticketNumbersSlice, nil
}

func extractJiraTicketNumber(commitMessage string) ([]string, error) {
	re, err := regexp.Compile(`([A-Z]+-\d+)`)
	if err != nil {
		return nil, err
	}

	return re.FindAllString(commitMessage, -1), nil
}

func getCommitMessages(startCommit, endCommit string) ([]string, error) {
	var revisionRange = endCommit
	if startCommit != "" {
		revisionRange = fmt.Sprintf("%s..%s", startCommit, endCommit)
	}
	cmd := exec.Command("git", "log", "--oneline", revisionRange)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error running git log command: %s", out.String())
	}

	commitMessages := strings.Split(strings.TrimSpace(out.String()), "\n")
	return commitMessages, nil
}
