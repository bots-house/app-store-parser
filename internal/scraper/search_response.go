package scraper

import (
	"strconv"

	"github.com/bots-house/app-store-parser/shared"
)

type searchResponse struct {
	Bubbles []searchBubbles `json:"bubbles"`
}

type searchBubbles struct {
	Results []searchResults `json:"results"`
}

type searchResults struct {
	ID string `json:"id"`
}

func (bubbles searchBubbles) paginated(count, page int) []int64 {
	page--

	start := count * page
	end := start + count
	if l := len(bubbles.Results); end > l {
		end = l
	}

	ids := bubbles.Results[start:end]

	return shared.Map(ids, func(entry searchResults) int64 {
		id, _ := strconv.ParseInt(entry.ID, 10, strconv.IntSize)
		return id
	})
}
