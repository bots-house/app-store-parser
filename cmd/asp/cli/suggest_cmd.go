package cli

import (
	"context"

	asp "github.com/bots-house/app-store-parser"
	"github.com/urfave/cli/v2"
)

var suggestCMD = &cli.Command{
	Name: "suggest",
	Flags: []cli.Flag{
		queryFlag,
		countryFlag,
	},
	Action: action(func(ctx context.Context, client *cli.Context, collector asp.Collector) ([]asp.Suggest, error) {
		return collector.Suggest(ctx, asp.SuggestSpec{
			Query:   queryFlag.Get(client),
			Country: countryFlag.Get(client),
		})
	}),
}
