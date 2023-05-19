package scraper

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/bots-house/app-store-parser/shared"
)

func Suggest(ctx context.Context, client shared.HTTPClient, spec shared.SuggestSpec) ([]shared.Suggest, error) {
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}

	body, err := rawRequest(ctx, client, requestSpec{
		url:    suggestURL,
		params: spec,
		headers: http.Header{
			"X-Apple-Store-Front": []string{shared.GetCountryHeader(spec.Country, "29")},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("suggest data not found: %w", err)
	}

	var data suggestXML

	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("suggest data parse: %w", err)
	}

	return shared.Map(data.Dict.Hints, func(entry suggestHint) shared.Suggest {
		return shared.Suggest{
			Term: entry.Values[0],
			URL:  entry.Values[1],
		}
	}), nil
}
