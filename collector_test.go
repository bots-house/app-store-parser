package asp

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/multierr"
)

func checkApp(app *App) error {
	errs := make([]error, 0, 33)

	if app.ID == 0 {
		errs = append(errs, fmt.Errorf("id missing"))
	}

	if app.AppID == "" {
		errs = append(errs, fmt.Errorf("app id missing"))
	}

	if app.Title == "" {
		errs = append(errs, fmt.Errorf("title missing"))
	}

	if app.Description == "" {
		errs = append(errs, fmt.Errorf("description missing"))
	}

	if app.Developer == "" {
		errs = append(errs, fmt.Errorf("developer missing"))
	}

	if app.Released.IsZero() {
		errs = append(errs, fmt.Errorf("released date missing"))
	}

	if app.RequiredOsVersion == "" {
		errs = append(errs, fmt.Errorf("required os version missing"))
	}

	if len(app.Languages) == 0 {
		errs = append(errs, fmt.Errorf("languages missing"))
	}

	return multierr.Combine(errs...)
}

func Test_Collector(t *testing.T) {
	collector := New()
	id := int64(553834731)

	ctx := context.Background()

	t.Run("App", func(t *testing.T) {
		app, err := collector.App(ctx, AppSpec{ID: id})
		if !assert.NoError(t, err) {
			return
		}

		assert.NoError(t, checkApp(app))
	})

	t.Run("App by app id", func(t *testing.T) {
		// "com.facebook.Messenger"
		app, err := collector.App(ctx, AppSpec{AppID: "com.facebook.Messenger"})
		if !assert.NoError(t, err) {
			return
		}

		assert.NoError(t, checkApp(app))
	})

	t.Run("App with ratings", func(t *testing.T) {
		app, err := collector.App(ctx, AppSpec{ID: id, Ratings: true})
		if !assert.NoError(t, err) {
			return
		}

		assert.NoError(t, checkApp(app))
		assert.NotEmpty(t, app.Ratings.Total)
		assert.NotEmpty(t, app.Ratings.Histogram)
	})

	t.Run("Similar", func(t *testing.T) {
		apps, err := collector.Similar(ctx, AppSpec{ID: id})
		if !assert.NoError(t, err) {
			return
		}

		for _, app := range apps {
			assert.NoError(t, checkApp(&app))
		}
	})

	t.Run("Ratings", func(t *testing.T) {
		ratings, err := collector.Ratings(ctx, RatingsSpec{ID: id})
		if !assert.NoError(t, err) {
			return
		}

		assert.NotEmpty(t, ratings.Total)
		assert.NotEmpty(t, ratings.Histogram)
	})

	t.Run("Developer", func(t *testing.T) {
		apps, err := collector.Developer(ctx, DeveloperSpec{ID: 284882218})
		if !assert.NoError(t, err) {
			return
		}

		for _, app := range apps {
			assert.NoError(t, checkApp(&app))
		}
	})
}
