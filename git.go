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

func extractJiraIssueNumbers(commitMessages []string) ([]string, []string, error) {
	issueNumbers := mapset.NewSet[string]()
	commitsWithoutIssueNumber := []string{}
	for _, commit := range commitMessages {
		issueNumbersInCommit, err := extractJiraIssueNumber(commit)
		if err != nil {
			return nil, nil, err
		}

		if issueNumbersInCommit == nil {
			commitsWithoutIssueNumber = append(commitsWithoutIssueNumber, commit)
			continue
		}

		for _, issueNumber := range issueNumbersInCommit {
			issueNumbers.Add(strings.ToUpper(issueNumber))
		}
	}
	issueNumbersSlice := issueNumbers.ToSlice()
	slices.Sort(issueNumbersSlice)
	return issueNumbersSlice, commitsWithoutIssueNumber, nil
}

func extractJiraIssueNumber(commitMessage string) ([]string, error) {
	re, err := regexp.Compile(`((?i)[A-Z]+-\d+)`)
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
