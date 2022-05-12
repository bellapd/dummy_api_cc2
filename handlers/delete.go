package handlers

import (
	"log"
	"net/http"
)

type Delete struct {
	Logger *log.Logger
}

func NewDelete(logger *log.Logger) *Get {
	return &Get{logger}
}

func (delete *Delete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	delete.Logger.Println("Post request received")
}