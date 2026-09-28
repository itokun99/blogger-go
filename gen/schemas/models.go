// Code generated from the Blogger API v3 discovery document (revision 20260924); DO NOT EDIT.

// Package schemas contains the model structs for every schema in the Blogger v3
// discovery document (revision 20260924): Blog, BlogList, BlogPerUserInfo,
// BlogUserInfo, Comment, CommentList, Page, PageList, Pageviews, Post, PostList,
// PostPerUserInfo, PostUserInfo, PostUserInfosList, and User, together with the
// nested object schemas those types reference.
//
// The structs mirror the canonical google.golang.org/api/blogger/v3 types field
// for field: the same Go field names, the same Go field types, and the same JSON
// tags - including omitempty and the string-encoded int64 fields - so the
// same values serialize to identical JSON in either package. Field names and
// tags come from the published API schema and are stable across versions.
//
// Only the schema fields are modeled; the SDK plumbing of the canonical types
// (ForceSendFields, NullFields, and the embedded googleapi.ServerResponse) is
// intentionally omitted. The discovery document declares no required properties
// on these schemas, so every field is optional and annotated as such.
package schemas

// Blog: the Blog schema from the Blogger v3 discovery document.
type Blog struct {
	// Optional. The JSON custom meta-data for the Blog.
	CustomMetaData string `json:"customMetaData,omitempty"`
	// Optional. The description of this blog. This is displayed underneath the title.
	Description string `json:"description,omitempty"`
	// Optional. The identifier for this resource.
	Id string `json:"id,omitempty"`
	// Optional. The kind of this entry. Always blogger#blog.
	Kind string `json:"kind,omitempty"`
	// Optional. The locale this Blog is set to.
	Locale *BlogLocale `json:"locale,omitempty"`
	// Optional. The name of this blog. This is displayed as the title.
	Name string `json:"name,omitempty"`
	// Optional. The container of pages in this blog.
	Pages *BlogPages `json:"pages,omitempty"`
	// Optional. The container of posts in this blog.
	Posts *BlogPosts `json:"posts,omitempty"`
	// Optional. RFC 3339 date-time when this blog was published.
	Published string `json:"published,omitempty"`
	// Optional. The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`
	// Optional. The status of the blog.
	//
	// Possible values:
	//   "LIVE"
	//   "DELETED"
	Status string `json:"status,omitempty"`
	// Optional. RFC 3339 date-time when this blog was last updated.
	Updated string `json:"updated,omitempty"`
	// Optional. The URL where this blog is published.
	Url string `json:"url,omitempty"`
}

// BlogLocale: The locale this Blog is set to.
type BlogLocale struct {
	// Optional. The country this blog's locale is set to.
	Country string `json:"country,omitempty"`
	// Optional. The language this blog is authored in.
	Language string `json:"language,omitempty"`
	// Optional. The language variant this blog is authored in.
	Variant string `json:"variant,omitempty"`
}

// BlogPages: The container of pages in this blog.
type BlogPages struct {
	// Optional. The URL of the container for pages in this blog.
	SelfLink string `json:"selfLink,omitempty"`
	// Optional. The count of pages in this blog.
	TotalItems int64 `json:"totalItems,omitempty"`
}

// BlogPosts: The container of posts in this blog.
type BlogPosts struct {
	// Optional. The List of Posts for this Blog.
	Items []*Post `json:"items,omitempty"`
	// Optional. The URL of the container for posts in this blog.
	SelfLink string `json:"selfLink,omitempty"`
	// Optional. The count of posts in this blog.
	TotalItems int64 `json:"totalItems,omitempty"`
}

// BlogList: the BlogList schema from the Blogger v3 discovery document.
type BlogList struct {
	// Optional. Admin level list of blog per-user information.
	BlogUserInfos []*BlogUserInfo `json:"blogUserInfos,omitempty"`
	// Optional. The list of Blogs this user has Authorship or Admin rights over.
	Items []*Blog `json:"items,omitempty"`
	// Optional. The kind of this entity. Always blogger#blogList.
	Kind string `json:"kind,omitempty"`
}

