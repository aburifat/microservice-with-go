package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aburifat/microservice-with-go/metadata/pkg/model"
	"github.com/aburifat/microservice-with-go/movie/internal/gateway"
)

type Gateway struct {
	addr string
}

func New(addr string) *Gateway {
	return &Gateway{
		addr: addr,
	}
}

func (g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error) {
	req, err := http.NewRequest(http.MethodGet, g.addr+"/metadata", nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", id)
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, gateway.ErrNotFound
	} else if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("non-2xx response: %v", resp)
	}
	var v *model.Metadata
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

func (g *Gateway) Put(ctx context.Context, id string, title string, description string, director string) error {
	req, err := http.NewRequest(http.MethodPut, g.addr+"/metadata/put", nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", id)
	values.Add("title", title)
	values.Add("description", description)
	values.Add("director", director)
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("non-2xx response: %v", resp)
	}
	return nil
}
