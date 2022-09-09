package models

import "context"

type RssFeed struct {
	Id    int64    `json:"rss_feed_id"`
	Url   string   `json:"url"`
	Label string   `json:"label"`
	Tags  []string `json:"tags"`
}

func SelectAllFeedsForUser(ctx context.Context, userId int64) ([]RssFeed, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT 
      r.rss_feed_id, 
      r.url,
			r.label,
      t.label
    FROM rss_feed r
    LEFT JOIN tag t ON t.rss_feed_id = r.rss_feed_id
    WHERE r.user_id = ?;`, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feeds []RssFeed
	tagMap := make(map[int64][]string)

	for rows.Next() {
		var feed RssFeed
		var tag string
		if err := rows.Scan(&feed.Id, &feed.Url, &feed.Label, &tag); err != nil {
			return feeds, err
		}
		feeds = append(feeds, feed)
		tagMap[feed.Id] = append(tagMap[feed.Id], tag)
	}

	if err = rows.Err(); err != nil {
		return feeds, err
	}

	for _, feed := range feeds {
		feed.Tags = tagMap[feed.Id]
	}

	return feeds, nil
}

func SelectFeedForUser(ctx context.Context, userId int64, rssFeedId int64) (RssFeed, error) {
	var feed RssFeed

	rows, err := db.QueryContext(ctx,
		`SELECT
  r.rss_feed_id,
  r.url,
	r.label,
  t.label
FROM rss_feed r
LEFT JOIN tag t on t.rss_feed_id = r.rss_feed_id
WHERE r.user_id = ? AND r.rss_feed_id = ?;`, userId, rssFeedId)
	if err != nil {
		return feed, err
	}
	defer rows.Close()

	tags := []string{}

	for rows.Next() {
		var tag string
		if err := rows.Scan(&feed.Id, &feed.Url, &feed.Label, &tag); err != nil {
			return feed, err
		}
		tags = append(tags, tag)
	}
	feed.Tags = tags

	return feed, nil
}

func InsertFeedForUser(ctx context.Context, userId int64, url string, label string) (int64, error) {
	res, err := db.ExecContext(ctx,
		`INSERT INTO rss_feed (user_id, url, label) VALUES (?, ?, ?);`, userId, url, label)

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func DeleteFeedForUser(ctx context.Context, userId int64, rssFeedId int64) (int64, error) {
	res, err := db.ExecContext(ctx,
		`DELETE FROM rss_feed WHERE user_id = ? AND rss_feed_id = ?;`,
		userId,
		rssFeedId,
	)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
