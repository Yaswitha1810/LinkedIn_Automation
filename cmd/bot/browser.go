package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func NewBrowser(headless bool) *rod.Browser {
	u := launcher.New().
		Leakless(false).
		Headless(headless).
		MustLaunch()

	return rod.New().ControlURL(u).MustConnect()
}
