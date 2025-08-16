package model

import "github.com/aburifat/microservice-with-go/metadata/pkg/model"

type MovieDetails struct {
	Rating   float64        `json:"rating,omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
