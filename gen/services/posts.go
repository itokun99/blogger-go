// Code generated for github.com/itokun99/blogger-go. DO NOT EDIT.
// Implements the Blogger API v3 posts resource.

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

// postsBasePath is the Blogger v3 API root the posts paths resolve against:
// GET /v3/blogs/{blogId}/posts.
const postsBasePath = "https://blogger.googleapis.com/"

// PostsService implements the Blogger v3 posts resource.
type PostsService struct {
	client builder.TransportProvider
}

// NewPostsService returns a PostsService bound to client, which supplies the
// transport and API base path for every request.
func NewPostsService(client builder.TransportProvider) *PostsService {
	return &PostsService{client: client}
}

// PostOption customises a posts request. Options are applied in order after the
// required path and query parameters.
type PostOption func(*builder.Builder)

// WithPageToken continues a listing from the cursor returned as nextPageToken.
func WithPageToken(token string) PostOption {
	return func(b *builder.Builder) {
		b.Param(builder.PageToken, token)
	}
}

// WithMaxResults caps the number of posts returned in a single page.
func WithMaxResults(maxResults int64) PostOption {
	return func(b *builder.Builder) {
		b.Param(builder.MaxResults, strconv.FormatInt(maxResults, 10))
	}
}

// WithStartDate restricts the listing to posts published no earlier than the
// given RFC 3339 date-time.
func WithStartDate(startDate string) PostOption {
	return func(b *builder.Builder) {
		b.Param("startDate", startDate)
	}
}

// WithEndDate restricts the listing to posts published no later than the given
// RFC 3339 date-time.
func WithEndDate(endDate string) PostOption {
	return func(b *builder.Builder) {
		b.Param("endDate", endDate)
	}
}

// WithStatus restricts the listing to posts in the given status, for example
// "LIVE", "DRAFT", "SCHEDULED", or "SOFT_TRASHED".
func WithStatus(status string) PostOption {
	return func(b *builder.Builder) {
		b.Param(builder.Status, status)
	}
}

// WithLabels restricts the listing to posts carrying the given labels.
func WithLabels(labels ...string) PostOption {
	joined := joinList(labels)
	return func(b *builder.Builder) {
		if joined == "" {
			return
		}
		b.Param(builder.Labels, joined)
	}
}

// WithFetchImages controls whether the images of each post are included.
func WithFetchImages(fetchImages bool) PostOption {
	return func(b *builder.Builder) {
		b.Param("fetchImages", strconv.FormatBool(fetchImages))
	}
}

// WithFetchBody controls whether the body content is included.
func WithFetchBody(fetchBody bool) PostOption {
	return func(b *builder.Builder) {
		b.Param("fetchBody", strconv.FormatBool(fetchBody))
	}
}

// WithOrderBy sets the order of results. Valid values are "PUBLISHED" or "UPDATED".
func WithOrderBy(orderBy string) PostOption {
	return func(b *builder.Builder) {
		b.Param(builder.OrderBy, orderBy)
	}
}

// WithSortOrder sets the sort direction. Valid values are "DESCENDING" or "ASCENDING".
func WithSortOrder(sortOrder string) PostOption {
	return func(b *builder.Builder) {
		b.Param("sortOption", sortOrder)
	}
}

// WithView sets the access level for the returned resource.
// Valid values are "READER", "AUTHOR", or "ADMIN".
func WithView(view string) PostOption {
	return func(b *builder.Builder) {
		b.Param("view", view)
	}
}

// WithMaxComments sets the maximum number of comments to return.
func WithMaxComments(maxComments int64) PostOption {
	return func(b *builder.Builder) {
		b.Param("maxComments", strconv.FormatInt(maxComments, 10))
	}
}

// WithPostDate selects the post published on the given RFC 3339 date when a
// blog path resolves to more than one post.
func WithPostDate(date string) PostOption {
	return func(b *builder.Builder) {
		b.Param("date", date)
	}
}

