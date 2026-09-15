package main

import (
	"github.com/dyn64/gator/internal/config"
	"github.com/dyn64/gator/internal/database"
)

type state struct {
	conf *config.Config
	db   *database.Queries
}

// maybe use ...string here?
type command struct {
	Name string
	Args []string
}
