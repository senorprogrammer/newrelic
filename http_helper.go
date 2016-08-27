package newrelic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (c *Client) doGet(path string, params fmt.Stringer, out interface{}) error {
	s := params.String()
	r := strings.NewReader(s)
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

func encodeGetParams(params map[string]interface{}) string {
	s := url.Values{}
	for k, v := range params {
		switch v.(type) {
		case string:
			val := v.(string)
			if val != "" {
				s.Add(k, val)
			}
		case int:
			val := v.(int)
			// TODO: Zero values versus not defined
			if val != 0 {
				s.Add(k, strconv.Itoa(val))
			}
		case []string:
			val := v.([]string)
			if len(val) != 0 {
				s.Add(k, strings.Join(val, ","))
			}
		}
	}
	return s.Encode()
}
