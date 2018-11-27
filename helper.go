package goinsta

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Ashwin-Rajeev/goerror"
	"golang.org/x/net/html"
)

// image is a image used to get the response from the image url
func image(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, goerror.GetErrorInfo(err)
	}
	return resp, nil
}

// getFileName retrives the file name
func (a *information) getFileName() (string, error) {
	r := strings.NewReader(string(a.content))
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return "", goerror.New("error token")
		case html.StartTagToken:
			t := z.Token()
			switch t.Data {
			case "title":
				z.Next()
				t = z.Token()
				return t.Data, nil
			}
		}
	}
}

// WriteFile write the image to aparticular path as you specified
func (a *information) writeFile(resp *http.Response, path string) error {
	name, err := a.getFileName()
	if err != nil {
		log.Fatal(goerror.GetErrorInfo(err))
	}
	f := path + strings.TrimSpace(name) + ".jpg"
	file, err := os.Create(f)
	if err != nil {
		log.Fatal(goerror.GetErrorInfo(err))
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(goerror.GetErrorInfo(err))
	}
	defer resp.Body.Close()
	fmt.Printf("Image downloaded successfully...,\nplease check %s\n", f)
	return nil
}
