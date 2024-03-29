package scraper

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/bots-house/app-store-parser/shared"
)

func App(ctx context.Context, client shared.HTTPClient, spec shared.AppSpec) (*shared.App, error) {
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}

	appsSpec := appsSpecFromApp(spec).applyIDs(spec.ID).applyAppIDs(spec.AppID)

	apps, err := getApps(ctx, client, appsSpec)
	if err != nil {
		return nil, err
	}

	app := &apps[0]

	if !spec.Ratings {
		return app, nil
	}

	ratings, err := Ratings(ctx, client, shared.RatingsSpec{ID: app.ID})
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("parse ratings")
		return app, nil
	}

	app.Ratings = ratings

	return app, nil
}

func getApps(ctx context.Context, client shared.HTTPClient, spec appsSpec) ([]shared.App, error) {
	result, err := request[lookupResponse[shared.App]](ctx, client, requestSpec{
		url:    lookupURL,
		params: spec,
	})
	if err != nil {
		return nil, err
	}

	if result.ResultCount == 0 || len(result.Results) == 0 {
		return nil, fmt.Errorf("apps not found")
	}

	apps := shared.Map(result.Results, func(app shared.App) shared.App {
		app.Sanitize()
		return app
	})

	apps = shared.Filter(apps, func(app shared.App) bool {
		return app.WrapperType == "software"
	})

	if len(apps) == 0 {
		return nil, fmt.Errorf("apps not found")
	}

	return apps, nil
}
