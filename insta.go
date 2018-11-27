// goinsta is a a complete instagram library

// NewInfo returns a new  info pointer
// By using this pointer you can access
// the functions implemented by Info

// func NewInfo(url, dest string) *Info

// GetImage is used to download an image from instagram link
// we need to provide image link, destination in which we
// need to download the image and the file name
// returned a success flag and an error message

// func GetImage(url, dest string) (bool, error)

// which returns a success flag and a error meassage
// if the operation works correctly, it will return a true and error will be nil
// if the operation failed, it will return false and the corresponding error message

// _, err := goinsta.GetImage("https://www.instagram.com/dyfgsuyswyer47477834982","/home/user/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// After the successful execution you will receive a
// success message aling with the path and name of the image downloaded

package goinsta

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/Ashwin-Rajeev/goerror"

	"golang.org/x/net/html"
)

// Info struct stores the image information
type Info struct {
	Crawler
}
type information struct {
	content []byte
	url     string
	dest    string
}

// ParseHTML is a function used to parse the url response
func (i *information) parseHTML() error {
	doc, err := html.Parse(strings.NewReader(string(i.content)))
	if err != nil {
		return err
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			for _, a := range n.Attr {
				if a.Key == "content" {
					matched, err := regexp.MatchString(
						"https://.*..jpg",
						a.Val,
					)
					if err != nil {
						log.Fatal(goerror.GetErrorInfo(err))
					}
					if matched {
						img, err := image(a.Val)
						if err != nil {
							log.Fatal(goerror.GetErrorInfo(err))
						}
						err = i.writeFile(
							img,
							i.dest,
						)
						if err != nil {
							log.Fatal(goerror.GetErrorInfo(err))
						}
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nil
}

func (i *information) get() error {
	resp, err := http.Get(i.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	i.content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// GetImage is used to download an image from instagram link
// we need to provide image link, destination in which we
// need to download the image and the file name
// returned a success flag and an error message
func (i *Info) GetImage() (bool, error) {
	err := i.get()
	if err != nil {
		return false, goerror.GetErrorInfo(err)
	}
	err = i.parseHTML()
	if err != nil {
		return false, goerror.GetErrorInfo(err)
	}
	return true, nil
}

// NewInfo returns a new  info pointer
// By using this pointer you can access
// the functions implemented by Info
func NewInfo(url, dest string) *Info {
	return &Info{
		Crawler: &information{
			url:  url,
			dest: dest,
		},
	}
}
