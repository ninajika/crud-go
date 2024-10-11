package controllers

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
