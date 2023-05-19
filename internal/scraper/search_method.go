package scraper

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bots-house/app-store-parser/shared"
)

func Search(ctx context.Context, client shared.HTTPClient, spec shared.SearchSpec) ([]shared.App, error) {
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}

	result, err := request[searchResponse](ctx, client, requestSpec{
		url:    searchURL,
		params: spec,
		headers: http.Header{
			"X-Apple-Store-Front": []string{shared.GetCountryHeader(spec.Country, "24 t:native")},
			"Accept-Language":     []string{spec.Lang},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("parse search data: %w", err)
	}

	if len(result.Bubbles) == 0 {
		return nil, fmt.Errorf("ids not found")
	}

	ids := result.Bubbles[0].paginated(spec.Count, spec.Page)
	if spec.IDsOnly {
		return shared.Map(ids, func(id int64) shared.App {
			return shared.App{ID: id}
		}), nil
	}

	return getApps(ctx, client, appsSpecFromSearch(spec).applyIDs(ids...))
}
