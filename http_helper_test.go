package newrelic

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDoGet(t *testing.T) {
	for _, tt := range doGetTests {
		t.Logf("Testing")
		f := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(tt.in.status)
			fmt.Fprintf(w, tt.in.data)
		}
		c, s := initHTTP(t, testAPIKey, f)
		defer s.Close()
		err := c.doGet(tt.in.path, tt.in.params, tt.in.out)
		t.Logf("Checking err...")
		expect(t, tt.out.err, err)
		t.Logf("Checkout output...")
		expect(t, tt.in.out, tt.out.data)
	}
}

func TestDoRequest(t *testing.T) {

}

func TestEncodeGetParams(t *testing.T) {
	t.Logf("Starting TestEncodeGetParams")
	for _, tt := range encodeGetParamsTests {
		t.Logf("Testing")
		expect(t, tt.out, encodeGetParams(tt.in))
	}
}
