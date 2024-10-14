package scraper

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
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

	apps = shared.Map(apps, func(app shared.App) shared.App {
		ok, err := foundInAppPurchase(ctx, client, app.ID)
		if err != nil {
			return app
		}

		app.InAppPurchase = ok

		return app
	})

	if len(apps) == 0 {
		return nil, fmt.Errorf("apps not found")
	}

	return apps, nil
}

func foundInAppPurchase(ctx context.Context, client shared.HTTPClient, appID int64) (bool, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(inAppPurchaseLink, appID), http.NoBody)
	if err != nil {
		return false, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return false, err
	}

	text := doc.Find(inAppPurchaseSelector).Text()

	return text != "", nil
}
