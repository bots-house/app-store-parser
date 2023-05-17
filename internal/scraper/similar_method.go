package scraper

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/bots-house/app-store-parser/shared"
)

func Similar(ctx context.Context, client shared.HTTPClient, spec shared.AppSpec) ([]shared.App, error) {
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}

	similarID, err := produceSimilarID(ctx, client, spec)
	if err != nil {
		return nil, err
	}

	result, err := request[[]string](ctx, client, requestSpec{
		url:    similarURL + similarID,
		params: spec,
		headers: http.Header{
			"X-Apple-Store-Front": []string{"143441,32"}, // todo: add country mapping
		},
		prepareResponse: parseSimilarResponse,
	})
	if err != nil {
		return nil, err
	}

	ids := shared.MapCheck(result, func(id string) (int64, bool) {
		entry, err := strconv.ParseInt(id, 10, strconv.IntSize)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("convert id from string")
			return 0, false
		}

		return entry, true
	})

	apps := shared.MapCheck(ids, func(id int64) (shared.App, bool) {
		spec := spec
		spec.ID = id

		app, err := App(ctx, client, spec)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Int64("app_id", id).Msg("parse app failed")
			return shared.App{}, false
		}

		return *app, true
	})

	if len(apps) == 0 {
		return nil, fmt.Errorf("similar apps not found")
	}

	return apps, nil
}

func parseSimilarResponse(body []byte) []byte {
	pattern := regexp.MustCompile(`customersAlsoBoughtApps":(.*?\])`)

	raw := pattern.FindSubmatch(body)

	if len(raw) > 1 {
		return raw[1]
	}

	return body
}

func produceSimilarID(ctx context.Context, client shared.HTTPClient, spec shared.AppSpec) (string, error) {
	if spec.ID != 0 {
		return strconv.FormatInt(spec.ID, 10), nil
	}

	app, err := App(ctx, client, spec)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(app.ID, 10), nil
}
