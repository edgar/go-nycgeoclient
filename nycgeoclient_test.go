package nycgeoclient

import (
	"net/url"
	"testing"
)

func TestNewClient(t *testing.T) {
	c, _ := NewClient(Config{})

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is \"%v\", want \"%v\"", got, want)
	}
	if got, want := c.UserAgent, defaultUserAgent; got != want {
		t.Errorf("NewClient UserAgent is \"%v\", want \"%v\"", got, want)
	}
	if got, want := c.Format, defaulFormat; got != want {
		t.Errorf("NewClient Format is \"%v\", want \"%v\"", got, want)
	}
	if got, want := c.AppID, ""; got != want {
		t.Errorf("NewClient AppID is \"%v\", want \"%v\"", got, want)
	}
	if got, want := c.AppKey, ""; got != want {
		t.Errorf("NewClient AppKey is \"%v\", want \"%v\"", got, want)
	}
}

func TestNewClient_withCustomValues(t *testing.T) {
	testURL := "http://test.com/path/"
	testUserAgent := "dumb-1.0"
	testAppID := "foo"
	testAppKey := "bar"
	testFormat := XML

	c, _ := NewClient(Config{BaseURL: testURL, UserAgent: testUserAgent, AppID: testAppID, AppKey: testAppKey, Format: testFormat})

	if got, want := c.BaseURL.String(), testURL; got != want {
		t.Errorf("NewClient BaseURL is \"%v\", want \"%v\"", got, want)
	}
	if got, want := c.UserAgent, testUserAgent; got != want {
		t.Errorf("NewClient UserAgent is \"%v\", want \"%v\"", got, want)
	}
	if got, want := c.Format, testFormat; got != want {
		t.Errorf("NewClient Format is \"%v\", want \"%v\"", got, want)
	}
	if got, want := c.AppID, testAppID; got != want {
		t.Errorf("NewClient AppID is \"%v\", want \"%v\"", got, want)
	}
	if got, want := c.AppKey, testAppKey; got != want {
		t.Errorf("NewClient AppKey is \"%v\", want \"%v\"", got, want)
	}
}

func TestNewClient_withURLWithoutTrailingSlash(t *testing.T) {
	testURL := "http://test.com/path"
	formattedTestURL := testURL + "/"

	c, _ := NewClient(Config{BaseURL: testURL})

	if got, want := c.BaseURL.String(), formattedTestURL; got != want {
		t.Errorf("NewClient BaseURL is \"%v\", want \"%v\"", got, want)
	}
}

func TestNewRequest(t *testing.T) {
	c, _ := NewClient(Config{})

	inURL, outURL := "foo", defaultBaseURL+"foo"
	req, _ := c.NewRequest("GET", inURL)

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that default user-agent is attached to the request
	if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
		t.Errorf("NewRequest() User-Agent is %v, want %v", got, want)
	}
}

func TestNewRequest_withAbsoluteURL(t *testing.T) {
	c, _ := NewClient(Config{})

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	req, _ := c.NewRequest("GET", inURL)

	// test that absolute URL is handled as a relative one
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}
}

func TestNewRequest_badURL(t *testing.T) {
	c, _ := NewClient(Config{})
	_, err := c.NewRequest("GET", ":")
	testURLParseError(t, err)
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}
