package builder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// APIError is the structured error returned by the Blogger API for any
// non-2xx response.
type APIError struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Reasons []string `json:"reasons"`
	Domain  string   `json:"domain"`
}

// Error renders the API error with its status code, reason and domain attached
// so the message stays informative when logged on its own.
func (e *APIError) Error() string {
	if e == nil {
		return "blogger: <nil> APIError"
	}

	parts := []string{}
	if e.Reasons != nil {
		parts = append(parts, fmt.Sprintf("reasons=%v", e.Reasons))
	}
	if e.Domain != "" {
		parts = append(parts, fmt.Sprintf("domain=%q", e.Domain))
	}

	summary := fmt.Sprintf("blogger: API error %d: %s", e.Code, e.Message)
	if len(parts) == 0 {
		return summary
	}
	return summary + " (" + strings.Join(parts, ", ") + ")"
}

// NewAPIError decodes an error response body into an APIError. It never returns
// a nil error: an unreadable or empty body yields an APIError carrying the HTTP
// status code.
func NewAPIError(resp *http.Response) (*APIError, error) {
	if resp == nil {
		return &APIError{Message: "nil response"}, fmt.Errorf("blogger: nil response")
	}

	apiErr := &APIError{Code: resp.StatusCode}
	payload, readErr := io.ReadAll(resp.Body)

	if err := decodeAPIError(payload, apiErr); err != nil {
		apiErr.Message = http.StatusText(resp.StatusCode)
		apiErr.Reasons = []string{}
		return apiErr, readErr
	}

	if apiErr.Code == 0 {
		apiErr.Code = resp.StatusCode
	}
	if apiErr.Reasons == nil {
		apiErr.Reasons = []string{}
	}
	if apiErr.Message == "" {
		apiErr.Message = http.StatusText(resp.StatusCode)
	}

	return apiErr, readErr
}

// apiErrorEnvelope mirrors the `{"error": {...}}` wrapper Google APIs use.
type apiErrorEnvelope struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Reason  string `json:"reason"`
			Domain  string `json:"domain"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"error"`
}

func decodeAPIError(payload []byte, apiErr *APIError) error {
	if len(payload) == 0 {
		return fmt.Errorf("blogger: empty error body")
	}

	var envelope apiErrorEnvelope
	if err := json.Unmarshal(payload, &envelope); err != nil {
		return err
	}

	apiErr.Code = envelope.Error.Code
	apiErr.Message = envelope.Error.Message
	apiErr.Reasons = []string{}

	for _, item := range envelope.Error.Errors {
		if item.Reason != "" {
			apiErr.Reasons = append(apiErr.Reasons, item.Reason)
		}
		if apiErr.Domain == "" {
			apiErr.Domain = item.Domain
		}
		if apiErr.Message == "" {
			apiErr.Message = item.Message
		}
	}

	return nil
}

// parseAPIError converts a non-2xx response into an error. The body is drained
// and closed here so callers never leak the response.
func parseAPIError(resp *http.Response) error {
	defer resp.Body.Close()

	apiErr, err := NewAPIError(resp)
	if err != nil && apiErr == nil {
		return fmt.Errorf("blogger: request failed with status %d: %w", resp.StatusCode, err)
	}
	return apiErr
}

// IsAPIError reports whether err is, or wraps, an *APIError.
func IsAPIError(err error) bool {
	var apiErr *APIError
	return asAPIError(err, &apiErr)
}

func asAPIError(err error, target **APIError) bool {
	for err != nil {
		if apiErr, ok := err.(*APIError); ok {
			*target = apiErr
			return true
		}
		unwrapper, ok := err.(interface{ Unwrap() error })
		if !ok {
			return false
		}
		err = unwrapper.Unwrap()
	}
	return false
}
