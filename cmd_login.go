package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	usr, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		log.Fatalf("%s not found in database\n%v", username, err)
	}
	//s.conf.CurrentUserName = username

	err = s.conf.SetUser(usr.Name)
	if err != nil {
		return err
	}

	fmt.Printf("CurrentUserName set to %s \n", username)
	return nil
}
