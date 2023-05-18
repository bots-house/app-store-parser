package cli

import (
	"context"

	asp "github.com/bots-house/app-store-parser"
	"github.com/urfave/cli/v2"
)

var privacyCMD = &cli.Command{
	Name: "privacy",
	Flags: []cli.Flag{
		idFlag,
	},
	Action: action(func(ctx context.Context, client *cli.Context, collector asp.Collector) ([]asp.Privacy, error) {
		return collector.Privacy(ctx, idFlag.Get(client))
	}),
}
