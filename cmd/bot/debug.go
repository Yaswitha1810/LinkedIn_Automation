package main

import "github.com/go-rod/rod"

// EnableMouseDebug draws a red dot that follows mouse movements (debug only)
func EnableMouseDebug(page *rod.Page) {
	page.MustEval(`() => {
		if (document.getElementById("__mouse_debug_dot")) return;

		const dot = document.createElement("div");
		dot.id = "__mouse_debug_dot";
		dot.style.position = "fixed";
		dot.style.width = "10px";
		dot.style.height = "10px";
		dot.style.background = "red";
		dot.style.borderRadius = "50%";
		dot.style.zIndex = "999999";
		dot.style.pointerEvents = "none";

		document.body.appendChild(dot);

		document.addEventListener("mousemove", (e) => {
			dot.style.left = e.clientX + "px";
			dot.style.top = e.clientY + "px";
		});
	}`)
}
