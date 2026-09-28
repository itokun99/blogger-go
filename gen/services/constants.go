package services

// PostStatus values accepted by posts.list.
const (
	PostStatusLive        = "LIVE"
	PostStatusDraft       = "DRAFT"
	PostStatusScheduled   = "SCHEDULED"
	PostStatusSoftTrashed = "SOFT_TRASHED"
)

// CommentStatus values accepted by comments.list.
const (
	CommentStatusLive    = "LIVE"
	CommentStatusEmptied = "EMPTIED"
	CommentStatusPending = "PENDING"
	CommentStatusSpam    = "SPAM"
)

// View values accepted by the view parameter of the list and get methods.
const (
	ViewTypeUnspecified = "VIEW_TYPE_UNSPECIFIED"
	ViewReader          = "READER"
	ViewAuthor          = "AUTHOR"
	ViewAdmin           = "ADMIN"
)

// OrderBy values accepted by posts.list.
const (
	OrderByPublished = "PUBLISHED"
	OrderByUpdated   = "UPDATED"
)

// SortOrder values accepted by posts.list.
const (
	SortOrderDescending = "DESCENDING"
	SortOrderAscending  = "ASCENDING"
)
