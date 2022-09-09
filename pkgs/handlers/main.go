package handlers

import "github.com/ec965/rss-server/pkgs/env"

var hmacSecret []byte

func init() {
	hmacSecret = []byte(env.Get("SECRET", "_super_secret_"))
}
