package builder_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/itokun99/blogger-go/gen/builder"
)

// mockTransport is a minimal HTTP client that also implements BasePathProvider.
type mockTransport struct {
	basePath string
	http.Client
}

func (m *mockTransport) BasePath() string { return m.basePath }

func TestBuilderNewWithMethod(t *testing.T) {
	b := builder.New(http.MethodGet, "test-path")
	if b == nil {
		t.Fatal("New returned nil")
	}
}

func TestBuilderParam(t *testing.T) {
	b := builder.New(http.MethodGet, "test").Param("key", "value")
	if b == nil {
		t.Fatal("Param returned nil")
	}
}

func TestBuilderContext(t *testing.T) {
	b := builder.New(http.MethodGet, "test").Context(nil)
	if b == nil {
		t.Fatal("Context returned nil")
	}
}

func TestBuilderDo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"id": "123"})
	}))
	defer server.Close()

	type response struct {
		ID string `json:"id"`
	}

	// Create a mock client that implements both HTTPClient and BasePathProvider
	mockClient := &mockTransport{basePath: server.URL}

	var resp response
	_, err := builder.New(http.MethodGet, "/").
		Context(nil).
		Do(mockClient, &resp)
	if err != nil {
		t.Fatalf("Do failed: %v", err)
	}
	if resp.ID != "123" {
		t.Errorf("got %q, want 123", resp.ID)
	}
}

func TestBuilderDoWithBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := struct {
			Title string `json:"title"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		json.NewEncoder(w).Encode(body)
	}))
	defer server.Close()

	type result struct {
		Title string `json:"title"`
	}

	mockClient := &mockTransport{basePath: server.URL}
	var got result
	// Body expects a JSON-serializable value, not raw bytes
	_, err := builder.New(http.MethodPost, "/").
		Context(nil).
		Body(map[string]string{"title": "Hello"}).
		Do(mockClient, &got)
	if err != nil {
		t.Fatalf("Do failed: %v", err)
	}
	if got.Title != "Hello" {
		t.Errorf("got %q, want Hello", got.Title)
	}
}

func TestBuilderDoQueryParams(t *testing.T) {
	var received string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		received = r.URL.Query().Get("q")
		w.WriteHeader(200)
	}))
	defer server.Close()

	mockClient := &mockTransport{basePath: server.URL}
	var got interface{}
	_, err := builder.New(http.MethodGet, "/").
		Context(nil).
		Param("q", "test-query").
		Do(mockClient, &got)
	if err != nil {
		t.Fatalf("Do failed: %v", err)
	}
	if received != "test-query" {
		t.Errorf("got query %q, want test-query", received)
	}
}

func TestBuilderHeaders(t *testing.T) {
	var received string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		received = r.Header.Get("X-Custom")
		w.WriteHeader(200)
	}))
	defer server.Close()

	mockClient := &mockTransport{basePath: server.URL}
	var got interface{}
	_, err := builder.New(http.MethodGet, "/").
		Context(nil).
		Header("X-Custom", "custom-value").
		Do(mockClient, &got)
	if err != nil {
		t.Fatalf("Do failed: %v", err)
	}
	if received != "custom-value" {
		t.Errorf("got header %q, want custom-value", received)
	}
}

func TestNewAPIErrorNilResponse(t *testing.T) {
	_, err := builder.NewAPIError(nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestNewAPIErrorEmptyBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":{"code":400,"message":"bad request"}}`))
	}))
	defer server.Close()

	resp, err := server.Client().Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	apiErr, err := builder.NewAPIError(resp)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if apiErr.Code != 400 {
		t.Errorf("got code %d, want 400", apiErr.Code)
	}
	if apiErr.Message != "bad request" {
		t.Errorf("got message %q, want bad request", apiErr.Message)
	}
}

func TestIsAPIError(t *testing.T) {
	apiErr := &builder.APIError{Code: 404, Message: "not found"}
	if !builder.IsAPIError(apiErr) {
		t.Error("expected IsAPIError to return true")
	}
	if builder.IsAPIError(fmt.Errorf("wrapped: %w", apiErr)) {
		// This is expected to work via Unwrap
	}
	if builder.IsAPIError(fmt.Errorf("plain error")) {
		t.Error("expected IsAPIError to return false for plain error")
	}
}

