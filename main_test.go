package main

import (
	"holo_check_CLI/model"
	"testing"
)

func TestFormatSearchResults(t *testing.T) {
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
			`Target: 湊あくあ
    upcoming: テスト配信
`,
		},
		{
			[]*model.SearchResult{
				&model.SearchResult{
					Owner:      "湊あくあ",
					Title:      "テスト配信",
					LiveStatus: "upcoming",
				},
				&model.SearchResult{
					Owner:      "桐生ココ",
					Title:      "【#桐生ココ】地獄耐久に備える！まったり準備作業雑談【#とまらないARK】",
					LiveStatus: "live",
				},
			},
			`Target: 湊あくあ
    upcoming: テスト配信
Target: 桐生ココ
    live: 【#桐生ココ】地獄耐久に備える！まったり準備作業雑談【#とまらないARK】
`,
		},
	}

	for _, tt := range cases {
		resultString := FormatSearchResults(tt.input)
		if resultString != tt.expected {
			t.Fatalf("wrong format.\n got: \n %v\n expected: \n %v\n", resultString, tt.expected)
		}
	}
}
