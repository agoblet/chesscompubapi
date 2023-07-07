package chesscompubapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	defaultBaseURL = "https://api.chess.com"
)

/*
Client handles requests to the chess.com PubAPI.
The Client should always be created via the NewClient function.
The zero-value of Client is invalid.

The Get* and List* functions of the Client, e.g. GetPlayerProfile, all behave in the same way.
They return an *HTTPError if the API returns a status code other than 200.
They an *url.Error if the API call failed for other reasons.
They return an error if the response cannot be decoded into the return type, e.g. PlayerProfile for function GetPlayerProfile.
*/
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// Option is the inferface used for functional options to configure the *Client.
type Option func(*Client)

// HTTPError is used when a status code other than 200 is returned by the chess.com PubAPI.
type HTTPError struct {
	StatusCode   int
	ResponseBody string
}

// NewClient creates a new *Client to send requests to the chess.com PubAPI.
// Accepts any number of Options to customize the *Client.
func NewClient(options ...Option) *Client {
	client := &Client{
		baseURL: defaultBaseURL, httpClient: http.DefaultClient}
	for _, option := range options {
		option(client)
	}
	return client
}

// WithBaseURL configures a custom base URL to send requests to.
// It must have the format protocol://host[:port].
// If this option is omitted, the default URL https://api.chess.com will be used.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithHTTPClient configures a custom *http.Client to use for requests to the chess.com PubAPI.
// It allows users to configure client properties such as timeouts and connection pooling.
// If this option is omitted, http.DefaultClient will be used.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func (c *Client) get(path string) ([]byte, error) {
	response, err := c.httpClient.Get(fmt.Sprintf("%s/pub/%s", c.baseURL, path))
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
