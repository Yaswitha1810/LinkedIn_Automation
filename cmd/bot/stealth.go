package main

import (
	"math/rand"
	"time"
)

func HumanDelay(minMs, maxMs int) {
	delay := rand.Intn(maxMs-minMs) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
