package models

import "context"

// Data of RSS feed as stringified json
type RSSFeed struct {
	Id   int64
	Data string
}

// URL at which to fetch an RSS feed
type RSSUrl struct {
	Id  int64
	Url string
}

// SelectAllRSSFeeds selects all rss feed data
func SelectAllRSSFeeds(ctx context.Context) ([]RSSFeed, error) {
	rows, err := db.QueryContext(ctx, "SELECT id, data FROM rss_feeds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feeds []RSSFeed

	for rows.Next() {
		var feed RSSFeed

		if err := rows.Scan(&feed.Id, &feed.Data); err != nil {
			return feeds, err
		}

		feeds = append(feeds, feed)
	}

	if err = rows.Err(); err != nil {
		return feeds, err
	}

	return feeds, nil
}

// InsertRSSItem inserts a new rss item into the database
func InsertRSSItem(ctx context.Context, url string, data []byte) (int64, error) {
	stmt, err := db.PrepareContext(ctx, "INSERT INTO rss_feeds (url, data) VALUES (?, ?)")

	var newId int64

	if err != nil {
		return newId, err
	}

	res, err := stmt.ExecContext(ctx, url, data)

	if err != nil {
		return newId, err
	}

	newId, err = res.LastInsertId()

	if err != nil {
		return newId, err
	}

	return newId, nil
}

// UpdateRSSFeedById updates an rss feed with new stringified json
func UpdateRSSFeedById(ctx context.Context, id int64, data string) error {
	stmt, err := db.PrepareContext(ctx, "UPDATE rss_feeds SET data = ? WHERE id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, data, id)

	if err != nil {
		return err
	}

	return nil
}

// SelectAllRSSUrls returns all rss urls
func SelectAllRSSUrls(ctx context.Context) ([]RSSUrl, error) {
	rows, err := db.QueryContext(ctx, "SELECT id, url FROM rss_feeds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []RSSUrl

	for rows.Next() {
		var url RSSUrl;
		if err := rows.Scan(&url.Id, &url.Url); err != nil {
			return urls, err
		}

		urls = append(urls, url)
	}
	if err = rows.Err(); err != nil {
		return urls, err
	}

	return urls, nil
}
