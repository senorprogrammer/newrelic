package newrelic

import (
	"time"
)

// Metric describes a New Relic metric.
type Metric struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}

// MetricsOptions options allow filtering when getting lists of metric names
// associated with an entity.
type MetricsOptions struct {
	Name string
	Page int
}

// MetricTimeslice describes the period to which a Metric pertains.
type MetricTimeslice struct {
	From   time.Time          `json:"from,omitempty"`
	To     time.Time          `json:"to,omitempty"`
	Values map[string]float64 `json:"values,omitempty"`
}

// MetricData describes the data for a particular metric.
type MetricData struct {
	Name       string            `json:"name,omitempty"`
	Timeslices []MetricTimeslice `json:"timeslices,omitempty"`
}

// MetricDataOptions allow filtering when getting data about a particular set
// of New Relic metrics.
type MetricDataOptions struct {
	Names     Array
	Values    Array
	From      time.Time
	To        time.Time
	Period    int
	Summarize bool
	Raw       bool
}

// MetricDataResponse is the response received from New Relic for any request
// for metric data.
type MetricDataResponse struct {
	From            time.Time    `json:"from,omitempty"`
	To              time.Time    `json:"to,omitempty"`
	MetricsNotFound []string     `json:"metrics_not_found,omitempty"`
	MetricsFound    []string     `json:"metrics_found,omitempty"`
	Metrics         []MetricData `json:"metrics,omitempty"`
}

func (o *MetricsOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"name": o.Name,
		"page": o.Page,
	})
}

func (o *MetricDataOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"names[]":   o.Names,
		"values[]":  o.Values,
		"from":      o.From,
		"to":        o.To,
		"period":    o.Period,
		"summarize": o.Summarize,
		"raw":       o.Raw,
	})
}
