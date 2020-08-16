package model

type SearchResult struct {
	Owner      string
	Title      string
	LiveStatus string
}

func NewSearchResult(owner string) *SearchResult {
	return &SearchResult{Owner: owner}
}
