package main

import (
	"fmt"
	"os"

	"github.com/dyn64/gator/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config\n%v\n", err)
		conf, err = config.DefaultConfig()
		if err != nil {
			fmt.Printf("Error creating default config: %v\n", err)
		}
		fmt.Printf("Using default config:\n dbUrl: %v\n current_username: %v\n", conf.DbURL, conf.CurrentUserName)

	}
	con := state{
		conf: &conf,
	}

	coms := commands{
		cmds: make(map[string]func(*state, command) error),
	}

	coms.register("login", handlerLogin)

	cmdArgs := os.Args
	if len(cmdArgs) < 2 {
		fmt.Printf("Error, missing missing command\n")
		os.Exit(1)
	}
	cmdName := cmdArgs[1]
	cmdArg := cmdArgs[2:]
	cmd := command{
		name: cmdName,
		args: cmdArg,
	}
	err = coms.run(&con, cmd)
	if err != nil {
		fmt.Printf("Error running command \"%s\"\n:", cmd.name)
		fmt.Println(err)
		os.Exit(1)
	}

}
