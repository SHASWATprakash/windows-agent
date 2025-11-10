package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/shaswatprakash/windows-agent/internal/collector"
	"github.com/shaswatprakash/windows-agent/internal/models"
	"github.com/shaswatprakash/windows-agent/internal/sender"
)

func main() {
	// Collect Hostname, Installed Apps, and CIS Checks
	hostname, _ := os.Hostname()
	apps, _ := collector.GetInstalledApps()
	checks := collector.RunCISChecks()

	data := models.HostData{
		Hostname:     hostname,
		Applications: apps,
		CISChecks:    checks,
	}

	// Print to console for debugging
	fmt.Println("=== Host Data ===")
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(jsonData))

	// Optional: Send to AWS (mock endpoint for now)
	endpoint := "https://your-api-endpoint.amazonaws.com/dev/ingest"
	if err := sender.SendToAWS(data, endpoint); err != nil {
		fmt.Println("❌ Error sending to AWS:", err)
	} else {
		fmt.Println("✅ Data sent to AWS successfully (mock)")
	}

	// REST API endpoint
	http.HandleFunc("/host", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") // allow frontend access
		json.NewEncoder(w).Encode(data)
	})

	// Start local server
	fmt.Println("🚀 Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("❌ Server error:", err)
	}
}
