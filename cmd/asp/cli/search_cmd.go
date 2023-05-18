package cli

import (
	"context"

	asp "github.com/bots-house/app-store-parser"
	"github.com/urfave/cli/v2"
)

var searchCMD = &cli.Command{
	Name:    "search",
	Aliases: []string{"s"},
	Flags: []cli.Flag{
		queryFlag,
		countFlag,
		pageFlag,
		langFlag,
		countryFlag,
		idsOnlyFlag,
	},
	Action: action(func(ctx context.Context, client *cli.Context, collector asp.Collector) ([]asp.App, error) {
		return collector.Search(ctx, asp.SearchSpec{
			Query:   queryFlag.Get(client),
			Count:   countFlag.Get(client),
			Page:    pageFlag.Get(client),
			Lang:    langFlag.Get(client),
			Country: countryFlag.Get(client),
			IDsOnly: idsOnlyFlag.Get(client),
		})
	}),
}
