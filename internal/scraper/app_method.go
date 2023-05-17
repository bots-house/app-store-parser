package scraper

import (
	"context"
	"fmt"

	"github.com/bots-house/app-store-parser/shared"
)

func App(ctx context.Context, client shared.HTTPClient, spec shared.AppSpec) (*shared.App, error) {
	appsSpec := newAppsSpec(spec).applyIDs(spec.ID).applyAppIDs(spec.AppID)

	apps, err := getApps(ctx, client, appsSpec)
	if err != nil {
		return nil, err
	}

	return &apps[0], nil
}

func getApps(ctx context.Context, client shared.HTTPClient, spec appsSpec) ([]shared.App, error) {
	if err := spec.validate(); err != nil {
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

	apps := shared.Map(result.Results, func(app shared.App) shared.App {
		app.Sanitize()
		return app
	})

	apps = shared.Filter(apps, func(app shared.App) bool {
		return app.WrapperType == "software"
	})

	return apps, nil
}
