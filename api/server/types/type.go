package types

import "time"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Post struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type PostType struct {
	ID        int64        `json:"id"`
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

type CreatePostInput struct {
	ID    int64    `json:"id" binding:"required"`
	Title string   `json:"title" binding:"required"`
	Body  string   `json:"body" binding:"required"`
	Tags  []string `json:"tags" binding:"required"`
}
