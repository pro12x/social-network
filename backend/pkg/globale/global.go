package globale

import (
	"backend/pkg/db/sqlite"
)

var DB *sqlite.Database

var Endpoints = map[string][]interface{}{
	"users": {
		"http://localhost:1111/api/v1/social-network/register",
		"http://localhost:1111/api/v1/social-network/login",
		"http://localhost:1111/api/v1/social-network/users",
		"http://localhost:1111/api/v1/social-network/logout",
		"http://localhost:1111/api/v1/social-network/is_online",
		"http://localhost:1111/api/v1/social-network/profile/{id}",
		"http://localhost:1111/api/v1/social-network/update-profile/{id}",
	},
	"posts": {
		"http://localhost:1111/api/v1/social-network/post",
		"http://localhost:1111/api/v1/social-network/update-post/{id}",
		"http://localhost:1111/api/v1/social-network/delete-post/{id}",
		"http://localhost:1111/api/v1/social-network/get-post/{id}",
		"http://localhost:1111/api/v1/social-network/posts",
		"http://localhost:1111/api/v1/social-network/count-posts",
		"http://localhost:1111/api/v1/social-network/user-posts/{id}&{privacy}",
		"http://localhost:1111/api/v1/social-network/category-posts/{id}",
	},
	"comments": {},
	"follows": {
		"http://localhost:1111/api/v1/social-network/follow",
		"http://localhost:1111/api/v1/social-network/unfollow",
		"http://localhost:1111/api/v1/social-network/accept/{id}",
		"http://localhost:1111/api/v1/social-network/decline/{id}",
		"http://localhost:1111/api/v1/social-network/pending/{id}",
		"http://localhost:1111/api/v1/social-network/followers/{id}",
		"http://localhost:1111/api/v1/social-network/followings/{id}",
		"http://localhost:1111/api/v1/social-network/follower-count/{id}",
		"http://localhost:1111/api/v1/social-network/following-count/{id}",
		"http://localhost:1111/api/v1/social-network/follow-count",
	},
	"actions": {},
	"groups":  {},
	"categories": {
		"http://localhost:1111/api/v1/social-network/category",
		"http://localhost:1111/api/v1/social-network/category-update/{id}",
		"http://localhost:1111/api/v1/social-network/category-delete/{id}",
		"http://localhost:1111/api/v1/social-network/category-get/{id}",
		"http://localhost:1111/api/v1/social-network/category-name/{name}",
		"http://localhost:1111/api/v1/social-network/categories",
		"http://localhost:1111/api/v1/social-network/categories-count",
	},
}
