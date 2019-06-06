package gorillastack

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
)

const DEFAULT_GORILLASTACK_API_URL = "https://api.gorillastack.com";

type Client struct {
	BaseURL   *url.URL
	apiKey		string
	UserAgent string

	Users      *UserService
	httpClient *http.Client
}

func getURL() string {
	if v := os.Getenv("GORILLASTACK_API_URL"); v != "" {
		return v
	}
	return DEFAULT_GORILLASTACK_API_URL

}

func NewClient(apiKey string) (*Client, error) {
	apiUrl, err := url.Parse(getURL())
	if (err != nil) {
		return nil, err
	}

	c := &Client{
		httpClient: http.DefaultClient,
		BaseURL: apiUrl,
		apiKey: apiKey,
	}
	c.Users = &UserService{c: c}

	return c, nil
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
