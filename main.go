package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	
	"github.com/shaswatprakash/windows-agent/internal/models"
	
)

var (
	mu       sync.Mutex
	lastData models.HostData
)

func main() {
	// Start local ingestion and API server
	http.HandleFunc("/ingest", handleIngest)
	http.HandleFunc("/host", handleHost)

	fmt.Println("🚀 Local ingestion server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("❌ Server error:", err)
	}
}

// POST /ingest → called by agent to send data
func handleIngest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data models.HostData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	lastData = data
	mu.Unlock()

	saveToFile(data)
	fmt.Println("✅ Data ingested locally")
	w.Write([]byte(`{"message":"data ingested successfully"}`))
}

// GET /host → frontend fetches data here
func handleHost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	mu.Lock()
	defer mu.Unlock()

	// If no live data, load from file
	if lastData.Hostname == "" {
		data, err := loadFromFile()
		if err == nil {
			lastData = data
		}
	}

	json.NewEncoder(w).Encode(lastData)
}

// Save latest host data locally
func saveToFile(data models.HostData) {
	os.MkdirAll("data", os.ModePerm)
	bytes, _ := json.MarshalIndent(data, "", "  ")
	ioutil.WriteFile("data/data.json", bytes, 0644)
}

// Load latest host data from local file
func loadFromFile() (models.HostData, error) {
	bytes, err := ioutil.ReadFile("data/data.json")
	if err != nil {
		return models.HostData{}, err
	}
	var data models.HostData
	json.Unmarshal(bytes, &data)
	return data, nil
}
