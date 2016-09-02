package newrelic

import (
	"fmt"
	"net/http"
	"testing"
)

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
	for _, tt := range applicationOptionsStringerTests {
		t.Logf("Testing")
		expect(t, tt.in.String(), tt.out)
	}
}
