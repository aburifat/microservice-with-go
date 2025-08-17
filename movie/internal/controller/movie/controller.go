package movie

import (
	"context"
	"errors"

	metadataModel "github.com/aburifat/microservice-with-go/metadata/pkg/model"
	"github.com/aburifat/microservice-with-go/movie/internal/gateway"
	"github.com/aburifat/microservice-with-go/movie/pkg/model"
	ratingModel "github.com/aburifat/microservice-with-go/rating/pkg/model"
)

var ErrNotFound = errors.New("movie not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordId ratingModel.RecordID, recordType ratingModel.RecordType) (float64, error)
	PutRating(ctx context.Context, recordId ratingModel.RecordID, recordType ratingModel.RecordType, rating *ratingModel.Rating) error
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadataModel.Metadata, error)
	Put(ctx context.Context, metadata *metadataModel.Metadata) error
}

type Controller struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

func New(ratingGateway ratingGateway, metadataGateway metadataGateway) *Controller {
	return &Controller{
		ratingGateway:   ratingGateway,
		metadataGateway: metadataGateway,
	}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gateway.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	details := &model.MovieDetails{
		Metadata: *metadata,
	}

	rating, err := c.ratingGateway.GetAggregatedRating(ctx, ratingModel.RecordID(id), ratingModel.RecordTypeMovie)
	if err != nil && !errors.Is(err, gateway.ErrNotFound) {
		// Do not return it, as we still want to return the metadata
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = rating
	}

	return details, nil
}
