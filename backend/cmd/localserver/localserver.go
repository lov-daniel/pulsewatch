package main

import (
	"log"
	"os"

	_ "daniellov.com/health-monitor/functions" // blank import registers HelloWorld via init()
	funcframework "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v", err)
	}
}
