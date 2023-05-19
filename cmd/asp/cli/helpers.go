package cli

import (
	"context"
	"encoding/json"
	"fmt"

	asp "github.com/bots-house/app-store-parser"
	"github.com/urfave/cli/v2"
)

func action[T any](do func(ctx context.Context, client *cli.Context, collector asp.Collector) (T, error)) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		collector := asp.New()

		result, err := do(ctx.Context, ctx, collector)
		if err != nil {
			return err
		}

		data, err := json.Marshal(result)
		if err != nil {
			return fmt.Errorf("encode: %w", err)
		}

		if _, err := ctx.App.Writer.Write(data); err != nil {
			return fmt.Errorf("write: %w", err)
		}

		return nil
	}
}
