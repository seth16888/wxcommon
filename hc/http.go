package hc

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultTimeout             = 10 * time.Second
	DefaultIdleConnTimeout     = 90 * time.Second
	DefaultMaxIdleConns        = 100
	DefaultMaxIdleConnsPerHost = 100
)

type CheckRedirectFunc func(req *http.Request, via []*http.Request) error

type Client struct {
	cli *http.Client
}

// NewClient is a constructor
func NewClient(timeout time.Duration, idleConnTimeout time.Duration,
	checkRedirect CheckRedirectFunc) *Client {
	return &Client{
		cli: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        DefaultMaxIdleConns,
				IdleConnTimeout:     idleConnTimeout,
				MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
			},
			CheckRedirect: checkRedirect,
			Jar:           nil,
			Timeout:       timeout,
		},
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.cli.Do(req)
}

func (c *Client) Get(url string) (*http.Response, error) {
	return c.cli.Get(url)
}

func (c *Client) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	return c.cli.Post(url, contentType, body)
}

func (c *Client) Put(url string, contentType string, reader io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("PUT", url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.cli.Do(req)
}

func (c *Client) Delete(uri string) (*http.Response, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	return c.cli.Do(&http.Request{Method: http.MethodDelete, URL: u})
}

func (c *Client) CloseIdleConnections() {
	c.cli.CloseIdleConnections()
}

// CommonCheckRedirect is a common check redirect function
//
// it will return ErrUseLastResponse
func CommonCheckRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}
