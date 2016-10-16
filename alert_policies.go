package newrelic

type AlertPolicyLinks struct {
	NotificationChannels []int `json:"notification_channels,omitempty"`
	Servers              []int `json:"servers,omitempty"`
}

type AlertPolicyCondition struct {
	ID             int     `json:"id,omitempty"`
	Enabled        bool    `json:"enabled,omitempty"`
	Severity       string  `json:"severity,omitempty"`
	Threshold      float64 `json:"threshold,omitempty"`
	TriggerMinutes int     `json:"trigger_minutes,omitempty"`
	Type           string  `json:"type,omitempty"`
}

// AlertPolicy describes a New Relic alert policy.
type AlertPolicy struct {
	Conditions         []AlertPolicyCondition `json:"conditions,omitempty"`
	Enabled            bool                   `json:"enabled,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	Links              AlertPolicyLinks       `json:"links,omitempty"`
	IncidentPreference string                 `json:"incident_preference,omitempty"`
	Name               string                 `json:"name,omitempty"`
}

// AlertPolicyFilter provides filters for AlertPolicyOptions.
type AlertPolicyFilter struct {
	Name string
}

// AlertPolicyOptions is an optional means of filtering when calling
// GetAlertPolicies.
type AlertPolicyOptions struct {
	Filter AlertPolicyFilter
	Page   int
}

func (o *AlertPolicyOptions) String() string {
	if o == nil {
		return ""
	}
	return encodeGetParams(map[string]interface{}{
		"filter[name]": o.Filter.Name,
		"page":         o.Page,
	})
}

// GetAlertPolices will return a slice of AlertPolicy items, optionally
// filtering by AlertPolicyOptions.
func (c *Client) GetAlertPolicies(options *AlertPolicyOptions) ([]AlertPolicy, error) {
	resp := &struct {
		AlertPolicies []AlertPolicy `json:"alert_policies,omitempty"`
	}{}
	err := c.doGet("alert_policies.json", options, resp)
	if err != nil {
		return nil, err
	}
	return resp.AlertPolicies, nil
}
