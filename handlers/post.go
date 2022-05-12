package handlers

import (
	"log"
	"net/http"
)

type Post struct {
	Logger *log.Logger
}

func NewPost(logger *log.Logger) *Get {
	return &Get{logger}
}

func (post *Post) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	post.Logger.Println("Post request received")
}