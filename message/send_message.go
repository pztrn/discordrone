package message

import (
	// stdlib
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	// local
	"go.dev.pztrn.name/discordrone/env"
)

// Sends message to Discord.
func sendMessage(message interface{}) error {
	webhookURL := fmt.Sprintf("https://discordapp.com/api/webhooks/%s/%s", env.Data.Plugin.Webhook.ID, env.Data.Plugin.Webhook.Token)
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(message); err != nil {
		return err
	}
	_, err := http.Post(webhookURL, "application/json; charset=utf-8", b)

	if err != nil {
		return err
	}

	return nil
}
