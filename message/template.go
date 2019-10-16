package message

import (
	// stdlib
	"strconv"
	"strings"

	// local
	"go.dev.pztrn.name/discordrone/env"
)

func createEmbed() EmbedObject {
	// Create initial payload object.
	// Description.
	// Color.
	embed := EmbedObject{
		Title: env.Data.Drone.Commit.Message,
		URL:   env.Data.Drone.Build.Link,
		Author: EmbedAuthorObject{
			Name:    env.Data.Drone.Commit.Author.Name,
			IconURL: env.Data.Drone.Commit.Author.Avatar,
		},
	}

	// Compose description.
	var description string
	// ToDo: promote/rollback?
	switch env.Data.Drone.Build.Event {
	case "pull_request":
		var branch string
		if env.Data.Drone.Commit.Ref != "" {
			branch = env.Data.Drone.Commit.Ref
		} else {
			branch = env.Data.Drone.Commit.Branch
		}
		description = env.Data.Drone.Commit.Author.Name + " updated pull request " + branch
	case "push":
		description = env.Data.Drone.Commit.Author.Name + " pushed to " + env.Data.Drone.Commit.Branch
	case "tag":
		description = env.Data.Drone.Commit.Author.Name + " pushed tag " + env.Data.Drone.Commit.Branch
	}

	embed.Description = description

	// Compose color.
	var color int64
	if env.Data.Plugin.Color != "" {
		env.Data.Plugin.Color = strings.Replace(env.Data.Plugin.Color, "#", "", -1)
		s, err := strconv.ParseInt(env.Data.Plugin.Color, 16, 32)
		if err == nil {
			color = s
		}
	} else {

		switch env.Data.Drone.Build.Status {
		case "success":
			// green
			color = 0x1ac600
		case "failure", "error", "killed":
			// red
			color = 0xff3232
		default:
			// yellow
			color = 0xffd930
		}
	}

	embed.Color = color

	return embed
}
