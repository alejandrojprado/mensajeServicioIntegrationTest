package models

import "time"

type Message struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	AuthorID  string    `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type RequestCreateMessage struct {
	Content string `json:"content"`
}

type RequestFollow struct {
	FollowingID string `json:"following_id"`
}
