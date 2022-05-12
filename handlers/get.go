package handlers

import (
	"log"
	"net/http"
)

type Get struct {
	Logger *log.Logger
}

func NewGet(logger *log.Logger) *Get {
	return &Get{logger}
}

func (get *Get) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	get.Logger.Println("Get request received")
}
