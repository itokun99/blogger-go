// Code generated for github.com/itokun99/blogger-go. DO NOT EDIT.
// Implements the Blogger API v3 pageViews resource.

package services

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/itokun99/blogger-go/gen/builder"
	"github.com/itokun99/blogger-go/gen/schemas"
)

// pageViewsBasePath is the Blogger v3 API root the pageViews paths resolve
// against: GET /v3/blogs/{blogId}/pageviews.
const pageViewsBasePath = "https://blogger.googleapis.com/"

// PageViewsService implements the Blogger v3 pageViews resource.
type PageViewsService struct {
	client builder.TransportProvider
}

// NewPageViewsService returns a PageViewsService bound to client, which
// supplies the transport and API base path for every request.
func NewPageViewsService(client builder.TransportProvider) *PageViewsService {
	return &PageViewsService{client: client}
}

// PageViewsOption customises a pageViews request. Options are applied in order
// after the required path parameters.
type PageViewsOption func(*builder.Builder)

// WithPageViewRange restricts the response to the given time range.
// Valid values are "all", "30days", "7days".
func WithPageViewRange(rangeVal string) PageViewsOption {
	return func(b *builder.Builder) {
		b.Param("range", rangeVal)
	}
}

// Get returns the page views of a blog, narrowed by opts.
//
// GET /v3/blogs/{blogId}/pageviews
func (s *PageViewsService) Get(ctx context.Context, blogId string, opts ...PageViewsOption) (*schemas.Pageviews, error) {
	if err := requireValues("blogId", blogId); err != nil {
		return nil, err
	}

	request := builder.New(http.MethodGet, pageViewsPath(blogId)).Context(ctx)
	applyPageViewsOptions(request, opts)

	pageViews := &schemas.Pageviews{}
	if err := s.do(request, pageViews); err != nil {
		return nil, err
	}

	return pageViews, nil
}

// do executes a pageViews request, failing before touching the network when
// the service has no usable client.
func (s *PageViewsService) do(b *builder.Builder, responseModel interface{}) error {
	if s.client == nil {
		return fmt.Errorf("blogger: pageViews service has no client")
	}

	_, err := b.Do(s.transport(), responseModel)
	return err
}

// transport is the HTTP client requests are sent through.
func (s *PageViewsService) transport() builder.HTTPClient {
	return pageViewsClient{
		httpClient: s.client.HTTPClient(),
		basePath:   s.client.BasePath(),
	}
}

// pageViewsClient adapts the transport to builder.HTTPClient.
type pageViewsClient struct {
	httpClient *http.Client
	basePath   string
}

// Do sends req through the transport, or the default client when absent.
func (c pageViewsClient) Do(req *http.Request) (*http.Response, error) {
	if c.httpClient == nil {
		return http.DefaultClient.Do(req)
	}
	return c.httpClient.Do(req)
}

// BasePath returns the API root the request path is resolved against.
func (c pageViewsClient) BasePath() string {
	if c.basePath != "" {
		return c.basePath
	}
	return pageViewsBasePath
}

// applyPageViewsOptions applies every non-nil option to the request.
func applyPageViewsOptions(b *builder.Builder, opts []PageViewsOption) {
	for _, opt := range opts {
		if opt != nil {
			opt(b)
		}
	}
}

// pageViewsPath is the request path of a blog's page views.
func pageViewsPath(blogId string) string {
	return fmt.Sprintf("v3/blogs/%s/pageviews", url.PathEscape(blogId))
}