// WithQuery sets the search query string.
func WithQuery(q string) PostOption {
	return func(b *builder.Builder) {
		b.Param("q", q)
	}
}

// Get returns the post with the given id.
//
// GET /v3/blogs/{blogId}/posts/{postId}
func (s *PostsService) Get(ctx context.Context, blogId, postId string) (*schemas.Post, error) {
	if err := requireValues("blogId", blogId, "postId", postId); err != nil {
		return nil, err
	}

	post := &schemas.Post{}
	request := builder.New(http.MethodGet, postPath(blogId, postId)).Context(ctx)

	if err := s.do(request, post); err != nil {
		return nil, err
	}

	return post, nil
}

// List returns the posts of a blog, narrowed by opts.
//
// GET /v3/blogs/{blogId}/posts
func (s *PostsService) List(ctx context.Context, blogId string, opts ...PostOption) (*schemas.PostList, error) {
	if err := requireValues("blogId", blogId); err != nil {
		return nil, err
	}

	path := fmt.Sprintf("v3/blogs/%s/posts", url.PathEscape(blogId))

	return s.list(ctx, path, opts)
}

// Search returns the posts of a blog matching the query string q.
//
// GET /v3/blogs/{blogId}/posts/search
func (s *PostsService) Search(ctx context.Context, blogId, q string, opts ...PostOption) (*schemas.PostList, error) {
	if err := requireValues("blogId", blogId, "q", q); err != nil {
		return nil, err
	}

	path := fmt.Sprintf("v3/blogs/%s/posts/search", url.PathEscape(blogId))

	return s.list(ctx, path, opts, WithQuery(q))
}

// GetByPath returns the post published at a blog path, narrowed by opts.
//
// GET /v3/blogs/{blogId}/posts/bypath?path={path}
func (s *PostsService) GetByPath(ctx context.Context, blogId, path string, opts ...PostOption) (*schemas.Post, error) {
	if err := requireValues("blogId", blogId, "path", path); err != nil {
		return nil, err
	}

	request := builder.New(http.MethodGet, fmt.Sprintf("v3/blogs/%s/posts/bypath", url.PathEscape(blogId))).
		Context(ctx).
		Param("path", path)
	applyOptions(request, opts)

	post := &schemas.Post{}
	if err := s.do(request, post); err != nil {
		return nil, err
	}

	return post, nil
}

// Insert creates a post on the blog and returns the stored post.
//
// POST /v3/blogs/{blogId}/posts
func (s *PostsService) Insert(ctx context.Context, blogId string, body *schemas.Post) (*schemas.Post, error) {
	if err := requireValues("blogId", blogId); err != nil {
		return nil, err
	}

	path := fmt.Sprintf("v3/blogs/%s/posts", url.PathEscape(blogId))

	return s.mutate(ctx, http.MethodPost, path, body)
}

// Update replaces the post and returns the stored post.
//
// PUT /v3/blogs/{blogId}/posts/{postId}
func (s *PostsService) Update(ctx context.Context, blogId, postId string, body *schemas.Post) (*schemas.Post, error) {
	if err := requireValues("blogId", blogId, "postId", postId); err != nil {
		return nil, err
	}

	return s.mutate(ctx, http.MethodPut, postPath(blogId, postId), body)
}

// Patch updates the provided fields of the post and returns the stored post.
//
// PATCH /v3/blogs/{blogId}/posts/{postId}
func (s *PostsService) Patch(ctx context.Context, blogId, postId string, body *schemas.Post) (*schemas.Post, error) {
	if err := requireValues("blogId", blogId, "postId", postId); err != nil {
		return nil, err
	}

	return s.mutate(ctx, http.MethodPatch, postPath(blogId, postId), body)
}

// Delete removes the post.
//
// DELETE /v3/blogs/{blogId}/posts/{postId}
func (s *PostsService) Delete(ctx context.Context, blogId, postId string) error {
	if err := requireValues("blogId", blogId, "postId", postId); err != nil {
		return err
	}

	request := builder.New(http.MethodDelete, postPath(blogId, postId)).Context(ctx)
	if err := s.do(request, nil); err != nil {
		return err
	}

	return nil
}

