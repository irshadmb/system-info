package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/irshadmb/system-info/internal/collectors"
)

func main() {
	// Initialize collector
	collector := collectors.NewSystemCollector()

	// Collect system information
	sysInfo, err := collector.CollectAll()
	if err != nil {
		log.Fatal("Error collecting system information:", err)
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(sysInfo, "", "    ")
	if err != nil {
		log.Fatal("Error converting to JSON:", err)
	}

	// Print JSON output
	fmt.Println(string(jsonData))
}