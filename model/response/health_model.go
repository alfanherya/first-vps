package response

type HealthResponse struct {
	Database CheckDBResponse `json:"database"`
}

type CheckDBResponse struct {
	Idle              string `json:"idle"`
	InUse             string `json:"in_use"`
	MaxIdleClosed     string `json:"max_idle_closed"`
	MaxLifetimeClosed string `json:"max_lifetime_closed"`
	Message           string `json:"message"`
	OpenConnections   string `json:"open_connections"`
	Status            string `json:"status"`
	WaitCount         string `json:"wait_count"`
	WaitDuration      string `json:"wait_duration"`
}
