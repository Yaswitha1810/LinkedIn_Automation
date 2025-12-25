package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// MoveMouseHumanLike moves the mouse in a curved, human-like path
func MoveMouseHumanLike(page *rod.Page, startX, startY, endX, endY float64) {
	steps := rand.Intn(20) + 30 // number of micro movements

	// Random control point for Bezier curve
	controlX := (startX+endX)/2 + rand.Float64()*100 - 50
	controlY := (startY+endY)/2 + rand.Float64()*100 - 50

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)

		// Quadratic Bezier curve formula
		x := math.Pow(1-t, 2)*startX +
			2*(1-t)*t*controlX +
			math.Pow(t, 2)*endX

		y := math.Pow(1-t, 2)*startY +
			2*(1-t)*t*controlY +
			math.Pow(t, 2)*endY

		// Move mouse to next point
		page.Mouse.MoveTo(proto.Point{
			X: x,
			Y: y,
		})

		// Tiny delay between movements
		time.Sleep(time.Duration(rand.Intn(8)+5) * time.Millisecond)
	}
}
