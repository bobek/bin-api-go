package binapi

import (
	"context"
	"fmt"
)

// UserPropertiesService handles communication with user property related API
type UserPropertiesService service

// UserProperty describes the API user property object
type UserProperty struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

func (u UserProperty) String() string {
	return Stringify(u)
}

// Get returns information about selected user property
func (s *UserPropertiesService) Get(ctx context.Context, propertyName string) (*UserProperty, *Response, error) {
	u := fmt.Sprintf("users/properties/%v", propertyName)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	up := new(UserProperty)
	resp, err := s.client.Do(ctx, req, up)
	if err != nil {
		return nil, resp, err
	}

	return up, resp, nil
}

// Get returns information about selected user property
func (s *UserPropertiesService) PrepareGet(propertyName string) (*Request, error) {
	u := fmt.Sprintf("users/properties/%v", propertyName)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	return &Request{
		HttpRequest: req,
		Target:      new(UserProperty),
	}, nil
}

// List returns all defined user properties
func (s *UserPropertiesService) List(ctx context.Context) ([]*UserProperty, *Response, error) {
	u := "users/properties/list/"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var properties []*UserProperty
	resp, err := s.client.Do(ctx, req, &properties)
	if err != nil {
		return nil, resp, err
	}

	return properties, resp, nil
}

func (s *UserPropertiesService) PrepareList() (*Request, error) {
	u := "users/properties/list/"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var properties []*UserProperty
	return &Request{
		HttpRequest: req,
		Target:      &properties,
	}, nil
}
