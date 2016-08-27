package newrelic

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

const (
	testAPIKey          = "test_api_key"
	fixtureDir          = "./test-fixtures"
	testApplicationJSON = `
{
  "application": {
    "application_summary": {
      "apdex_score": 1.0,
      "apdex_target": 0.5,
      "error_rate": 0.0,
      "host_count": 1,
      "instance_count": 1,
      "response_time": 0.263,
      "throughput": 12.3
    },
    "end_user_summary": {
      "response_time": 0.263,
      "throughput": 12.3,
      "apdex_target": 0.5,
      "apdex_score": 1
    },
    "health_status": "green",
    "id": 12345,
    "language": "java",
    "last_reported_at": "2016-01-20T20:29:38+00:00",
    "links": {
      "alert_policy": 123,
      "application_hosts": [
        1234567
      ],
      "application_instances": [
        1234568
      ],
      "servers": [
        54321
      ]
    },
    "name": "test.example.com",
    "reporting": true,
    "settings": {
      "app_apdex_threshold": 0.5,
      "enable_real_user_monitoring": true,
      "end_user_apdex_threshold": 1.0,
      "use_server_side_config": false
    }
  },
  "links": {
    "application.alert_policy": "/v2/alert_policies/{alert_policy_id}",
    "application.application_host": "/v2/application/{application_id}/hosts/{host_id}",
    "application.application_hosts": "/v2/application/{application_id}/hosts?ids={host_ids}",
    "application.application_instance": "/v2/application/{application_id}/instances/{instance_id}",
    "application.application_instances": "/v2/application/{application_id}/instances?ids={instance_ids}",
    "application.server": "/v2/servers/{server_id}",
    "application.servers": "/v2/servers?ids={server_ids}"
  }
}
`
)

var (
	testTime, _     = time.Parse(time.RFC3339, "2016-01-20T20:29:38+00:00")
	testApplication = &Application{
		ID:             12345,
		Name:           "test.example.com",
		Language:       "java",
		HealthStatus:   "green",
		Reporting:      true,
		LastReportedAt: Time{testTime},
		ApplicationSummary: ApplicationSummary{
			ResponseTime:            0.263,
			Throughput:              12.3,
			ErrorRate:               0,
			ApdexTarget:             0.5,
			ApdexScore:              1,
			HostCount:               1,
			InstanceCount:           1,
			ConcurrentInstanceCount: 0,
		},
		EndUserSummary: EndUserSummary{
			ResponseTime: 0.263,
			Throughput:   12.3,
			ApdexTarget:  0.5,
			ApdexScore:   1,
		},
		Settings: Settings{
			AppApdexThreshold:        0.5,
			EndUserApdexThreshold:    1,
			EnableRealUserMonitoring: true,
			UseServerSideConfig:      false,
		},
		Links: Links{
			Servers:              []int{54321},
			ApplicationHosts:     []int{1234567},
			ApplicationInstances: []int{1234568},
			AlertPolicy:          123,
		},
	}
)

type getApplicationInput struct {
	id   int
	data string
}

type getApplicationOutput struct {
	data *Application
	err  error
}

var getApplicationTests = []struct {
	in  getApplicationInput
	out getApplicationOutput
}{
	{
		getApplicationInput{
			id:   12345,
			data: testApplicationJSON,
		},
		getApplicationOutput{
			data: testApplication,
		},
	},
}

func TestGetApplication(t *testing.T) {
	t.Logf("Starting TestGetApplication")
	for _, tt := range getApplicationTests {
		t.Logf("Testing")
		h := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			fmt.Fprintf(w, tt.in.data)
		}
		c, s := initHTTP(t, testAPIKey, h)
		defer s.Close()
		resp, err := c.GetApplication(tt.in.id)
		t.Logf("Checking err...")
		expect(t, tt.out.err, err)
		t.Logf("Checking output...")
		expect(t, tt.out.data, resp)
	}
}
