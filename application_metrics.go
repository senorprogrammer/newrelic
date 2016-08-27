package newrelic

import "strconv"

type ApplicationMetric struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}

type ApplicationMetricOptions struct {
	Name string
	Page int
}

func (c *Client) GetApplicationMetrics(id int, options *ApplicationMetricOptions) ([]ApplicationMetric, error) {
	resp := &struct {
		Metrics []ApplicationMetric `json:"metrics,omitempty"`
	}{}
	path := "/applications/" + strconv.Itoa(id) + "/metrics.json"
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
