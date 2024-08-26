package globale

import (
	"backend/pkg/db/sqlite"
)

var DB *sqlite.Database

type Format struct {
	Method string
	URL    string
}

var Endpoint = map[string][]Format{
	"users": {
		{"POST", "/api/v1/social-network/register"},
		{"POST", "/api/v1/social-network/login"},
		{"GET", "/api/v1/social-network/users"},
		{"POST", "/api/v1/social-network/logout"},
		{"POST", "/api/v1/social-network/is_online"},
		{"GET", "/api/v1/social-network/profile/{id}"},
		{"PUT", "/api/v1/social-network/update-profile/{id}"},
	},
	"posts": {
		{"POST", "/api/v1/social-network/post"},
		{"PUT", "/api/v1/social-network/update-post/{id}"},
		{"DELETE", "/api/v1/social-network/delete-post/{id}"},
		{"GET", "/api/v1/social-network/get-post/{id}"},
		{"GET", "/api/v1/social-network/posts"},
		{"GET", "/api/v1/social-network/count-posts"},
		{"GET", "/api/v1/social-network/user-posts/{id}&{privacy}"},
		{"GET", "/api/v1/social-network/category-posts/{id}"},
	},
	"comments": {
		{"POST", "/api/v1/social-network/comment"},
		{"PUT", "/api/v1/social-network/comment-update/{id}"},
		{"DELETE", "/api/v1/social-network/comment-delete/{id}"},
		{"GET", "/api/v1/social-network/comment-get/{id}"},
		{"GET", "/api/v1/social-network/comments-post/{postID}"},
		{"GET", "/api/v1/social-network/comments-count"},
	},
	"follows": {
		{"POST", "/api/v1/social-network/follow"},
		{"DELETE", "/api/v1/social-network/unfollow"},
		{"PUT", "/api/v1/social-network/accept/{id}"},
		{"DELETE", "/api/v1/social-network/decline/{id}"},
		{"GET", "/api/v1/social-network/pending/{id}"},
		{"GET", "/api/v1/social-network/followers/{id}"},
		{"GET", "/api/v1/social-network/followings/{id}"},
		{"GET", "/api/v1/social-network/follower-count/{id}"},
		{"GET", "/api/v1/social-network/following-count/{id}"},
		{"GET", "/api/v1/social-network/follow-count"},
	},
	"actions": {},
	"groups":  {},
	"categories": {
		{"POST", "/api/v1/social-network/category"},
		{"PUT", "/api/v1/social-network/category-update/{id}"},
		{"DELETE", "/api/v1/social-network/category-delete/{id}"},
		{"GET", "/api/v1/social-network/category-get/{id}"},
		{"GET", "/api/v1/social-network/category-name/{name}"},
		{"GET", "/api/v1/social-network/categories"},
		{"GET", "/api/v1/social-network/categories-count"},
	},
}

var Endpoints = map[string][]interface{}{
	"users": {
		"http://localhost:{PORT}/api/v1/social-network/register",
		"http://localhost:{PORT}/api/v1/social-network/login",
		"http://localhost:{PORT}/api/v1/social-network/users",
		"http://localhost:{PORT}/api/v1/social-network/logout",
		"http://localhost:{PORT}/api/v1/social-network/is_online",
		"http://localhost:{PORT}/api/v1/social-network/profile/{id}",
		"http://localhost:{PORT}/api/v1/social-network/update-profile/{id}",
	},
	"posts": {
		"http://localhost:{PORT}/api/v1/social-network/post",
		"http://localhost:{PORT}/api/v1/social-network/update-post/{id}",
		"http://localhost:{PORT}/api/v1/social-network/delete-post/{id}",
		"http://localhost:{PORT}/api/v1/social-network/get-post/{id}",
		"http://localhost:{PORT}/api/v1/social-network/posts",
		"http://localhost:{PORT}/api/v1/social-network/count-posts",
		"http://localhost:{PORT}/api/v1/social-network/user-posts/{id}&{privacy}",
		"http://localhost:{PORT}/api/v1/social-network/category-posts/{id}",
	},
	"comments": {
		"http://localhost:{PORT}/api/v1/social-network/comment",
		"http://localhost:{PORT}/api/v1/social-network/comment-update/{id}",
		"http://localhost:{PORT}/api/v1/social-network/comment-delete/{id}",
		"http://localhost:{PORT}/api/v1/social-network/comment-get/{id}",
		"http://localhost:{PORT}/api/v1/social-network/comments-post/{postID}",
		"http://localhost:{PORT}/api/v1/social-network/comments-count",
	},
	"follows": {
		"http://localhost:{PORT}/api/v1/social-network/follow",
		"http://localhost:{PORT}/api/v1/social-network/unfollow",
		"http://localhost:{PORT}/api/v1/social-network/accept/{id}",
		"http://localhost:{PORT}/api/v1/social-network/decline/{id}",
		"http://localhost:{PORT}/api/v1/social-network/pending/{id}",
		"http://localhost:{PORT}/api/v1/social-network/followers/{id}",
		"http://localhost:{PORT}/api/v1/social-network/followings/{id}",
		"http://localhost:{PORT}/api/v1/social-network/follower-count/{id}",
		"http://localhost:{PORT}/api/v1/social-network/following-count/{id}",
		"http://localhost:{PORT}/api/v1/social-network/follow-count",
	},
	"actions": {},
	"groups":  {},
	"categories": {
		"http://localhost:{PORT}/api/v1/social-network/category",
		"http://localhost:{PORT}/api/v1/social-network/category-update/{id}",
		"http://localhost:{PORT}/api/v1/social-network/category-delete/{id}",
		"http://localhost:{PORT}/api/v1/social-network/category-get/{id}",
		"http://localhost:{PORT}/api/v1/social-network/category-name/{name}",
		"http://localhost:{PORT}/api/v1/social-network/categories",
		"http://localhost:{PORT}/api/v1/social-network/categories-count",
	},
}
