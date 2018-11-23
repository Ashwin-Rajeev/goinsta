# goinsta

`goinsta` is a a complete instagram library purely built in `GO`

To use the package import the package into your machine using the command:

> go get 

##GetImage
GetImage is used to download an image from instagram link
we need to provide image link, destination in which we
need to download the image and the file name
returnd a success flag and an error message

```func GetImage(url, dest string) (bool, error)```

which returns a success flag and a error meassage
if the operation works correctly, it will return a true and error will be nil
if the operation failed, it will return false and the corresponding error message

```_, err := goinsta.GetImage("https://www.instagram.com/dyfgsuyswyer47477834982","/home/user/")
	if err != nil {
		log.Fatal(err)
	}
```

After the successful execution you will receive a
success message aling with the path and name of the image downloaded


## Contributing

We welcome pull requests, bug fixes and issue reports. With that said, the bar for adding new symbols to this package is intentionally set high. Before proposing a change, please discuss your change by raising an issue.

## License

BSD-2-Clause