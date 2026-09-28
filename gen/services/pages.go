// Code generated for github.com/itokun99/blogger-go. DO NOT EDIT.

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

// pagesBasePath is the Blogger v3 API root the pages paths resolve against:
// GET /v3/blogs/{blogId}/pages.
const pagesBasePath = "https://blogger.googleapis.com/"

// PagesService implements the Blogger v3 pages resource.
type PagesService struct {
	client builder.TransportProvider
}

// NewPagesService returns a PagesService bound to client, which supplies the
// transport and API base path for every request.
func NewPagesService(client builder.TransportProvider) *PagesService {
	return &PagesService{client: client}
}

// PageListOption customises a pages list request. Options are applied in order
// after the required path parameters.
type PageListOption func(*builder.Builder)

// WithPageListToken continues a listing from the cursor returned as
// nextPageToken.
func WithPageListToken(token string) PageListOption {
	return func(b *builder.Builder) {
		b.Param(builder.PageToken, token)
	}
}

// WithPageListMaxResults caps the number of pages returned in a single request.
func WithPageListMaxResults(maxResults int64) PageListOption {
	return func(b *builder.Builder) {
		b.Param(builder.MaxResults, strconv.FormatInt(maxResults, 10))
	}
}

// PageStatus values accepted by pages.list.
const (
	PageStatusLive        = "LIVE"
	PageStatusDraft       = "DRAFT"
	PageStatusSoftTrashed = "SOFT_TRASHED"
)

// WithPageListStatus restricts the listing to pages in the given status.
func WithPageListStatus(status string) PageListOption {
	return func(b *builder.Builder) {
		b.Param(builder.Status, status)
	}
}

// WithPageView sets the access level for the returned resource.
// Valid values are "READER", "AUTHOR", or "ADMIN".
func WithPageView(view string) PageListOption {
	return func(b *builder.Builder) {
		b.Param("view", view)
	}
}

// WithPageFetchBodies controls whether page bodies are included.
func WithPageFetchBodies(fetchBodies bool) PageListOption {
	return func(b *builder.Builder) {
		b.Param("fetchBodies", strconv.FormatBool(fetchBodies))
	}
}

// Get returns the page with the given id.
//
// GET /v3/blogs/{blogId}/pages/{pageId}
func (s *PagesService) Get(ctx context.Context, blogId, pageId string) (*schemas.Page, error) {
	if err := requireValues("blogId", blogId, "pageId", pageId); err != nil {
		return nil, err
	}

	page := &schemas.Page{}
	request := builder.New(http.MethodGet, pagePath(blogId, pageId)).Context(ctx)

	if err := s.do(request, page); err != nil {
		return nil, err
	}

	return page, nil
}

// List returns the pages of a blog, narrowed by opts.
//
// GET /v3/blogs/{blogId}/pages
func (s *PagesService) List(ctx context.Context, blogId string, opts ...PageListOption) (*schemas.PageList, error) {
	if err := requireValues("blogId", blogId); err != nil {
		return nil, err
	}

	path := fmt.Sprintf("v3/blogs/%s/pages", url.PathEscape(blogId))
	request := builder.New(http.MethodGet, path).Context(ctx)
	applyPageOptions(request, opts)

	list := &schemas.PageList{Items: []*schemas.Page{}}
	if err := s.do(request, list); err != nil {
		return nil, err
	}

	return list, nil
}

// Insert creates a page on the blog and returns the stored page.
//
// POST /v3/blogs/{blogId}/pages
func (s *PagesService) Insert(ctx context.Context, blogId string, body *schemas.Page) (*schemas.Page, error) {
	if err := requireValues("blogId", blogId); err != nil {
		return nil, err
	}

	path := fmt.Sprintf("v3/blogs/%s/pages", url.PathEscape(blogId))
	return s.mutate(ctx, http.MethodPost, path, body)
}

// Update replaces the page and returns the stored page.
//
// PUT /v3/blogs/{blogId}/pages/{pageId}
func (s *PagesService) Update(ctx context.Context, blogId, pageId string, body *schemas.Page) (*schemas.Page, error) {
	if err := requireValues("blogId", blogId, "pageId", pageId); err != nil {
		return nil, err
	}

	return s.mutate(ctx, http.MethodPut, pagePath(blogId, pageId), body)
}

