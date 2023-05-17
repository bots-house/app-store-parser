package scraper

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/bots-house/app-store-parser/shared"
)

const (
	lookupURL  = "https://itunes.apple.com/lookup"
	similarURL = "https://itunes.apple.com/us/app/app/id"
	ratingsURL = "https://itunes.apple.com/%s/customer-reviews/id%d"
)

type lookupResponse[T any] struct {
	ResultCount int `json:"resultCount"`
	Results     []T `json:"results"`
}

type appsSpec struct {
	ids     []int64
	appIDs  []string
	lang    string
	country string
	ratings bool
}

func (spec appsSpec) Encode() string {
	values := url.Values{
		"entity":  []string{"software"},
		"country": []string{spec.country},
	}

	idKey := "id"
	idValue := strings.Join(shared.Map(spec.ids, func(id int64) string {
		return strconv.FormatInt(id, 10)
	}), ",")

	if len(spec.ids) == 0 {
		idKey = "bundleId"
		idValue = strings.Join(spec.appIDs, ",")
	}

	values.Set(idKey, idValue)

	if spec.lang != "" {
		values.Set("lang", spec.lang)
	}

	return values.Encode()
}

func (spec *appsSpec) validate() error {
	if spec.country == "" {
		spec.country = "us"
	}

	if len(spec.ids) == 0 && len(spec.appIDs) == 0 {
		return fmt.Errorf("ids required")
	}

	return nil
}

func (spec appsSpec) applyIDs(ids ...int64) appsSpec {
	if len(ids) == 0 {
		return spec
	}

	spec.ids = append(spec.ids, ids...)

	return spec
}

func (spec appsSpec) applyAppIDs(ids ...string) appsSpec {
	if len(ids) == 0 {
		return spec
	}

	spec.appIDs = append(spec.appIDs, ids...)

	return spec
}

func newAppsSpec(spec shared.AppSpec) appsSpec {
	return appsSpec{
		lang:    spec.Lang,
		country: spec.Country,
		ratings: spec.Ratings,
	}
}
