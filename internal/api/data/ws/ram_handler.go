// Package http websocket handlers
package http

import (
	"fmt"
	"net/http"
)

func RAM(url string) http.Response {
	var res http.Response
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("failed to fetch data")
		}
	}()

	return res
}
