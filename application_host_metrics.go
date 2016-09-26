package newrelic

import (
	"strconv"
)

// GetApplicationHostMetrics will return a slice of Metric items for a
// particular Application ID's Host ID, optionally filtering by
// MetricsOptions.
func (c *Client) GetApplicationHostMetrics(appID, hostID int, options *MetricsOptions) ([]Metric, error) {
	resp := &struct {
		Metrics []Metric `json:"metrics,omitempty"`
	}{}
	path := "applications/" + strconv.Itoa(appID) + "/hosts/" + strconv.Itoa(hostID) + "/metrics.json"
	err := c.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}
	return resp.Metrics, nil
}

// GetApplicationHostMetricData will return all metric data for a particular
// application's host and slice of metric names, optionally filtered by
// MetricDataOptions.
func (c *Client) GetApplicationHostMetricData(appID, hostID int, names []string, options *MetricDataOptions) (*MetricDataResponse, error) {
	resp := &struct {
		MetricData MetricDataResponse `json:"metric_data,omitempty"`
	}{}
	if options == nil {
		options = &MetricDataOptions{}
	}
	options.Names = Array{names}
	path := "applications/" + strconv.Itoa(appID) + "/hosts/" + strconv.Itoa(hostID) + "/metrics/data.json"
	err := c.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}
	return &resp.MetricData, nil
}
