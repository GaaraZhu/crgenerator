package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

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
