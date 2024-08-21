package repository

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/entity"
	"database/sql"
	"errors"
	"strings"
)

type ActionRepoImpl struct {
	db sqlite.Database
}

func NewActionRepoImpl(db sqlite.Database) *ActionRepoImpl {
	return &ActionRepoImpl{db: db}
}

func (a *ActionRepoImpl) IsUserAlreadyLiked(userID, actionID uint, nature string) (bool, bool, bool, error) {
	action := new(entity.Action)
	var query string

	switch strings.ToLower(nature) {
	case "post":
		query = `SELECT * FROM actions WHERE user_id = ? AND post_id = ? AND comment_id < 0`
	case "comment":
		query = `SELECT * FROM actions WHERE user_id = ? AND comment_id = ? AND post_id < 0`
	default:
		return false, false, false, errors.New("invalid nature")
	}

	err := a.db.GetDB().QueryRow(query, userID, actionID).Scan(&action.ID, &action.Like, &action.Dislike, &action.PostID, &action.CommentID, &action.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, action.Like, action.Dislike, nil // user has not liked
		}
		return false, action.Like, action.Dislike, err
	}

	return true, action.Like, action.Dislike, nil
}

func (a *ActionRepoImpl) LikePost(userID, postID uint, typ string) error {
	check, liked, disliked, err := a.IsUserAlreadyLiked(userID, postID, "post")
	if err != nil {
		return errors.New("error checking if user has already liked")
	}

	if check {
		if !liked && !disliked {
			if typ == "like" {
				_, err = a.db.GetDB().Exec(`UPDATE actions SET like = ?, dislike = ? WHERE user_id = ? AND post_id = ?`, true, false, userID, postID)
			} else if typ == "dislike" {
				_, err = a.db.GetDB().Exec(`UPDATE actions SET like = ?, dislike = ? WHERE user_id = ? AND post_id = ?`, false, true, userID, postID)
			}
		} else if liked && !disliked {
			if typ == "like" {
				_, err = a.db.GetDB().Exec(`UPDATE actions SET like = ?, dislike = ? WHERE user_id = ? AND post_id = ?`, false, false, userID, postID)
			} else if typ == "dislike" {
				_, err = a.db.GetDB().Exec(`UPDATE actions SET like = ?, dislike = ? WHERE user_id = ? AND post_id = ?`, false, true, userID, postID)
			}
		} else if !liked && disliked {
			if typ == "like" {
				_, err = a.db.GetDB().Exec(`UPDATE actions SET like = ?, dislike = ? WHERE user_id = ? AND post_id = ?`, true, false, userID, postID)
			} else if typ == "dislike" {
				_, err = a.db.GetDB().Exec(`UPDATE actions SET like = ?, dislike = ? WHERE user_id = ? AND post_id = ?`, false, false, userID, postID)
			}
		} else {
			return errors.New("error liking post")
		}
	} else {
		if typ == "like" {
			_, err = a.db.GetDB().Exec(`INSERT INTO actions (like, post_id, user_id) VALUES (?, ?, ?)`, true, postID, userID)
		} else if typ == "dislike" {
			_, err = a.db.GetDB().Exec(`INSERT INTO actions (dislike, post_id, user_id) VALUES (?, ?, ?)`, true, postID, userID)
		}
	}
	return nil
}

func (a *ActionRepoImpl) LikeComment(userID, commentID uint, typ string) error {
	check, liked, disliked, err := a.IsUserAlreadyLiked(userID, commentID, "comment")
	if err != nil {
		return errors.New("error checking if user has already liked")
	}

	if check {
		if !liked && !disliked {
			if typ == "like" {
				_, err = a.db.GetDB().Exec(`UPDATE actions SET like = ?, dislike = ? WHERE user_id = ? AND comment_id = ?`, true, false, userID, commentID)
			} else if typ == "dislike" {
				_, err = a.db.GetDB().Exec(`UPDATE actions SET like = ?, dislike = ? WHERE user_id = ? AND comment_id = ?`, false, true, userID, commentID)
			}
		} else if liked && !disliked {
			if typ == "like" {
				_, err = a.db.GetDB().Exec(`UPDATE actions SET like = ?, dislike = ? WHERE user_id = ? AND comment_id = ?`, false, false, userID, commentID)
			} else if typ == "dislike" {
				_, err = a.db.GetDB().Exec(`UPDATE actions SET like = ?, dislike = ? WHERE user_id = ? AND comment_id = ?`, false, true, userID, commentID)
			}
		} else if !liked && disliked {
			if typ == "like" {
				_, err = a.db.GetDB().Exec(`UPDATE actions SET like = ?, dislike = ? WHERE user_id = ? AND comment_id = ?`, true, false, userID, commentID)
			} else if typ == "dislike" {
				_, err = a.db.GetDB().Exec(`UPDATE actions SET like = ?, dislike = ? WHERE user_id = ? AND comment_id = ?`, false, false, userID, commentID)
			}
		} else {
			return errors.New("error liking comment")
		}
	} else {
		if typ == "like" {
			_, err = a.db.GetDB().Exec(`INSERT INTO actions (like, comment_id, user_id) VALUES (?, ?, ?)`, true, commentID, userID)
		} else if typ == "dislike" {
			_, err = a.db.GetDB().Exec(`INSERT INTO actions (dislike, comment_id, user_id) VALUES (?, ?, ?)`, true, commentID, userID)
		}
	}
	return nil
}
