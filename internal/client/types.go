package client

type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Obj     T      `json:"obj"`
}

// Inbound represents a single inbound entry.
type Inbound struct {
	ID             int           `json:"id"`
	Up             int64         `json:"up"`
	Down           int64         `json:"down"`
	Total          int64         `json:"total"`
	Remark         string        `json:"remark"`
	Enable         bool          `json:"enable"`
	ExpiryTime     int64         `json:"expiryTime"`
	ClientStats    []ClientStats `json:"clientStats"`
	Listen         string        `json:"listen"`
	Port           int           `json:"port"`
	Protocol       string        `json:"protocol"`
	Settings       string        `json:"settings"`
	StreamSettings string        `json:"streamSettings"`
	Tag            string        `json:"tag"`
	Sniffing       string        `json:"sniffing"`
	Allocate       string        `json:"allocate"`
}

// ClientStats represents individual client statistics within an inbound.
type ClientStats struct {
	ID         int    `json:"id"`
	InboundID  int    `json:"inboundId"`
	Enable     bool   `json:"enable"`
	Email      string `json:"email"`
	Up         int64  `json:"up"`
	Down       int64  `json:"down"`
	ExpiryTime int64  `json:"expiryTime"`
	Total      int64  `json:"total"`
	Reset      int64  `json:"reset"`
}

// OnlinesResponse represents the `inbound/onlines` response object.
type OnlinesResponse []string
