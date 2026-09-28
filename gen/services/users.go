// Code generated for github.com/itokun99/blogger-go. DO NOT EDIT.
// Implements the Blogger API v3 users resource.

package services

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/itokun99/blogger-go/gen/builder"
	"github.com/itokun99/blogger-go/gen/schemas"
)

// usersBasePath is the Blogger v3 API root the users paths resolve against:
// GET /v3/users/{userId}.
const usersBasePath = "https://blogger.googleapis.com/"

// UsersService implements the Blogger v3 users resource.
type UsersService struct {
	client builder.TransportProvider
}

// NewUsersService returns a UsersService bound to client, which supplies the
// transport and API base path for every request.
func NewUsersService(client builder.TransportProvider) *UsersService {
	return &UsersService{client: client}
}

// Get returns the user with the given id.
//
// GET /v3/users/{userId}
func (s *UsersService) Get(ctx context.Context, userId string) (*schemas.User, error) {
	if err := requireValues("userId", userId); err != nil {
		return nil, err
	}

	request := builder.New(http.MethodGet, userPath(userId)).Context(ctx)

	user := &schemas.User{}
	if _, err := request.Do(s.transport(), user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UsersService) transport() builder.HTTPClient {
	return usersClient{
		httpClient: s.client.HTTPClient(),
		basePath:   s.client.BasePath(),
	}
}

// usersClient adapts the transport to builder.HTTPClient.
type usersClient struct {
	httpClient *http.Client
	basePath   string
}

// Do sends req through the transport, or the default client when absent.
func (c usersClient) Do(req *http.Request) (*http.Response, error) {
	if c.httpClient == nil {
		return http.DefaultClient.Do(req)
	}
	return c.httpClient.Do(req)
}

// BasePath returns the API root the request path is resolved against.
func (c usersClient) BasePath() string {
	if c.basePath != "" {
		return c.basePath
	}
	return usersBasePath
}

// userPath is the request path of a single user.
func userPath(userId string) string {
	return fmt.Sprintf("v3/users/%s", url.PathEscape(userId))
}
