package main

import (
	"log"
	"net/http"

	"github.com/aburifat/microservice-with-go/metadata/internal/controller/metadata"
	httpHandler "github.com/aburifat/microservice-with-go/metadata/internal/handler/http"
	"github.com/aburifat/microservice-with-go/metadata/internal/repository/memory"
)

func main() {
	log.Printf("Starting metadata service...")
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httpHandler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
