package sender

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/shaswatprakash/windows-agent/internal/models"
)

func SendToAWS(data models.HostData, endpoint string) error {
	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
