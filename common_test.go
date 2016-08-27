package newrelic

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func expect(t *testing.T, exp interface{}, got interface{}) {
	if !reflect.DeepEqual(got, exp) {
		e := pretty.Sprint(exp)
		g := pretty.Sprint(got)
		t.Errorf("Want (type %v): \n---\n\"%s\"\n---\n", reflect.TypeOf(exp), e)
		t.Errorf("Got (type %v): \n---\n\"%s\"\n---\n", reflect.TypeOf(got), g)
	} else {
		t.Logf("Expectation matches.")
	}
}

type RewriteTransport struct {
	URL *url.URL
}

func (t RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Maintain the opaque path, but rewrite the base URL
	req.URL.Scheme = t.URL.Scheme
	req.URL.Host = t.URL.Host
	req.URL.Path = path.Join(t.URL.Path, req.URL.Path)
	return http.DefaultTransport.RoundTrip(req)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func handlerFactory(code int) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		fmt.Fprintf(w, "%s ", r.Method)
		fmt.Fprintf(w, "%s %s %s", r.Host, r.URL.Path, r.URL.Query().Encode())
	}
}

func initHTTP(t *testing.T, apiKey string, f HandlerFunc) (*Client, *httptest.Server) {
	mockServer := httptest.NewServer(http.HandlerFunc(f))
	u, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Log("Problem initiating mock http server: ", err)
	}
	c := NewClient(apiKey)
	c.httpClient = &http.Client{Transport: RewriteTransport{URL: u}}
	return c, mockServer
}
