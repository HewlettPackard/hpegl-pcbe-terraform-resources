package defaults

import "time"

const (
	// Time in seconds between polling for updates of resources
	PollInterval = float32(5.0)
	// Limit on number of polls for a resource
	MaxPolls = int32(100)
	// http client timeout
	HTTPTimeout = 60 * time.Second
)
