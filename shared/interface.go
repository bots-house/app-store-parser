package shared

import (
	"net/http"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type QueryParams interface {
	Encode() string
}
