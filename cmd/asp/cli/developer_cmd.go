package cli

import (
	"context"

	asp "github.com/bots-house/app-store-parser"
	"github.com/urfave/cli/v2"
)

var developerCMD = &cli.Command{
	Name:    "developer",
	Aliases: []string{"d", "dev"},
	Flags: []cli.Flag{
		idFlag,
		langFlag,
		countryFlag,
	},
	Action: action(func(ctx context.Context, client *cli.Context, collector asp.Collector) ([]asp.App, error) {
		return collector.Developer(ctx, asp.DeveloperSpec{
			ID:      idFlag.Get(client),
			Lang:    langFlag.Get(client),
			Country: countryFlag.Get(client),
		})
	}),
}
