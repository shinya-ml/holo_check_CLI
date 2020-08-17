package model

import "google.golang.org/api/youtube/v3"

type SearchResult struct {
	Owner      string
	Title      string
	LiveStatus string
}

func NewSearchResult(owner string) *SearchResult {
	return &SearchResult{Owner: owner}
}

func BindSearchResults(name string, items []*youtube.SearchResult) []*SearchResult {
	searchResults := make([]*SearchResult, 0)
	for _, item := range items {
		sr := &SearchResult{
			Owner:      name,
			Title:      item.Snippet.Title,
			LiveStatus: item.Snippet.LiveBroadcastContent,
		}
		searchResults = append(searchResults, sr)
	}
	return searchResults
}
