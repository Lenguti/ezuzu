package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	v1 "github.com/lenguti/ezuzu/app/services/property/api/handlers/v1"
)

// Client - represents our client container.
type Client struct {
	cfg Config
	hc  *http.Client
}

// NewClient - intialized a new client with provided config.
func NewClient(cfg Config) *Client {
	return &Client{
		cfg: cfg,
		hc:  http.DefaultClient,
	}
}

// GetProperty - will call property service to fetch a property by the provided id.
func (c *Client) GetProperty(managerID, propertyID uuid.UUID) (v1.ClientProperty, error) {
	url := fmt.Sprintf("http://%s:%s/v1/managers/%s/properties/%s", c.cfg.Host, c.cfg.Port, managerID, propertyID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return v1.ClientProperty{}, fmt.Errorf("get property: unable to make request: %w", err)
	}
	resp, err := c.hc.Do(req)
	if err != nil {
		return v1.ClientProperty{}, fmt.Errorf("get property: unable to execute request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return v1.ClientProperty{}, fmt.Errorf("get property: invalid status code returned [%d]: %w", resp.StatusCode, err)
	}
	var out v1.GetPropertyResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return v1.ClientProperty{}, fmt.Errorf("get property: unable to decode reponse body: %w", err)
	}
	return out.Property, nil
}
