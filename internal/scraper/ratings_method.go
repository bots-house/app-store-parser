package scraper

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"

	"github.com/bots-house/app-store-parser/shared"
)

func Ratings(ctx context.Context, client shared.HTTPClient, spec shared.RatingsSpec) (shared.Ratings, error) {
	if err := spec.Validate(); err != nil {
		return shared.Ratings{}, fmt.Errorf("validation: %w", err)
	}

	body, err := rawRequest(ctx, client, requestSpec{
		url: fmt.Sprintf(ratingsURL, spec.Country, spec.ID),
		headers: http.Header{
			"X-Apple-Store-Front": []string{getCountryHeader(spec.Country, "12")},
		},
		params: url.Values{
			"displayable-kind": []string{"11"},
		},
	})
	if err != nil {
		return shared.Ratings{}, err
	}

	return parseRatings(body)
}

func parseRatings(body []byte) (shared.Ratings, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return shared.Ratings{}, fmt.Errorf("load html: %w", err)
	}

	countSelection := doc.Find(".rating-count").Text()
	countString := regexp.MustCompile(`\d+`).FindString(countSelection)

	count, err := strconv.ParseInt(countString, 10, strconv.IntSize)
	if err != nil {
		log.Error().Err(err).Msg("ratings not found")
	}

	ratings := doc.Find(".vote .total").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})

	lRatings := len(ratings)

	histogram := make(map[int]int64, lRatings)

	for idx, entry := range ratings {
		value, err := strconv.ParseInt(entry, 10, strconv.IntSize)
		if err != nil {
			log.Error().Err(err).Msg("histogram entry")
			continue
		}

		histogram[lRatings-idx] = value
	}

	return shared.Ratings{
		Total:     count,
		Histogram: histogram,
	}, nil
}
