package task5

import (
	"fmt"
	"net/http"
	"time"
)

func AsyncHTTPRequest(url string, client *http.Client) (string, error) {
	respChan := make(chan *http.Response)
	errChan := make(chan error)

	go func() {
		resp, err := client.Get(url)
		if err != nil {
			errChan <- err
			return
		}
		respChan <- resp
	}()

	select {
	case resp := <-respChan:
		return resp.Status, nil
	case err := <-errChan:
		return "", err
	case <-time.After(2 * time.Second):
		return "", fmt.Errorf("request timed out")
	}
}
