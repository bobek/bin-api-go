package binapi

import (
	"context"
	"fmt"
)

type UsersService service

type User struct {
	Login *string `json:"login,omitempty"`
	ID    *int64  `json:"id,omitempty"`
}

func (u User) String() string {
	return Stringify(u)
}

// Get fetches a user. Passing the empty string will fetch the authenticated
// user.
//
// GitHub API docs: https://docs.github.com/en/free-pro-team@latest/rest/reference/users/#get-the-authenticated-user
// GitHub API docs: https://docs.github.com/en/free-pro-team@latest/rest/reference/users/#get-a-user
func (s *UsersService) Get(ctx context.Context, user string) (*User, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v", user)
	} else {
		u = "user"
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(User)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// GetByID fetches a user.
//
// Note: GetByID uses the undocumented GitHub API endpoint /user/:id.
func (s *UsersService) GetByID(ctx context.Context, id int64) (*User, *Response, error) {
	u := fmt.Sprintf("user/%d", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	resp, err := s.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// Edit the authenticated user.
//
// GitHub API docs: https://docs.github.com/en/free-pro-team@latest/rest/reference/users/#update-the-authenticated-user
func (s *UsersService) Edit(ctx context.Context, user *User) (*User, *Response, error) {
	u := "user"
	req, err := s.client.NewRequest("PATCH", u, user)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(User)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// HovercardOptions specifies optional parameters to the UsersService.GetHovercard
// method.
type HovercardOptions struct {
	// SubjectType specifies the additional information to be received about the hovercard.
	// Possible values are: organization, repository, issue, pull_request. (Required when using subject_id.)
	SubjectType string `url:"subject_type"`

	// SubjectID specifies the ID for the SubjectType. (Required when using subject_type.)
	SubjectID string `url:"subject_id"`
}

// Hovercard represents hovercard information about a user.
type Hovercard struct {
	Contexts []*UserContext `json:"contexts,omitempty"`
}

// UserContext represents the contextual information about user.
type UserContext struct {
	Message *string `json:"message,omitempty"`
	Octicon *string `json:"octicon,omitempty"`
}

// GetHovercard fetches contextual information about user. It requires authentication
// via Basic Auth or via OAuth with the repo scope.
//
// GitHub API docs: https://docs.github.com/en/free-pro-team@latest/rest/reference/users/#get-contextual-information-for-a-user
func (s *UsersService) GetHovercard(ctx context.Context, user string, opts *HovercardOptions) (*Hovercard, *Response, error) {
	u := fmt.Sprintf("users/%v/hovercard", user)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	hc := new(Hovercard)
	resp, err := s.client.Do(ctx, req, hc)
	if err != nil {
		return nil, resp, err
	}

	return hc, resp, nil
}

// UserListOptions specifies optional parameters to the UsersService.ListAll
// method.
type UserListOptions struct {
	// ID of the last user seen
	Since int64 `url:"since,omitempty"`
}

// ListAll lists all GitHub users.
//
// To paginate through all users, populate 'Since' with the ID of the last user.
//
// GitHub API docs: https://docs.github.com/en/free-pro-team@latest/rest/reference/users/#list-users
func (s *UsersService) ListAll(ctx context.Context, opts *UserListOptions) ([]*User, *Response, error) {
	u, err := addOptions("users", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*User
	resp, err := s.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

// AcceptInvitation accepts the currently-open repository invitation for the
// authenticated user.
//
// GitHub API docs: https://docs.github.com/en/free-pro-team@latest/rest/reference/repos/#accept-a-repository-invitation
func (s *UsersService) AcceptInvitation(ctx context.Context, invitationID int64) (*Response, error) {
	u := fmt.Sprintf("user/repository_invitations/%v", invitationID)
	req, err := s.client.NewRequest("PATCH", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeclineInvitation declines the currently-open repository invitation for the
// authenticated user.
//
// GitHub API docs: https://docs.github.com/en/free-pro-team@latest/rest/reference/repos/#decline-a-repository-invitation
func (s *UsersService) DeclineInvitation(ctx context.Context, invitationID int64) (*Response, error) {
	u := fmt.Sprintf("user/repository_invitations/%v", invitationID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
