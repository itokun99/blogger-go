package blogger

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Scopes are the OAuth2 scopes required by the Blogger API v3.
var Scopes = []string{
	"https://www.googleapis.com/auth/blogger",
	"https://www.googleapis.com/auth/blogger.readonly",
}

// TokenSource is the OAuth2 token source used to authenticate requests.
type TokenSource = oauth2.TokenSource

// NewHTTPClient returns an HTTP client that authenticates every request with ts.
func NewHTTPClient(ctx context.Context, ts TokenSource) *http.Client {
	return oauth2.NewClient(ctx, ts)
}

// TokenSourceFromJSON builds a TokenSource from a Google OAuth2 client secrets
// file (the "credentials.json" downloaded from Google Cloud Console). If
// tokenFile already holds a cached token it is reused, and refreshed tokens are
// written back to it.
func TokenSourceFromJSON(credentialsFile, tokenFile string) (TokenSource, error) {
	data, err := os.ReadFile(credentialsFile)
	if err != nil {
		return nil, fmt.Errorf("blogger: read credentials file: %w", err)
	}

	cfg, err := google.ConfigFromJSON(data, Scopes...)
	if err != nil {
		return nil, fmt.Errorf("blogger: parse credentials file: %w", err)
	}

	if tokenFile == "" {
		return cfg.TokenSource(context.Background(), nil), nil
	}

	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		return nil, fmt.Errorf("blogger: read token file: %w", err)
	}
	return cfg.TokenSource(context.Background(), tok), nil
}

func tokenFromFile(path string) (*oauth2.Token, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tok := &oauth2.Token{}
	if err := json.NewDecoder(f).Decode(tok); err != nil {
		return nil, err
	}
	return tok, nil
}
