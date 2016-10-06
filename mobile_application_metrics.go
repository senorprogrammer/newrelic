package newrelic

import (
	"strconv"
)

// GetMobileAplicationMetrics will return a slice of Metric items for a
// particular MobileAplication ID, optionally filtering by
// MetricsOptions.
func (c *Client) GetMobileApplicationMetrics(id int, options *MetricsOptions) ([]Metric, error) {
	resp := &struct {
		Metrics []Metric `json:"metrics,omitempty"`
	}{}
	path := "mobile_applications/" + strconv.Itoa(id) + "/metrics.json"
	err := c.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}
	return resp.Metrics, nil
}

// GetMobileAplicationMetricData will return all metric data for a particular
// MobileAplication and slice of metric names, optionally filtered by
// MetricDataOptions.
func (c *Client) GetMobileApplicationMetricData(id int, names []string, options *MetricDataOptions) (*MetricDataResponse, error) {
	resp := &struct {
		MetricData MetricDataResponse `json:"metric_data,omitempty"`
	}{}
	if options == nil {
		options = &MetricDataOptions{}
	}
	options.Names = Array{names}
	path := "mobile_applications/" + strconv.Itoa(id) + "/metrics/data.json"
	err := c.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}
	return &resp.MetricData, nil
}
