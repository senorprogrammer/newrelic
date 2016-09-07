package newrelic

import (
	"time"
)

const (
	testAPIKey     = "test_api_key"
	testTimeString = "2016-01-20T20:29:38+00:00"
)

var (
	testTime, _ = time.Parse(time.RFC3339, testTimeString)
)
