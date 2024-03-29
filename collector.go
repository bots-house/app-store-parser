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

func (collector collector) Developer(ctx context.Context, spec DeveloperSpec) ([]App, error) {
	apps, err := scraper.Developer(ctx, collector.client, shared.DeveloperSpec(spec))
	if err != nil {
		return nil, err
	}

	return newApps(apps...), nil
}

func (collector collector) List(ctx context.Context, spec ListSpec) ([]App, error) {
	apps, err := scraper.List(ctx, collector.client, shared.ListSpec(spec))
	if err != nil {
		return nil, err
	}

	return newApps(apps...), nil
}

func (collector collector) Search(ctx context.Context, spec SearchSpec) ([]App, error) {
	apps, err := scraper.Search(ctx, collector.client, shared.SearchSpec(spec))
	if err != nil {
		return nil, err
	}

	return newApps(apps...), nil
}

func (collector collector) Reviews(ctx context.Context, spec ReviewsSpec) ([]Review, error) {
	reviews, err := scraper.Reviews(ctx, collector.client, shared.ReviewsSpec(spec))
	if err != nil {
		return nil, err
	}

	return shared.Map(reviews, func(review shared.Review) Review {
		return Review(review)
	}), nil
}

func (collector collector) Privacy(ctx context.Context, id int64) ([]Privacy, error) {
	privacies, err := scraper.Privacy(ctx, collector.client, id)
	if err != nil {
		return nil, err
	}

	return shared.Map(privacies, func(entry shared.Privacy) Privacy {
		return Privacy(entry)
	}), nil
}

func (collector collector) Suggest(ctx context.Context, spec SuggestSpec) ([]Suggest, error) {
	result, err := scraper.Suggest(ctx, collector.client, shared.SuggestSpec(spec))
	if err != nil {
		return nil, err
	}

	return shared.Map(result, func(suggest shared.Suggest) Suggest {
		return Suggest(suggest)
	}), nil
}
