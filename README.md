# goinsta [![Build Status](https://travis-ci.com/Ashwin-Rajeev/goinsta.svg?branch=master)](https://travis-ci.com/Ashwin-Rajeev/goinsta) [![Go Report Card](https://goreportcard.com/badge/github.com/Ashwin-Rajeev/goinsta)](https://goreportcard.com/report/github.com/Ashwin-Rajeev/goinsta) [![GoDoc](https://godoc.org/github.com/Ashwin-Rajeev/goinsta?status.svg)](https://godoc.org/github.com/Ashwin-Rajeev/goinsta)

`goinsta` is a a complete instagram library purely built in `GO`

To use the package import the package into your machine using the command:

> go get github.com/Ashwin-Rajeev/goinsta

## NewInfo
NewInfo returns a new  info pointer
By using this pointer you can access
the functions implemented by Info

``` func NewInfo(url, dest string) *Info```

## GetImage
GetImage is used to download an image from instagram link
we need to provide image link, destination in which we
need to download the image and the file name
returnd a success flag and an error message

```func GetImage(url, dest string) (bool, error)```

which returns a success flag and a error meassage
if the operation works correctly, it will return a true and error will be nil
if the operation failed, it will return false and the corresponding error message

```
info := goinsta.NewInfo("https://www.instagram.com/dyfgsuyswyer47477834982","/home/user/")

 _,err:= info.GetImage()
		if err != nil {
			log.Fatal(err)
		}
		
```

After the successful execution you will receive a
success message along with the path and name of the image downloaded


## Contributing

We welcome pull requests, bug fixes and issue reports. With that said, the bar for adding new symbols to this package is intentionally set high. Before proposing a change, please discuss your change by raising an issue.

## License

BSD-2-Clause
