package scraper

const (
	lookupURL  = "https://itunes.apple.com/lookup"
	similarURL = "https://itunes.apple.com/us/app/app/id"
)

type lookupResponse[T any] struct {
	ResultCount int `json:"resultCount"`
	Results     []T `json:"results"`
}
