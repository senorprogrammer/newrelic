package newrelic

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

const (
	testAPIKey          = "test_api_key"
	testApplicationJSON = `
  {
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
  }
`
)

var (
	testTime, _     = time.Parse(time.RFC3339, "2016-01-20T20:29:38+00:00")
	testApplication = Application{
		ID:             12345,
		Name:           "test.example.com",
		Language:       "java",
		HealthStatus:   "green",
		Reporting:      true,
		LastReportedAt: testTime,
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
	testApplications = []Application{
		testApplication,
		testApplication,
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
			data: `{ "application":` + testApplicationJSON + `}`,
		},
		getApplicationOutput{
			data: &testApplication,
		},
	},
}

type getApplicationsInput struct {
	options *ApplicationOptions
	data    string
}

type getApplicationsOutput struct {
	data []Application
	err  error
}

var getApplicationsTests = []struct {
	in  getApplicationsInput
	out getApplicationsOutput
}{
	{
		getApplicationsInput{
			options: nil,
			data:    `{"applications":[` + testApplicationJSON + "," + testApplicationJSON + "]}",
		},
		getApplicationsOutput{
			data: testApplications,
			err:  nil,
		},
	},
}

var testApplicationOptionsStringer = []struct {
	in  *ApplicationOptions
	out string
}{
	{
		&ApplicationOptions{},
		"",
	},
	{
		&ApplicationOptions{
			Filter: ApplicationFilter{
				Name:     "testName",
				Host:     "testHost",
				Ids:      []string{"test1", "test2"},
				Language: "java",
			},
			Page: 5,
		},
		`filter%5Bhost%5D=testHost` +
			`&filter%5Bids%5D=test1%2Ctest2` +
			`&filter%5Blanguage%5D=java` +
			`&filter%5Bname%5D=testName` +
			`&page=5`,
	},
	{
		nil,
		"",
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

func TestGetApplications(t *testing.T) {
	t.Logf("Starting TestGetApplication")
	for _, tt := range getApplicationsTests {
		t.Logf("Testing")
		h := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			fmt.Fprintf(w, tt.in.data)
		}
		c, s := initHTTP(t, testAPIKey, h)
		defer s.Close()
		resp, err := c.GetApplications(tt.in.options)
		t.Logf("Checking err...")
		expect(t, tt.out.err, err)
		t.Logf("Checking output...")
		expect(t, tt.out.data, resp)
	}
}

func TestApplicationOptionsStringer(t *testing.T) {
	t.Logf("Starting TestApplicationOptionsStringer")
	for _, tt := range testApplicationOptionsStringer {
		t.Logf("Testing")
		expect(t, tt.in.String(), tt.out)
	}
}
