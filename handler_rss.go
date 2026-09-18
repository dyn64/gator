package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dyn64/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	fetchURL := "https://www.wagslane.dev/index.xml"
	rss, err := fetchFeed(context.Background(), fetchURL)
	if err != nil {
		return fmt.Errorf("agg-error(s):\n%w\n", err)
	}

	fmt.Printf("rssfeed:\n%+v\n", rss)

	return nil
}

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage %s <name> <url>\n", cmd.Name)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return err
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	feedparams := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	}
	feed, err := s.db.AddFeed(context.Background(), feedparams)
	if err != nil {
		return err
	}
	printFeed(feed)

	err = handlerFollow(s, command{Name: "", Args: []string{url}})
	if err != nil {
		return err
	}

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	listfeeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting feeds: %w\n", err)
	}

	if len(listfeeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n\n", len(listfeeds))
	for i, row := range listfeeds {
		fmt.Printf("- %d ----\n", i)
		fmt.Printf("Name:\t%s\n", row.Name)
		fmt.Printf("URL:\t%s\n", row.Url)
		fmt.Printf("Owner:\t%s\n", row.Username)
		fmt.Printf("--------\n\n")
	}
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("-------New feed created-------------------------------\n")
	fmt.Printf(" * ID:\t\t%v\n", feed.ID)
	fmt.Printf(" * Created at:\t%v\n", feed.CreatedAt)
	fmt.Printf(" * Updated at:\t%v\n", feed.UpdatedAt)
	fmt.Printf(" * Name:\t%v\n", feed.Name)
	fmt.Printf(" * URL:\t\t%v\n", feed.Url)
	fmt.Printf(" * CreatorID:\t%v\n", feed.UserID)
	fmt.Printf("------------------------------------------------------\n")
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <url>\n", cmd.Name)
	}

	user := s.conf.CurrentUserName
	url := cmd.Args[0]

	user_data, err := s.db.GetUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("user %s not found\nerror:%w\n", user, err)
	}
	feed_data, err := s.db.ListFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Cannot find feed with url: %s\nerror:%w\n", url, err)
	}

	user_id := user_data.ID
	feed_id := feed_data.ID

	followArgs := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user_id,
		FeedID:    feed_id,
	}
	res, err := s.db.CreateFeedFollow(context.Background(), followArgs)
	if err != nil {
		return err
	}

	fmt.Printf("Current user: %s\n", res[0].UserName)
	fmt.Printf("Followed feed name: %s\n", res[0].FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return err
	}
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	if len(feeds) == 0 {
		fmt.Printf("User %s is not following any feeds\n", user.Name)
	}

	fmt.Printf("Feeds followed by %s:\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf(" * %s\n", feed)
	}

	return nil
}
