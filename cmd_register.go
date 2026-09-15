package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/dyn64/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	dbParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}
	usr, err := s.db.CreateUser(context.Background(), dbParams)
	if err != nil {
		log.Fatal(err)
	}
	err = s.conf.SetUser(usr.Name)
	fmt.Printf("New user %s created\n", usr.Name)
	fmt.Printf("ID: %v\nCreatedAt: %v\nUpdatedAt: %v\nName: %s\n", usr.ID, usr.CreatedAt, usr.UpdatedAt, usr.Name)

	return nil
}
