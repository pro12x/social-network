package utils

import "backend/pkg/entity"

func CheckPost(post *entity.Post) bool {
	if post.Title == "" || post.Content == "" || post.Image == "" {
		return false
	}
	return true
}

func CheckUser(user *entity.User) bool {
	if user.Email == "" || user.Password == "" || user.Firstname == "" || user.Lastname == "" || user.DateOfBirth == "" {
		return false
	}
	return true
}

func CheckFollow(follow *entity.Follow) bool {
	if follow.FollowerID == 0 || follow.FolloweeID == 0 {
		return false
	}
	return true
}

/*func CheckComment(comment entity.Comment) bool {
	if comment.Content == "" {
		return false
	}
	return true
}*/

func CheckCategory(category *entity.Category) bool {
	if category.Name == "" {
		return false
	}
	return true
}
