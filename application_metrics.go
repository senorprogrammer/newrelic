package newrelic

import "strconv"

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

func (o *ApplicationMetricOptions) String() string {
	return encodeGetParams(map[string]interface{}{
		"name": o.Name,
		"page": o.Page,
	})
}
