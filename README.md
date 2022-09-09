# RSS Server

Simple RSS feed consumer server written in Go and deployed with [Fly.io](https://fly.io).

## Design

### v0

1. Users can add new RSS urls

### v1

1. Implement Authentication and users

### v2

2. users can tag their feeds with custom labels

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
