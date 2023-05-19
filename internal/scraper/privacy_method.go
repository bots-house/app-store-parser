package scraper

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/bots-house/app-store-parser/shared"
)

func Privacy(ctx context.Context, client shared.HTTPClient, id int64) ([]shared.Privacy, error) {
	if id == 0 {
		return nil, fmt.Errorf("validation: id required")
	}

	body, err := rawRequest(ctx, client, requestSpec{
		url: appTokenURL + strconv.FormatInt(id, 10),
	})
	if err != nil {
		return nil, fmt.Errorf("token request: %w", err)
	}

	token, err := parseToken(body)
	if err != nil {
		return nil, err
	}

	return parsePrivacy(ctx, client, id, token)
}

func parseToken(body []byte) (string, error) {
	pattern := regexp.MustCompile(`token%22%3A%22([^%]+)%22%7D`)

	result := pattern.FindStringSubmatch(string(body))

	if len(result) < 2 {
		return "", fmt.Errorf("token not found")
	}

	return result[1], nil
}

func parsePrivacy(ctx context.Context, client shared.HTTPClient, id int64, token string) ([]shared.Privacy, error) {
	result, err := request[privacy](ctx, client, requestSpec{
		url: appsPrivacyURL + strconv.FormatInt(id, 10),
		headers: http.Header{
			"Origin":        []string{"https://apps.apple.com"},
			"Authorization": []string{"Bearer " + token},
		},
		params: url.Values{
			"platform": []string{"web"},
			"fields":   []string{"privacyDetails"},
			"l":        []string{"en-us"},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("app privacy not found: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("privacy not parsed")
	}

	return shared.Map(result.Data[0].Attributes.PrivacyDetails.PrivacyTypes, func(privacy privacyType) shared.Privacy {
		return shared.Privacy{
			Categories: shared.Map(privacy.DataCategories, func(entry dataCategory) shared.DataCategory {
				return shared.DataCategory{
					Category:   entry.DataCategory,
					Types:      append([]string{}, entry.DataTypes...),
					Identifier: entry.Identifier,
				}
			}),
			Description: privacy.Description,
			Identifier:  privacy.Identifier,
			Type:        privacy.PrivacyType,
			Purposes: shared.Map(privacy.Purposes, func(entry purpose) shared.Purpose {
				return shared.Purpose{
					Categories: shared.Map(entry.DataCategories, func(entry dataCategory) shared.DataCategory {
						return shared.DataCategory{
							Category:   entry.DataCategory,
							Types:      append([]string{}, entry.DataTypes...),
							Identifier: entry.Identifier,
						}
					}),
					Identifier: entry.Identifier,
					Purpose:    entry.Purpose,
				}
			}),
		}
	}), nil
}
