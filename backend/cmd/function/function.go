package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"daniellov.com/health-monitor/internal/monitor"
	"daniellov.com/health-monitor/internal/types"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloWorld", helloWorld)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
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
			fmt.Fprintf(w, "Error when creating monitor for endpoint: %s", endpoint)
			return
		}

		monitors = append(monitors, new_mon)
	}

	var wg sync.WaitGroup
	results := make(chan types.Result, len(endpoints))

	for _, mon := range monitors {
		wg.Add(1)

		go func(m *monitor.Monitor) {
			defer wg.Done()
			result, err := m.Run()

			if err {
				fmt.Printf("Error when running monitor: %d", m.Id())
				return
			}

			results <- result
		}(mon)
	}

	wg.Wait()
	close(results)

	for result := range results {
		json_val, err := json.Marshal(result)
		if err != nil {
			fmt.Fprint(w, "Error marshalling json")
			return
		}
		w.Write(json_val)
	}

	fmt.Println("Endpoints successfully written")
}
