package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration format (1s, 1m, 1h, etc.): %w", err)
	}

	fmt.Printf("Collecting feeds every %v...\n", timeBetweenRequests)

	// Run once immediately
	scrapeFeeds(s)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	// Loop: run scrapeFeeds every tick
	for range ticker.C {
		scrapeFeeds(s)
	}

	return nil
}
