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
	apiKey, appKey string
	HttpClient     *http.Client
}

func NewClient(apiKey, appKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		appKey:     appKey,
		HttpClient: http.DefaultClient,
	}
}
