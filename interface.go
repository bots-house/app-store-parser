package asp

import "context"

type Collector interface {
	App(context.Context, AppSpec) (*App, error)
	Similar(context.Context, AppSpec) ([]App, error)
}
