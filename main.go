package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dyn64/gator/internal/config"
	"github.com/dyn64/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	conf *config.Config
	db   *database.Queries
}

func main() {
	// read config, if it doesnt exist. create a new one from Defaultconfig
	conf, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config\n%v\n", err)
		conf, err = config.DefaultConfig()
		if err != nil {
			fmt.Printf("Error creating default config: %v\n", err)
		}
		fmt.Printf("Using default config:\n dbUrl: %v\n current_username: %v\n", conf.DbURL, conf.CurrentUserName)

	}

	// open a connection to the database using the dburl string
	db, err := sql.Open("postgres", conf.DbURL)
	if err != nil {
		log.Fatalf("Error opening db %v", err)
	}
	defer db.Close()

	// save it to the state struct
	dbQueries := database.New(db)

	// saves the config+db into the state struct, with a very descriptive name 'con'
	con := state{
		conf: &conf,
		db:   dbQueries,
	}

	// initializes the commands map
	coms := commands{
		cmds: make(map[string]func(*state, command) error),
	}
	// register commands
	coms.register("login", handlerLogin)
	coms.register("register", handlerRegister)
	coms.register("reset", handlerReset)
	coms.register("users", handlerUsers)

	// grabs the command line arguments ex "go run . blabla xx" -> blabla xx
	cmdArgs := os.Args
	if len(cmdArgs) < 2 {
		fmt.Printf("Error, missing missing command\n")
		os.Exit(1)
	}
	cmdName := cmdArgs[1]
	cmdArg := cmdArgs[2:]

	// tries to execute the command
	err = coms.run(&con, command{
		Name: cmdName,
		Args: cmdArg,
	})
	// failure
	if err != nil {
		fmt.Printf("Error running command \"%s\"\n:", cmdName)
		fmt.Println(err)
		os.Exit(1)
	}

}
