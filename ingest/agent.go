package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shaswatprakash/windows-agent/internal/collector"
	"github.com/shaswatprakash/windows-agent/internal/models"
	"github.com/shaswatprakash/windows-agent/internal/sender"
)

func main() {
	hostname, _ := os.Hostname()
	apps, _ := collector.GetInstalledApps()
	checks := collector.RunCISChecks()

	data := models.HostData{
		Hostname:     hostname,
		Applications: apps,
		CISChecks:    checks,
	}

	fmt.Println("=== Sending data to local ingestion ===")
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(jsonData))

	if err := sender.SendToLocalIngestion(data); err != nil {
		fmt.Println("❌ Failed to send data:", err)
	} else {
		fmt.Println("✅ Data sent to local ingestion successfully")
	}
}
