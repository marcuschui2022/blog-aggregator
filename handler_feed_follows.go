package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcuschui2022/blog-aggregator/internal/database"
	"time"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	url := cmd.Args[0]

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	followFeed, err := s.db.CreateFeedFollow(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	fmt.Println("feed follow created successfully!")
	printFeedFollow(followFeed.UserName, followFeed.FeedName)

	return nil
}

func handlerListFeedFollows(s *state, cmd command) error {
	_ = cmd
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	ctx := context.Background()
	feedFollows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("feed follows for user %s:\n", user.Name)
	for _, feed := range feedFollows {
		fmt.Printf("* %s\n", feed.FeedName)
	}

	return nil
}

func printFeedFollow(userName, feedName string) {
	fmt.Printf("* User:          %s\n", userName)
	fmt.Printf("* Feed:          %s\n", feedName)
}
