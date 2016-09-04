package newrelic

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetApplicationMetrics(t *testing.T) {
	t.Logf("Starting TestGetApplication")
	for _, tt := range getApplicaitonMetricsTests {
		t.Logf("Testing")
		f := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			fmt.Fprintf(w, tt.in.data)
		}
		c, s := initHTTP(t, testAPIKey, f)
		defer s.Close()
		resp, err := c.GetApplicationMetrics(tt.in.id, tt.in.options)
		t.Logf("Checking err...")
		expect(t, tt.out.err, err)
		t.Logf("Checking output...")
		expect(t, tt.out.data, resp)
	}
}

func TestGetApplicationMetricData(t *testing.T) {
	t.Log("Starting TestGetApplicationMetricData")
	// TODO
}

func TestApplicationMetricOptionsStringer(t *testing.T) {
	t.Logf("Starting TestApplicationMetricOptionsStringer")
	for _, tt := range applicationMetricOptionsStringerTests {
		t.Logf("Testing")
		expect(t, tt.out, tt.in.String())
	}
}

func TestApplicationMetricDataOptionsStringer(t *testing.T) {
	t.Logf("Starting TestApplicationMetricDataOptionsStringer")
	// TODO
}
