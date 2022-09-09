CREATE TABLE rss_feed (
  rss_feed_id INTEGER PRIMARY KEY,
  url TEXT NOT NULL UNIQUE,
  user_id INTEGER NOT NULL,
  label TEXT NOT NULL,
  FOREIGN KEY(user_id) REFERENCES user(user_id)
);

CREATE TABLE user (
  user_id INTEGER PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL UNIQUE
);

CREATE TABLE tag (
  tag_id INTEGER PRIMARY KEY,
  label TEXT NOT NULL,
  rss_feed_id INTEGER NOT NULL,
  FOREIGN KEY(rss_feed_id) REFERENCES rss_feed(rss_feed_id)
)