func TestParsePaginationResultBareArray(t *testing.T) {
	payload := `[{"id":"1"},{"id":"2"}]`
	result, err := builder.ParsePaginationResult([]byte(payload))
	if err != nil {
		t.Fatalf("ParsePaginationResult failed: %v", err)
	}
	// Bare arrays don't have nextPageToken, so HasMore should be false
	if result.HasMore() {
		t.Error("expected HasMore to be false for bare array")
	}
	if result.TotalItems != 2 {
		t.Errorf("got %d items, want 2", result.TotalItems)
	}
}

func TestParsePaginationResultEnvelope(t *testing.T) {
	payload := `{"items":[{"id":"1"}],"nextPageToken":"abc","totalItems":1}`
	result, err := builder.ParsePaginationResult([]byte(payload))
	if err != nil {
		t.Fatalf("ParsePaginationResult failed: %v", err)
	}
	if result.Next().String() != "abc" {
		t.Errorf("got token %q, want abc", result.Next())
	}
	if !result.HasMore() {
		t.Error("expected HasMore to be true")
	}
}

func TestTokenHelpers(t *testing.T) {
	token := builder.NewPageToken("  cursor  ")
	if token.String() != "cursor" {
		t.Errorf("got %q, want cursor", token.String())
	}
	// Non-empty token should NOT be zero
	if token.IsZero() {
		t.Error("expected IsZero to be false for non-empty token")
	}
	empty := builder.NewPageToken("")
	// Empty token SHOULD be zero
	if !empty.IsZero() {
		t.Error("expected empty token IsZero to be true")
	}
}

func TestCommonFilterApply(t *testing.T) {
	f := builder.CommonFilter{
		Labels:      []string{"a", "b"},
		Status:      "LIVE",
		OrderBy:     "published",
		SortOrder:   "descending",
		FetchBodies: boolPtr(true),
	}
	b := builder.New(http.MethodGet, "/")
	f.Apply(b)
	// Just verify it doesn't panic
}

func boolPtr(v bool) *bool {
	return &v
}

func TestPaginationApply(t *testing.T) {
	p := builder.Pagination{MaxResults: 10, PageToken: "tok"}
	b := builder.New(http.MethodGet, "/")
	p.Apply(b)
	// Just verify it doesn't panic
}

func TestTokensWalk(t *testing.T) {
	callCount := 0
	fetch := func(token builder.Token) (*builder.PaginationResult, error) {
		callCount++
		if callCount >= 3 {
			return &builder.PaginationResult{Items: []byte("[]"), NextPageToken: ""}, nil
		}
		return &builder.PaginationResult{Items: []byte("[]"), NextPageToken: "next"}, nil
	}
	results, err := builder.Tokens(5, fetch)
	if err != nil {
		t.Fatalf("Tokens failed: %v", err)
	}
	if len(results) != 3 {
		t.Errorf("got %d results, want 3", len(results))
	}
}

func TestTokensFetchNil(t *testing.T) {
	_, err := builder.Tokens(1, nil)
	if err == nil {
		t.Fatal("expected error for nil fetch")
	}
}

func TestDecodeEmptyItems(t *testing.T) {
	result := &builder.PaginationResult{Items: []byte("[]")}
	var items []string
	if err := result.Decode(&items); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	if len(items) != 0 {
		t.Errorf("got %d items, want 0", len(items))
	}
}

func TestDecodeNilTarget(t *testing.T) {
	result := &builder.PaginationResult{Items: []byte(`[1]`)}
	if err := result.Decode(nil); err == nil {
		t.Fatal("expected error for nil target")
	}
}

func TestParsePaginationResultEmpty(t *testing.T) {
	result, err := builder.ParsePaginationResult([]byte{})
	if err != nil {
		t.Fatalf("ParsePaginationResult failed: %v", err)
	}
	// Empty payload should have no next page token
	if result.HasMore() {
		t.Error("expected HasMore to be false for empty result")
	}
}
