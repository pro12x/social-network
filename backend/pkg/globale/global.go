package globale

import (
	"backend/pkg/db/sqlite"
)

var DB *sqlite.Database

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
