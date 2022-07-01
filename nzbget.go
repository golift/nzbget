package nzbget

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/rpc/json"
)

// Package defaults.
const (
	DefaultTimeout = 1 * time.Minute
)

// maxBodyError limits the size of the body that gets appended to an error.
// ie. if a bad URL is used, the body content that shows the error gets appended to the error.
// sometimes the body cotent is huge, so we limit it to 1000 charaters.
const maxBodyError = 1000

// Config is the input data needed to return a NZBGet struct.
// This is setup to allow you to easily pass this data in from a config file.
type Config struct {
	URL       string   `json:"url" toml:"url" xml:"url" yaml:"url"`
	User      string   `json:"user" toml:"user" xml:"user" yaml:"user"`
	Pass      string   `json:"pass" toml:"pass" xml:"pass" yaml:"pass"`
	Timeout   Duration `json:"timeout" toml:"timeout" xml:"timeout" yaml:"timeout"`
	VerifySSL bool     `json:"verify_ssl" toml:"verify_ssl" xml:"verify_ssl" yaml:"verify_ssl"`
}

// NZBGet is what you get in return for passing in a valid Config to New().
type NZBGet struct {
	client *client
	url    string
}

// Duration is used to parse durations from a config file.
type Duration struct{ time.Duration }

type client struct {
	auth string
	*http.Client
}

// UnmarshalText parses a duration type from a config file.
func (d *Duration) UnmarshalText(data []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(data))
	return err // nolint:wrapcheck,wsl
}

func New(config *Config) *NZBGet {
	timeout := config.Timeout.Duration
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	// Set username and password if one's configured.
	auth := config.User + ":" + config.Pass
	if auth != ":" {
		auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	} else {
		auth = ""
	}

	nzb := &NZBGet{
		url: strings.TrimSuffix(config.URL, "/jsonrpc") + "/jsonrpc",
		client: &client{
			auth: auth,
			Client: &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: config.VerifySSL}, //nolint:gosec
				},
				Timeout: timeout,
			},
		},
	}

	return nzb
}

// GetInto is a helper method to make a JSON-RPC request and turn the response into structured data.
func (n *NZBGet) GetInto(method string, output interface{}, args ...interface{}) (int, error) {
	message, err := json.EncodeClientRequest(method, args)
	if err != nil {
		return 0, fmt.Errorf("encoding request: %w", err)
	}

	req, err := http.NewRequest("POST", n.url, bytes.NewBuffer(message))
	if err != nil {
		return 0, fmt.Errorf("creating request: %w", err)
	}

	resp, err := n.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	defer buf.Reset()

	tee := io.TeeReader(resp.Body, &buf)
	if err := json.DecodeClientResponse(tee, &output); err != nil {
		return buf.Len(), fmt.Errorf("parsing response: %w: %s %s", err, resp.Status, limitBuf(&buf))
	}

	return buf.Len(), nil
}

// Do allows overriding the http request parameters in aggregate.
func (c *client) Do(req *http.Request) (*http.Response, error) {
	if c.auth != "" {
		req.Header.Set("Authorization", c.auth)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return resp, fmt.Errorf("making request: %w", err)
	}

	return resp, nil
}

func limitBuf(buf *bytes.Buffer) string {
	str := buf.String()
	if len(str) > maxBodyError {
		return str[:maxBodyError]
	}

	return str
}
