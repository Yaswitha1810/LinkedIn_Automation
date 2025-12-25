package main

import (
	"fmt"
	"strings"
	"time"
)

type MessageTemplate struct {
	Name    string
	Content string
}

var initialMessageTemplate = MessageTemplate{
	Name: "initial_connection",
	Content: "Hi {{name}}, I came across your profile and noticed your work as a {{role}} at {{company}}. Would love to connect!",
}

var followUpTemplate = MessageTemplate{
	Name: "follow_up",
	Content: "Hi {{name}}, just following up on my earlier message. Looking forward to connecting!",
}

// Render template with variables
func renderTemplate(t MessageTemplate, vars map[string]string) string {
	msg := t.Content
	for k, v := range vars {
		msg = strings.ReplaceAll(msg, "{{"+k+"}}", v)
	}
	return msg
}

// Send message (demo-safe)
func SendMessageDemo(profileURL string, message string, state *BotState) {
	fmt.Println("Sending message to:", profileURL)
	fmt.Println("Message:", message)

	// Simulate delay
	HumanDelay(800, 1400)

	// Record message in state
	state.SentMessages[profileURL] = time.Now()
	state.DailyCount++
}
