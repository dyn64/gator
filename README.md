Gator is a guided project from Boot.dev -- "Build a Blog Aggregator"

Gator needs postgresql, go and goose (to set up the database)

Installation:
  go install https://github.com/dyn64/gator

Database:
The default name for the database is just "gator", just create one and run the migration scripts with goose.
Goose installation:
  go install github.com/pressly/goose/v3/cmd/goose@latest

After goose is installed, navigate to the sql/schemas directory and run the migrations:
  goose <DbURL> up


Gator uses a config file in the users home directory called .gatorconfig.json

If the file does not exist, a default config will be created:
	DbURL:           "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
	CurrentUserName: "secret",

The DbURL is the connection string for the database. The default one assumes the database is installed locally and is named 'gator'.

Usage:
go run . <command> <parameter>

Commands:
  register <username>
    Registers <username> in the database.
    Sets the current user as <username>

  login <username>
    Changes the current user to <username>

  reset
    Resets the database
    
  users
    Displays the registered users
    
  agg <interval>
    Starts the aggregator. Refreshes after <interval>
    Interval is in 1s/1m/1h etc
    
  addfeed <name> <url>
    Adds a feed with <name> and <url> to the database. Automatically follows the feed for the logged in user.
    
  feeds
    Displays all the feeds in the database
    
  follow <url>
    Follows the feed with <url>
    
  following
    Displays the feeds that the logged in user is following.
    
  unfollow <url>
    Unfollows the feed with <url>
  
  browse
    Displays all the posts the logged in user is following.
