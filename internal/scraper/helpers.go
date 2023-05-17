package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bots-house/app-store-parser/shared"
)

type requestSpec struct {
	method          string
	url             string
	params          shared.QueryParams
	headers         http.Header
	prepareResponse func([]byte) []byte
}

func (spec *requestSpec) sanitize() {
	if spec.method == "" {
		spec.method = http.MethodGet
	}
}

func (spec *requestSpec) validate() error {
	spec.sanitize()

	if spec.url == "" {
		return fmt.Errorf("url required")
	}

	return nil
}

func request[T any](ctx context.Context, client shared.HTTPClient, spec requestSpec) (result T, _ error) {
	body, err := rawRequest(ctx, client, spec)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return result, fmt.Errorf("decode: %w", err)
	}

	return result, nil
}

func rawRequest(ctx context.Context, client shared.HTTPClient, spec requestSpec) ([]byte, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := spec.validate(); err != nil {
		return nil, fmt.Errorf("prepare request: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, spec.method, spec.url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("prepare request: %w", err)
	}

	if spec.params != nil {
		request.URL.RawQuery = spec.params.Encode()
	}

	if spec.headers != nil {
		request.Header = spec.headers
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	if spec.prepareResponse != nil {
		body = spec.prepareResponse(body)
	}

	return body, nil
}

func getCountryHeader(country, separator string) string {
	code, ok := countryMap[strings.ToLower(country)]
	if !ok {
		code = countryMap["us"]
	}

	if separator != "" {
		code += "," + separator
	}

	return code
}
