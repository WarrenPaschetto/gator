package main

import (
	"context"
	"fmt"
	"time"

	"github.com/WarrenPaschetto/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <feed-url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not find feed: %w", err)
	}

	id := uuid.New()
	now := time.Now()

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not follow feed: %w", err)
	}

	fmt.Printf("%v is now following %v\n", follow.UserName, follow.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not retrieve follows: %w", err)
	}

	if len(follows) == 0 {
		fmt.Println("You are not currently following any feeds.")
		return nil
	}

	fmt.Println("You are following:")
	for _, f := range follows {
		fmt.Printf("* %s\n", f.FeedName)
	}

	return nil
}
