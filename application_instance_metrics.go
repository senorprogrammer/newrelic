package newrelic

import (
	"strconv"
)

// GetApplicationInstanceMetrics will return a slice of Metric items for a
// particular Application ID's instance ID, optionally filtering by
// MetricsOptions.
func (c *Client) GetApplicationInstanceMetrics(appID, instanceID int, options *MetricsOptions) ([]Metric, error) {
	resp := &struct {
		Metrics []Metric `json:"metrics,omitempty"`
	}{}
	path := "applications/" + strconv.Itoa(appID) + "/instances/" + strconv.Itoa(instanceID) + "/metrics.json"
	err := c.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}
	return resp.Metrics, nil
}

// GetApplicationInstanceMetricData will return all metric data for a
// particular application's instance and slice of metric names, optionally
// filtered by MetricDataOptions.
func (c *Client) GetApplicationInstanceMetricData(appID, instanceID int, names []string, options *MetricDataOptions) (*MetricDataResponse, error) {
	resp := &struct {
		MetricData MetricDataResponse `json:"metric_data,omitempty"`
	}{}
	if options == nil {
		options = &MetricDataOptions{}
	}
	options.Names = Array{names}
	path := "applications/" + strconv.Itoa(appID) + "/instances/" + strconv.Itoa(instanceID) + "/metrics/data.json"
	err := c.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}
	return &resp.MetricData, nil
}
