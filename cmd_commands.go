package main

import "fmt"

type commands struct {
	cmds map[string]func(*state, command) error
}

// run given command if it exists
func (c *commands) run(s *state, cmd command) error {
	runme, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("%s not found\n", cmd.name)
	}
	err := runme(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

// register a new handler function for a given command
func (c *commands) register(name string, f func(*state, command) error) error {
	_, ok := c.cmds[name]
	if ok {
		return fmt.Errorf("Command \"%s\" exists already\n", name)
	}
	c.cmds[name] = f

	return nil
}
