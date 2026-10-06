package types

import (
	"time"
)

type Result struct {
	Endpoint  string
	Status    string
	LatencyMs int64
	Timestamp time.Time
	Err       error
}
