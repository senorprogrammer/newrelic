/*
 * NewRelic API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2016 by authors and contributors.
 */

package newrelic

import (
	"net/http"
	"net/url"
	"time"
)

const DEFAULT_API_URL = "https://api.newrelic.com/v2"

type Client struct {
	apiKey     string
	httpClient *http.Client
	url        *url.URL
}

func NewClient(apiKey string) *Client {
	u, err := url.Parse(DEFAULT_API_URL)
	if err != nil {
		panic(err)
	}
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: 5 * time.Second},
		url:        u,
	}
}
