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

// DefaultAPIURL is the default base URL for New Relic's latest API.
const DefaultAPIURL = "https://api.newrelic.com/v2/"

// Client provides a set of methods to interact with the New Relic API.
type Client struct {
	apiKey     string
	httpClient *http.Client
	url        *url.URL
}

// NewClient returns a new Client object for interfacing with the New Relic API.
func NewClient(apiKey string) *Client {
	u, err := url.Parse(DefaultAPIURL)
	if err != nil {
		panic(err)
	}
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: 5 * time.Second},
		url:        u,
	}
}
