package newrelic

import (
	"time"
)

// ServersFilter is the filtering component of ServersOptions.
type ServersFilter struct {
	Name     string
	Host     string
	IDs      []int
	Labels   []string
	Reported bool
}

// ServersOptions provides a filtering mechanism for GetServers.
type ServersOptions struct {
	Filter ServersFilter
	Page   int
}

// ServerSummary describes the summary component of a Server.
type ServerSummary struct {
	CPU             float64 `json:"cpu,omitempty"`
	CPUStolen       float64 `json:"cpu_stolen,omitempty"`
	DiskIO          float64 `json:"disk_io,omitempty"`
	Memory          float64 `json:"memory,omitempty"`
	MemoryUsed      int64   `json:"memory_used,omitempty"`
	MemoryTotal     int64   `json:"memory_total,omitempty"`
	FullestDisk     float64 `json:"fullest_disk,omitempty"`
	FullestDiskFree int64   `json:"fullest_disk_free,omitempty"`
}

// ServerLinks link Servers to the objects to which they pertain.
type ServerLinks struct {
	AlertPolicy int `json:"alert_policy,omitempty"`
}

// Server represents a New Relic Server.
type Server struct {
	ID             int           `json:"id,omitempty"`
	AccountID      int           `json:"account_id,omitempty"`
	Name           string        `json:"name,omitempty"`
	Host           string        `json:"host,omitempty"`
	HealthStatus   string        `json:"health_status,omitempty"`
	Reporting      bool          `json:"reporting,omitempty"`
	LastReportedAt time.Time     `json:"last_reported_at,omitempty"`
	Summary        ServerSummary `json:"summary,omitempty"`
	Links          ServerLinks   `json:"links,omitempty"`
}

// GetServers will return a slice of New Relic Servers, optionally filtered by
// ServerOptions.
func (c *Client) GetServers(opt *ServersOptions) ([]Server, error) {
	resp := &struct {
		Servers []Server `json:"servers,omitempty"`
	}{}
	path := "servers.json"
	err := c.doGet(path, opt, resp)
	if err != nil {
		return nil, err
	}
	return resp.Servers, nil
}

func (o *ServersOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"filter[name]":     o.Filter.Name,
		"filter[host]":     o.Filter.Host,
		"filter[ids]":      o.Filter.IDs,
		"filter[labels]":   o.Filter.Labels,
		"filter[reported]": o.Filter.Reported,
		"page":             o.Page,
	})
}
