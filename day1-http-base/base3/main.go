package main

import (
	"Gee/day1-http-base/base3/gee"
	"fmt"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.path = %v\n", r.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%v] = %v\n", k, v)
		}
	})

	r.Run(":9090")
}
