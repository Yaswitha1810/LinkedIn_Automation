package main

import (
	"fmt"

	"github.com/go-rod/rod/lib/proto"
)

func main() {
	// --------------------------------------------------
	// DEMO LOGIN (SAFE, VISUAL)
	// --------------------------------------------------
	fmt.Println("Starting demo login flow...")

	br := NewBrowser(false)
	page := br.MustPage("https://example.com")
	page.MustWaitLoad()

	// Inject fake login UI
	page.MustEval(`() => {
		const box = document.createElement("div");
		box.innerHTML =
			"<h2>Demo Login</h2>" +
			"<input id='email' placeholder='Email' /> <br/><br/>" +
			"<input id='password' type='password' placeholder='Password' /> <br/><br/>" +
			"<button id='login'>Sign In</button>";
		box.style.margin = "40px";
		box.style.fontSize = "18px";
		document.body.appendChild(box);
	}`)

	email := page.MustElement("#email")
	email.Click(proto.InputMouseButtonLeft, 1)
	email.Focus()
	HumanType(email, "demo@example.com")

	password := page.MustElement("#password")
	password.Click(proto.InputMouseButtonLeft, 1)
	password.Focus()
	HumanType(password, "********")

	page.MustElement("#login").Click(proto.InputMouseButtonLeft, 1)

	fmt.Println("Demo login completed.\n")

	// --------------------------------------------------
	// LOAD STATE
	// --------------------------------------------------
	state, err := LoadState()
	if err != nil {
		panic(err)
	}

	fmt.Println("State loaded")
	fmt.Println("Daily count:", state.DailyCount)

	// --------------------------------------------------
	// SEARCH (DEMO)
	// --------------------------------------------------
	criteria := SearchCriteria{
		Keyword:  "Software Engineer",
		Location: "India",
		Company:  "DemoCorp",
		MaxPages: 2,
	}

	profiles := RunSearchDemo(criteria, state)

	// --------------------------------------------------
	// MESSAGING
	// --------------------------------------------------
	maxMessagesPerDay := 3

	for _, profileURL := range profiles {
		if state.DailyCount >= maxMessagesPerDay {
			fmt.Println("Daily messaging limit reached.")
			break
		}

		if _, sent := state.SentMessages[profileURL]; sent {
			fmt.Println("Already messaged:", profileURL)
			continue
		}

		vars := map[string]string{
			"name":    "Demo User",
			"role":    "Engineer",
			"company": "DemoCorp",
		}

		message := renderTemplate(initialMessageTemplate, vars)
		SendMessageDemo(profileURL, message, state)
		fmt.Println("Message sent.\n")
	}

	// --------------------------------------------------
	// SAVE STATE
	// --------------------------------------------------
	if err := SaveState(state); err != nil {
		panic(err)
	}

	fmt.Println("Demo completed. State saved.")
	select {} // keep browser open for demo
}
