package newrelic

import (
	"strconv"
	"time"
)

type ApplicationHostMetric struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}

type ApplicationHostMetricsOptions struct {
	Name string `json:"name,omitempty"`
	Page int    `json:"name,omitempty"`
}

type ApplicationHostMetricDataOptions struct {
	names     Array
	Values    Array
	From      time.Time
	To        time.Time
	Period    int
	Summarize bool
	Raw       bool
}

type ApplicationHostTimeslice struct {
	From   time.Time          `json:"from,omitempty"`
	To     time.Time          `json:"to,omitempty"`
	Values map[string]float64 `json:"values,omitempty"`
}

type ApplicationHostMetricData struct {
	Name       string                     `json:"name,omitempty"`
	Timeslices []ApplicationHostTimeslice `json:"timeslices,omitempty"`
}

type ApplicationHostMetricDataResp struct {
	From            time.Time                   `json:"from,omitempty"`
	To              time.Time                   `json:"from,omitempty"`
	MetricsNotFound []string                    `json:"metrics_not_found,omitempty"`
	MetricsFound    []string                    `json:"metrics_found,omitempty"`
	Metrics         []ApplicationHostMetricData `json:"metrics,omitempty"`
}

func (c *Client) GetApplicationHostMetrics(aId, hId int, options *ApplicationHostMetricsOptions) ([]ApplicationHostMetric, error) {
	resp := &struct {
		Metrics []ApplicationHostMetric `json:"metrics,omitempty"`
	}{}
	path := "applications/" + strconv.Itoa(aId) + "/hosts/" + strconv.Itoa(hId) + "/metrics.json"
	err := c.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}
	return resp.Metrics, nil
}

func (c *Client) GetApplicationHostMetricData(aId, hId int, names []string, options *ApplicationHostMetricDataOptions) (*ApplicationHostMetricDataResp, error) {
	resp := &struct {
		MetricData ApplicationHostMetricDataResp `json:"metric_data,omitempty"`
	}{}
	if options == nil {
		options = &ApplicationHostMetricDataOptions{}
	}
	options.names = Array{names}
	path := "applications/" + strconv.Itoa(aId) + "/hosts/" + strconv.Itoa(hId) + "/metrics/data.json"
	err := c.doGet(path, options, resp)
	if err != nil {
		return nil, err
	}
	return &resp.MetricData, nil
}

func (o *ApplicationHostMetricsOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"name": o.Name,
		"page": o.Page,
	})
}

func (o *ApplicationHostMetricDataOptions) String() string {
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
