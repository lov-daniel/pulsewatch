package monitor

import (
	"encoding/json"
	"fmt"

	"daniellov.com/health-monitor/internal/ping"
	"daniellov.com/health-monitor/internal/types"
)

type Monitor struct {
	id       int64
	endpoint string
	history  []types.Result
}

var id int64 = 0

func SetMonitorConfig(encoder *json.Encoder) {
	encoder.SetIndent("", "    ")
}

func New(endpoint string) (*Monitor, bool) {
	id += 1

	return &Monitor{
		id:       id,
		endpoint: endpoint,
	}, false
}

// return value: error
func (m *Monitor) Run() (types.Result, bool) {
	result, err := ping.CheckEndpoint(m.endpoint)

	fmt.Printf("monitor (%d) @ (%v): [%s] %s in %d ms, err: %v \n", m.id, result.Timestamp, result.Status, result.Endpoint, result.LatencyMs, result.Err)
	if err {
		return types.Result{}, true
	}

	m.history = append(m.history, result)
	return result, false
}

func (m *Monitor) History() []types.Result {
	return m.history
}
