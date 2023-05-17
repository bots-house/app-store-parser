package scraper

import (
	"context"
	"fmt"

	"github.com/bots-house/app-store-parser/shared"
)

func App(ctx context.Context, client shared.HTTPClient, spec shared.AppSpec) (*shared.App, error) {
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}

	result, err := request[lookupResponse[shared.App]](ctx, client, requestSpec{
		url:    lookupURL,
		params: spec,
	})
	if err != nil {
		return nil, err
	}

	if result.ResultCount == 0 || len(result.Results) == 0 {
		return nil, fmt.Errorf("app not found")
	}

	app := result.Results[0]
	app.Sanitize()

	return &app, nil
}
