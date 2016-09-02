package newrelic

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAlertEvents(t *testing.T) {
	t.Logf("Starting TestGetAlertEvents")
	for _, tt := range getAlertEventsTests {
		t.Logf("Testing")
		f := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			fmt.Fprintf(w, tt.in.data)
		}
		c, s := initHTTP(t, testAPIKey, f)
		defer s.Close()
		resp, err := c.GetAlertEvents(tt.in.options)
		t.Logf("Checking err...")
		expect(t, tt.out.err, err)
		t.Logf("Checking output...")
		expect(t, tt.out.data, resp)
	}
}

func TestAlertEventOptionsStringer(t *testing.T) {
	t.Logf("Starting TestAlertEventOptionsStringer")
	for _, tt := range alertEventOptionsStringerTests {
		t.Logf("Testing")
		expect(t, tt.out, tt.in.String())
	}
}
