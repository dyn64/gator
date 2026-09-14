package main

import "github.com/dyn64/gator/internal/config"

type state struct {
	conf *config.Config
}

// maybe use ...string here?
type command struct {
	name string
	args []string
}
