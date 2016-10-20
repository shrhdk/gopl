package main

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

func fetch(url *url.URL, dst io.Writer, cancel <-chan struct{}) error {
	// Create Request
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return err
	}
	req.Cancel = cancel

	// Send Request
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Copy Response
	_, err = io.Copy(dst, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
