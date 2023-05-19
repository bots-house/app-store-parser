package asp

import (
	"time"

	"github.com/bots-house/app-store-parser/shared"
)

type App struct {
	ID                    int64     `json:"id"`
	AppID                 string    `json:"app_id,omitempty"`
	Title                 string    `json:"title,omitempty"`
	URL                   string    `json:"url,omitempty"`
	Description           string    `json:"description,omitempty"`
	Genres                []string  `json:"genres,omitempty"`
	GenreIDs              []string  `json:"genre_ids,omitempty"`
	PrimaryGenre          string    `json:"primary_genre,omitempty"`
	PrimaryGenreID        int64     `json:"primary_genre_id,omitempty"`
	ContentRating         string    `json:"content_rating,omitempty"`
	Languages             []string  `json:"languages,omitempty"`
	Size                  string    `json:"size,omitempty"`
	RequiredOsVersion     string    `json:"required_os_version,omitempty"`
	Released              time.Time `json:"released,omitempty"`
	Updated               time.Time `json:"updated,omitempty"`
	ReleaseNotes          string    `json:"release_notes,omitempty"`
	Version               string    `json:"version,omitempty"`
	Price                 float64   `json:"price,omitempty"`
	Currency              string    `json:"currency,omitempty"`
	Free                  bool      `json:"free,omitempty"`
	DeveloperID           int64     `json:"developer_id,omitempty"`
	Developer             string    `json:"developer,omitempty"`
	DeveloperURL          string    `json:"developer_url,omitempty"`
	DeveloperWebsite      string    `json:"developer_website,omitempty"`
	Score                 float64   `json:"score,omitempty"`
	Reviews               int64     `json:"reviews,omitempty"`
	CurrentVersionScore   float64   `json:"current_version_score,omitempty"`
	CurrentVersionReviews int64     `json:"current_version_reviews,omitempty"`
	Screenshots           []string  `json:"screenshots,omitempty"`
	IpadScreenshots       []string  `json:"ipad_screenshots,omitempty"`
	AppleTVScreenshots    []string  `json:"apple_tv_screenshots,omitempty"`
	SupportedDevices      []string  `json:"supported_devices,omitempty"`
	Icon                  string    `json:"icon,omitempty"`
	Ratings               Ratings   `json:"ratings,omitempty"`
}

func newApp(app *shared.App) *App {
	return &App{
		ID:                    app.ID,
		AppID:                 app.AppID,
		Title:                 app.Title,
		URL:                   app.URL,
		Description:           app.Description,
		Genres:                app.Genres,
		GenreIDs:              app.GenreIDs,
		PrimaryGenre:          app.PrimaryGenre,
		PrimaryGenreID:        app.PrimaryGenreID,
		ContentRating:         app.ContentRating,
		Languages:             app.Languages,
		Size:                  app.Size,
		RequiredOsVersion:     app.RequiredOsVersion,
		Released:              app.Released,
		Updated:               app.Updated,
		ReleaseNotes:          app.ReleaseNotes,
		Version:               app.Version,
		Price:                 app.Price,
		Currency:              app.Currency,
		Free:                  app.Free,
		DeveloperID:           app.DeveloperID,
		Developer:             app.Developer,
		DeveloperURL:          app.DeveloperURL,
		DeveloperWebsite:      app.DeveloperWebsite,
		Score:                 app.Score,
		Reviews:               app.Reviews,
		CurrentVersionScore:   app.CurrentVersionScore,
		CurrentVersionReviews: app.CurrentVersionReviews,
		Screenshots:           app.Screenshots,
		IpadScreenshots:       app.IpadScreenshots,
		AppleTVScreenshots:    app.AppleTVScreenshots,
		SupportedDevices:      app.SupportedDevices,
		Icon:                  app.Icon,
		Ratings:               Ratings(app.Ratings),
	}
}

func newApps(apps ...shared.App) []App {
	return shared.Map(apps, func(app shared.App) App {
		return *newApp(&app)
	})
}

type AppSpec shared.AppSpec

type Ratings shared.Ratings
type RatingsSpec shared.RatingsSpec
type DeveloperSpec shared.DeveloperSpec
type ListSpec shared.ListSpec
type SearchSpec shared.SearchSpec
type Review shared.Review
type ReviewsSpec shared.ReviewsSpec
type Privacy shared.Privacy
type Suggest shared.Suggest
type SuggestSpec shared.SuggestSpec
