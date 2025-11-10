package sender

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/shaswatprakash/windows-agent/internal/models"
)

func SendToLocalIngestion(data models.HostData) error {
	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "http://localhost:8080/ingest", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
