package shared

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type App struct {
	ID                    int64     `json:"trackId"`
	AppID                 string    `json:"bundleId"`
	Title                 string    `json:"trackName"`
	URL                   string    `json:"trackViewUrl"`
	Description           string    `json:"description"`
	Genres                []string  `json:"genres"`
	GenreIDs              []string  `json:"genreIds"`
	PrimaryGenre          string    `json:"primaryGenreName"`
	PrimaryGenreID        int64     `json:"primaryGenreId"`
	ContentRating         string    `json:"contentAdvisoryRating"`
	Languages             []string  `json:"languageCodesISO2A"`
	Size                  string    `json:"fileSizeBytes"`
	RequiredOsVersion     string    `json:"minimumOsVersion"`
	Released              time.Time `json:"releaseDate"`
	Updated               time.Time `json:"currentVersionReleaseDate"`
	ReleaseNotes          string    `json:"releaseNotes"`
	Version               string    `json:"version"`
	Price                 float64   `json:"price"`
	Currency              string    `json:"currency"`
	Free                  bool      `json:"free"`
	DeveloperID           int64     `json:"artistId"`
	Developer             string    `json:"artistName"`
	DeveloperURL          string    `json:"artistViewUrl"`
	DeveloperWebsite      string    `json:"sellerUrl"`
	Score                 float64   `json:"averageUserRating"`
	Reviews               int64     `json:"userRatingCount"`
	CurrentVersionScore   float64   `json:"averageUserRatingForCurrentVersion"`
	CurrentVersionReviews int64     `json:"userRatingCountForCurrentVersion"`
	Screenshots           []string  `json:"screenshotUrls"`
	IpadScreenshots       []string  `json:"ipadScreenshotUrls"`
	AppleTVScreenshots    []string  `json:"appletvScreenshotUrls"`
	SupportedDevices      []string  `json:"supportedDevices"`
	// Icon
	Icon          string
	ArtworkURL512 string `json:"artworkUrl512"`
	ArtworkURL100 string `json:"artworkUrl100"`
	ArtworkURL60  string `json:"artworkUrl60"`

	// Tech fields
	WrapperType string  `json:"wrapperType"`
	Ratings     Ratings `json:"-"`
}

func (app *App) Sanitize() {
	icon := app.ArtworkURL512
	if icon == "" && app.ArtworkURL100 != "" {
		icon = app.ArtworkURL100
	}

	if icon == "" && app.ArtworkURL60 != "" {
		icon = app.ArtworkURL60
	}

	app.Icon = icon
	app.Free = app.Price == 0

	if app.Updated.IsZero() {
		app.Updated = app.Released
	}
}

type AppSpec struct {
	ID      int64
	AppID   string
	Lang    string
	Country string
	Ratings bool
}

func (spec *AppSpec) sanitize() {
	if spec.Country == "" {
		spec.Country = "us"
	}
}

func (spec *AppSpec) Validate() error {
	spec.sanitize()

	if spec.ID == 0 && spec.AppID == "" {
		return fmt.Errorf("id or app_id required")
	}

	return nil
}

func (spec AppSpec) Encode() string {
	values := url.Values{
		"entity":  []string{"software"},
		"country": []string{spec.Country},
	}

	idKey := "id"
	idValue := strconv.FormatInt(spec.ID, 10)

	if spec.ID == 0 {
		idKey = "bundleId"
		idValue = spec.AppID
	}

	values.Set(idKey, idValue)

	if spec.Lang != "" {
		values.Set("lang", spec.Lang)
	}

	return values.Encode()
}

type RatingsSpec struct {
	ID      int64
	Lang    string
	Country string
}

func (spec *RatingsSpec) sanitize() {
	if spec.Country == "" {
		spec.Country = "us"
	}
}

func (spec *RatingsSpec) Validate() error {
	spec.sanitize()

	if spec.ID == 0 {
		return fmt.Errorf("id required")
	}

	return nil
}

type Ratings struct {
	Total     int64
	Histogram map[int]int64
}

type DeveloperSpec struct {
	ID      int64
	Lang    string
	Country string
}

func (spec DeveloperSpec) Encode() string {
	values := url.Values{
		"entity":  []string{"software"},
		"country": []string{spec.Country},
	}

	values.Set("id", strconv.FormatInt(spec.ID, 10))

	if spec.Lang != "" {
		values.Set("lang", spec.Lang)
	}

	return values.Encode()
}

func (spec *DeveloperSpec) Validate() error {
	if spec.Country == "" {
		spec.Country = "us"
	}

	if spec.ID == 0 {
		return fmt.Errorf("ids required")
	}

	return nil
}

type ListSpec struct {}