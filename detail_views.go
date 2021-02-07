package binapi

import (
	"context"
	"fmt"
)

// DetailViewsService handles communication with user property related API
type DetailViewsService service

// DetailView describes the API user property object
type DetailView struct {
	Timestamp     string `json:"timestamp,omitempty"`
	Duration      int64  `json:"duration,omitempty"`
	CascadeCreate bool   `json:"cascadeCreate,omitempty"`
	RecommId      string `json:"recommId,omitempty"`
	UserId        string `json:"userId,omitempty"`
	ItemId        string `json:"itemId,omitempty"`
}

func (u DetailView) String() string {
	return Stringify(u)
}

func (s *DetailViewsService) Post(ctx context.Context, dv *DetailView) (*Response, error) {
	u := "detailviews/"

	req, err := s.client.NewRequest("POST", u, dv)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

func (s *DetailViewsService) PreparePost(dv *DetailView) (*Request, error) {
	u := "detailviews/"

	req, err := s.client.NewRequest("POST", u, dv)
	if err != nil {
		return nil, err
	}

	return &Request{
		HttpRequest: req,
		Target:      nil,
	}, nil
}
