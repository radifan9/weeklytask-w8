package utils

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Simulates fetching data from a URL with random delay
// chan<- : means send-only channel (one-directional)
func WebFetcher(url string, wg *sync.WaitGroup, urlResults chan<- string) {
	// Ensures wg is decremented when functions exit
	defer wg.Done()

	// Simulate network latency with random delay (1-5 seconds)
	time.Sleep(time.Duration((rand.IntN(5) + 1)) * time.Second)

	// Create result message
	result := fmt.Sprintf("Fetched: %s", url)

	// Send result through channel
	urlResults <- result
}
