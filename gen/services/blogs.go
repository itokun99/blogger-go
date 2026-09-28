// Code generated for github.com/itokun99/blogger-go. DO NOT EDIT.
// Implements the Blogger API v3 blogs resource.

package services

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/itokun99/blogger-go/gen/builder"
	"github.com/itokun99/blogger-go/gen/schemas"
)

// blogsBasePath is the Blogger v3 API root the blogs paths resolve against:
// GET /v3/blogs/{blogId}.
const blogsBasePath = "https://blogger.googleapis.com/"

// BlogsService implements the Blogger v3 blogs resource.
type BlogsService struct {
	client builder.TransportProvider
}

// NewBlogsService returns a BlogsService bound to client, which supplies the
// transport and API base path for every request.
func NewBlogsService(client builder.TransportProvider) *BlogsService {
	return &BlogsService{client: client}
}

// BlogsOption customises a blogs request.
type BlogsOption func(*builder.Builder)

// WithBlogView sets the access level for the returned resource.
// Valid values are "READER", "AUTHOR", or "ADMIN".
func WithBlogView(view string) BlogsOption {
	return func(b *builder.Builder) {
		b.Param("view", view)
	}
}

// WithBlogMaxPosts sets the maximum number of posts to include in the response.
func WithBlogMaxPosts(maxPosts int64) BlogsOption {
	return func(b *builder.Builder) {
		b.Param("maxPosts", strconv.FormatInt(maxPosts, 10))
	}
}

// Get returns the Blog with the given id.
//
// GET /v3/blogs/{blogId}
func (s *BlogsService) Get(ctx context.Context, blogId string, opts ...BlogsOption) (*schemas.Blog, error) {
	if err := requireValues("blogId", blogId); err != nil {
		return nil, err
	}

	request := builder.New(http.MethodGet, blogPath(blogId)).Context(ctx)
	applyBlogsOptions(request, opts)

	blog := &schemas.Blog{}
	if err := s.do(request, blog); err != nil {
		return nil, err
	}
	return blog, nil
}

// GetByUrl returns the Blog that has the given url.
//
// GET /v3/blogs/byurl?url={url}
func (s *BlogsService) GetByUrl(ctx context.Context, blogURL string, opts ...BlogsOption) (*schemas.Blog, error) {
	if err := requireValues("url", blogURL); err != nil {
		return nil, err
	}

	request := builder.New(http.MethodGet, "v3/blogs/byurl").
		Context(ctx).
		Param("url", blogURL)
	applyBlogsOptions(request, opts)

	blog := &schemas.Blog{}
	if err := s.do(request, blog); err != nil {
		return nil, err
	}
	return blog, nil
}

// ListByUser lists blogs by the given user.
//
// GET /v3/users/{userId}/blogs
func (s *BlogsService) ListByUser(ctx context.Context, userId string, opts ...BlogsOption) (*schemas.BlogList, error) {
	if err := requireValues("userId", userId); err != nil {
		return nil, err
	}

	request := builder.New(http.MethodGet, userBlogsPath(userId)).Context(ctx)
	applyBlogsOptions(request, opts)

	list := &schemas.BlogList{}
	if err := s.do(request, list); err != nil {
		return nil, err
	}
	return list, nil
}

// do executes a blogs request.
func (s *BlogsService) do(b *builder.Builder, responseModel interface{}) error {
	if s.client == nil {
		return fmt.Errorf("blogger: blogs service has no client")
	}
	_, err := b.Do(s.transport(), responseModel)
	return err
}

// transport is the HTTP client requests are sent through.
func (s *BlogsService) transport() builder.HTTPClient {
	return blogsClient{
		httpClient: s.client.HTTPClient(),
		basePath:   s.client.BasePath(),
	}
}

// blogsClient adapts the transport to builder.HTTPClient.
type blogsClient struct {
	httpClient *http.Client
	basePath   string
}

// Do sends req through the transport, or the default client when absent.
func (c blogsClient) Do(req *http.Request) (*http.Response, error) {
	if c.httpClient != nil {
		return c.httpClient.Do(req)
	}
	return http.DefaultClient.Do(req)
}

// BasePath returns the API root the request path is resolved against.
func (c blogsClient) BasePath() string {
	return c.basePath
}

// applyBlogsOptions applies every non-nil option to the request.
func applyBlogsOptions(b *builder.Builder, opts []BlogsOption) {
	for _, opt := range opts {
		if opt != nil {
			opt(b)
		}
	}
}

// blogPath is the request path of a single blog.
func blogPath(blogId string) string {
	return fmt.Sprintf("v3/blogs/%s", url.PathEscape(blogId))
}

// userBlogsPath is the request path for listing blogs by user.
func userBlogsPath(userId string) string {
	return fmt.Sprintf("v3/users/%s/blogs", url.PathEscape(userId))
}
