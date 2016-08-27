package newrelic

type AlertEvent struct {
	Id            int    `json:"id,omitempty"`
	EventType     string `json:"event_type,omitempty"`
	Product       string `json:"product,omitempty"`
	EntityType    string `json:"entity_type,omitempty"`
	EntityGroupId int    `json:"entity_group_id,omitempty"`
	EntityId      int    `json:"entity_id,omitempty"`
	Priority      string `json:"priority,omitempty"`
	Description   string `json:"description,omitempty"`
	Timestamp     int    `json:"timestamp,omitempty"`
	IncidentId    int    `json:"incident_id"`
}

type AlertEventFilter struct {
	Product       string
	EntityType    string
	EntityGroupId int
	EntityId      int
	EventType     string
}

type AlertEventOptions struct {
	Filter AlertEventFilter
	Page   int
}

func (o *AlertEventOptions) String() string {
	return encodeGetParams(map[string]interface{}{
		"filter[product]":         o.Filter.Product,
		"filter[entity_type]":     o.Filter.EntityType,
		"filter[entity_group_id]": o.Filter.EntityGroupId,
		"filter[entity_id]":       o.Filter.EntityId,
		"filter[event_type]":      o.Filter.EventType,
		"page":                    o.Page,
	})
}

func (c *Client) GetAlertEvents(options *AlertEventOptions) ([]AlertEvent, error) {
	resp := &struct {
		RecentEvents []AlertEvent `json:"recent_events,omitempty"`
	}{}
	err := c.doGet("alerts_events.json", options, resp)
	if err != nil {
		return nil, err
	}
	return resp.RecentEvents, nil
}
