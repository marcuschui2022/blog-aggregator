package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcuschui2022/blog-aggregator/internal/database"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]
	ctx := context.Background()
	now := time.Now().UTC()

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to find create feed's user: %w", err)
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(ctx, params)

	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	fmt.Println("feed created successfully!")
	fmt.Println("feed:", feed)
	return nil
}
