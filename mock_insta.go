package goinsta

import "net/http"

// MockInsta is a mock implementation of struct
type MockInsta struct {
	MockParseHTML func() error
	// MockGetImage    func(url, dest string) (bool, error)
	MockGetFileName func() (string, error)
	MockWriteFile   func(resp *http.Response, path string) error
	MockGet         func() error
}

// parseHTML mock
func (m *MockInsta) parseHTML() error {
	return m.MockParseHTML()
}

// GetImage mock
// func (m *MockInsta) GetImage(url, dest string) (bool, error) {
// 	return m.MockGetImage(url, dest)
// }

// getFileName mock
func (m *MockInsta) getFileName() (string, error) {
	return m.MockGetFileName()
}

// writeFile mock
func (m *MockInsta) writeFile(resp *http.Response, path string) error {
	return m.MockWriteFile(resp, path)
}

// get mock
func (m *MockInsta) get() error {
	return m.MockGet()
}
