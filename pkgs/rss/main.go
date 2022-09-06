package rss

import (
	"context"
	"log"

	"encoding/json"

	"github.com/ec965/rss-server/pkgs/models"
	"github.com/mmcdole/gofeed"
)

// GetFeed gets an RSS feed from the given url and returns it as a json string
func GetFeed(ctx context.Context, url string) ([]byte, error) {
	feed, err := gofeed.NewParser().ParseURLWithContext(url, ctx)

	if err != nil {
		return nil, err
	}

	j, err := json.Marshal(feed)

	return j, err
}

// UpdateAllFeeds updates all rss feeds in the database
func UpdateAllFeeds() error {
	urls, err := models.SelectAllRSSUrls(context.TODO())

	if err != nil {
		return err
	}

	for _, url := range urls {
		go func(ctx context.Context, url *models.RSSUrl) {

			j, err := GetFeed(ctx, url.Url)
			if err != nil {
				log.Println(err)
				return
			}

			if err = models.UpdateRSSFeedById(ctx, url.Id, string(j)); err != nil {
				log.Println(err)
			}
		}(context.TODO(), &url)
	}

	return nil
}
