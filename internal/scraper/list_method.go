package scraper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/bots-house/app-store-parser/shared"
)

func List(ctx context.Context, client shared.HTTPClient, spec shared.ListSpec) ([]shared.App, error) {
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}

	result, err := request[appListResult](ctx, client, requestSpec{
		url: spec.Path(listURL),
	})
	if err != nil {
		return nil, fmt.Errorf("app ids not found: %w", err)
	}

	appsSpec, err := parseListResult(result)
	if err != nil {
		return nil, err
	}

	return getApps(ctx, client, appsSpec)
}

func parseListResult(result appListResult) (appsSpec, error) {
	ids := shared.MapCheck(result.Feed.Entry, func(entry feedEntry) (int64, bool) {
		id, err := strconv.ParseInt(entry.ID.Attributes.ID, 10, strconv.IntSize)
		if err != nil {
			log.Error().Err(err).Msg("id not found")
			return 0, false
		}

		return id, true
	})

	if len(ids) == 0 {
		return appsSpec{}, fmt.Errorf("apps data not found")
	}

	appsSpec := appsSpec{
		ids: append([]int64{}, ids...),
	}

	return appsSpec, nil
}
