package builder

import "net/http"

// TransportProvider exposes the HTTP transport and base path used by services.
// Implementing this interface allows the root Client to be passed to services
// without creating an import cycle.
type TransportProvider interface {
	HTTPClient() *http.Client
	BasePath() string
}
