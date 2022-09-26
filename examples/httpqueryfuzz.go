package main

import (
	"flag"
	"fmt"
	"net/url"
	"xparams"
)

func main() {
	rawURL := flag.String("u", "http://127.0.0.1/index?a=1&b=2", "URL with query parameters")
	flag.Parse()
	u, _ := url.Parse(*rawURL)

	params := xparams.DefaultQueryFormExtractor.Extract(u.RawQuery, xparams.LocationQuery)
	for _, param := range params {
		query := param.Replace("test")
		u.RawQuery = query
		fmt.Println(u)
	}
}
