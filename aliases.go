package blogger

import (
	"github.com/itokun99/blogger-go/gen/schemas"
	"github.com/itokun99/blogger-go/gen/services"
)

type (
	PostsOption     = services.PostOption
	CommentsOption  = services.CommentListOption
	PagesOption     = services.PageListOption
	BlogsOption     = services.BlogsOption
	PageViewsOption = services.PageViewsOption
)

var (
	WithMaxResults  = services.WithMaxResults
	WithStatus      = services.WithStatus
	WithOrderBy     = services.WithOrderBy
	WithView        = services.WithView
	WithSortOrder   = services.WithSortOrder
	WithPageToken   = services.WithPageToken
	WithStartDate   = services.WithStartDate
	WithEndDate     = services.WithEndDate
	WithLabels      = services.WithLabels
	WithFetchImages = services.WithFetchImages
	WithFetchBody   = services.WithFetchBody
	WithMaxComments = services.WithMaxComments
	WithPostDate    = services.WithPostDate
	WithQuery       = services.WithQuery
)

var (
	WithCommentPageToken   = services.WithCommentPageToken
	WithCommentMaxResults  = services.WithCommentMaxResults
	WithCommentStatus      = services.WithCommentStatus
	WithCommentPost        = services.WithCommentPost
	WithCommentView        = services.WithCommentView
	WithCommentStartDate   = services.WithCommentStartDate
	WithCommentEndDate     = services.WithCommentEndDate
	WithCommentFetchBodies = services.WithCommentFetchBodies
)

var (
	WithPageListToken      = services.WithPageListToken
	WithPageListMaxResults = services.WithPageListMaxResults
	WithPageListStatus     = services.WithPageListStatus
	WithPageView           = services.WithPageView
	WithPageFetchBodies    = services.WithPageFetchBodies
)

var (
	WithBlogView     = services.WithBlogView
	WithBlogMaxPosts = services.WithBlogMaxPosts
)

var (
	WithPageViewRange = services.WithPageViewRange
)

const (
	PostStatusLive        = services.PostStatusLive
	PostStatusDraft       = services.PostStatusDraft
	PostStatusScheduled   = services.PostStatusScheduled
	PostStatusSoftTrashed = services.PostStatusSoftTrashed

	CommentStatusLive    = services.CommentStatusLive
	CommentStatusEmptied = services.CommentStatusEmptied
	CommentStatusPending = services.CommentStatusPending
	CommentStatusSpam    = services.CommentStatusSpam

	PageStatusLive        = services.PageStatusLive
	PageStatusDraft       = services.PageStatusDraft
	PageStatusSoftTrashed = services.PageStatusSoftTrashed

	ViewTypeUnspecified = services.ViewTypeUnspecified
	ViewReader          = services.ViewReader
	ViewAuthor          = services.ViewAuthor
	ViewAdmin           = services.ViewAdmin

	OrderByPublished = services.OrderByPublished
	OrderByUpdated   = services.OrderByUpdated

	SortOrderDescending = services.SortOrderDescending
	SortOrderAscending  = services.SortOrderAscending
)

type (
	Post    = schemas.Post
	Comment = schemas.Comment
	Page    = schemas.Page
	Blog    = schemas.Blog
	User    = schemas.User

	PostList          = schemas.PostList
	CommentList       = schemas.CommentList
	PageList          = schemas.PageList
	BlogList          = schemas.BlogList
	PostUserInfo      = schemas.PostUserInfo
	PostUserInfosList = schemas.PostUserInfosList
)
