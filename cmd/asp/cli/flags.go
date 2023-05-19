package cli

import "github.com/urfave/cli/v2"

var idFlag = &cli.Int64Flag{
	Name: "id",
}

var appIDFlag = &cli.StringFlag{
	Name: "app-id",
}

var langFlag = &cli.StringFlag{
	Name: "lang",
}

var countryFlag = &cli.StringFlag{
	Name: "country",
}

var ratingsFlag = &cli.BoolFlag{
	Name:    "ratings",
	Aliases: []string{"r"},
}

var countFlag = &cli.IntFlag{
	Name: "count",
}

var collectionFlag = &cli.StringFlag{
	Name: "collection",
}

var categoryFlag = &cli.StringFlag{
	Name: "category",
}

var queryFlag = &cli.StringFlag{
	Name:    "query",
	Aliases: []string{"q"},
}

var pageFlag = &cli.IntFlag{
	Name: "page",
}

var idsOnlyFlag = &cli.BoolFlag{
	Name: "ids-only",
}

var sortFlag = &cli.StringFlag{
	Name: "sort",
}
