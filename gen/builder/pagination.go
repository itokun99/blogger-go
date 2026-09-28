package builder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Token is the opaque cursor Blogger returns as nextPageToken and accepts as
// the pageToken query parameter. The constant of the same wire name lives in
// options.go as PageToken, so the type is named Token to avoid the clash.
type Token string

// String returns the token as a plain string.
func (t Token) String() string {
	return string(t)
}

// IsZero reports whether the token is empty, meaning there is no next page.
func (t Token) IsZero() bool {
	return t == ""
}

// Query returns the token encoded for use in a URL query string.
func (t Token) Query() string {
	return url.QueryEscape(string(t))
}

// PaginationResult is the page envelope every Blogger list response is decoded
// into. Items stays raw so callers can decode it into their concrete model.
type PaginationResult struct {
	Items         []byte `json:"items"`
	NextPageToken string `json:"nextPageToken"`
	TotalItems    int64  `json:"totalItems"`
	ETag          string `json:"etag"`
	Kind          string `json:"kind"`
}

// NewPageToken validates a raw cursor and wraps it as a Token.
func NewPageToken(raw string) Token {
	return Token(strings.TrimSpace(raw))
}

// Next returns the cursor for the following page, or an empty token when the
// current page was the last one.
func (r PaginationResult) Next() Token {
	return NewPageToken(r.NextPageToken)
}

// HasMore reports whether another page can be requested.
func (r PaginationResult) HasMore() bool {
	return r.NextPageToken != ""
}

// Decode unmarshals the raw Items payload into v, returning an empty slice
// rather than nil when the page carried no items.
func (r PaginationResult) Decode(v interface{}) error {
	if v == nil {
		return fmt.Errorf("blogger: decode target is nil")
	}

	payload := bytes.TrimSpace(r.Items)
	if len(payload) == 0 {
		payload = []byte("[]")
	}

	return json.Unmarshal(payload, v)
}

// ParsePaginationResult decodes a raw list response body into a
// PaginationResult, tolerating a response that is itself a bare items array.
func ParsePaginationResult(payload []byte) (*PaginationResult, error) {
	result := &PaginationResult{Items: []byte("[]")}

	trimmed := bytes.TrimSpace(payload)
	if len(trimmed) == 0 {
		return result, nil
	}

	if trimmed[0] == '[' {
		result.Items = append([]byte{}, trimmed...)
		var items []json.RawMessage
		if err := json.Unmarshal(trimmed, &items); err != nil {
			return nil, err
		}
		result.TotalItems = int64(len(items))
		return result, nil
	}

	var envelope struct {
		Items         json.RawMessage `json:"items"`
		NextPageToken string          `json:"nextPageToken"`
		TotalItems    int64           `json:"totalItems"`
		ETag          string          `json:"etag"`
		Kind          string          `json:"kind"`
	}
	if err := json.Unmarshal(trimmed, &envelope); err != nil {
		return nil, err
	}

	if len(envelope.Items) > 0 {
		result.Items = append([]byte{}, envelope.Items...)
	}
	result.NextPageToken = envelope.NextPageToken
	result.TotalItems = envelope.TotalItems
	result.ETag = envelope.ETag
	result.Kind = envelope.Kind

	return result, nil
}

// PageTokenFromResponse extracts nextPageToken from a raw list response body.
func PageTokenFromResponse(payload []byte) (Token, error) {
	result, err := ParsePaginationResult(payload)
	if err != nil {
		return "", err
	}
	return result.Next(), nil
}

// Tokens walks a paginated endpoint by repeatedly calling fetch until the API
// stops returning a nextPageToken or maxPages pages have been collected.
func Tokens(maxPages int, fetch func(token Token) (*PaginationResult, error)) ([]*PaginationResult, error) {
	if fetch == nil {
		return nil, fmt.Errorf("blogger: fetch function is nil")
	}

	pages := []*PaginationResult{}
	token := Token("")

	for {
		page, err := fetch(token)
		if err != nil {
			return pages, err
		}
		if page == nil {
			return pages, nil
		}

		pages = append(pages, page)

		if !page.HasMore() {
			return pages, nil
		}
		if maxPages > 0 && len(pages) >= maxPages {
			return pages, nil
		}

		token = page.Next()
	}
}

func joinList(values []string) string {
	cleaned := []string{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" {
			cleaned = append(cleaned, value)
		}
	}
	return strings.Join(cleaned, ",")
}

func boolParam(value bool) string {
	return strconv.FormatBool(value)
}

func itoa(value int64) string {
	return strconv.FormatInt(value, 10)
}
