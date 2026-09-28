package builder

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// HTTPClient is the transport the builder sends requests through. Every
// generated client satisfies it.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// BasePathProvider exposes the API root the builder resolves relative paths
// against.
type BasePathProvider interface {
	BasePath() string
}

// Builder accumulates the request state shared by every generated API call and
// executes it on demand. The zero value is not usable; create one with New.
type Builder struct {
	basePath    string
	method      string
	queryParams map[string]string
	headers     map[string]string
	body        interface{}
	ctx         context.Context
	client      HTTPClient
}

// New returns a Builder for the given HTTP method and path. The path may be
// either absolute or relative to the client's base path.
func New(method, path string) *Builder {
	return &Builder{
		basePath:    path,
		method:      strings.ToUpper(method),
		queryParams: map[string]string{},
		headers:     map[string]string{},
		ctx:         context.Background(),
	}
}

// Context attaches a context to the request. A nil context is ignored so the
// chain never loses its default.
func (b *Builder) Context(ctx context.Context) *Builder {
	if ctx != nil {
		b.ctx = ctx
	}
	return b
}

// Param adds a query parameter to the request. An empty key is ignored.
func (b *Builder) Param(key, value string) *Builder {
	if key == "" {
		return b
	}
	b.queryParams[key] = value
	return b
}

// Body sets the request payload that Do marshals as JSON.
func (b *Builder) Body(body interface{}) *Builder {
	b.body = body
	return b
}

// Header adds an HTTP header to the request. An empty key is ignored.
func (b *Builder) Header(key, value string) *Builder {
	if key == "" {
		return b
	}
	b.headers[key] = value
	return b
}

// Do builds the HTTP request, sends it through client, and unmarshals a
// successful response into responseModel.
func (b *Builder) Do(client HTTPClient, responseModel interface{}) (interface{}, error) {
	req, err := b.newRequest(client)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, parseAPIError(resp)
	}

	if responseModel == nil {
		return nil, nil
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(bytes.TrimSpace(payload)) == 0 {
		return responseModel, nil
	}

	if err := json.Unmarshal(payload, responseModel); err != nil {
		return nil, err
	}

	return responseModel, nil
}

func (b *Builder) newRequest(client HTTPClient) (*http.Request, error) {
	endpoint, err := b.resolveURL(client)
	if err != nil {
		return nil, err
	}

	var payload io.Reader
	if b.body != nil && b.method != http.MethodGet && b.method != http.MethodDelete {
		encoded, err := json.Marshal(b.body)
		if err != nil {
			return nil, err
		}
		payload = bytes.NewReader(encoded)
	}

	req, err := http.NewRequestWithContext(b.ctx, b.method, endpoint, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	for key, value := range b.headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

func (b *Builder) resolveURL(client HTTPClient) (string, error) {
	path := b.basePath
	if !strings.HasPrefix(path, "http://") && !strings.HasPrefix(path, "https://") {
		path = strings.TrimRight(clientBasePath(client), "/") + "/" + strings.TrimLeft(path, "/")
	}

	parsed, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	query := parsed.Query()
	for key, value := range b.queryParams {
		query.Set(key, value)
	}
	parsed.RawQuery = query.Encode()

	return parsed.String(), nil
}

func clientBasePath(client HTTPClient) string {
	provider, ok := client.(BasePathProvider)
	if !ok {
		return ""
	}
	return provider.BasePath()
}
