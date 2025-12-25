package main

import (
	"fmt"
	"time"
)

// Simulated search runner
func RunSearchDemo(criteria SearchCriteria, state *BotState) []string {
	fmt.Println("Starting search with criteria:")
	fmt.Printf("Keyword=%s | Location=%s | Company=%s\n",
		criteria.Keyword, criteria.Location, criteria.Company)

	collected := []string{}

	for page := 1; page <= criteria.MaxPages; page++ {
		fmt.Printf("\nProcessing page %d...\n", page)

		// Simulate delay between pages
		HumanDelay(800, 1500)

		// Fake profiles returned per page
		pageResults := fakeProfilesForPage(page)

		for _, profileURL := range pageResults {
			// Deduplication
			if _, seen := state.SentConnections[profileURL]; seen {
				fmt.Println("Skipping duplicate:", profileURL)
				continue
			}

			fmt.Println("Found profile:", profileURL)
			collected = append(collected, profileURL)

			// Mark as seen in state (even before sending)
			state.SentConnections[profileURL] = time.Now()
		}
	}

	return collected
}

// Fake paginated results
func fakeProfilesForPage(page int) []string {
	switch page {
	case 1:
		return []string{
			"https://linkedin.com/in/demo-1",
			"https://linkedin.com/in/demo-2",
			"https://linkedin.com/in/demo-3",
		}
	case 2:
		return []string{
			"https://linkedin.com/in/demo-2", // duplicate
			"https://linkedin.com/in/demo-4",
			"https://linkedin.com/in/demo-5",
		}
	case 3:
		return []string{
			"https://linkedin.com/in/demo-5", // duplicate
			"https://linkedin.com/in/demo-6",
		}
	default:
		return []string{}
	}
}
