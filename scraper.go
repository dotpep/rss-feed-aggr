package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/dotpep/rss-feed-aggr/internal/database"
	"github.com/google/uuid"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)

	// make once immediately
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedListToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds: ", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			// decrements counter of wg by one
			// and spawn len(feed) goroutines to scrape all of them in same time
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}
}

func scrapeFeed(
	db *database.Queries,
	wg *sync.WaitGroup,
	feed database.Feed,
) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched: ", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed: ", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		//log.Println("Fount post", item.Title, "on feed", feed.Name)

		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		formattedPubAt, err := parsePubDate(item.PubDate)
		if err != nil {
			log.Printf(
				"couldn't parse date %v of %v post with error %v",
				item.PubDate, item.Title, err,
			)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: formattedPubAt,
			Url:         item.Link,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("failed to create post: ", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}

var (
	formatTimeLayouts = []string{
		time.RFC1123Z, // "Mon, 29 Aug 2022 00:00:00 +0000"
		time.RFC1123,  // "Sat, 01 Jun 2024 14:24:27 GMT"
		// TODO: add other formats as you need in feed aggr
	}
)

func parsePubDate(pubDate string) (time.Time, error) {
	var parsedTime time.Time
	var err error

	for _, layout := range formatTimeLayouts {
		parsedTime, err = time.Parse(layout, pubDate)
		if err == nil {
			return parsedTime, nil
		}
	}
	return time.Time{}, err
}
