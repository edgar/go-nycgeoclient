package nycgeoclient

import (
	// "encoding/json"
	// "io"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type ResponseType int

const (
	JSON ResponseType = 1 + iota
	XML
)

const (
	libraryVersion   = "0.0.1"
	defaultBaseURL   = "https://api.cityofnewyork.us/geoclient/v1/"
	defaultUserAgent = "go-nycgeoclient/" + libraryVersion
	defaulFormat     = JSON
)

type Config struct {
	HttpClient *http.Client

	// Base URL for API requests.
	BaseURL string

	// App ID
	AppID string

	// App key
	AppKey string

	// Format
	Format ResponseType

	// User Agent
	UserAgent string
}

// Client manages communication with NYC Geoclient V1 API.
type Client struct {
	// HTTP client used to communicate with the NYC Geoclient API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent for client
	UserAgent string

	// App ID
	AppID string

	// App key
	AppKey string

	// Format
	Format ResponseType
}

// NewClient returns a new NYC GeoClient API client.
func NewClient(config Config) (*Client, error) {
	httpClient := config.HttpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	baseEndpoint, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}
	format := config.Format
	if format == 0 {
		format = defaulFormat
	}
	userAgent := config.UserAgent
	if userAgent == "" {
		userAgent = defaultUserAgent
	}
	c := &Client{client: httpClient, BaseURL: baseEndpoint, UserAgent: userAgent, AppID: config.AppID, AppKey: config.AppKey, Format: format}

	return c, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// It only accepts relative URLs, if the URL provided starts with an "/" the first "/" will be removed.
// The NYC GeoClient API is read-only, no need to send a body
func (c *Client) NewRequest(method, urlStr string) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	if strings.HasPrefix(urlStr, "/") {
		urlStr = urlStr[1:]
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}
