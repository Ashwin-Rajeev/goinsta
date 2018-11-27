package goinsta

import "net/http"

// mockInsta is a mock implementation of struct
type mockInsta struct {
	MockParseHTML func() error
	MockGetFileName func() (string, error)
	MockWriteFile   func(resp *http.Response, path string) error
	MockGet         func() error
}

// parseHTML mock
func (m *mockInsta) parseHTML() error {
	return m.MockParseHTML()
}

// getFileName mock
func (m *mockInsta) getFileName() (string, error) {
	return m.MockGetFileName()
}

// writeFile mock
func (m *mockInsta) writeFile(resp *http.Response, path string) error {
	return m.MockWriteFile(resp, path)
}

// get mock
func (m *mockInsta) get() error {
	return m.MockGet()
}
