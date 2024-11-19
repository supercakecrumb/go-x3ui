package client

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// AddClient adds a new client to the specified inbound
func (c *Client) AddClient(inboundID int, clientConfig InboundClient) error {
	c.Logger.Info("Adding client to inbound", "inboundID", inboundID, "email", clientConfig.Email)

	// Marshal the clients into JSON
	clientsJSON, err := json.Marshal(AddInboundClientConfig{
		Clients: []InboundClient{clientConfig},
	})
	if err != nil {
		c.Logger.Error("Failed to marshal clients to JSON", "error", err)
		return fmt.Errorf("failed to marshal clients: %w", err)
	}

	// Create the payload with `settings` as a JSON-encoded string
	payload := map[string]interface{}{
		"id":       inboundID,
		"settings": string(clientsJSON),
	}

	// Send the request
	resp, err := c.Resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/panel/inbound/addClient")

	if err != nil {
		c.Logger.Error("Failed to add client", "error", err)
		return fmt.Errorf("failed to add client: %w", err)
	}

	// Unmarshal the response
	var response APIResponse[interface{}]
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		c.Logger.Error("Failed to unmarshal response", "error", err, "body", string(resp.Body()))
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Handle non-success responses
	if !response.Success {
		c.Logger.Error("Add client failed", "message", response.Msg)
		return fmt.Errorf("add client failed: %s", response.Msg)
	}

	c.Logger.Info("Client successfully added to inbound", "inboundID", inboundID, "email", clientConfig.Email)
	return nil
}

// GenerateDefaultConfig creates a default client configuration for adding a client
func (c *Client) GenerateDefaultConfig(email string, tgID int64) InboundClient {
	return InboundClient{
		ID:         uuid.NewString(),
		Flow:       "xtls-rprx-vision",
		Email:      email,
		LimitIP:    0,
		TotalGB:    0,
		ExpiryTime: 0,
		Enable:     true,
		TgID:       tgID,
		SubID:      uuid.NewString(),
		Reset:      0,
	}
}