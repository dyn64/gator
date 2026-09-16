package main

import (
	"context"
	"fmt"
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
