package main

import (
	"holo_check_CLI/model"
	"testing"
)

func TestPrintSearchResults(t *testing.T) {
	var cases = []struct {
		input    []*model.SearchResult
		expected string
	}{
		{
			[]*model.SearchResult{
				&model.SearchResult{
					Owner:      "湊あくあ",
					Title:      "テスト配信",
					LiveStatus: "upcoming",
				},
			},
			`--------------------------------------
 | owner name |   title    |  status  |
 |------------+------------+----------|
 |  湊あくあ  | テスト配信 | upcoming |
 |------------+------------+----------|`,
		},
	}

	for _, tt := range cases {
		resultString := PrintSearchResults(tt.input)
		if resultString != tt.expected {
			t.Fatalf("wrong format.\n got: \n %v\n expected: \n %v\n", resultString, tt.expected)
		}
	}
}
