// Code generated for github.com/itokun99/blogger-go. DO NOT EDIT.
// Implements the Blogger API v3 comments resource.

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

// commentsBasePath is the Blogger v3 API root the comments paths resolve
// against: GET /v3/blogs/{blogId}/posts/{postId}/comments/{commentId}.
const commentsBasePath = "https://blogger.googleapis.com/"

// CommentsService implements the Blogger v3 comments resource.
type CommentsService struct {
	client builder.TransportProvider
}

// NewCommentsService returns a CommentsService bound to client, which supplies
// the transport and API base path for every request.
func NewCommentsService(client builder.TransportProvider) *CommentsService {
	return &CommentsService{client: client}
}

// CommentListOption customises a comments list request. Options are applied in
// order after the required path parameters.
type CommentListOption func(*builder.Builder)

// WithCommentPageToken continues a listing from the cursor returned as
// nextPageToken.
func WithCommentPageToken(token string) CommentListOption {
	return func(b *builder.Builder) {
		b.Param(builder.PageToken, token)
	}
}

// WithCommentMaxResults caps the number of comments returned in a single page.
func WithCommentMaxResults(maxResults int64) CommentListOption {
	return func(b *builder.Builder) {
		b.Param(builder.MaxResults, strconv.FormatInt(maxResults, 10))
	}
}

// WithCommentStatus restricts the listing to comments in the given status, for
// example "LIVE" or "SPAM".
func WithCommentStatus(status string) CommentListOption {
	return func(b *builder.Builder) {
		b.Param(builder.Status, status)
	}
}

// WithCommentPost restricts the listing to the comments of a single post.
func WithCommentPost(postId string) CommentListOption {
	return func(b *builder.Builder) {
		b.Param("post", postId)
	}
}

// WithCommentView sets the access level for the returned resource.
// Valid values are "READER", "AUTHOR", or "ADMIN".
func WithCommentView(view string) CommentListOption {
	return func(b *builder.Builder) {
		b.Param("view", view)
	}
}

// WithCommentStartDate restricts the listing to comments published no earlier
// than the given RFC 3339 date-time.
func WithCommentStartDate(startDate string) CommentListOption {
	return func(b *builder.Builder) {
		b.Param("startDate", startDate)
	}
}

// WithCommentEndDate restricts the listing to comments published no later than
// the given RFC 3339 date-time.
func WithCommentEndDate(endDate string) CommentListOption {
	return func(b *builder.Builder) {
		b.Param("endDate", endDate)
	}
}

// WithCommentFetchBodies controls whether the comment bodies are included.
func WithCommentFetchBodies(fetchBodies bool) CommentListOption {
	return func(b *builder.Builder) {
		b.Param("fetchBodies", strconv.FormatBool(fetchBodies))
	}
}

