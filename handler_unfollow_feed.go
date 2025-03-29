package main

import (
	"context"
	"fmt"

	"github.com/WarrenPaschetto/gator/internal/database"
)

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <feed-url>", cmd.Name)
	}

	url := cmd.Args[0]

	err := s.db.DeleteFeedFollowByUserAndURL(context.Background(), database.DeleteFeedFollowByUserAndURLParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return fmt.Errorf("could not unfollow feed: %w", err)
	}

	fmt.Printf("You have unfollowed %s\n", url)

	return nil
}
