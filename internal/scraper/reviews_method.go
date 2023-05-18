package scraper

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/bots-house/app-store-parser/shared"
)

func Reviews(ctx context.Context, client shared.HTTPClient, spec shared.ReviewsSpec) ([]shared.Review, error) {
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}

	id, err := produceReviewsID(ctx, client, spec)
	if err != nil {
		return nil, err
	}

	spec.ID = id

	result, err := request[appListResult[feedEntryLink]](ctx, client, requestSpec{
		url: spec.Path(reviewsURL),
	})
	if err != nil {
		return nil, fmt.Errorf("reviews data not found: %w", err)
	}

	if len(result.Feed.Entry) == 0 {
		return nil, fmt.Errorf("reviews data not found")
	}

	return shared.Map(result.Feed.Entry, func(entry feedEntry[feedEntryLink]) shared.Review {
		t, err := time.Parse(time.RFC3339, entry.Updated.Label)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("parse updated time")
		}
		return shared.Review{
			ID:       entry.ID.Label,
			Title:    entry.Title.Label,
			Content:  entry.Content.Label,
			UserName: entry.Author.Name.Label,
			UserURL:  entry.Author.URI.Label,
			Version:  entry.Version.Label,
			Score:    entry.Rating.Label,
			URL:      entry.URL.Attributes.Href,
			Updated:  t,
		}
	}), nil
}

func produceReviewsID(ctx context.Context, client shared.HTTPClient, spec shared.ReviewsSpec) (int64, error) {
	if spec.ID != 0 {
		return spec.ID, nil
	}

	app, err := App(ctx, client, shared.AppSpec{AppID: spec.AppID, Country: spec.Country})
	if err != nil {
		return 0, err
	}

	return app.ID, nil
}
