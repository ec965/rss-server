package rss

import (
	"context"

	"github.com/mmcdole/gofeed"
  "encoding/json"
)

var fp *gofeed.Parser

func init() {
  fp = gofeed.NewParser()
}

func GetFeed(ctx context.Context, url string) ([]byte, error) {
  feed, err := fp.ParseURLWithContext(url, ctx)

  if err != nil {
    return nil, err
  }

  j, err := json.Marshal(feed)

  return j, err
}
