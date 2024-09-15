package main

import (
	"reflect"
	"testing"
)

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
