package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	_ = cmd
	ctx := context.Background()
	err := s.db.DeleteUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to reset users: %w", err)
	}

	fmt.Println("database reset successfully!")
	return nil
}
