package services_test

import (
	"testing"

	"github.com/itokun99/blogger-go/gen/services"
)

func TestStatusAndViewConstantValues(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"PostStatusLive", services.PostStatusLive, "LIVE"},
		{"PostStatusDraft", services.PostStatusDraft, "DRAFT"},
		{"PostStatusScheduled", services.PostStatusScheduled, "SCHEDULED"},
		{"PostStatusSoftTrashed", services.PostStatusSoftTrashed, "SOFT_TRASHED"},

		{"CommentStatusLive", services.CommentStatusLive, "LIVE"},
		{"CommentStatusEmptied", services.CommentStatusEmptied, "EMPTIED"},
		{"CommentStatusPending", services.CommentStatusPending, "PENDING"},
		{"CommentStatusSpam", services.CommentStatusSpam, "SPAM"},

		{"PageStatusLive", services.PageStatusLive, "LIVE"},
		{"PageStatusDraft", services.PageStatusDraft, "DRAFT"},
		{"PageStatusSoftTrashed", services.PageStatusSoftTrashed, "SOFT_TRASHED"},

		{"ViewTypeUnspecified", services.ViewTypeUnspecified, "VIEW_TYPE_UNSPECIFIED"},
		{"ViewReader", services.ViewReader, "READER"},
		{"ViewAuthor", services.ViewAuthor, "AUTHOR"},
		{"ViewAdmin", services.ViewAdmin, "ADMIN"},

		{"OrderByPublished", services.OrderByPublished, "PUBLISHED"},
		{"OrderByUpdated", services.OrderByUpdated, "UPDATED"},

		{"SortOrderDescending", services.SortOrderDescending, "DESCENDING"},
		{"SortOrderAscending", services.SortOrderAscending, "ASCENDING"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("%s = %q, want %q", tt.name, tt.got, tt.want)
			}
		})
	}
}
