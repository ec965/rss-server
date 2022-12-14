CREATE TABLE rssFeed (
  rssFeedId INTEGER PRIMARY KEY,
  url TEXT NOT NULL UNIQUE,
  userId INTEGER NOT NULL,
  label TEXT NOT NULL,
  FOREIGN KEY(userId) REFERENCES user(userId)
);

CREATE TABLE user (
  userId INTEGER PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL UNIQUE
);

CREATE TABLE tag (
  tagId INTEGER PRIMARY KEY,
  label TEXT NOT NULL,
  rssFeedId INTEGER NOT NULL,
  FOREIGN KEY(rssFeedId) REFERENCES rssFeed(rssFeedId)
);
