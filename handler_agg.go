package main

import (
	"context"
	"fmt"
	"time"
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

	for _, title := range rss.Channel.Item {
		fmt.Printf("Title: %s\n", title.Title)
	}

	return nil
}
