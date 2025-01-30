package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("username is required")
	}
	username := cmd.Args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set current user: %w", err)
	}

	fmt.Println("user switched successfully!")
	return nil
}
