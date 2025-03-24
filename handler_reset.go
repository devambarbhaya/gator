package main

import (
	"context"
	"fmt"
)

func handlerReset( s* state, cmd command) error {
	err := s.db.DeleteUser(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("Database reset successful!")
	return nil
}