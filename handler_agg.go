package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/marcuschui2022/blog-aggregator/internal/database"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	_ = s
	_ = cmd

	feedUrl, err := scrapeFeed(s)

	ctx := context.Background()
	feed, err := fetchFeed(ctx, feedUrl)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	fmt.Printf("Feed Title: %s\n", feed.Channel.Title)
	fmt.Printf("Feed Link: %s\n", feed.Channel.Link)
	fmt.Printf("Feed Description: %s\n", feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Printf("Item Title: %s\n", item.Title)
		fmt.Printf("Item Link: %s\n", item.Link)
		fmt.Printf("Item Description: %s\n", item.Description)
		fmt.Printf("Item PubDate: %s\n", item.PubDate)
	}
	return nil
}

func scrapeFeed(s *state) (string, error) {
	nextFeedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get next feed to fetch: %w", err)
	}

	now := time.Now().UTC()
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: now, Valid: true},
		UpdatedAt:     now,
		ID:            nextFeedToFetch.ID,
	})

	if err != nil {
		return "", fmt.Errorf("failed to mark feed fetched: %w", err)
	}
	return nextFeedToFetch.Url, nil

}
