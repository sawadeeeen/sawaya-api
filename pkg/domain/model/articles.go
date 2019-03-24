package model

import (
	"time"
)

type Article struct {
	Id           int8      `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Published    bool      `json:"boolean"`
	Published_at time.Time `json:"timestamp"`
}

type Articles Article