// BlogPerUserInfo: the BlogPerUserInfo schema from the Blogger v3 discovery document.
type BlogPerUserInfo struct {
	// Optional. ID of the Blog resource.
	BlogId string `json:"blogId,omitempty"`
	// Optional. True if the user has Admin level access to the blog.
	HasAdminAccess bool `json:"hasAdminAccess,omitempty"`
	// Optional. The kind of this entity. Always blogger#blogPerUserInfo.
	Kind string `json:"kind,omitempty"`
	// Optional. The Photo Album Key for the user when adding photos to the blog.
	PhotosAlbumKey string `json:"photosAlbumKey,omitempty"`
	// Optional. Access permissions that the user has for the blog (ADMIN, AUTHOR, or READER).
	//
	// Possible values:
	//   "VIEW_TYPE_UNSPECIFIED"
	//   "READER"
	//   "AUTHOR"
	//   "ADMIN"
	Role string `json:"role,omitempty"`
	// Optional. ID of the User.
	UserId string `json:"userId,omitempty"`
}

// BlogUserInfo: the BlogUserInfo schema from the Blogger v3 discovery document.
type BlogUserInfo struct {
	// Optional. The Blog resource.
	Blog *Blog `json:"blog,omitempty"`
	// Optional. Information about a User for the Blog.
	BlogUserInfo *BlogPerUserInfo `json:"blog_user_info,omitempty"`
	// Optional. The kind of this entity. Always blogger#blogUserInfo.
	Kind string `json:"kind,omitempty"`
}

// Comment: the Comment schema from the Blogger v3 discovery document.
type Comment struct {
	// Optional. The author of this Comment.
	Author *CommentAuthor `json:"author,omitempty"`
	// Optional. Data about the blog containing this comment.
	Blog *CommentBlog `json:"blog,omitempty"`
	// Optional. The actual content of the comment. May include HTML markup.
	Content string `json:"content,omitempty"`
	// Optional. The identifier for this resource.
	Id string `json:"id,omitempty"`
	// Optional. Data about the comment this is in reply to.
	InReplyTo *CommentInReplyTo `json:"inReplyTo,omitempty"`
	// Optional. The kind of this entry. Always blogger#comment.
	Kind string `json:"kind,omitempty"`
	// Optional. Data about the post containing this comment.
	Post *CommentPost `json:"post,omitempty"`
	// Optional. RFC 3339 date-time when this comment was published.
	Published string `json:"published,omitempty"`
	// Optional. The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`
	// Optional. The status of the comment (only populated for admin users).
	//
	// Possible values:
	//   "LIVE"
	//   "EMPTIED"
	//   "PENDING"
	//   "SPAM"
	Status string `json:"status,omitempty"`
	// Optional. RFC 3339 date-time when this comment was last updated.
	Updated string `json:"updated,omitempty"`
}

// CommentAuthor: The author of this Comment.
type CommentAuthor struct {
	// Optional. The display name.
	DisplayName string `json:"displayName,omitempty"`
	// Optional. The identifier of the creator.
	Id string `json:"id,omitempty"`
	// Optional. The creator's avatar.
	Image *CommentAuthorImage `json:"image,omitempty"`
	// Optional. The URL of the creator's Profile page.
	Url string `json:"url,omitempty"`
}

// CommentAuthorImage: The creator's avatar.
type CommentAuthorImage struct {
	// Optional. The creator's avatar URL.
	Url string `json:"url,omitempty"`
}

// CommentBlog: Data about the blog containing this comment.
type CommentBlog struct {
	// Optional. The identifier of the blog containing this comment.
	Id string `json:"id,omitempty"`
}

// CommentInReplyTo: Data about the comment this is in reply to.
type CommentInReplyTo struct {
	// Optional. The identified of the parent of this comment.
	Id string `json:"id,omitempty"`
}

// CommentPost: Data about the post containing this comment.
type CommentPost struct {
	// Optional. The identifier of the post containing this comment.
	Id string `json:"id,omitempty"`
}

