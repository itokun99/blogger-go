// Code generated for github.com/itokun99/blogger-go. DO NOT EDIT.
// Implements the Blogger API v3 postUserInfos resource.

package services

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/itokun99/blogger-go/gen/builder"
	"github.com/itokun99/blogger-go/gen/schemas"
)

// postUserInfosBasePath is the Blogger v3 API root the postUserInfos paths
// resolve against: GET /v3/users/{userId}/blogs/{blogId}/posts/{postId}.
const postUserInfosBasePath = "https://blogger.googleapis.com/"

// PostUserInfosService implements the Blogger v3 postUserInfos resource.
type PostUserInfosService struct {
	client builder.TransportProvider
}

// NewPostUserInfosService returns a PostUserInfosService bound to client,
// which supplies the transport and API base path for every request.
func NewPostUserInfosService(client builder.TransportProvider) *PostUserInfosService {
	return &PostUserInfosService{client: client}
}

// Get returns the post and user info pair identified by the blog, post, and
// user ids.
//
// GET /v3/users/{userId}/blogs/{blogId}/posts/{postId}
func (s *PostUserInfosService) Get(ctx context.Context, userId, blogId, postId string) (*schemas.PostUserInfo, error) {
	if err := requireValues("userId", userId, "blogId", blogId, "postId", postId); err != nil {
		return nil, err
	}

	info := &schemas.PostUserInfo{}
	request := builder.New(http.MethodGet, postUserInfoPath(userId, blogId, postId)).Context(ctx)

	if err := s.do(request, info); err != nil {
		return nil, err
	}

	return info, nil
}

// List returns the list of posts with user info for the given user and blog.
//
// GET /v3/users/{userId}/blogs/{blogId}/posts
func (s *PostUserInfosService) List(ctx context.Context, userId, blogId string) (*schemas.PostUserInfosList, error) {
	if err := requireValues("userId", userId, "blogId", blogId); err != nil {
		return nil, err
	}

	list := &schemas.PostUserInfosList{}
	request := builder.New(http.MethodGet, postUserInfosListPath(userId, blogId)).Context(ctx)

	if err := s.do(request, list); err != nil {
		return nil, err
	}

	return list, nil
}

// do executes a postUserInfos request, failing before touching the network when
// the service has no usable client.
func (s *PostUserInfosService) do(b *builder.Builder, responseModel interface{}) error {
	if s.client == nil {
		return fmt.Errorf("blogger: postUserInfos service has no client")
	}

	_, err := b.Do(s.transport(), responseModel)
	return err
}

// transport is the HTTP client requests are sent through.
func (s *PostUserInfosService) transport() builder.HTTPClient {
	return postUserInfosClient{
		httpClient: s.client.HTTPClient(),
		basePath:   s.client.BasePath(),
	}
}

// postUserInfosClient adapts the transport to builder.HTTPClient.
type postUserInfosClient struct {
	httpClient *http.Client
	basePath   string
}

// Do sends req through the transport, or the default client when absent.
func (c postUserInfosClient) Do(req *http.Request) (*http.Response, error) {
	if c.httpClient == nil {
		return http.DefaultClient.Do(req)
	}
	return c.httpClient.Do(req)
}

// BasePath returns the API root the request path is resolved against.
func (c postUserInfosClient) BasePath() string {
	if c.basePath != "" {
		return c.basePath
	}
	return postUserInfosBasePath
}

// postUserInfoPath is the request path of a single post and user info pair.
func postUserInfoPath(userId, blogId, postId string) string {
	return fmt.Sprintf(
		"v3/users/%s/blogs/%s/posts/%s",
		url.PathEscape(userId),
		url.PathEscape(blogId),
		url.PathEscape(postId),
	)
}

// postUserInfosListPath is the request path for listing post user infos.
func postUserInfosListPath(userId, blogId string) string {
	return fmt.Sprintf(
		"v3/users/%s/blogs/%s/posts",
		url.PathEscape(userId),
		url.PathEscape(blogId),
	)
}
