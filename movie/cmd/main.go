package main

import (
	"log"
	"net/http"

	"github.com/aburifat/microservice-with-go/movie/internal/controller/movie"
	metadataGateway "github.com/aburifat/microservice-with-go/movie/internal/gateway/metadata/http"
	ratingGateway "github.com/aburifat/microservice-with-go/movie/internal/gateway/rating/http"
	httpHandler "github.com/aburifat/microservice-with-go/movie/internal/handler/http"
)

func main() {
	log.Printf("Starting movie gateway...")
	log.Printf("Movie gateway is listening to http://localhost:8083 ...")

	metadataGateway := metadataGateway.New("http://localhost:8081")
	ratingGateway := ratingGateway.New("http://localhost:8082")
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httpHandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
