package scraper

type appListResult struct {
	Feed feed `json:"feed"`
}

type feed struct {
	Entry []feedEntry `json:"entry"`
}

type feedEntry struct {
	ID feedEntryID `json:"id"`
}

type feedEntryID struct {
	Label      string              `json:"label"`
	Attributes feedEntryAttributes `json:"attributes"`
}

type feedEntryAttributes struct {
	ID    string `json:"im:id"`
	AppID string `json:"im:bundleId"`
}
