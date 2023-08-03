package chesscompubapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

/*
Client handles HTTP requests to the chess.com PubAPI.

The Get* and List* functions of the Client, e.g. GetPlayerProfile, all behave in the same way:
  - Get* functions return a single item.
  - List* functions return a slice of items.
  - They return an *HTTPError if the API returns a status code other than 200.
  - They an *url.Error if the API call failed for other reasons.
  - They return an error if the response cannot be decoded into the return type, e.g. PlayerProfile for function GetPlayerProfile.
*/
type Client struct {
	customBaseURL    string
	hasCustomBaseURL bool
	httpClient       *http.Client
}

// Option is the inferface used for functional options to configure the *Client.
// To apply an Option, pass it to the NewClient function.
type Option interface {
	apply(*Client)
}

type option struct {
	applyFn func(*Client)
}

func (o *option) apply(c *Client) {
	o.applyFn(c)
}

// HTTPError is used when a status code other than 200 is returned by the chess.com PubAPI.
type HTTPError struct {
	StatusCode   int
	ResponseBody string
}

// NewClient creates a new *Client to send requests to the chess.com PubAPI.
// Accepts any number of Options to customize the *Client.
func NewClient(options ...Option) *Client {
	client := &Client{
		httpClient: &http.Client{
			Transport: http.DefaultTransport.(*http.Transport).Clone(),
		},
	}
	for _, option := range options {
		option.apply(client)
	}
	return client
}

// WithBaseURL configures a custom base URL to send requests to.
// It must have the format protocol://host[:port].
// If this option is omitted, the default URL https://api.chess.com will be used.
// To apply this Option, pass it to the NewClient function.
func WithBaseURL(url string) Option {
	return &option{func(c *Client) {
		c.customBaseURL = url
		c.hasCustomBaseURL = true
	}}
}

// WithTimeout configures a timeout for requests to the chess.com PubAPI.
// If this Option is omitted or a timeout of 0 is passed, there will be no timeout.
// To apply this Option, pass it to the NewClient function.
func WithTimeout(timeout time.Duration) Option {
	return &option{func(c *Client) {
		c.httpClient.Timeout = timeout
	}}
}

func (c *Client) getBaseURL() string {
	baseURL := "https://api.chess.com"
	if c.hasCustomBaseURL {
		baseURL = c.customBaseURL
	}
	return baseURL
}

func (c *Client) getHTTPClient() *http.Client {
	if c.httpClient != nil {
		return c.httpClient
	}
	return http.DefaultClient
}

func (c *Client) get(path string) ([]byte, error) {
	response, err := c.getHTTPClient().Get(fmt.Sprintf("%s/pub/%s", c.getBaseURL(), path))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, &HTTPError{
			StatusCode:   response.StatusCode,
			ResponseBody: string(body),
		}
	}

	return body, nil
}

func (c *Client) getInto(path string, v any) error {
	body, err := c.get(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

// Error returns a formatted string representing the HTTPError.
func (e *HTTPError) Error() string {
	const template = "%d %s"
	return fmt.Sprintf(template, e.StatusCode, e.ResponseBody)
}
