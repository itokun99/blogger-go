// Code generated for github.com/itokun99/blogger-go. DO NOT EDIT.
// Implements the Blogger API v3 blogUserInfos resource.

package services

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/itokun99/blogger-go/gen/builder"
	"github.com/itokun99/blogger-go/gen/schemas"
)

// blogUserInfosBasePath is the Blogger v3 API root the blogUserInfos paths
// resolve against: GET /v3/users/{userId}/blogs/{blogId}.
const blogUserInfosBasePath = "https://blogger.googleapis.com/"

// BlogUserInfosService implements the Blogger v3 blogUserInfos resource.
type BlogUserInfosService struct {
	client builder.TransportProvider
}

// NewBlogUserInfosService returns a BlogUserInfosService bound to client, which
// supplies the transport and API base path for every request.
func NewBlogUserInfosService(client builder.TransportProvider) *BlogUserInfosService {
	return &BlogUserInfosService{client: client}
}

// Get returns the blog and user info pair identified by the blog and user ids.
//
// GET /v3/users/{userId}/blogs/{blogId}
func (s *BlogUserInfosService) Get(ctx context.Context, userId, blogId string) (*schemas.BlogUserInfo, error) {
	if err := requireValues("userId", userId, "blogId", blogId); err != nil {
		return nil, err
	}

	info := &schemas.BlogUserInfo{}
	request := builder.New(http.MethodGet, blogUserInfoPath(userId, blogId)).Context(ctx)

	if err := s.do(request, info); err != nil {
		return nil, err
	}

	return info, nil
}

// do executes a blogUserInfos request, failing before touching the network when
// the service has no usable client.
func (s *BlogUserInfosService) do(b *builder.Builder, responseModel interface{}) error {
	if s.client == nil {
		return fmt.Errorf("blogger: blogUserInfos service has no client")
	}

	_, err := b.Do(s.transport(), responseModel)
	return err
}

// transport is the HTTP client requests are sent through.
func (s *BlogUserInfosService) transport() builder.HTTPClient {
	return blogUserInfosClient{
		httpClient: s.client.HTTPClient(),
		basePath:   s.client.BasePath(),
	}
}

// blogUserInfosClient adapts the transport to builder.HTTPClient.
type blogUserInfosClient struct {
	httpClient *http.Client
	basePath   string
}

func (c blogUserInfosClient) Do(req *http.Request) (*http.Response, error) {
	if c.httpClient == nil {
		return http.DefaultClient.Do(req)
	}
	return c.httpClient.Do(req)
}

func (c blogUserInfosClient) BasePath() string {
	if c.basePath != "" {
		return c.basePath
	}
	return blogUserInfosBasePath
}

// blogUserInfoPath is the request path of a single blog and user info pair.
func blogUserInfoPath(userId, blogId string) string {
	return fmt.Sprintf(
		"v3/users/%s/blogs/%s",
		url.PathEscape(userId),
		url.PathEscape(blogId),
	)
}
