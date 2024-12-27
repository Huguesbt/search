package search

import (
	_ "modernc.org/sqlite"
)

type Document struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
	Tags        string `json:"tags"`
}
