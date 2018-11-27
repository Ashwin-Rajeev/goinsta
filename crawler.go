package goinsta

import "net/http"

// Crawler is a abstract interface
type Crawler interface {
	// insta functions
	get() error
	parseHTML() error

	// helper functions
	getFileName() (string, error)
	writeFile(resp *http.Response, path string) error
}
