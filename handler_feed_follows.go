package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcuschui2022/blog-aggregator/internal/database"
	"time"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	url := cmd.Args[0]

	ctx := context.Background()

	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("failed to get feed: %w", err)
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

func handlerListFeedFollows(s *state, cmd command, user database.User) error {
	_ = cmd

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("feed follows for user %s:\n", user.Name)
	for _, feed := range feedFollows {
		fmt.Printf("* %s\n", feed.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <feed_url>", cmd.Name)
	}

	ctx := context.Background()
	feedURL := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	err = s.db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}
	fmt.Printf("%s unfollowed successfully!\n", feed.Name)

	return nil
}

func printFeedFollow(userName, feedName string) {
	fmt.Printf("* User:          %s\n", userName)
	fmt.Printf("* Feed:          %s\n", feedName)
}
