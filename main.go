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
		os.Exit(0)
	}
	err = conf.SetUser("")
	if err != nil {
		fmt.Printf("Error setting/writing username\n%v\n", err)
		os.Exit(0)
	}

	conf, err = config.Read()
	if err != nil {
		fmt.Printf("Error reading config again\n%v\n", err)
		os.Exit(0)
	}
	fmt.Printf("dbUrl: %v\n", conf.DbURL)
	fmt.Printf("current_username: %v\n", conf.CurrentUserName)
}