// Patch updates the provided fields of the page and returns the stored page.
//
// PATCH /v3/blogs/{blogId}/pages/{pageId}
func (s *PagesService) Patch(ctx context.Context, blogId, pageId string, body *schemas.Page) (*schemas.Page, error) {
	if err := requireValues("blogId", blogId, "pageId", pageId); err != nil {
		return nil, err
	}

	return s.mutate(ctx, http.MethodPatch, pagePath(blogId, pageId), body)
}

// Delete removes the page.
//
// DELETE /v3/blogs/{blogId}/pages/{pageId}
func (s *PagesService) Delete(ctx context.Context, blogId, pageId string) error {
	if err := requireValues("blogId", blogId, "pageId", pageId); err != nil {
		return err
	}

	request := builder.New(http.MethodDelete, pagePath(blogId, pageId)).Context(ctx)

	return s.do(request, nil)
}

// Publish makes the page publicly visible and returns it.
//
// POST /v3/blogs/{blogId}/pages/{pageId}/publish
func (s *PagesService) Publish(ctx context.Context, blogId, pageId string) (*schemas.Page, error) {
	return s.action(ctx, "publish", blogId, pageId)
}

// Revert discards the unpublished changes of the page and returns the stored
// page.
//
// POST /v3/blogs/{blogId}/pages/{pageId}/revert
func (s *PagesService) Revert(ctx context.Context, blogId, pageId string) (*schemas.Page, error) {
	return s.action(ctx, "revert", blogId, pageId)
}

// mutate sends body to a page endpoint and decodes the stored page.
func (s *PagesService) mutate(ctx context.Context, method, path string, body *schemas.Page) (*schemas.Page, error) {
	if body == nil {
		return nil, fmt.Errorf("blogger: page body is required")
	}

	request := builder.New(method, path).Context(ctx).Body(body)

	page := &schemas.Page{}
	if err := s.do(request, page); err != nil {
		return nil, err
	}

	return page, nil
}

// action calls a page transition endpoint such as publish or revert.
func (s *PagesService) action(ctx context.Context, name, blogId, pageId string) (*schemas.Page, error) {
	if err := requireValues("blogId", blogId, "pageId", pageId); err != nil {
		return nil, err
	}

	request := builder.New(http.MethodPost, pagePath(blogId, pageId)+"/"+name).Context(ctx)

	page := &schemas.Page{}
	if err := s.do(request, page); err != nil {
		return nil, err
	}

	return page, nil
}

// do executes a pages request, failing before touching the network when the
// service has no usable client.
func (s *PagesService) do(b *builder.Builder, responseModel interface{}) error {
	if s.client == nil {
		return fmt.Errorf("blogger: pages service has no client")
	}

	_, err := b.Do(s.transport(), responseModel)
	return err
}

// transport is the HTTP client requests are sent through. It prefers the
// authenticated transport the generated service was built with, falling back
// to the default client so a service without credentials still resolves paths
// against the API root.
func (s *PagesService) transport() builder.HTTPClient {
	return pagesClient{
		httpClient: s.client.HTTPClient(),
		basePath:   s.client.BasePath(),
	}
}

// pagesClient adapts the transport to builder.HTTPClient.
type pagesClient struct {
	httpClient *http.Client
	basePath   string
}

// Do sends req through the transport, or the default client when absent.
func (c pagesClient) Do(req *http.Request) (*http.Response, error) {
	if c.httpClient == nil {
		return http.DefaultClient.Do(req)
	}
	return c.httpClient.Do(req)
}

// BasePath returns the API root the request path is resolved against.
func (c pagesClient) BasePath() string {
	if c.basePath == "" {
		return pagesBasePath
	}
	return c.basePath
}

// applyPageOptions applies every non-nil option to the request.
func applyPageOptions(b *builder.Builder, opts []PageListOption) {
	for _, opt := range opts {
		if opt != nil {
			opt(b)
		}
	}
}

// pagePath is the request path of a single page, or of one of its
// sub-resources once a suffix is appended.
func pagePath(blogId, pageId string) string {
	return fmt.Sprintf(
		"v3/blogs/%s/pages/%s",
		url.PathEscape(blogId),
		url.PathEscape(pageId),
	)
}
