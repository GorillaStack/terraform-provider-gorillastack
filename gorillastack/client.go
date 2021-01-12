package gorillastack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const DEFAULT_GORILLASTACK_API_URL = "https://api.gorillastack.com"

type Client struct {
	baseURL    *url.URL
	apiKey     string
	UserAgent  string
	Users      *UserService
	httpClient *http.Client
}

func getURL() string {
	if v := os.Getenv("GORILLASTACK_API_URL"); v != "" {
		return v
	}
	return DEFAULT_GORILLASTACK_API_URL

}

func getPrefix() string {
	if v := os.Getenv("PREFIX_OVERRIDE"); v != "" {
		return v 
	}
	return "/v1"
}

func NewClient(apiKey string) (*Client, error) {
	apiUrl, err := url.Parse(getURL())
	if err != nil {
		return nil, err
	}

	c := &Client{
		httpClient: http.DefaultClient,
		baseURL:    apiUrl,
		apiKey:     apiKey,
	}
	c.Users = &UserService{c: c}

	return c, nil
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: getPrefix() + path}
	u := c.baseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != "" {
		sBody := fmt.Sprintf("%s", body)
		buf = bytes.NewBufferString(fmt.Sprintf("%s", sBody))
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("[%d] Error: %s \n %+v", resp.StatusCode, body, resp))
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
