package cli

import (
	"context"

	asp "github.com/bots-house/app-store-parser"
	"github.com/urfave/cli/v2"
)

var listCMD = &cli.Command{
	Name: "list",
	Flags: []cli.Flag{
		countFlag,
		countryFlag,
		collectionFlag,
		categoryFlag,
	},
	Action: action(func(ctx context.Context, client *cli.Context, collector asp.Collector) ([]asp.App, error) {
		return collector.List(ctx, asp.ListSpec{
			Count:      countFlag.Get(client),
			Country:    countryFlag.Get(client),
			Collection: collectionFlag.Get(client),
			Category:   categoryFlag.Get(client),
		})
	}),
}
