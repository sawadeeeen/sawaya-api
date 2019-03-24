package model

import (
	"time"
)

type Articles struct {
	Id           int8      `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Published    bool      `json:"boolean"`
	Published_at time.Time `json:"timestamp"`
}
