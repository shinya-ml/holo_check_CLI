package model

import "google.golang.org/api/youtube/v3"

type SearchResult struct {
	Owner      string
	Title      string
	LiveStatus string
	URL        string
}

func NewSearchResult(owner string, item *youtube.SearchResult) *SearchResult {
	return &SearchResult{
		Owner:      owner,
		Title:      item.Snippet.Title,
		LiveStatus: item.Snippet.LiveBroadcastContent,
		URL:        "https://youtube.com/watch?v=" + item.Id.VideoId,
	}
}

func BindSearchResults(name string, status string, items []*youtube.SearchResult) []*SearchResult {
	searchResults := make([]*SearchResult, 0)
	switch status {
	case "all":
		for _, item := range items {
			sr := NewSearchResult(name, item)
			searchResults = append(searchResults, sr)
		}
	case "live":
		for _, item := range items {
			if item.Snippet.LiveBroadcastContent == "live" {
				sr := NewSearchResult(name, item)
				searchResults = append(searchResults, sr)
			}
		}
	case "upcoming":
		for _, item := range items {
			if item.Snippet.LiveBroadcastContent == "upcoming" {
				sr := NewSearchResult(name, item)
				searchResults = append(searchResults, sr)
			}
		}
	}
	return searchResults
}
