package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ec965/rss-server/pkgs/models"
	"github.com/ec965/rss-server/pkgs/rss"
)

type UrlFeedBody struct {
	Url string `json:"url"`
}

func UpdateFeeds(w http.ResponseWriter, r *http.Request) {
	err := rss.UpdateAllFeeds()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	GetFeeds(w, r)
}

func GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := models.SelectAllRSSFeeds(context.TODO())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json := "["

	for i, feed := range feeds {
		json += feed.Data

		if i != len(feeds)-1 {
			json += ","
		}
	}

	json += "]"

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(json))
}

func AddFeed(w http.ResponseWriter, r *http.Request) {
	var body UrlFeedBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := rss.GetFeed(context.TODO(), body.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	_, err = models.InsertRSSItem(context.TODO(), body.Url, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteFeed(w http.ResponseWriter, r *http.Request) {
	var body UrlFeedBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := models.DeleteRSSItemByUrl(context.TODO(), body.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
