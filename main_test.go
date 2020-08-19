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
				{
					Owner:      "湊あくあ",
					Title:      "テスト配信",
					LiveStatus: "upcoming",
					URL:        "https://youtube.com/watch?v=hoge",
				},
			},
			`Target: 湊あくあ
    upcoming: テスト配信
    URL: https://youtube.com/watch?v=hoge
`,
		},
		{
			[]*model.SearchResult{
				{
					Owner:      "湊あくあ",
					Title:      "テスト配信",
					LiveStatus: "upcoming",
					URL:        "https://youtube.com/watch?v=hoge",
				},
				{
					Owner:      "桐生ココ",
					Title:      "【#桐生ココ】地獄耐久に備える！まったり準備作業雑談【#とまらないARK】",
					LiveStatus: "live",
					URL:        "https://youtube.com/watch?v=hoge",
				},
			},
			`Target: 湊あくあ
    upcoming: テスト配信
    URL: https://youtube.com/watch?v=hoge
Target: 桐生ココ
    live: 【#桐生ココ】地獄耐久に備える！まったり準備作業雑談【#とまらないARK】
    URL: https://youtube.com/watch?v=hoge
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
