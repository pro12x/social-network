package utils

import (
	"backend/pkg/entity"
	"regexp"
	"strings"
)

func CheckPost(post *entity.Post) bool {
	if strings.TrimSpace(post.Title) == "" || strings.TrimSpace(post.Content) == "" {
		return false
	}

	if len(strings.TrimSpace(post.Title)) < 15 || len(strings.TrimSpace(post.Title)) > 255 || len(strings.TrimSpace(post.Content)) < 100 || len(strings.TrimSpace(post.Content)) > 1000 {
		return false
	}
	return true
}

func CheckEmail(email string) bool {
	if len(email) >= 10 && len(email) <= 35 {
		parts := strings.Split(email, "@")
		if len(parts) != 2 || len(strings.TrimSpace(parts[0])) == 0 || len(strings.TrimSpace(parts[1])) == 0 {
			return false
		}

		username := parts[0]
		regUsername := regexp.MustCompile(`^[a-z0-9]+$`)
		if len(username) < 3 || len(username) > 20 || !regUsername.MatchString(username) {
			return false
		}

		domainName := strings.Split(parts[1], ".")
		if len(domainName) != 2 || (len(strings.TrimSpace(domainName[0])) < 3 || len(strings.TrimSpace(domainName[0])) > 10) || (len(strings.TrimSpace(domainName[1])) < 2 || len(strings.TrimSpace(domainName[1])) > 3) {
			return false
		}

		return true
	}
	return false
}

func CheckPassword(password string) bool {
	if len(password) >= 8 && len(password) <= 20 {
		hasUppercase := regexp.MustCompile(`[A-Z]+?`).MatchString(password)
		hasLowercase := regexp.MustCompile(`[a-z]+?`).MatchString(password)
		hasDigit := regexp.MustCompile(`[0-9]+?`).MatchString(password)
		hasSpecialChar := regexp.MustCompile(`[\W_]+?`).MatchString(password)
		return hasUppercase && hasLowercase && hasDigit && hasSpecialChar
	}
	return false
}

func CheckNickname(nickname string) bool {
	return regexp.MustCompile(`^[a-z0-9]{3,20}$`).MatchString(nickname)
}

func CheckUser(user *entity.User) bool {
	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" || strings.TrimSpace(user.Firstname) == "" || strings.TrimSpace(user.Lastname) == "" || strings.TrimSpace(user.DateOfBirth) == "" {
		return false
	}

	if !CheckEmail(strings.TrimSpace(user.Email)) || !CheckPassword(strings.TrimSpace(user.Password)) || !regexp.MustCompile(`^[a-zA-Z\s]{3,20}$`).MatchString(user.Firstname) || !regexp.MustCompile(`^[a-zA-Z\s]{3,20}$`).MatchString(user.Lastname) {
		return false
	}

	if len(strings.TrimSpace(user.Nickname)) != 0 {
		if !CheckNickname(strings.TrimSpace(user.Nickname)) {
			return false
		}
	}

	return true
}

func CheckFollow(follow *entity.Follow) bool {
	if follow.FollowerID == 0 || follow.FolloweeID == 0 || follow.FollowerID == follow.FolloweeID {
		return false
	}
	return true
}

func CheckComment(comment *entity.Comment) bool {
	if strings.TrimSpace(comment.Content) == "" {
		return false
	}

	if len(strings.TrimSpace(comment.Content)) < 10 || len(strings.TrimSpace(comment.Content)) > 255 {
		return false
	}

	if comment.UserID == 0 && comment.PostID == 0 {
		return false
	}

	return true
}

func CheckCategory(category *entity.Category) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9\s\W_]{3,20}$`).MatchString(category.Name)
}
