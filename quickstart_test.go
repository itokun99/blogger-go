package blogger_test

import (
	"context"
	"fmt"

	blogger "github.com/itokun99/blogger-go"
)

// Example_postsListOptions mirrors the Quick Start snippet in README.md: the
// root package must expose the option constructors and status constants used
// there, so the documented call has to compile as written.
func Example_postsListOptions() {
	client := blogger.NewRawClient(nil)

	posts, _ := client.Posts().List(context.Background(), "1234567890",
		blogger.WithMaxResults(10),
		blogger.WithStatus(blogger.PostStatusLive),
		blogger.WithOrderBy(blogger.OrderByPublished),
	)
	fmt.Println(posts == nil)
	// Output: true
}
