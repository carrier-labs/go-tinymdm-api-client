// Package client provides the core HTTP client, authentication, and configuration for TinyMDM API.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/carrier-labs/go-tinymdm-api-client/debug"
)

const DefaultBaseAPI = "https://www.tinymdm.net/api/v1/"

type Config struct {
	PublicKey string
	SecretKey string
	BaseAPI   string        // Optional; if empty, DefaultBaseAPI is used
	Timeout   time.Duration // Optional; if zero, 10s is used
}

type Client struct {
	publicKey  string
	secretKey  string
	baseAPI    string
	httpClient *http.Client
}

// New creates a new TinyMDM API client using the provided Config.
func New(cfg Config) *Client {
	baseAPI := cfg.BaseAPI
	if baseAPI == "" {
		baseAPI = DefaultBaseAPI
	}
	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}
	return &Client{
		publicKey:  cfg.PublicKey,
		secretKey:  cfg.SecretKey,
		baseAPI:    baseAPI,
		httpClient: &http.Client{Timeout: timeout},
	}
}

// DoRequest performs an HTTP request with context and returns the response body.
func (c *Client) DoRequest(ctx context.Context, method, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	var bodyLog string
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyLog = string(b)
		reqBody = bytes.NewBuffer(b)
	}
	url := c.baseAPI + endpoint
	debug.Debug("[TinyMDM] Request", "method", method, "url", url)
	if bodyLog != "" {
		debug.Debug("[TinyMDM] Request body", "body", bodyLog)
	}
	if c.publicKey != "" {
		debug.Debug("[TinyMDM] Using public key", "prefix", c.publicKey[:4])
	}
	if c.secretKey != "" {
		debug.Debug("[TinyMDM] Using secret key", "prefix", c.secretKey[:4])
	}
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Tinymdm-Apikey-Public", c.publicKey)
	req.Header.Set("X-Tinymdm-Apikey-Secret", c.secretKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	debug.Debug("[TinyMDM] Response body", "body", string(respBody))
	debug.Debug("[TinyMDM] Response status", "status", resp.Status)
	if resp.StatusCode >= 400 {
		debug.Debug("[TinyMDM] Error response body", "body", string(respBody))
		return nil, fmt.Errorf("API error: %s", respBody)
	}
	return respBody, nil
}

// HttpClient returns the underlying http.Client for advanced use.
func (c *Client) HttpClient() *http.Client {
	return c.httpClient
}
