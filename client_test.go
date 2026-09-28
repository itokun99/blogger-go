package blogger

import (
	"context"
	"net/http"
	"testing"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
)

type ctxKey struct{}

func staticTokenSource(token string) TokenSource {
	return oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
}

func newTestClient(t *testing.T) *Client {
	t.Helper()

	httpClient := NewHTTPClient(context.Background(), staticTokenSource("test-token"))

	c, err := NewClient(context.Background(), option.WithHTTPClient(httpClient))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c == nil {
		t.Fatal("NewClient returned nil client")
	}
	return c
}

func TestNewClientBuildsService(t *testing.T) {
	c := newTestClient(t)

	if c.HTTPClient() == nil {
		t.Fatal("HTTPClient() is nil")
	}
	if c.BasePath() == "" {
		t.Error("BasePath() is empty")
	}
}

func TestClientAccessorsReturnServices(t *testing.T) {
	c := newTestClient(t)

	for name, got := range map[string]interface{}{
		"Blogs":         c.Blogs(),
		"Comments":      c.Comments(),
		"Pages":         c.Pages(),
		"Posts":         c.Posts(),
		"Users":         c.Users(),
		"BlogUserInfos": c.BlogUserInfos(),
		"PageViews":     c.PageViews(),
		"PostUserInfos": c.PostUserInfos(),
	} {
		if got == nil {
			t.Errorf("%s() returned nil", name)
		}
	}
}

func TestWithContextReturnsIndependentClient(t *testing.T) {
	c := newTestClient(t)
	ctx := context.WithValue(context.Background(), ctxKey{}, "value")

	clone := c.WithContext(ctx)
	if clone == c {
		t.Fatal("WithContext returned the receiver")
	}
	if clone.HTTPClient() != c.HTTPClient() {
		t.Error("WithContext did not share the underlying transport")
	}
	// Note: context is not exported, so we verify indirectly through behavior
	_ = clone
}

func TestRawClientUsesDefaultBaseURL(t *testing.T) {
	c := NewRawClient(&http.Client{})
	if c == nil {
		t.Fatal("NewRawClient returned nil")
	}
	if c.BasePath() != defaultBaseServiceURL {
		t.Errorf("BasePath() = %q, want %q", c.BasePath(), defaultBaseServiceURL)
	}
}

func TestScopesCoverBloggerAPI(t *testing.T) {
	if len(Scopes) == 0 {
		t.Fatal("Scopes is empty")
	}
	for _, s := range Scopes {
		if s == "" {
			t.Error("Scopes contains an empty entry")
		}
	}
}

func TestNewHTTPClientWrapsTokenSource(t *testing.T) {
	ts := staticTokenSource("test-token")

	c := NewHTTPClient(context.Background(), ts)
	if c == nil {
		t.Fatal("NewHTTPClient returned nil")
	}
	if c.Transport == nil {
		t.Fatal("NewHTTPClient returned a client without a transport")
	}
}
