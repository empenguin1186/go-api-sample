package model

type Favorite struct {
	// ツイートID
	TweetId string `json:"tweet_id,omitempty"`
	// お気に入りに登録した日時
	RegisteredAt string `json:"registered_at,omitempty"`
}
