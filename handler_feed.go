package main

import (
	"context"
	"fmt"
	"time"

	"github.com/WarrenPaschetto/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not auto-follow feed: %w", err)
	}

	fmt.Printf("Auto-followed feed: %s\n", follow.FeedName)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:            %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:            %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:            %s\n", feed.Name)
	fmt.Printf("* URL:            %s\n", feed.Url)
	fmt.Printf("* UserID:            %s\n", feed.UserID)
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not retrieve feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		fmt.Printf("* Name:        %s\n", feed.FeedName)
		fmt.Printf("* URL:        %s\n", feed.FeedUrl)
		fmt.Printf("* User:        %s\n", feed.UserName)
	}

	return nil
}
