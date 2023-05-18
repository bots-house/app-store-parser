package scraper

type appListResult[T any] struct {
	Feed feed[T] `json:"feed"`
}

type feed[T any] struct {
	Entry []feedEntry[T] `json:"entry"`
}

type feedEntry[T any] struct {
	ID      feedEntryID     `json:"id"`
	Title   feedEntryLabel  `json:"title"`
	Content feedEntryLabel  `json:"content"`
	Author  feedEntryAuthor `json:"author"`
	Version feedEntryLabel  `json:"im:version"`
	Rating  feedEntryLabel  `json:"im:rating"`
	URL     T               `json:"link,omitempty"`
	Updated feedEntryLabel  `json:"updated"`
}

type feedEntryID struct {
	feedEntryLabel
	Attributes feedEntryIDAttributes `json:"attributes"`
}

type feedEntryIDAttributes struct {
	ID    string `json:"im:id"`
	AppID string `json:"im:bundleId"`
}

type feedEntryAuthor struct {
	feedEntryLabel
	Name feedEntryLabel `json:"name"`
	URI  feedEntryLabel `json:"uri"`
}

type feedEntryLink struct {
	Attributes feedEntryLinkAttributes `json:"attributes"`
}

type feedEntryLinkAttributes struct {
	Href string `json:"href"`
}

type feedEntryLabel struct {
	Label string `json:"label"`
}
