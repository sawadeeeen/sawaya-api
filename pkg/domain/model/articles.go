package model

type Articles struct {
	id           int8   `json:"id"`
	title        string `json:"title"`
	content      string `json:"content"`
	published    bool   `json:"boolean"`
	published_at string `json:"timestamp"`
}
