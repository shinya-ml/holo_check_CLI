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

func BindSearchResults(name string, status string, items []*youtube.SearchResult) []*SearchResult {
	searchResults := make([]*SearchResult, 0)
	switch status {
	case "all":
		for _, item := range items {
			sr := &SearchResult{
				Owner:      name,
				Title:      item.Snippet.Title,
				LiveStatus: item.Snippet.LiveBroadcastContent,
			}
			searchResults = append(searchResults, sr)
		}
	case "live":
		for _, item := range items {
			if item.Snippet.LiveBroadcastContent == "live" {
				sr := &SearchResult{
					Owner:      name,
					Title:      item.Snippet.Title,
					LiveStatus: item.Snippet.LiveBroadcastContent,
				}
				searchResults = append(searchResults, sr)
			}
		}
	case "upcoming":
		for _, item := range items {
			if item.Snippet.LiveBroadcastContent == "upcoming" {
				sr := &SearchResult{
					Owner:      name,
					Title:      item.Snippet.Title,
					LiveStatus: item.Snippet.LiveBroadcastContent,
				}
				searchResults = append(searchResults, sr)
			}
		}
	}
	return searchResults
}
