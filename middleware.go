package main

import (
	"context"
	"fmt"

	"github.com/WarrenPaschetto/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		if s.cfg.CurrentUserName == "" {
			return fmt.Errorf("no user is currently logged in")
		}

		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("could not find user: %w", err)
		}

		return handler(s, cmd, user)
	}
}
