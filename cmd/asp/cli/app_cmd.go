package cli

import (
	"context"

	asp "github.com/bots-house/app-store-parser"
	"github.com/urfave/cli/v2"
)

var appCMD = &cli.Command{
	Name: "app",
	Flags: []cli.Flag{
		idFlag,
		appIDFlag,
		langFlag,
		countryFlag,
		ratingsFlag,
	},
	Action: action(func(ctx context.Context, client *cli.Context, collector asp.Collector) (*asp.App, error) {
		return collector.App(ctx, asp.AppSpec{
			ID:      idFlag.Get(client),
			AppID:   appIDFlag.Get(client),
			Lang:    langFlag.Get(client),
			Country: countryFlag.Get(client),
			Ratings: ratingsFlag.Get(client),
		})
	}),
}
