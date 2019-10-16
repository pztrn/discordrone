package message

import (
	// stdlib
	"errors"
	"log"

	// local
	"go.dev.pztrn.name/discordrone/env"

	// other
	"github.com/drone/drone-template-lib/template"
)

// Process starts message(s) processing.
func Process() error {
	if env.Data.Debug {
		log.Println("Preparing message for sending...")
	}

	// Webhook data should present.
	if env.Data.Plugin.Webhook.ID == "" && env.Data.Plugin.Webhook.Token == "" {
		return errors.New("webhook id or token is missing")
	}

	// Payload.
	p := &payload{
		Embeds: []EmbedObject{
			createEmbed(),
		},
	}

	// If no additional message was passed - send what we got.
	if len(env.Data.Plugin.Message) == 0 {
		err := sendMessage(p)
		if err != nil {
			return err
		}
	}

	// If message was set - format it.
	if len(env.Data.Plugin.Message) > 0 {
		text, err := template.RenderTrim(env.Data.Plugin.Message, env.Data.Drone)
		if err != nil {
			return err
		}
		p.Content = text
		err1 := sendMessage(p)
		if err1 != nil {
			return err1
		}
	}

	return nil
}
