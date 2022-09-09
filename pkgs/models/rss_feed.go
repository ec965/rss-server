package models

import (
	"context"
	"database/sql"
	"fmt"
)

type RssFeed struct {
	Id    int64    `json:"rssFeedId"`
	Url   string   `json:"url"`
	Label string   `json:"label"`
	Tags  []string `json:"tags"`
}

func rssFeedMapToSlice(m *map[int64]*RssFeed) []*RssFeed {
	res := make([]*RssFeed, 0, len(*m))
	for _, val := range *m {
		res = append(res, val)
	}
	return res
}

func SelectAllFeedsForUser(ctx context.Context, userId int64) ([]*RssFeed, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT 
      r.rss_feed_id, 
      r.url,
			r.label,
      t.label as tag
    FROM rss_feed r
    LEFT JOIN tag t ON t.rss_feed_id = r.rss_feed_id
    WHERE r.user_id = ?;`, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	m := make(map[int64]*RssFeed)

	for rows.Next() {
		var feed RssFeed
		var tag sql.NullString
		if err := rows.Scan(&feed.Id, &feed.Url, &feed.Label, &tag); err != nil {
			res := rssFeedMapToSlice(&m)
			return res, err
		}
		fmt.Println("feed", feed)
		fmt.Println("tag", tag)

		_, ok := m[feed.Id]
		if ok {
			if tag.Valid {
				m[feed.Id].Tags = append(m[feed.Id].Tags, tag.String)
			}
		} else {
			// instantiate the slice
			if tag.Valid {
				feed.Tags = []string{tag.String}
			} else {
				feed.Tags = []string{}
			}
			m[feed.Id] = &feed
		}
	}

	if err = rows.Err(); err != nil {
		res := rssFeedMapToSlice(&m)
		return res, err
	}

	res := rssFeedMapToSlice(&m)
	fmt.Println("res", res)
	return res, nil
}

func SelectFeedForUser(ctx context.Context, userId int64, rssFeedId int64) (RssFeed, error) {
	var feed RssFeed

	rows, err := db.QueryContext(ctx,
		`SELECT
  r.rss_feed_id,
  r.url,
	r.label,
  t.label as tag
FROM rss_feed r
LEFT JOIN tag t on t.rss_feed_id = r.rss_feed_id
WHERE r.user_id = ? AND r.rss_feed_id = ?;`, userId, rssFeedId)
	if err != nil {
		return feed, err
	}
	defer rows.Close()

	tags := []string{}

	for rows.Next() {
		var tag sql.NullString
		if err := rows.Scan(&feed.Id, &feed.Url, &feed.Label, &tag); err != nil {
			return feed, err
		}
		if tag.Valid {
			tags = append(tags, tag.String)
		}
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
