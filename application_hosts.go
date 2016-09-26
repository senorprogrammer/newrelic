package newrelic

import (
	"strconv"
)

type ApplicationHostSummary struct {
	ApdexScore    float64 `json:"apdex_score,omitempty"`
	ErrorRate     float64 `json:"error_rate,omitempty"`
	InstanceCount int     `json:"instance_count,omitempty"`
	ResponseTime  float64 `json:"response_time,omitempty"`
	Throughput    float64 `json:"throughput,omitempty"`
}

type ApplicationHostEndUserSummary struct {
	ResponseTime float64 `json:"response_time,omitempty"`
	Throughput   float64 `json:"throughput,omitempty"`
	ApdexScore   float64 `json:"apdex_score,omitempty"`
}

type ApplicationHostLinks struct {
	Application          int   `json:"application,omitempty"`
	ApplicationInstances []int `json:"application_instances,omitempty"`
	Server               int   `json:"server,omitempty"`
}

type ApplicationHost struct {
	ApplicationName    string                        `json:"application_name,omitempty"`
	ApplicationSummary ApplicationHostSummary        `json:"application_summary,omitempty"`
	HealthStatus       string                        `json:"health_status,omitempty"`
	Host               string                        `json:"host,omitempty"`
	ID                 int                           `json:"idomitempty"`
	Language           string                        `json:"language,omitempty"`
	Links              ApplicationHostLinks          `json:"links,omitempty"`
	EndUserSummary     ApplicationHostEndUserSummary `json:"end_user_summary,omitempty"`
}

type ApplicationHostsFilter struct {
	Hostname string
	Ids      []string
}

type ApplicationHostsOptions struct {
	Filter ApplicationHostsFilter
	Page   int
}

func (c *Client) GetApplicationHosts(id int, options *ApplicationHostsOptions) ([]ApplicationHost, error) {
	resp := &struct {
		ApplicationHosts []ApplicationHost `json:"application_hosts,omitempty"`
	}{}
	path := "applications/" + strconv.Itoa(id) + "/hosts.json"
	err := c.doGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp.ApplicationHosts, nil
}

func (c *Client) GetApplicationHost(aId, hId int) (*ApplicationHost, error) {
	resp := &struct {
		ApplicationHost ApplicationHost `json:"application_host,omitempty"`
	}{}
	path := "applications/" + strconv.Itoa(aId) + "/hosts/" + strconv.Itoa(hId) + ".json"
	err := c.doGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return &resp.ApplicationHost, nil
}

func (o *ApplicationHostsOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"filter[hostname]": o.Filter.Hostname,
		"filter[ids]":      o.Filter.Ids,
		"page":             o.Page,
	})
}
