package asp

import (
	"time"

	"github.com/bots-house/app-store-parser/shared"
)

type App struct {
	ID                    int64     `json:"id"`
	AppID                 string    `json:"app_id"`
	Title                 string    `json:"title"`
	URL                   string    `json:"url"`
	Description           string    `json:"description"`
	Genres                []string  `json:"genres"`
	GenreIDs              []string  `json:"genre_ids"`
	PrimaryGenre          string    `json:"primary_genre"`
	PrimaryGenreID        int64     `json:"primary_genre_id"`
	ContentRating         string    `json:"content_rating"`
	Languages             []string  `json:"languages"`
	Size                  string    `json:"size"`
	RequiredOsVersion     string    `json:"required_os_version"`
	Released              time.Time `json:"released"`
	Updated               time.Time `json:"updated"`
	ReleaseNotes          string    `json:"release_notes"`
	Version               string    `json:"version"`
	Price                 float64   `json:"price"`
	Currency              string    `json:"currency"`
	Free                  bool      `json:"free"`
	DeveloperID           int64     `json:"developer_id"`
	Developer             string    `json:"developer"`
	DeveloperURL          string    `json:"developer_url"`
	DeveloperWebsite      string    `json:"developer_website"`
	Score                 float64   `json:"score"`
	Reviews               int64     `json:"reviews"`
	CurrentVersionScore   float64   `json:"current_version_score"`
	CurrentVersionReviews int64     `json:"current_version_reviews"`
	Screenshots           []string  `json:"screenshots"`
	IpadScreenshots       []string  `json:"ipad_screenshots"`
	AppleTVScreenshots    []string  `json:"apple_tv_screenshots"`
	SupportedDevices      []string  `json:"supported_devices"`
	Icon                  string    `json:"icon"`
	Ratings               Ratings   `json:"ratings"`
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
