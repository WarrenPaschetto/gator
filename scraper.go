package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/WarrenPaschetto/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("could not get next feed: %w", err)
	}

	fmt.Printf("Fetching feed: %s (%s)\n", feed.Name, feed.Url)

	rss, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching/parsing feed: %w", err)
	}

	for _, item := range rss.Channel.Item {
		publishedAt, err := parsePublishedTime(item.PubDate)
		if err != nil {
			fmt.Printf("Could not parse publish time for '%s': %v\n", item.Title, err)
			publishedAt = nil
		}

		var nullTime sql.NullTime
		if publishedAt != nil {
			nullTime = sql.NullTime{
				Time:  *publishedAt,
				Valid: true,
			}
		} else {
			nullTime = sql.NullTime{
				Valid: false,
			}
		}

		err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			PublishedAt: nullTime,
			FeedID:      feed.ID,
		})
		if err != nil {
			// If it's a unique conflict, we skip it quietly
			if isUniqueViolation(err) {
				continue
			}
			fmt.Printf("Error saving post '%s': %v\n", item.Title, err)
		}
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched %w", err)
	}

	fmt.Println("Done fetching feed.")
	return nil

}

func parsePublishedTime(dateStr string) (*time.Time, error) {
	layouts := []string{
		time.RFC1123Z,                     // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC1123,                      // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC822Z,                      // "02 Jan 06 15:04 -0700"
		time.RFC822,                       // "02 Jan 06 15:04 MST"
		time.RFC3339,                      // "2006-01-02T15:04:05Z07:00"
		"Mon, 02 Jan 2006 15:04:05 -0700", // custom fallback
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, dateStr); err == nil {
			return &t, nil
		}
	}

	return nil, fmt.Errorf("unsupported date format: %s", dateStr)
}

func isUniqueViolation(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
		return true
	}
	return false
}