// CommentList: the CommentList schema from the Blogger v3 discovery document.
type CommentList struct {
	// Optional. Etag of the response.
	Etag string `json:"etag,omitempty"`
	// Optional. The List of Comments for a Post.
	Items []*Comment `json:"items,omitempty"`
	// Optional. The kind of this entry. Always blogger#commentList.
	Kind string `json:"kind,omitempty"`
	// Optional. Pagination token to fetch the next page, if one exists.
	NextPageToken string `json:"nextPageToken,omitempty"`
	// Optional. Pagination token to fetch the previous page, if one exists.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

// Page: the Page schema from the Blogger v3 discovery document.
type Page struct {
	// Optional. The author of this Page.
	Author *PageAuthor `json:"author,omitempty"`
	// Optional. Data about the blog containing this Page.
	Blog *PageBlog `json:"blog,omitempty"`
	// Optional. The body content of this Page, in HTML.
	Content string `json:"content,omitempty"`
	// Optional. Etag of the resource.
	Etag string `json:"etag,omitempty"`
	// Optional. The identifier for this resource.
	Id string `json:"id,omitempty"`
	// Optional. The kind of this entity. Always blogger#page.
	Kind string `json:"kind,omitempty"`
	// Optional. RFC 3339 date-time when this Page was published.
	Published string `json:"published,omitempty"`
	// Optional. The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`
	// Optional. The status of the page for admin resources (either LIVE or DRAFT).
	//
	// Possible values:
	//   "LIVE"
	//   "DRAFT"
	//   "SOFT_TRASHED"
	Status string `json:"status,omitempty"`
	// Optional. The title of this entity. This is the name displayed in the Admin user interface.
	Title string `json:"title,omitempty"`
	// Optional. RFC 3339 date-time when this Page was trashed.
	Trashed string `json:"trashed,omitempty"`
	// Optional. RFC 3339 date-time when this Page was last updated.
	Updated string `json:"updated,omitempty"`
	// Optional. The URL that this Page is displayed at.
	Url string `json:"url,omitempty"`
}

// PageAuthor: The author of this Page.
type PageAuthor struct {
	// Optional. The display name.
	DisplayName string `json:"displayName,omitempty"`
	// Optional. The identifier of the creator.
	Id string `json:"id,omitempty"`
	// Optional. The creator's avatar.
	Image *PageAuthorImage `json:"image,omitempty"`
	// Optional. The URL of the creator's Profile page.
	Url string `json:"url,omitempty"`
}

// PageAuthorImage: The creator's avatar.
type PageAuthorImage struct {
	// Optional. The creator's avatar URL.
	Url string `json:"url,omitempty"`
}

// PageBlog: Data about the blog containing this Page.
type PageBlog struct {
	// Optional. The identifier of the blog containing this page.
	Id string `json:"id,omitempty"`
}

// PageList: the PageList schema from the Blogger v3 discovery document.
type PageList struct {
	// Optional. Etag of the response.
	Etag string `json:"etag,omitempty"`
	// Optional. The list of Pages for a Blog.
	Items []*Page `json:"items,omitempty"`
	// Optional. The kind of this entity. Always blogger#pageList.
	Kind string `json:"kind,omitempty"`
	// Optional. Pagination token to fetch the next page, if one exists.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// Pageviews: the Pageviews schema from the Blogger v3 discovery document.
type Pageviews struct {
	// Optional. Blog Id.
	BlogId string `json:"blogId,omitempty"`
	// Optional. The container of posts in this blog.
	Counts []*PageviewsCounts `json:"counts,omitempty"`
	// Optional. The kind of this entry. Always blogger#page_views.
	Kind string `json:"kind,omitempty"`
}

// PageviewsCounts: the PageviewsCounts schema from the Blogger v3 discovery document.
type PageviewsCounts struct {
	// Optional. Count of page views for the given time range.
	Count int64 `json:"count,omitempty,string"`
	// Optional. Time range the given count applies to.
	//
	// Possible values:
	//   "ALL_TIME"
	//   "THIRTY_DAYS"
	//   "SEVEN_DAYS"
	TimeRange string `json:"timeRange,omitempty"`
}

// Post: the Post schema from the Blogger v3 discovery document.
type Post struct {
	// Optional. The author of this Post.
	Author *PostAuthor `json:"author,omitempty"`
	// Optional. Data about the blog containing this Post.
	Blog *PostBlog `json:"blog,omitempty"`
	// Optional. The content of the Post. May contain HTML markup.
	Content string `json:"content,omitempty"`
	// Optional. The JSON meta-data for the Post.
	CustomMetaData string `json:"customMetaData,omitempty"`
	// Optional. Etag of the resource.
	Etag string `json:"etag,omitempty"`
	// Optional. The identifier of this Post.
	Id string `json:"id,omitempty"`
	// Optional. Display image for the Post.
	Images []*PostImages `json:"images,omitempty"`
	// Optional. The kind of this entity. Always blogger#post.
	Kind string `json:"kind,omitempty"`
	// Optional. The list of labels this Post was tagged with.
	Labels []string `json:"labels,omitempty"`
	// Optional. The location for geotagged posts.
	Location *PostLocation `json:"location,omitempty"`
	// Optional. RFC 3339 date-time when this Post was published.
	Published string `json:"published,omitempty"`
	// Optional. Comment control and display setting for readers of this post.
	//
	// Possible values:
	//   "ALLOW"
	//   "DONT_ALLOW_SHOW_EXISTING"
	//   "DONT_ALLOW_HIDE_EXISTING"
	ReaderComments string `json:"readerComments,omitempty"`
	// Optional. The container of comments on this Post.
	Replies *PostReplies `json:"replies,omitempty"`
	// Optional. The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`
	// Optional. Status of the post. Only set for admin-level requests.
	//
	// Possible values:
	//   "LIVE"
	//   "DRAFT"
	//   "SCHEDULED"
	//   "SOFT_TRASHED"
	Status string `json:"status,omitempty"`
	// Optional. The title of the Post.
	Title string `json:"title,omitempty"`
	// Optional. The title link URL, similar to atom's related link.
	TitleLink string `json:"titleLink,omitempty"`
	// Optional. RFC 3339 date-time when this Post was last trashed.
	Trashed string `json:"trashed,omitempty"`
	// Optional. RFC 3339 date-time when this Post was last updated.
	Updated string `json:"updated,omitempty"`
	// Optional. The URL where this Post is displayed.
	Url string `json:"url,omitempty"`
}

// PostAuthor: The author of this Post.
type PostAuthor struct {
	// Optional. The display name.
	DisplayName string `json:"displayName,omitempty"`
	// Optional. The identifier of the creator.
	Id string `json:"id,omitempty"`
	// Optional. The creator's avatar.
	Image *PostAuthorImage `json:"image,omitempty"`
	// Optional. The URL of the creator's Profile page.
	Url string `json:"url,omitempty"`
}

// PostAuthorImage: The creator's avatar.
type PostAuthorImage struct {
	// Optional. The creator's avatar URL.
	Url string `json:"url,omitempty"`
}

// PostBlog: Data about the blog containing this Post.
type PostBlog struct {
	// Optional. The identifier of the Blog that contains this Post.
	Id string `json:"id,omitempty"`
}

// PostImages: the PostImages schema from the Blogger v3 discovery document.
type PostImages struct {
	// Optional.
	Url string `json:"url,omitempty"`
}

// PostLocation: The location for geotagged posts.
type PostLocation struct {
	// Optional. Location's latitude.
	Lat float64 `json:"lat,omitempty"`
	// Optional. Location's longitude.
	Lng float64 `json:"lng,omitempty"`
	// Optional. Location name.
	Name string `json:"name,omitempty"`
	// Optional. Location's viewport span. Can be used when rendering a map preview.
	Span string `json:"span,omitempty"`
}

// PostReplies: The container of comments on this Post.
type PostReplies struct {
	// Optional. The List of Comments for this Post.
	Items []*Comment `json:"items,omitempty"`
	// Optional. The URL of the comments on this post.
	SelfLink string `json:"selfLink,omitempty"`
	// Optional. The count of comments on this post.
	TotalItems int64 `json:"totalItems,omitempty,string"`
}

// PostList: the PostList schema from the Blogger v3 discovery document.
type PostList struct {
	// Optional. Etag of the response.
	Etag string `json:"etag,omitempty"`
	// Optional. The list of Posts for this Blog.
	Items []*Post `json:"items,omitempty"`
	// Optional. The kind of this entity. Always blogger#postList.
	Kind string `json:"kind,omitempty"`
	// Optional. Pagination token to fetch the next page, if one exists.
	NextPageToken string `json:"nextPageToken,omitempty"`
	// Optional. Pagination token to fetch the previous page, if one exists.
	PrevPageToken string `json:"prevPageToken,omitempty"`
}

// PostPerUserInfo: the PostPerUserInfo schema from the Blogger v3 discovery document.
type PostPerUserInfo struct {
	// Optional. ID of the Blog that the post resource belongs to.
	BlogId string `json:"blogId,omitempty"`
	// Optional. True if the user has Author level access to the post.
	HasEditAccess bool `json:"hasEditAccess,omitempty"`
	// Optional. The kind of this entity. Always blogger#postPerUserInfo.
	Kind string `json:"kind,omitempty"`
	// Optional. ID of the Post resource.
	PostId string `json:"postId,omitempty"`
	// Optional. ID of the User.
	UserId string `json:"userId,omitempty"`
}

// PostUserInfo: the PostUserInfo schema from the Blogger v3 discovery document.
type PostUserInfo struct {
	// Optional. The kind of this entity. Always blogger#postUserInfo.
	Kind string `json:"kind,omitempty"`
	// Optional. The Post resource.
	Post *Post `json:"post,omitempty"`
	// Optional. Information about a User for the Post.
	PostUserInfo *PostPerUserInfo `json:"post_user_info,omitempty"`
}

// PostUserInfosList: the PostUserInfosList schema from the Blogger v3 discovery document.
type PostUserInfosList struct {
	// Optional. The list of Posts with User information for the post, for this Blog.
	Items []*PostUserInfo `json:"items,omitempty"`
	// Optional. The kind of this entity. Always blogger#postList.
	Kind string `json:"kind,omitempty"`
	// Optional. Pagination token to fetch the next page, if one exists.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// User: the User schema from the Blogger v3 discovery document.
type User struct {
	// Optional. Profile summary information.
	About string `json:"about,omitempty"`
	// Optional. The container of blogs for this user.
	Blogs *UserBlogs `json:"blogs,omitempty"`
	// Optional. The timestamp of when this profile was created, in seconds since epoch.
	Created string `json:"created,omitempty"`
	// Optional. The display name.
	DisplayName string `json:"displayName,omitempty"`
	// Optional. The identifier for this User.
	Id string `json:"id,omitempty"`
	// Optional. The kind of this entity. Always blogger#user.
	Kind string `json:"kind,omitempty"`
	// Optional. This user's locale
	Locale *UserLocale `json:"locale,omitempty"`
	// Optional. The API REST URL to fetch this resource from.
	SelfLink string `json:"selfLink,omitempty"`
	// Optional. The user's profile page.
	Url string `json:"url,omitempty"`
}

// UserBlogs: The container of blogs for this user.
type UserBlogs struct {
	// Optional. The URL of the Blogs for this user.
	SelfLink string `json:"selfLink,omitempty"`
}

// UserLocale: This user's locale
type UserLocale struct {
	// Optional. The country this blog's locale is set to.
	Country string `json:"country,omitempty"`
	// Optional. The language this blog is authored in.
	Language string `json:"language,omitempty"`
	// Optional. The language variant this blog is authored in.
	Variant string `json:"variant,omitempty"`
}
