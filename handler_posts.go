package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dyn64/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32
	if len(cmd.Args) != 1 {
		limit = 2
	} else {
		limi, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
		limit = int32(limi)
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}
	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		fmt.Printf("Error getting posts: %v\n", err)
		return err
	}

	for _, post := range posts {
		printPost(post)
	}
	return nil
}

func printPost(post database.Post) {
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("URL: %s\n\n", post.Url)
	if post.PublishedAt.Valid {
		fmt.Printf("Posted at: %v\n", post.PublishedAt.Time)
	}
	if post.Description.Valid {
		fmt.Printf("%s\n\n", post.Description.String)
	}
	fmt.Printf("Last updated: %v\n\n", post.UpdatedAt)
}
