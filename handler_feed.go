package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcuschui2022/blog-aggregator/internal/database"
	"time"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	now := time.Now().UTC()

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

	feedFollow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}
	fmt.Println("feed created successfully!")
	printFeed(feed, user)
	fmt.Println()

	fmt.Println("feed followed successfully:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	_ = cmd
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("failed to get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("no feeds found")
		return nil
	}

	for _, feed := range feeds {
		printLiseFeedRow(feed)
	}
	return nil
}

func printLiseFeedRow(feed database.GetFeedsRow) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", feed.Name)
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
}
