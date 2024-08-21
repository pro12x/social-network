package interfaces

type ActionRepo interface {
	IsUserAlreadyLiked(userID, actionID uint, nature string) (bool, bool, bool, error)
	//IsUserLikedComment(userID, commentID uint) (bool, error)
	//LikePost(userID, postID uint) error
	//UnlikePost(userID, postID uint) error
	//LikeComment(userID, commentID uint) error
	//UnlikeComment(userID, commentID uint) error
	//LikeReply(userID, replyID uint) error
	//UnlikeReply(userID, replyID uint) error
	//LikeSubcomment(userID, subcommentID uint) error
	//UnlikeSubcomment(userID, subcommentID uint) error
	//CountAllLikes() (uint, error)
	//CountAllDislikes() (uint, error)
	//CountAllActions() (uint, error)
	//FindLikeByUserIDAndPostID(userID, postID uint) (bool, error)
	//FindLikeByUserIDAndCommentID(userID, commentID uint) (bool, error)
	//FindLikeByUserIDAndReplyID(userID, replyID uint) (bool, error)
	//FindLikeByUserIDAndSubcommentID(userID, subcommentID uint) (bool, error)
	//FindDislikeByUserIDAndPostID(userID, postID uint) (bool, error)
}
