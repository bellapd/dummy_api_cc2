package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Post struct {
	Logger *log.Logger
}

type Get struct {
	Logger *log.Logger
}

type Delete struct {
	Logger *log.Logger
}

func NewPost(logger *log.Logger) *Post {
	return &Post{logger}
}

func NewGet(logger *log.Logger) *Get {
	return &Get{logger}
}

func NewDelete(logger *log.Logger) *Delete {
	return &Delete{logger}
}

func (post *Post) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write([]byte("Received request for 'POST'\n"))
	post.Logger.Println("Received request for 'POST'")
}

func (get *Get) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write([]byte("Received request for 'GET'\n"))
	get.Logger.Println("Received request for 'GET'")
}

func (delete *Delete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write([]byte("Received request for 'DELETE'\n"))
	delete.Logger.Println("Received request for 'DELETE'")
}

func main() {
	logger := log.New(os.Stdout, "vodascheduler ", log.LstdFlags)

	serveMux := http.NewServeMux()
	serveMux.Handle("/post", NewPost(logger))
	serveMux.Handle("/get", NewGet(logger))
	serveMux.Handle("/delete", NewDelete(logger))

	server := &http.Server{
		Handler:      serveMux,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}() // run in background
	logger.Println("API up and running")

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	signal := <-signalChannel
	logger.Println("Received termination, gracefully shutdown", signal)

	tc, _ := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	server.Shutdown(tc)
}