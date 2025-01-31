package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	//_ = s
	//if len(cmd.Args) != 1 {
	//	return fmt.Errorf("usage: %v <url>", cmd.Name)
	//
	//}
	//_ = cmd.Args[0]

	ctx := context.Background()
	feed, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)

	return nil
}
