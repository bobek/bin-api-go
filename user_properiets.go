package binapi

import (
	"context"
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
