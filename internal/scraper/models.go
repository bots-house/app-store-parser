package scraper

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/bots-house/app-store-parser/shared"
)

const (
	lookupURL         = "https://itunes.apple.com/lookup"
	similarURL        = "https://itunes.apple.com/us/app/app/id"
	ratingsURL        = "https://itunes.apple.com/%s/customer-reviews/id%d"
	listURL           = "http://ax.itunes.apple.com/WebObjects/MZStoreServices.woa/ws/RSS/%s/%s/limit=%d/json"
	searchURL         = "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search"
	reviewsURL        = "https://itunes.apple.com/%s/rss/customerreviews/page=%d/id=%d/sortby=%s/json"
	appTokenURL       = "https://apps.apple.com/us/app/id"
	appsPrivacyURL    = "https://amp-api.apps.apple.com/v1/catalog/US/apps/"
	suggestURL        = "https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints"
	collyScrapperLink = "https://apps.apple.com/us/app/id%d"

	collyScrapperSelector = ".app-header__list__item--in-app-purchase"
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

func (spec appsSpec) applyIDs(ids ...int64) appsSpec {
	spec.ids = shared.MapCheck(ids, func(id int64) (int64, bool) {
		return id, id != 0
	})

	return spec
}

func (spec appsSpec) applyAppIDs(ids ...string) appsSpec {
	spec.appIDs = shared.MapCheck(ids, func(id string) (string, bool) {
		return id, id != ""
	})

	return spec
}

func appsSpecFromApp(spec shared.AppSpec) appsSpec {
	return appsSpec{
		lang:    spec.Lang,
		country: spec.Country,
		ratings: spec.Ratings,
	}
}

func appsSpecFromDev(spec shared.DeveloperSpec) appsSpec {
	ids := make([]int64, 1)
	if spec.ID != 0 {
		ids[0] = spec.ID
	}

	return appsSpec{
		ids:     ids,
		country: spec.Country,
		lang:    spec.Lang,
	}
}

func appsSpecFromSearch(spec shared.SearchSpec) appsSpec {
	return appsSpec{
		lang:    spec.Lang,
		country: spec.Country,
	}
}
