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

1. Add folders - RSS feeds can be organized into folders of depth 1

## DB Tables

**rss_feeds**

| Type   | Label | Description                          |
| ------ | ----- | ------------------------------------ |
| ID     | id    | id                                   |
| String | url   | RSS feed url                         |
| String | data | most up to date parsed rss feed json |

**user**

| Type   | Label    | Description |
| ------ | -------- | ----------- |
| ID     | id       | id          |
| String | email    | email       |
| String | password | password    |

**user_feeds**

| Type | Label       | Description                               |
| ---- | ----------- | ----------------------------------------- |
| ID   | id          | id                                        |
| ID   | user_id     | fk to user id of user that owns this feed |
| ID   | rss_feed_id | fk to rss_feed id                         |
| ID   | folder_id   | fk to folder                              |

**folder**

| Type   | Label   | Description                               |
| ------ | ------- | ----------------------------------------- |
| ID     | id      | id                                        |
| ID     | user_id | fk to user id of user who owns the folder |
| String | label   | folder label                              |
