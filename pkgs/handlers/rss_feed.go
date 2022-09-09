package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ec965/rss-server/pkgs/models"
	"github.com/go-chi/chi/v5"
)

// all handlers should pass through the AuthMiddleware

func GetFeeds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := ctx.Value(AuthCtxUserId).(int64)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	feeds, err := models.SelectAllFeedsForUser(context.TODO(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feeds)
}

func GetFeed(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := ctx.Value(AuthCtxUserId).(int64)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	feedUrlParam := chi.URLParam(r, "feedId")
	rssFeedId, err := strconv.Atoi(feedUrlParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feed, err := models.SelectFeedForUser(context.TODO(), userId, int64(rssFeedId))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}

type PostAddFeedBody struct {
	Url   string `json:"url"`
	Label string `json:"label"`
}

func PostAddFeed(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := ctx.Value(AuthCtxUserId).(int64)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	var b PostAddFeedBody
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newId, err := models.InsertFeedForUser(context.TODO(), userId, b.Url, b.Label)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	feed, err := models.SelectFeedForUser(context.TODO(), userId, newId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}
