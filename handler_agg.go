package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dyn64/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage %s <time_between_reqs> (1s/1m/1h..)\n", cmd.Name)
	}

	time_between_reqs := cmd.Args[0]
	interval, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)

	scrapeFeeds(s)

	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

	// rss, err := fetchFeed(context.Background(), fetchURL)
	// if err != nil {
	// 	return fmt.Errorf("agg-error(s):\n%w\n", err)
	// }

	// fmt.Printf("rssfeed:\n%+v\n", rss)

	// return nil
}

func scrapeFeeds(s *state) error {
	fmt.Println("Start of scrapeFeeds")
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("getnextfail\n %w\n", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("markfeeedfail\n %w\n", err)
	}

	rss, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("fetchfail\n %w\n", err)
	}

	for _, post := range rss.Channel.Item {
		pubAt, err := time.Parse(time.RFC1123Z, post.PubDate)
		if err != nil {
			fmt.Printf("Error in published_at field. Skipping title: %s\nError:\n%v\n", post.Title, err)
			continue
		}
		pubTime := sql.NullTime{Time: pubAt, Valid: true}
		var postDescription sql.NullString
		if len(post.Description) != 0 {
			postDescription = sql.NullString{String: post.Description, Valid: true}
		} else {
			postDescription = sql.NullString{String: "", Valid: false}
		}

		postArgs := database.AddPostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       post.Title,
			Url:         post.Link,
			Description: postDescription,
			PublishedAt: pubTime,
			FeedID:      feed.ID,
		}

		err = s.db.AddPost(context.Background(), postArgs)
		if err != nil {
			//fmt.Printf("Error adding post: %v", err)
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code {
				case "23505":
					fmt.Printf("Duplicate entry\n")
					continue
				}
			}
			continue
		}
	}

	return nil
}
