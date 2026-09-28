// Package blogger provides a client for the Blogger API v3.
package blogger

import (
	"context"
	"net/http"

	"google.golang.org/api/blogger/v3"
	"google.golang.org/api/option"

	"github.com/itokun99/blogger-go/gen/builder"
	"github.com/itokun99/blogger-go/gen/services"
)

const defaultBaseServiceURL = "https://blogger.googleapis.com/"

// Client is the entry point for talking to the Blogger API v3.
type Client struct {
	Service  *blogger.Service
	http     *http.Client
	basePath string
	ctx      context.Context
}

// NewClient creates a Client using the supplied options. When no HTTP client
// is supplied through opts, a default client is used.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	svc, err := blogger.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		Service:  svc,
		ctx:      ctx,
		basePath: svc.BasePath,
	}, nil
}

// NewRawClient creates a Client backed by a caller-provided http.Client. It
// never fails and requires no credentials; use it when you manage
// authentication and transport yourself. The supplied client must not be nil.
func NewRawClient(httpClient *http.Client) *Client {
	return &Client{
		http:     httpClient,
		basePath: defaultBaseServiceURL,
		ctx:      context.Background(),
	}
}

// WithContext returns a copy of the Client bound to ctx. The underlying
// transport is shared, so nothing else about the Client changes.
func (c *Client) WithContext(ctx context.Context) *Client {
	clone := *c
	clone.ctx = ctx
	return &clone
}

// HTTPClient returns the transport the Client sends through.
func (c *Client) HTTPClient() *http.Client {
	if c.http != nil {
		return c.http
	}
	return http.DefaultClient
}

// BasePath returns the API root the Client resolves paths against.
func (c *Client) BasePath() string {
	if c.Service != nil {
		return c.Service.BasePath
	}
	if c.basePath != "" {
		return c.basePath
	}
	return defaultBaseServiceURL
}

// Blogs returns a builder for the blogs resource.
func (c *Client) Blogs() *services.BlogsService {
	return services.NewBlogsService(c)
}

// Comments returns a builder for the comments resource.
func (c *Client) Comments() *services.CommentsService {
	return services.NewCommentsService(c)
}

// Pages returns a builder for the pages resource.
func (c *Client) Pages() *services.PagesService {
	return services.NewPagesService(c)
}

// Posts returns a builder for the posts resource.
func (c *Client) Posts() *services.PostsService {
	return services.NewPostsService(c)
}

// Users returns a builder for the users resource.
func (c *Client) Users() *services.UsersService {
	return services.NewUsersService(c)
}

// BlogUserInfos returns a builder for the blogUserInfos resource.
func (c *Client) BlogUserInfos() *services.BlogUserInfosService {
	return services.NewBlogUserInfosService(c)
}

// PageViews returns a builder for the pageViews resource.
func (c *Client) PageViews() *services.PageViewsService {
	return services.NewPageViewsService(c)
}

// PostUserInfos returns a builder for the postUserInfos resource.
func (c *Client) PostUserInfos() *services.PostUserInfosService {
	return services.NewPostUserInfosService(c)
}

var _ builder.TransportProvider = (*Client)(nil)
