package main

import (
	"reflect"
	"testing"
)

func TestExtractJiraIssueNumber(t *testing.T) {
	ts := []struct {
		Input          string
		ExpectedOutput []string
	}{
		{
			Input:          "abcesfe feat: ABC-1234: adds function to extract jira issue from commit message",
			ExpectedOutput: []string{"ABC-1234"},
		},
		{
			Input:          "sdfadsf fix: 1BCD-1234: solves NPE",
			ExpectedOutput: []string{"BCD-1234"},
		},
		{
			Input:          "asdfadf chore: updates sequence diagram ABC-345",
			ExpectedOutput: []string{"ABC-345"},
		},
		{
			Input:          "sdsddwq fix: solves NPEs ABC-345 and ABC-992",
			ExpectedOutput: []string{"ABC-345", "ABC-992"},
		},
		{
			Input:          "aasdfavt merge pull reqeust #12 from feature/abc-019-support-no-end-commit",
			ExpectedOutput: []string{"abc-019"},
		},
	}

	for _, test := range ts {
		o, err := extractJiraIssueNumber(test.Input)
		if err != nil {
			t.Fatalf("Failed to extract jira issue number")
		}

		if !reflect.DeepEqual(o, test.ExpectedOutput) {
			t.Fatalf("want %v, got %v", test.ExpectedOutput, o)
		}
	}
}

func TestGetCommitMessages(t *testing.T) {
	commitMessages, err := getCommitMessages("0a649cc62d77ae5f92b0ca8df25b5b51793bb0a7", "333510f3469bbb417c2b4ae001a4d9294eb6fc90")
	if err != nil {
		t.Fatalf("Failed to get commit messages")
	}

	expectedCommitMessages := []string{"333510f test: ABC-012: fixes broken tests", "6575921 feat: ABC-012: adds function to extract jira ticket from commit message", "23b72ee feat: ABC-011: setup go project"}
	if len(commitMessages) == 0 || !reflect.DeepEqual(commitMessages, expectedCommitMessages) {
		t.Fatalf("want %v, got %v", expectedCommitMessages, commitMessages)
	}
}
