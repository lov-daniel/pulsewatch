package event_logger

import (
	"encoding/json"
	"fmt"
	"os"

	"daniellov.com/health-monitor/internal/types"
)

type EventLogger struct {
	id      int64
	encoder *json.Encoder
	file    *os.File
	history []string
}

var LOG_FILE = "event_logs.json"
var id int64 = 0

func (e *EventLogger) SetupLogConfigs() {
	e.encoder.SetIndent("", "    ")
}

func New(filename string) (*EventLogger, bool) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Printf("Error when creating file: %e", err)
		return &EventLogger{}, true
	}
	id += 1

	encoder := json.NewEncoder(file)

	event_logger := EventLogger{
		id:      id,
		encoder: encoder,
		file:    file,
	}

	event_logger.SetupLogConfigs()
	return &event_logger, false
}

func (e *EventLogger) Write(message types.Result) bool {
	json, err := json.Marshal(message)
	fmt.Printf("event logger (%d): %s \n", e.id, json)
	err = e.encoder.Encode(message)

	if err != nil {
		fmt.Printf("event logger (%d): %e", e.id, err)
		return true
	}

	return false
}

func (e *EventLogger) Clean() {
	e.file.Close()
}
