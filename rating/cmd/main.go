package main

import (
	"log"
	"net/http"

	"github.com/aburifat/microservice-with-go/rating/internal/controller/rating"
	httpHandler "github.com/aburifat/microservice-with-go/rating/internal/handler/http"
	"github.com/aburifat/microservice-with-go/rating/internal/repository/memory"
)

func main() {
	log.Printf("Starting rating service...")
	log.Printf("rating service is listening to http://localhost:8082 ...")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httpHandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
