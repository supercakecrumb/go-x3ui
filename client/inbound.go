package client

import (
	"fmt"
)

// ListInbounds fetches the list of inbounds.
func (c *Client) ListInbounds() ([]Inbound, error) {
	c.Logger.Info("Fetching inbound list")
	var response APIResponse[[]Inbound]

	// Send the request
	_, err := c.Resty.R().
		SetHeader("Accept", "application/json").
		SetResult(&response).
		Post("/panel/inbound/list")

	if err != nil {
		c.Logger.Error("Failed to fetch inbound list", "error", err)
		return nil, err
	}

	// Handle non-success responses
	if !response.Success {
		c.Logger.Error("Failed to fetch inbounds", "message", response.Msg)
		return nil, fmt.Errorf("failed to fetch inbounds: %s", response.Msg)
	}

	c.Logger.Info("Successfully fetched inbound list")
	return response.Obj, nil
}

// GetOnlineClients fetches the list of online clients.
func (c *Client) GetOnlineClients() (OnlinesResponse, error) {
	c.Logger.Info("Fetching online clients")
	var response APIResponse[OnlinesResponse]

	// Send the request
	_, err := c.Resty.R().
		SetHeader("Accept", "application/json").
		SetResult(&response).
		Post("/panel/inbound/onlines")

	if err != nil {
		c.Logger.Error("Failed to fetch online clients", "error", err)
		return nil, err
	}

	// Handle non-success responses
	if !response.Success {
		c.Logger.Error("Failed to fetch online clients", "message", response.Msg)
		return nil, fmt.Errorf("failed to fetch online clients: %s", response.Msg)
	}

	c.Logger.Info("Successfully fetched online clients")
	return response.Obj, nil
}