// Publish makes the post publicly visible and returns it.
//
// POST /v3/blogs/{blogId}/posts/{postId}/publish
func (s *PostsService) Publish(ctx context.Context, blogId, postId string) (*schemas.Post, error) {
	return s.action(ctx, "publish", blogId, postId)
}

// Revert discards the unpublished changes of the post and returns the stored
// post.
//
// POST /v3/blogs/{blogId}/posts/{postId}/revert
func (s *PostsService) Revert(ctx context.Context, blogId, postId string) (*schemas.Post, error) {
	return s.action(ctx, "revert", blogId, postId)
}

// list sends a listing request and decodes the returned page of posts.
func (s *PostsService) list(ctx context.Context, path string, opts []PostOption, required ...PostOption) (*schemas.PostList, error) {
	request := builder.New(http.MethodGet, path).Context(ctx)
	applyOptions(request, append(required, opts...))

	list := &schemas.PostList{Items: []*schemas.Post{}}
	if err := s.do(request, list); err != nil {
		return nil, err
	}

	return list, nil
}

// mutate sends body to a post endpoint and decodes the stored post.
func (s *PostsService) mutate(ctx context.Context, method, path string, body *schemas.Post) (*schemas.Post, error) {
	if body == nil {
		return nil, fmt.Errorf("blogger: post body is required")
	}

	request := builder.New(method, path).Context(ctx).Body(body)

	post := &schemas.Post{}
	if err := s.do(request, post); err != nil {
		return nil, err
	}

	return post, nil
}

// action calls a post transition endpoint such as publish or revert.
func (s *PostsService) action(ctx context.Context, name, blogId, postId string) (*schemas.Post, error) {
	if err := requireValues("blogId", blogId, "postId", postId); err != nil {
		return nil, err
	}

	request := builder.New(http.MethodPost, postPath(blogId, postId)+"/"+name).Context(ctx)

	post := &schemas.Post{}
	if err := s.do(request, post); err != nil {
		return nil, err
	}

	return post, nil
}

// do executes a posts request, failing before touching the network when the
// service has no usable client.
func (s *PostsService) do(b *builder.Builder, responseModel interface{}) error {
	if s.client == nil {
		return fmt.Errorf("blogger: posts service has no client")
	}

	_, err := b.Do(s.transport(), responseModel)
	return err
}

// transport is the HTTP client requests are sent through. It carries the
// transport of the generated service so authenticated requests keep working;
// the base path resolves against the Blogger API root.
func (s *PostsService) transport() builder.HTTPClient {
	return postsClient{
		httpClient: s.client.HTTPClient(),
		basePath:   s.client.BasePath(),
	}
}

// postsClient resolves relative builder paths against the Blogger API root.
// postsClient adapts the transport to builder.HTTPClient.
type postsClient struct {
	httpClient *http.Client
	basePath   string
}

// Do sends req through the transport, or the default client when absent.
func (c postsClient) Do(req *http.Request) (*http.Response, error) {
	if c.httpClient != nil {
		return c.httpClient.Do(req)
	}
	return http.DefaultClient.Do(req)
}

// BasePath returns the API root the request path is resolved against.
func (c postsClient) BasePath() string {
	if c.basePath != "" {
		return c.basePath
	}
	return postsBasePath
}

// applyOptions applies every non-nil option to the request.
func applyOptions(b *builder.Builder, opts []PostOption) {
	for _, opt := range opts {
		if opt != nil {
			opt(b)
		}
	}
}

// postPath is the request path of a single post, or of one of its
// sub-resources once a suffix is appended.
func postPath(blogId, postId string) string {
	return fmt.Sprintf(
		"v3/blogs/%s/posts/%s",
		url.PathEscape(blogId),
		url.PathEscape(postId),
	)
}
