package cli

import (
	"context"

	asp "github.com/bots-house/app-store-parser"
	"github.com/urfave/cli/v2"
)

var similarCMD = &cli.Command{
	Name: "similar",
	Flags: []cli.Flag{
		idFlag,
		appIDFlag,
		langFlag,
		countryFlag,
		ratingsFlag,
		idsOnlyFlag,
	},
	Action: action(func(ctx context.Context, client *cli.Context, collector asp.Collector) ([]asp.App, error) {
		return collector.Similar(ctx, asp.AppSpec{
			ID:      idFlag.Get(client),
			AppID:   appIDFlag.Get(client),
			Lang:    langFlag.Get(client),
			Country: countryFlag.Get(client),
			Ratings: ratingsFlag.Get(client),
		})
	}),
}
