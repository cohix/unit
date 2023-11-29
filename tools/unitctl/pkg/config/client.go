package config

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

// Client is a Unit config client
type Client struct {
	client http.Client
}

// PutConfigResp is the response returned when configuring Unit
type PutConfigResp struct {
	Success string `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

// NewClient creates a new Client
func NewClient(unitSocket string) *Client {
	// create an HTTP client that uses Unix domain sockets
	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", unitSocket)
			},
		},
	}

	c := &Client{
		client: client,
	}

	return c
}

// GetConfig gets the config from Unit
func (c *Client) GetConfig() (*Config, error) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost/config", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to http.NewRequest: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to client.Do: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET /config returned non-200: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to io.ReadAll: %w", err)
	}

	configResp := Config{}
	if err := json.Unmarshal(respBody, &configResp); err != nil {
		return nil, fmt.Errorf("failed to json.Unmarshal: %w", err)
	}

	return &configResp, nil
}

func (c *Client) ApplyConfig(config Config) (*PutConfigResp, error) {
	configJSON, err := json.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to json.Marshal: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, "http://localhost/config", bytes.NewBuffer(configJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to http.NewRequest: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to client.Do: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("PUT /config returned non-200: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to io.ReadAll: %w", err)
	}

	putResp := PutConfigResp{}
	if err := json.Unmarshal(respBody, &putResp); err != nil {
		return nil, fmt.Errorf("failed to json.Unmarshal: %w", err)
	}

	return &putResp, nil
}
