package newrelic

import (
	"strconv"
	"time"
)

// ApplicationMetric describes a Metric for a particular Application.
type ApplicationMetric struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}

// ApplicationMetricOptions defines filters for GetApplicationMetrics.
type ApplicationMetricOptions struct {
	Name string
	Page int
}

type ApplicationTimeslice struct {
	From   time.Time          `json:"from,omitempty"`
	To     time.Time          `json:"to,omitempty"`
	Values map[string]float64 `json:"values,omitempty"`
}

type ApplicationMetricData struct {
	Name       string                 `json:"name,omitempty"`
	Timeslices []ApplicationTimeslice `json:"timeslices,omitempty"`
}

type ApplicationMetricDataResp struct {
	From            time.Time               `json:"from,omitempty"`
	To              time.Time               `json:"to,omitempty"`
	MetricsNotFound []string                `json:"metrics_not_found,omitempty"`
	MetricsFound    []string                `json:"metrics_found,omitempty"`
	Metrics         []ApplicationMetricData `json:"metrics,omitempty"`
}

type ApplicationMetricDataOptions struct {
	names     Array
	Values    Array
	From      time.Time
	To        time.Time
	Period    int
	Summarize bool
	Raw       bool
}

// GetApplicationMetrics will return a slice of ApplicationMetric items for a
// particular Application ID, optionally filtering by
// ApplicationMetricOptions.
func (c *Client) GetApplicationMetrics(id int, options *ApplicationMetricOptions) ([]ApplicationMetric, error) {
	resp := &struct {
		Metrics []ApplicationMetric `json:"metrics,omitempty"`
	}{}
	path := "applications/" + strconv.Itoa(id) + "/metrics.json"
	err := c.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}
	return resp.Metrics, nil
}

func (c *Client) GetApplicationMetricData(id int, names []string, options *ApplicationMetricDataOptions) (*ApplicationMetricDataResp, error) {
	resp := &struct {
		MetricData ApplicationMetricDataResp `json:"metric_data",omitempty`
	}{}
	options.names = Array{names}
	path := "applications/" + strconv.Itoa(id) + "/metrics/data.json"
	err := c.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}
	return &resp.MetricData, nil
}

func (o *ApplicationMetricOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"name": o.Name,
		"page": o.Page,
	})
}

func (o *ApplicationMetricDataOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"names[]":   o.names,
		"values[]":  o.Values,
		"from":      o.From,
		"to":        o.To,
		"period":    o.Period,
		"summarize": o.Summarize,
		"raw":       o.Raw,
	})
}
