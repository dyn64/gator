package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dyn64/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <url>\n", cmd.Name)
	}

	url := cmd.Args[0]

	feed_data, err := s.db.ListFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Cannot find feed with url: %s\nerror:%w\n", url, err)
	}

	followArgs := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed_data.ID,
	}
	res, err := s.db.CreateFeedFollow(context.Background(), followArgs)
	if err != nil {
		return err
	}

	fmt.Printf("Current user: %s\n", res[0].UserName)
	fmt.Printf("Followed feed name: %s\n", res[0].FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
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

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <url>", cmd.Name)
	}
	unfollowparams := database.UnFollowByURLParams{
		Url:  cmd.Args[0],
		Name: user.Name,
	}

	err := s.db.UnFollowByURL(context.Background(), unfollowparams)
	if err != nil {
		return err
	}
	return nil
}
