package github

import (
	"encoding/json"
	"github.com/bandev/lux/api/general"
)

func GetLatestRelease() Release {
	var arr []Release
	var conn general.Connection
	conn.Base = "https://api.github.com"
	conn.Key = "n/a"
	_ = json.Unmarshal(conn.Get("/repos/jackdevey/lux/releases"), &arr)
	return arr[0]
}

type Release struct {
	Tag  string `json:"tag_name"`
	Date string `json:"published_at"`
}
