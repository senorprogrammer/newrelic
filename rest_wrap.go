package newrelic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func (c *Client) doRequest(method, path string, body, out interface{}) error {
	var r io.Reader
	if method != "GET" && body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		r = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, c.url.String()+path, r)
	if err != nil {
		return err
	}
	if r != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("X-Api-Key", c.apiKey)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("newrelic http error (%s): %s", resp.Status, b)
	}
	if len(b) == 0 {
		b = []byte{'{', '}'}
	}
	err = json.Unmarshal(b, &out)
	if err != nil {
		return err
	}
	return nil
}
