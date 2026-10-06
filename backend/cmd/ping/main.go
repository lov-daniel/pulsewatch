package main

import (
	"fmt"
	"sync"

	"daniellov.com/health-monitor/internal/event_logger"
	"daniellov.com/health-monitor/internal/monitor"
	"daniellov.com/health-monitor/internal/types"
)

func main() {

	var monitors []*monitor.Monitor

	endpoints := []string{
		"https://www.google.com",
		"https://www.caida.org",
		"https://www.daniellov.com",
		"https://fakeurl.com",
	}

	for _, endpoint := range endpoints {

		new_mon, err := monitor.New(endpoint)

		if err {
			return
		}

		monitors = append(monitors, new_mon)
	}

	logger, err := event_logger.New("event_logs.json")

	if err {
		return
	}

	defer logger.Clean()

	var wg sync.WaitGroup
	results := make(chan types.Result, len(endpoints))

	for _, mon := range monitors {
		wg.Add(1)

		go func(m *monitor.Monitor) {
			defer wg.Done()
			result, err := m.Run()

			if err {
				return
			}

			results <- result
		}(mon)

		if err {
			return
		}
	}

	wg.Wait()
	close(results)

	for result := range results {
		logger.Write(result)
	}

	fmt.Println("Endpoints successfully written")
}
