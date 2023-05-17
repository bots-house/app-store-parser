package asp

import "context"

type Collector interface {
	App(context.Context, AppSpec) (*App, error)
	Similar(context.Context, AppSpec) ([]App, error)
	Ratings(context.Context, RatingsSpec) (Ratings, error)
	Developer(context.Context, DeveloperSpec) ([]App, error)
	List(context.Context, ListSpec) ([]App, error)
	Search(context.Context, SearchSpec) ([]App, error)
}
