package model

import (
	"testing"

	"google.golang.org/api/youtube/v3"
)

func TestBindSearchReults(t *testing.T) {
	var cases = []struct {
		items    []*youtube.SearchResult
		name     string
		status   string
		expected []*SearchResult
	}{
		{
			items: []*youtube.SearchResult{
				{
					Snippet: &youtube.SearchResultSnippet{
						LiveBroadcastContent: "upcoming",
						Title:                "aqua test",
					},
				},
			},
			name:   "湊あくあ",
			status: "upcoming",
			expected: []*SearchResult{
				{
					Owner:      "湊あくあ",
					Title:      "aqua test",
					LiveStatus: "upcoming",
				},
			},
		},
		{
			items: []*youtube.SearchResult{
				{
					Snippet: &youtube.SearchResultSnippet{
						LiveBroadcastContent: "upcoming",
						Title:                "aqua upcoming",
					},
				},
				{
					Snippet: &youtube.SearchResultSnippet{
						LiveBroadcastContent: "live",
						Title:                "aqua live",
					},
				},
			},
			name:   "湊あくあ",
			status: "upcoming",
			expected: []*SearchResult{
				{
					Owner:      "湊あくあ",
					Title:      "aqua upcoming",
					LiveStatus: "upcoming",
				},
			},
		},
		{
			items: []*youtube.SearchResult{
				{
					Snippet: &youtube.SearchResultSnippet{
						LiveBroadcastContent: "live",
						Title:                "aqua live1",
					},
				},
				{
					Snippet: &youtube.SearchResultSnippet{
						LiveBroadcastContent: "live",
						Title:                "aqua live2",
					},
				},
			},
			name:   "湊あくあ",
			status: "live",
			expected: []*SearchResult{
				{
					Owner:      "湊あくあ",
					Title:      "aqua live1",
					LiveStatus: "live",
				},
				{
					Owner:      "湊あくあ",
					Title:      "aqua live2",
					LiveStatus: "live",
				},
			},
		},
		{
			items: []*youtube.SearchResult{
				{
					Snippet: &youtube.SearchResultSnippet{
						LiveBroadcastContent: "live",
						Title:                "aqua live1",
					},
				},
				{
					Snippet: &youtube.SearchResultSnippet{
						LiveBroadcastContent: "live",
						Title:                "aqua live2",
					},
				},
			},
			name:     "湊あくあ",
			status:   "upcoming",
			expected: []*SearchResult{},
		},
	}

	for _, tt := range cases {
		results := BindSearchResults(tt.name, tt.status, tt.items)
		if len(results) != len(tt.expected) {
			t.Fatalf("length of results is wrong. got = %d, want = %d", len(results), len(tt.expected))
		}
		for idx, result := range results {
			switch {
			case result.Owner != tt.expected[idx].Owner:
				t.Fatalf("result.Owner is wrong. got = %s, want = %s", result.Owner, tt.expected[idx].Owner)
			}
		}
	}
}
