package scraper

import (
	"context"
	"fmt"

	"github.com/bots-house/app-store-parser/shared"
)

func Developer(ctx context.Context, client shared.HTTPClient, spec shared.DeveloperSpec) ([]shared.App, error) {
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}

	appsSpec := appsSpecFromDev(spec)

	apps, err := getApps(ctx, client, appsSpec)
	if err != nil {
		return nil, err
	}

	return apps, nil
}
