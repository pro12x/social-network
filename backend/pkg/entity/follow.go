package entity

type Follow struct {
	ID         uint   `json:"id" db:"id"`
	FollowerID uint   `json:"follower_id" db:"follower_id"`
	FolloweeID uint   `json:"followee_id" db:"followee_id"`
	CreatedAt  string `json:"created_at" db:"created_at"`
}
