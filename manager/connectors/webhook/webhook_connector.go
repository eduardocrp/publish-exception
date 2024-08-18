package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"publish-expcetion/manager/connectors"
)

const webhookUrl = "https://hooks.slack.com/services/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"

func SendMessage(url string, message string) error {
	requestBody := connectors.Message{
		Content: message,
	}
	messageBytes, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("erro ao marshalling a mensagem: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(messageBytes))
	if err != nil {
		return fmt.Errorf("erro ao enviar a mensagem: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro ao enviar a mensagem, status: %s", resp.Status)
	}

	return nil
}