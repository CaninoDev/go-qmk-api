package qmk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	defaultBaseURL   = "https://api.qmk.fm"
	apiVersion       = "v1"
	defaultMediaType = "application/json"
	defaultTimeOut   = time.Second * 2
)

// Client holds information about the API that will be consume and contain a reference to http.Client used to make request to the API
type Client struct {
	*http.Client
	BaseURL *url.URL
}

// New constructs a new Client embedding with the user supplied http.Client. If none are provided, sane defaults are used. Returns an initialized Client.
func New(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: defaultTimeOut,
		}
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	baseURL.Path = path.Join(baseURL.Path, apiVersion)
	client := &Client{
		httpClient,
		baseURL,
	}
	return client
}

func (c *Client) newRequest(method, endpoint string, reqType interface{}, body interface{}) (*http.Response, error) {
	c.BaseURL.Path = path.Join(c.BaseURL.Path, endpoint)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.BaseURL.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	switch reqType.(type) {
	case []byte:
		req.Header.Set("Accept", "application/text")
	default:
		req.Header.Set("Accept", "application/json")
	}

	resp, err := c.do(req, &reqType)
	return resp, err
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Don't try to decode on 204s
	if resp.StatusCode == http.StatusNoContent {
		return resp, nil
	}
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		err = json.NewDecoder(resp.Body).Decode(v)
		return resp, err
	}
	return resp, errors.New(string(resp.StatusCode))
}
