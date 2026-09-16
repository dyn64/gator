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

	name := cmd.Args[0]
	dbParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}
	usr, err := s.db.CreateUser(context.Background(), dbParams)
	if err != nil {
		log.Fatal(err)
	}
	err = s.conf.SetUser(usr.Name)
	printUser(usr)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	usr, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		log.Fatalf("%s not found in database\n%v", username, err)
	}

	err = s.conf.SetUser(usr.Name)
	if err != nil {
		return err
	}

	fmt.Printf("CurrentUserName set to %s \n", username)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error listing users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.conf.CurrentUserName {
			fmt.Printf(" * %s (current)\n", user.Name)
			continue
		}
		fmt.Printf(" * %s\n", user.Name)

	}

	return nil
}

func printUser(user database.User) {
	fmt.Printf("------New user created--------------------------------\n")
	fmt.Printf(" * ID:\t\t%v\n", user.ID)
	fmt.Printf(" * Created at:\t%v\n", user.CreatedAt)
	fmt.Printf(" * Updated at:\t%v\n", user.UpdatedAt)
	fmt.Printf(" * Username:\t%v\n", user.Name)
	fmt.Printf("------------------------------------------------------\n")
}
