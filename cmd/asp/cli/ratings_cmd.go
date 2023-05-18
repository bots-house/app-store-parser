package cli

import (
	"context"

	asp "github.com/bots-house/app-store-parser"
	"github.com/urfave/cli/v2"
)

var ratingsCMD = &cli.Command{
	Name: "ratings",
	Flags: []cli.Flag{
		idFlag,
		countryFlag,
	},
	Action: action(func(ctx context.Context, client *cli.Context, collector asp.Collector) (asp.Ratings, error) {
		return collector.Ratings(ctx, asp.RatingsSpec{
			ID:      idFlag.Get(client),
			Country: countryFlag.Get(client),
		})
	}),
}
