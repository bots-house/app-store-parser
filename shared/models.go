package shared

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
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

	if !In(spec.Country, Keys(countryMap)...) {
		return fmt.Errorf("invalid country")
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

	if !In(spec.Country, Keys(countryMap)...) {
		return fmt.Errorf("invalid country")
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

	if !In(spec.Country, Keys(countryMap)...) {
		return fmt.Errorf("invalid country")
	}

	return nil
}

type ListSpec struct {
	Count      int
	Country    string
	Collection string
	Category   string
}

func (spec *ListSpec) sanitize() {
	if spec.Collection == "" {
		spec.Collection = "TOP_FREE_IOS"
	}

	if spec.Count <= 0 {
		spec.Count = 50
	}

	if spec.Country == "" {
		spec.Country = "us"
	}

	spec.Country = strings.ToLower(spec.Country)
	spec.Category = strings.ToUpper(spec.Category)
	spec.Collection = strings.ToUpper(spec.Collection)
}

func (spec *ListSpec) Validate() error {
	spec.sanitize()

	if spec.Count > 200 {
		return fmt.Errorf("invalid count")
	}

	if !In(spec.Country, Keys(countryMap)...) {
		return fmt.Errorf("invalid country")
	}

	if spec.Category != "" && !In(spec.Category, Keys(categoryMap)...) {
		return fmt.Errorf("invalid category")
	}

	if !In(spec.Collection, Keys(collectionMap)...) {
		return fmt.Errorf("invalid collection")
	}

	return nil
}

func (spec ListSpec) Path(path string) string {
	collection := collectionMap[spec.Collection]
	category := ""
	if spec.Category != "" {
		categoryNumber := strconv.FormatInt(int64(categoryMap[spec.Category]), 10)
		category = "/genre=" + categoryNumber
	}

	return fmt.Sprintf(path, collection, category, spec.Count)
}

type SearchSpec struct {
	Query   string
	Lang    string
	Country string
	Count   int
	Page    int
	IDsOnly bool // if true apps will not parsed.
}

func (spec *SearchSpec) sanitize() {
	if spec.Count <= 0 {
		spec.Count = 50
	}

	if spec.Page <= 0 {
		spec.Page = 1
	}

	if spec.Lang == "" {
		spec.Lang = "en-us"
	}

	if spec.Country == "" {
		spec.Country = "us"
	}

	spec.Lang = strings.ToLower(spec.Lang)
	spec.Country = strings.ToLower(spec.Country)
}

func (spec *SearchSpec) Validate() error {
	spec.sanitize()

	if spec.Query == "" {
		return fmt.Errorf("query is required")
	}

	return nil
}

func (spec SearchSpec) Encode() string {
	values := url.Values{
		"clientApplication": []string{"Software"},
		"media":             []string{"software"},
		"term":              []string{spec.Query},
	}

	return values.Encode()
}

type Review struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UserName string    `json:"user_name"`
	UserURL  string    `json:"user_url"`
	Version  string    `json:"version"`
	Score    string    `json:"score"`
	URL      string    `json:"url"`
	Updated  time.Time `json:"updated"`
}

type ReviewsSpec struct {
	ID      int64
	AppID   string
	Page    int
	Sort    string
	Country string
}

func (spec *ReviewsSpec) sanitize() {
	if spec.Page <= 0 {
		spec.Page = 1
	}

	if spec.Sort == "" {
		spec.Sort = "RECENT"
	}

	if spec.Country == "" {
		spec.Country = "us"
	}

	spec.Country = strings.ToLower(spec.Country)
	spec.Sort = strings.ToUpper(spec.Sort)
}

func (spec *ReviewsSpec) Validate() error {
	spec.sanitize()

	if spec.ID == 0 && spec.AppID == "" {
		return fmt.Errorf("id or app_id required")
	}

	if !In(spec.Sort, Keys(sortMap)...) {
		return fmt.Errorf("invalid sort")
	}

	return nil
}

func (spec ReviewsSpec) Path(path string) string {
	return fmt.Sprintf(path, spec.Country, spec.Page, spec.ID, sortMap[spec.Sort])
}

type Privacy struct {
	Categories  []DataCategory `json:"categories"`
	Description string         `json:"description"`
	Identifier  string         `json:"identifier"`
	Type        string         `json:"type"`
	Purposes    []Purpose      `json:"purposes"`
}

type DataCategory struct {
	Category   string   `json:"category"`
	Types      []string `json:"types"`
	Identifier string   `json:"identifier"`
}

type Purpose struct {
	Categories []DataCategory `json:"categories"`
	Identifier string         `json:"identifier"`
	Purpose    string         `json:"purpose"`
}

type Suggest struct {
	Term string `json:"term"`
	URL  string `json:"url"`
}

type SuggestSpec struct {
	Query   string
	Country string
}

func (spec *SuggestSpec) sanitize() {
	if spec.Country == "" {
		spec.Country = "us"
	}

	spec.Country = strings.ToLower(spec.Country)
}

func (spec *SuggestSpec) Validate() error {
	spec.sanitize()

	if spec.Query == "" {
		return fmt.Errorf("query required")
	}

	return nil
}

func (spec SuggestSpec) Encode() string {
	values := url.Values{
		"clientApplication": []string{"Software"},
		"term":              []string{spec.Query},
	}

	return values.Encode()
}
