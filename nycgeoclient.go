package nycgeoclient

import (
	"encoding/json"
	"io"
	"net/url"
)

const (
	libraryVersion = "0.0.1"
	defaultBaseURL = "https://api.cityofnewyork.us/geoclient/v1/"
	userAgent      = "go-nycgeoclient/" + libraryVersion
	format         = "json"
)

/ Client manages communication with NYC Geoclient V1 API.
type Client struct {
  // HTTP client used to communicate with the NYC Geoclient API.
  client *http.Client

  // Base URL for API requests.
  BaseURL *url.URL

  // User agent for client
  UserAgent string

  // App ID
  AppId string

  // App key
  AppKey string

  // Format (json or xml)
  Format string
}
