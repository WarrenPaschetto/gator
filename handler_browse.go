package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/WarrenPaschetto/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2)

	if len(cmd.Args) >= 1 {
		parsed, err := strconv.Atoi(cmd.Args[0])
		if err != nil || parsed < 1 {
			return fmt.Errorf("limit must be a positive number")
		}
		limit = int32(parsed)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("could not get posts: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts available")
		return nil
	}

	fmt.Printf("Showing up to %d recent posts:\n\n", limit)
	for _, post := range posts {
		fmt.Println("- ", post.Title)
		fmt.Println("- ", post.Url)
		fmt.Println()
	}

	return nil
}
