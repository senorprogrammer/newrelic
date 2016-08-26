/*
 * NewRelic API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2016 by authors and contributors.
 */

package newrelic

import "net/http"

type Client struct {
	apiKey     string
	HttpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		HttpClient: http.DefaultClient,
	}
}
