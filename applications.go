package newrelic

import (
	"time"
)

type ApplicationSummary struct {
	ResponseTime            float64 `json:"response_time,omitempty"`
	Throughput              float64 `json:"throughput,omitempty"`
	ErrorRate               float64 `json:"error_rate,omitempty"`
	ApdexTarget             float64 `json:"apdex_target,omitempty"`
	ApdexScore              float64 `json:"apdex_score,omitempty"`
	HostCount               int     `json:"host_count,omitempty"`
	InstanceCount           int     `json:"instance_count,omitempty"`
	ConcurrentInstanceCount int     `json:"concurrent_instance_count,omitempty"`
}

type EndUserSummary struct {
	ResponseTime float64 `json:"response_time,omitempty"`
	Throughput   float64 `json:"throughput,omitempty"`
	ApdexTarget  float64 `json:"apdex_target,omitempty"`
	ApdexScore   float64 `json:"apdex_score,omitempty"`
}

type Settings struct {
	AppApdexThreshold        float64 `json:"app_apdex_threshold,omitempty"`
	EndUserApdexThreshold    float64 `json:"end_user_apdex_threshold,omitempty"`
	EnableRealUserMonitoring bool    `json:"enable_real_user_monitoring,omitempty"`
	UseServerSideConfig      bool    `json:"use_server_side_config,omitempty"`
}

type Links struct {
	Servers              []int `json:"servers,omitempty"`
	ApplicationHosts     []int `json:"application_hosts,omitempty"`
	ApplicationInstances []int `json:"application_instances,omitempty"`
	AlertPolicy          int   `json:"alert_policy,omitempty"`
}

type Application struct {
	Id                 int                `json:"id,omitempty"`
	Name               string             `json:"name,omitempty"`
	Language           string             `json:"language,omitempty"`
	HealthStatus       string             `json:"health_status,omitempty"`
	Reporting          bool               `json:"reporting,omitempty"`
	LastReportedAt     time.Time          `json:"last_reported_at,omitempty"`
	ApplicationSummary ApplicationSummary `json:"application_summary,omitempty"`
	EndUserSummary     EndUserSummary     `json:"end_user_summary,omitempty"`
	Settings           Settings           `json:"settings,omitempty"`
	Links              Links              `json:"links,omitempty"`
}

type ApplicationFilter struct {
	Name     string   `json:"name,omitempty"`
	Host     string   `json:"host,omitempty"`
	Ids      []string `json:"host,omitempty"`
	Language string   `json:"language,omitempty"`
}

type ApplicationOptions struct {
	Filter ApplicationFilter `json:"filter,omitempty"`
	Page   int               `json:"page,omitempty"`
}

type ApplicationResponse struct {
	Applications []Application `json:"applications,omitempty"`
}

func (c *Client) GetApplications(options *ApplicationOptions) (*ApplicationResponse, error) {
	apps := &ApplicationResponse{}
	err := c.doGet("/applications.json", options.encode(), apps)
	if err != nil {
		return nil, err
	}
	return apps, nil
}

func (o *ApplicationOptions) encode() string {
	return encodeGetParams(map[string]interface{}{
		"filter[name]":     o.Filter.Name,
		"filter[host]":     o.Filter.Host,
		"filter[ids]":      o.Filter.Ids,
		"filter[language]": o.Filter.Language,
		"page":             o.Page,
	})
}
