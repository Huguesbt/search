package search

import (
	_ "modernc.org/sqlite"
)

type Document struct {
	Id          int64
	Title       string
	Text        string
	Description string
	Notes       string
	Tags        string
}
