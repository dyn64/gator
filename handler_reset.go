package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	//	err := s.db.TruncUsers(context.Background())
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("This house is clean\n")
	return nil
}
