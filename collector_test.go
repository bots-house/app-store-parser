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

	ctx := context.Background()

	t.Run("App", func(t *testing.T) {
		app, err := collector.App(ctx, AppSpec{ID: 553834731})
		if !assert.NoError(t, err) {
			return
		}

		assert.NoError(t, checkApp(app))
	})

	t.Run("Similar", func(t *testing.T) {
		apps, err := collector.Similar(ctx, AppSpec{ID: 553834731})
		if !assert.NoError(t, err) {
			return
		}

		for _, app := range apps {
			assert.NoError(t, checkApp(&app))
		}
	})
}
