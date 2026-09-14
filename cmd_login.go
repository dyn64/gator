package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Error: handlerLogin needs 1 argument")
	}
	username := cmd.args[0]
	s.conf.CurrentUserName = username
	fmt.Printf("CurrentUserName set to %s \n", username)
	s.conf.SetUser(username)

	return nil
}
