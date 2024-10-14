package asp

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/multierr"

	"github.com/bots-house/app-store-parser/shared"
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

func checkApps(t *testing.T, apps ...App) {
	for idx := range apps {
		assert.NoError(t, checkApp(&apps[idx]))
	}
}

func checkReview(review *Review) error {
	errs := make([]error, 0, 9)

	if review.ID == "" {
		errs = append(errs, fmt.Errorf("id missing"))
	}

	if review.Title == "" {
		errs = append(errs, fmt.Errorf("title missing"))
	}

	if review.Content == "" {
		errs = append(errs, fmt.Errorf("content missing"))
	}

	if review.Score == "" {
		errs = append(errs, fmt.Errorf("score missing"))
	}

	if review.Version == "" {
		errs = append(errs, fmt.Errorf("version missing"))
	}

	if review.UserName == "" {
		errs = append(errs, fmt.Errorf("user_name missing"))
	}

	if review.UserURL == "" {
		errs = append(errs, fmt.Errorf("user_url missing"))
	}

	return multierr.Combine(errs...)
}

func checkReviews(t *testing.T, reviews ...Review) {
	for idx := range reviews {
		assert.NoError(t, checkReview(&reviews[idx]))
	}
}

// TODO: check functionality some tests fail
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

	t.Run("AppWithRatings", func(t *testing.T) {
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

		checkApps(t, apps...)
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

		checkApps(t, apps...)
	})

	t.Run("ListWithCount", func(t *testing.T) {
		apps, err := collector.List(ctx, ListSpec{Count: 10})
		if !assert.NoError(t, err) {
			return
		}

		assert.Len(t, apps, 10)
		checkApps(t, apps...)
	})

	t.Run("ListWithCategory", func(t *testing.T) {
		apps, err := collector.List(ctx, ListSpec{Count: 10, Category: "games"})
		if !assert.NoError(t, err) {
			return
		}

		assert.Len(t, apps, 10)
		checkApps(t, apps...)

		for _, app := range apps {
			assert.Equal(t, "games", strings.ToLower(app.PrimaryGenre))
		}
	})

	t.Run("Search", func(t *testing.T) {
		apps, err := collector.Search(ctx, SearchSpec{
			Query: "netflix",
			Count: 1,
		})
		if !assert.NoError(t, err) {
			return
		}

		checkApps(t, apps...)

		assert.Len(t, apps, 1)
		assert.Equal(t, "com.netflix.Netflix", apps[0].AppID)
	})

	t.Run("Reviews", func(t *testing.T) {
		reviews, err := collector.Reviews(ctx, ReviewsSpec{ID: id})
		if !assert.NoError(t, err) {
			return
		}

		checkReviews(t, reviews...)
	})

	t.Run("Privacy", func(t *testing.T) {
		privacies, err := collector.Privacy(ctx, id)
		if !assert.NoError(t, err) {
			return
		}

		assert.NotEmpty(t, privacies)
	})

	t.Run("Suggest", func(t *testing.T) {
		result, err := collector.Suggest(ctx, SuggestSpec{Query: "p"})
		if !assert.NoError(t, err) {
			return
		}

		assert.NotEmpty(t, result)

		assert.True(t, shared.In("paypal", shared.Map(result, func(suggest Suggest) string {
			return suggest.Term
		})...))
	})
}

func Test_InAppPurchase(t *testing.T) {
	tests := []struct {
		id       int64
		expected bool
	}{
		{
			id:       525818839,
			expected: true,
		},

		{
			id:       479516143,
			expected: true,
		},

		{
			id: 311395856,
		},
	}

	collector := New()

	for _, test := range tests {
		app, err := collector.App(context.Background(), AppSpec{ID: test.id})
		if !assert.NoError(t, err) {
			return
		}

		assert.Equal(t, test.expected, app.InAppPurchase)
	}
}
