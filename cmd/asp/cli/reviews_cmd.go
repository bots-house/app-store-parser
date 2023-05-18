package cli

import (
	"context"

	asp "github.com/bots-house/app-store-parser"
	"github.com/urfave/cli/v2"
)

var reviewsCMD = &cli.Command{
	Name: "reviews",
	Flags: []cli.Flag{
		idFlag,
		appIDFlag,
		sortFlag,
		pageFlag,
		countryFlag,
	},
	Action: action(func(ctx context.Context, client *cli.Context, collector asp.Collector) ([]asp.Review, error) {
		return collector.Reviews(ctx, asp.ReviewsSpec{
			ID:      idFlag.Get(client),
			AppID:   appIDFlag.Get(client),
			Sort:    sortFlag.Get(client),
			Page:    pageFlag.Get(client),
			Country: countryFlag.Get(client),
		})
	}),
}
