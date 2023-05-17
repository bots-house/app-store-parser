package asp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Collector(t *testing.T) {
	collector := New()

	ctx := context.Background()

	t.Run("App", func(t *testing.T) {
		app, err := collector.App(ctx, AppSpec{ID: 553834731})
		if !assert.NoError(t, err) {
			return
		}

		t.Log(app)
	})

	t.Run("Similar", func(t *testing.T) {
		apps, err := collector.Similar(ctx, AppSpec{ID: 553834731})
		if !assert.NoError(t, err) {
			return
		}

		t.Log(len(apps))
	})
}
