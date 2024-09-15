package main

import (
	"reflect"
	"testing"
)

func TestExtractJiraTicketNumbers(t *testing.T) {
	ticketNumbers, err := extractJiraTicketNumbers("0a649cc62d77ae5f92b0ca8df25b5b51793bb0a7", "1336a69fbc390a81f5ba152da92dcb28ad4b3901")
	if err != nil {
		t.Fatalf("Failed to extract jira ticket number")
	}

	expectedTicketNumbers := []string{"ABC-011", "ABC-012", "ABC-013", "ABC-014", "ABC-015"}
	if len(ticketNumbers) == 0 || !reflect.DeepEqual(ticketNumbers, expectedTicketNumbers) {
		t.Fatalf("want %v, got %v", expectedTicketNumbers, ticketNumbers)
	}
}

func TestExtractJiraTicketNumber(t *testing.T) {
	ts := []struct {
		Input          string
		ExpectedOutput []string
	}{
		{
			Input:          "feat: ABC-1234: adds function to extract jira ticket from commit message",
			ExpectedOutput: []string{"ABC-1234"},
		},
		{
			Input:          "fix: 1BCD-1234: solves NPE",
			ExpectedOutput: []string{"BCD-1234"},
		},
		{
			Input:          "chore: updates sequence diagram ABC-345",
			ExpectedOutput: []string{"ABC-345"},
		},
		{
			Input:          "fix: solves NPEs ABC-345 and ABC-992",
			ExpectedOutput: []string{"ABC-345", "ABC-992"},
		},
	}

	for _, test := range ts {
		o, err := extractJiraTicketNumber(test.Input)
		if err != nil {
			t.Fatalf("Failed to extract jira ticket number")
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
