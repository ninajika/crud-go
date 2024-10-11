package models

import "time"

type Post struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type PostType struct {
	ID        int          `json:"id"`
	Title     string       `json:"title"`
	Body      string       `json:"body"`
	Tags      []string     `json:"tags"`
	Reactions PostReaction `json:"reactions"`
	Views     int          `json:"views"`
	UserId    int          `json:"userId"`
}

type PostReaction struct {
	Likes   int `json:"likes"`
	Disikes int `json:"dislikes"`
}
