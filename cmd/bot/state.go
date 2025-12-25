package main

import (
	"encoding/json"
	"os"
	"time"
)

const stateFile = "state.json"

type BotState struct {
	SentConnections map[string]time.Time `json:"sent_connections"`
	SentMessages    map[string]time.Time `json:"sent_messages"`
	DailyCount      int                  `json:"daily_count"`
	LastReset       string               `json:"last_reset"`
}

// Load state from disk or create new
func LoadState() (*BotState, error) {
	if _, err := os.Stat(stateFile); os.IsNotExist(err) {
		return newState(), nil
	}

	data, err := os.ReadFile(stateFile)
	if err != nil {
		return nil, err
	}

	var state BotState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}

	state.resetIfNewDay()
	return &state, nil
}

// Save state to disk
func SaveState(state *BotState) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(stateFile, data, 0644)
}

func newState() *BotState {
	return &BotState{
		SentConnections: make(map[string]time.Time),
		SentMessages:    make(map[string]time.Time),
		DailyCount:      0,
		LastReset:       today(),
	}
}

func (s *BotState) resetIfNewDay() {
	if s.LastReset != today() {
		s.DailyCount = 0
		s.LastReset = today()
	}
}

func today() string {
	return time.Now().Format("2006-01-02")
}
