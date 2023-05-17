package asp

import (
	"context"
	"net/http"

	"github.com/bots-house/app-store-parser/internal/scraper"
	"github.com/bots-house/app-store-parser/shared"
)

type CollectorOption func(*collector)

func WithClient(client shared.HTTPClient) CollectorOption {
	return func(c *collector) {
		c.client = client
	}
}

var _ Collector = &collector{}

type collector struct {
	client shared.HTTPClient
}

func New(opts ...CollectorOption) Collector {
	collector := &collector{
		client: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(collector)
	}

	return collector
}

func (collector collector) App(ctx context.Context, spec AppSpec) (*App, error) {
	app, err := scraper.App(ctx, collector.client, shared.AppSpec(spec))
	if err != nil {
		return nil, err
	}

	return newApp(app), nil
}

func (collector collector) Similar(ctx context.Context, spec AppSpec) ([]App, error) {
	apps, err := scraper.Similar(ctx, collector.client, shared.AppSpec(spec))
	if err != nil {
		return nil, err
	}

	return newApps(apps...), nil
}

func (collector collector) Ratings(ctx context.Context, spec RatingsSpec) (Ratings, error) {
	ratings, err := scraper.Ratings(ctx, collector.client, shared.RatingsSpec(spec))
	if err != nil {
		return Ratings{}, err
	}

	return Ratings(ratings), nil
}
