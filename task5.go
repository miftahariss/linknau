package main

import (
	"fmt"
	"net/http"
	"time"

	"linknau/task5"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	url := "https://www.example.com"

	status, err := task5.AsyncHTTPRequest(url, client)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Status: %s\n", status)
	}
}
