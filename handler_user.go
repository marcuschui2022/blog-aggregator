package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcuschui2022/blog-aggregator/internal/database"
	"time"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)

	}
	username := cmd.Args[0]

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, username)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set current user: %w", err)
	}

	fmt.Println("user switched successfully!")
	fmt.Println("current user:", user.Name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	username := cmd.Args[0]
	ctx := context.Background()
	now := time.Now().UTC()

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      username,
	}

	user, err := s.db.CreateUser(ctx, params)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("failed to set current user: %w", err)
	}

	fmt.Println("user created successfully!")
	printUser(user)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}

func handlerGetUsers(s *state, cmd command) error {
	_ = cmd
	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
