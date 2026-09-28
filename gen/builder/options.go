package builder

// Xgafv selects the API version of the Google API surface being called.
const (
	Xgafv = "xgafv"
)

// Authentication and standard Google API System Parameters. Any of these can be
// passed to Builder.Param.
const (
	AccessToken = "access_token"
	Timestamp   = "timestamp"
	OauthToken  = "oauth_token"
)

// Output-shaping standard parameters.
const (
	PrettyPrint = "prettyPrint"
	Callback    = "callback"
	Fields      = "fields"
	Key         = "key"
)

// CommonBloggerParams are the query parameters shared by most Blogger v3 list
// endpoints.
const (
	MaxResults  = "maxResults"
	PageToken   = "pageToken"
	FetchBodies = "fetchBodies"
	StartIndex  = "startIndex"
	Status      = "status"
	OrderBy     = "orderBy"
	SortOrder   = "sortOrder"
	Labels      = "labels"
	View        = "view"
	Alt         = "alt"
)

// MaxResultsRange bounds the per-page item count accepted by Blogger list
// endpoints.
type MaxResultsRange struct {
	Min int64
	Max int64
}

// BloggerListRange is the documented 1..100 window for maxResults.
var BloggerListRange = MaxResultsRange{Min: 1, Max: 100}

// Clamp constrains n to the range, returning the nearest bound.
func (r MaxResultsRange) Clamp(n int64) int64 {
	if n < r.Min {
		return r.Min
	}
	if n > r.Max {
		return r.Max
	}
	return n
}

// Contains reports whether n falls inside the range.
func (r MaxResultsRange) Contains(n int64) bool {
	return n >= r.Min && n <= r.Max
}

// CommonFilter describes the optional narrowing parameters that Blogger list
// endpoints accept.
type CommonFilter struct {
	Labels      []string
	Status      string
	OrderBy     string
	SortOrder   string
	FetchBodies *bool
}

// Apply writes the non-empty filter fields onto the builder as query
// parameters. Empty values are skipped so unset filters never reach the wire.
func (f CommonFilter) Apply(b *Builder) *Builder {
	if b == nil {
		return b
	}

	if len(f.Labels) > 0 {
		b.Param(Labels, joinList(f.Labels))
	}
	if f.Status != "" {
		b.Param(Status, f.Status)
	}
	if f.OrderBy != "" {
		b.Param(OrderBy, f.OrderBy)
	}
	if f.SortOrder != "" {
		b.Param(SortOrder, f.SortOrder)
	}
	if f.FetchBodies != nil {
		b.Param(FetchBodies, boolParam(*f.FetchBodies))
	}

	return b
}

// Pagination describes the page window requested from a list endpoint.
type Pagination struct {
	MaxResults int64
	PageToken  string
}

// Apply writes the pagination window onto the builder.
func (p Pagination) Apply(b *Builder) *Builder {
	if b == nil {
		return b
	}

	if p.MaxResults > 0 {
		b.Param(MaxResults, itoa(BloggerListRange.Clamp(p.MaxResults)))
	}
	if p.PageToken != "" {
		b.Param(PageToken, p.PageToken)
	}

	return b
}
