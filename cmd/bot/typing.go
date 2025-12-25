package main

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// HumanType types text like a real human (variable speed + typos)
func HumanType(el *rod.Element, text string) {
	for _, ch := range text {

		// ~1% chance to make a typo
		if rand.Float64() < 0.01 {
			el.Input(string(randomChar()))
			time.Sleep(randomDelay(80, 150))

			// Backspace (Rod-compatible)
			el.Input("\b")
			time.Sleep(randomDelay(80, 150))
		}

		// Type correct character
		el.Input(string(ch))

		// Random delay between keystrokes
		time.Sleep(randomDelay(60, 180))
	}
}

func randomChar() rune {
	return rune('a' + rand.Intn(26))
}

func randomDelay(minMs, maxMs int) time.Duration {
	return time.Duration(rand.Intn(maxMs-minMs)+minMs) * time.Millisecond
}
