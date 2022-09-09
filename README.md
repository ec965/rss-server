# RSS Server

Simple RSS feed consumer server written in Go and deployed with [Fly.io](https://fly.io).

## Design

### v0

1. Users can add new RSS urls
1. RSS urls are stored in a table
1. RSS feeds are refreshed by cron job
1. Feeds are parsed into consumable JSON, stringified, and saved in a table
1. Users can request RSS payloads

### v1

1. Implement Authentication and users
2. All RSS urls and data is still saved globally but users can own a list of RSS urls

### v2

1. Add tags - RSS feeds can be organized into folders of depth 1

## DB Tables

**rssFeed**

Holds the RSS feed url.

| Type   | Label     | Description           |
| ------ | --------- | --------------------- |
| ID     | rssFeedId | id                    |
| String | url       | RSS feed url          |
| ID     | userId    | user id of feed owner |

**tag**

Each RSS feed can have many user specific tags.
Tags are used to organize feeds.

| Type   | Label     | Description    |
| ------ | --------- | -------------- |
| ID     | tag_id    | id             |
| String | label     | tag label      |
| String | rssFeedId | fk to rss feed |

**user**

| Type   | Label    | Description |
| ------ | -------- | ----------- |
| ID     | userId   | id          |
| String | email    | email       |
| String | password | password    |
