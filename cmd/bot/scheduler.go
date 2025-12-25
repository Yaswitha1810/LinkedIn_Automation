package main

import (
	"math/rand"
	"time"
)

// Check if current time is within working hours
func withinBusinessHours() bool {
	now := time.Now()
	hour := now.Hour()
	return hour >= 0 && hour <= 18
}

// Random cooldown between actions
func cooldown() {
	delay := rand.Intn(10)+10 // 10–20 seconds
	time.Sleep(time.Duration(delay) * time.Second)
}
