package newrelic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) doGet(path, params string, out interface{}) error {
	r := strings.NewReader(params)
	req, err := http.NewRequest("GET", c.url.String()+path, r)
	if err != nil {
		return err
	}
	req.Header.Add("X-Api-Key", c.apiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return c.doRequest(req, out)
}

func (c *Client) doRequest(req *http.Request, out interface{}) error {
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
