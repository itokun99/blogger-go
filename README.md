# blogger-go

A Go SDK for the [Blogger API v3](https://developers.google.com/blogger/docs/3.0/getting_started),
generated from Google's discovery document and layered over the official
[`google.golang.org/api/blogger/v3`](https://pkg.go.dev/google.golang.org/api/blogger/v3) client.

## What's Included

- **8 resources** covering all Blogger API endpoints
- **33 methods** matching the canonical discovery doc (revision 20260924)
- **15 schema models** derived from the discovery document
- **Builder pattern** for chainable request construction
- **Typed errors** with structured API error handling
- **Pagination helpers** for cursor-based listings
- **OAuth2 integration** via `golang.org/x/oauth2`

## Install

```sh
go get github.com/itokun99/blogger-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    blogger "github.com/itokun99/blogger-go"
)

func main() {
    ctx := context.Background()

    // Auth with OAuth2 credentials
    ts, err := blogger.TokenSourceFromJSON("credentials.json", "token.json")
    if err != nil {
        log.Fatal(err)
    }

    httpClient := blogger.NewHTTPClient(ctx, ts)
    client := blogger.NewRawClient(httpClient)

    // Get blog info
    blog, err := client.Blogs().Get(ctx, "1234567890")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(blog.Name)

    // List posts with options
    posts, err := client.Posts().List(ctx, "1234567890",
        blogger.WithMaxResults(10),
        blogger.WithStatus(blogger.PostStatusLive),
        blogger.WithOrderBy(blogger.OrderByPublished),
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d posts\n", len(posts.Items))
}
```

## API Surface

### Client

| Method | Description |
|--------|-------------|
| `NewClient(ctx, opts...)` | Creates client with Google auth options |
| `NewRawClient(httpClient)` | Creates client with provided HTTP client |
| `WithContext(ctx)` | Returns a clone with the given context |
| `HTTPClient()` | Returns the transport client |
| `BasePath()` | Returns the API base URL |

### Service Accessors

```go
client.Blogs()           *BlogsService
client.Comments()        *CommentsService
client.Pages()           *PagesService
client.Posts()           *PostsService
client.Users()           *UsersService
client.BlogUserInfos()   *BlogUserInfosService
client.PageViews()       *PageViewsService
client.PostUserInfos()   *PostUserInfosService
```

### Services and Methods

| Resource | Methods | Options |
|----------|---------|---------|
| **Blogs** | `Get`, `GetByUrl`, `ListByUser` | `WithBlogView`, `WithBlogMaxPosts` |
| **Comments** | `Get`, `List`, `ListByBlog`, `Approve`, `Delete`, `MarkAsSpam`, `RemoveContent` | `WithCommentPageToken`, `WithCommentMaxResults`, `WithCommentStatus`, `WithCommentView`, `WithCommentStartDate`, `WithCommentEndDate`, `WithCommentFetchBodies`, `WithCommentPost` |
| **Pages** | `Get`, `List`, `Insert`, `Update`, `Patch`, `Delete`, `Publish`, `Revert` | `WithPageListToken`, `WithPageListMaxResults`, `WithPageListStatus`, `WithPageView`, `WithPageFetchBodies` |
| **Posts** | `Get`, `List`, `Search`, `GetByPath`, `Insert`, `Update`, `Patch`, `Delete`, `Publish`, `Revert` | `WithPageToken`, `WithMaxResults`, `WithStartDate`, `WithEndDate`, `WithStatus`, `WithLabels`, `WithFetchImages`, `WithFetchBody`, `WithOrderBy`, `WithSortOrder`, `WithView`, `WithMaxComments`, `WithPostDate`, `WithQuery` |
| **Users** | `Get` | — |
| **BlogUserInfos** | `Get` | — |
| **PageViews** | `Get` | `WithPageViewRange` |
| **PostUserInfos** | `Get`, `List` | — |

### Status Constants

The same values are exported from the root package and from `gen/services`, so
`blogger.PostStatusLive` and `services.PostStatusLive` are interchangeable.

```go
// Post status values
blogger.PostStatusLive         = "LIVE"
blogger.PostStatusDraft        = "DRAFT"
blogger.PostStatusScheduled    = "SCHEDULED"
blogger.PostStatusSoftTrashed  = "SOFT_TRASHED"

// Comment status values
blogger.CommentStatusLive      = "LIVE"
blogger.CommentStatusEmptied   = "EMPTIED"
blogger.CommentStatusPending   = "PENDING"
blogger.CommentStatusSpam      = "SPAM"

// Page status values
blogger.PageStatusLive         = "LIVE"
blogger.PageStatusDraft        = "DRAFT"
blogger.PageStatusSoftTrashed  = "SOFT_TRASHED"

// View types
blogger.ViewTypeUnspecified    = "VIEW_TYPE_UNSPECIFIED"
blogger.ViewReader             = "READER"
blogger.ViewAuthor             = "AUTHOR"
blogger.ViewAdmin              = "ADMIN"

// Order and sort values
blogger.OrderByPublished       = "PUBLISHED"
blogger.OrderByUpdated         = "UPDATED"
blogger.SortOrderDescending    = "DESCENDING"
blogger.SortOrderAscending     = "ASCENDING"
```

### Schema Models

| Model | Description |
|-------|-------------|
| `Blog` | Blog resource |
| `BlogList` | List of blogs with optional blogUserInfos |
| `BlogPerUserInfo` | User access info for a blog |
| `BlogUserInfo` | Blog and user info pair |
| `Comment` | Comment resource |
| `CommentList` | List of comments |
| `Page` | Page resource |
| `PageList` | List of pages |
| `Pageviews` | Page view counts |
| `Post` | Post resource |
| `PostList` | List of posts |
| `PostPerUserInfo` | User access info for a post |
| `PostUserInfo` | Post and user info pair |
| `PostUserInfosList` | List of post/user info pairs |
| `User` | User profile resource |

## Requirements

- Go 1.21 or newer
- OAuth2 credentials (optional for read-only access with API key)

## Testing

```sh
go test ./...
go vet ./...
```

## License

MIT