// Get returns the comment identified by commentId on the given post.
//
// GET /v3/blogs/{blogId}/posts/{postId}/comments/{commentId}
func (s *CommentsService) Get(ctx context.Context, blogId, postId, commentId string) (*schemas.Comment, error) {
	comment := &schemas.Comment{}

	_, err := s.do(
		builder.New(http.MethodGet, commentPath(blogId, postId, commentId)).
			Context(ctx),
		comment,
	)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// List returns the comments of a post, narrowed by opts.
//
// GET /v3/blogs/{blogId}/posts/{postId}/comments
func (s *CommentsService) List(ctx context.Context, blogId, postId string, opts ...CommentListOption) (*schemas.CommentList, error) {
	return s.list(ctx, commentListPath(blogId, postId), opts)
}

// ListByBlog returns the comments of every post in the blog, narrowed by opts.
// Use WithCommentPost to restrict the result to a single post.
//
// GET /v3/blogs/{blogId}/comments
func (s *CommentsService) ListByBlog(ctx context.Context, blogId string, opts ...CommentListOption) (*schemas.CommentList, error) {
	return s.list(ctx, blogCommentListPath(blogId), opts)
}

// Approve marks the comment as visible and returns it.
//
// POST /v3/blogs/{blogId}/posts/{postId}/comments/{commentId}/approve
func (s *CommentsService) Approve(ctx context.Context, blogId, postId, commentId string) (*schemas.Comment, error) {
	return s.action(ctx, "approve", blogId, postId, commentId)
}

// Delete removes the comment.
//
// DELETE /v3/blogs/{blogId}/posts/{postId}/comments/{commentId}
func (s *CommentsService) Delete(ctx context.Context, blogId, postId, commentId string) error {
	_, err := s.do(
		builder.New(http.MethodDelete, commentPath(blogId, postId, commentId)).
			Context(ctx),
		nil,
	)
	return err
}

// MarkAsSpam marks the comment as spam and returns it.
//
// POST /v3/blogs/{blogId}/posts/{postId}/comments/{commentId}/spam
func (s *CommentsService) MarkAsSpam(ctx context.Context, blogId, postId, commentId string) (*schemas.Comment, error) {
	return s.action(ctx, "spam", blogId, postId, commentId)
}

// RemoveContent removes the comment content while keeping the comment itself
// in place, and returns the updated comment.
//
// POST /v3/blogs/{blogId}/posts/{postId}/comments/{commentId}/removecontent
func (s *CommentsService) RemoveContent(ctx context.Context, blogId, postId, commentId string) (*schemas.Comment, error) {
	return s.action(ctx, "removecontent", blogId, postId, commentId)
}

// list runs a comments listing against path, applying opts on top of the
// pagination defaults.
func (s *CommentsService) list(ctx context.Context, path string, opts []CommentListOption) (*schemas.CommentList, error) {
	request := builder.New(http.MethodGet, path).Context(ctx)

	for _, opt := range opts {
		if opt != nil {
			opt(request)
		}
	}

	// Items is initialized explicitly so an empty page serializes as [] rather
	// than null.
	list := &schemas.CommentList{Items: []*schemas.Comment{}}
	if _, err := s.do(request, list); err != nil {
		return nil, err
	}

	return list, nil
}

// action sends a POST to a comment sub-resource such as approve or spam.
func (s *CommentsService) action(ctx context.Context, name, blogId, postId, commentId string) (*schemas.Comment, error) {
	comment := &schemas.Comment{}

	_, err := s.do(
		builder.New(http.MethodPost, commentPath(blogId, postId, commentId)+"/"+name).
			Context(ctx),
		comment,
	)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// do executes a comments request, failing before touching the network when the
// service has no usable client.
func (s *CommentsService) do(b *builder.Builder, responseModel interface{}) (interface{}, error) {
	if s.client == nil {
		return nil, fmt.Errorf("blogger: comments service has no client")
	}
	return b.Do(s.transport(), responseModel)
}

// transport is the HTTP client requests are sent through. It prefers the
// caller-supplied HTTP client so its transport is reused, falling back to the
// default client so a service built without credentials still resolves paths
// against the API root.
func (s *CommentsService) transport() builder.HTTPClient {
	return commentsClient{
		httpClient: s.client.HTTPClient(),
		basePath:   s.client.BasePath(),
	}
}

// commentsClient adapts the transport to builder.HTTPClient.
type commentsClient struct {
	httpClient *http.Client
	basePath   string
}

// Do sends req through the client transport, or the default client when absent.
func (c commentsClient) Do(req *http.Request) (*http.Response, error) {
	if c.httpClient != nil {
		return c.httpClient.Do(req)
	}
	return http.DefaultClient.Do(req)
}

// BasePath returns the API root the request path is resolved against.
func (c commentsClient) BasePath() string {
	if c.basePath != "" {
		return c.basePath
	}
	return commentsBasePath
}

// commentPath is the request path of a single comment, or of one of its
// sub-resources once a suffix is appended.
func commentPath(blogId, postId, commentId string) string {
	return commentListPath(blogId, postId) + "/" + url.PathEscape(commentId)
}

// commentListPath is the request path of the comments of a single post.
func commentListPath(blogId, postId string) string {
	return blogCommentListPath(blogId) + "/posts/" + url.PathEscape(postId) + "/comments"
}

// blogCommentListPath is the request path of every comment in a blog.
func blogCommentListPath(blogId string) string {
	return "v3/blogs/" + url.PathEscape(blogId) + "/comments"
}
