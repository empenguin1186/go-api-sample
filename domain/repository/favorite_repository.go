package repository

import "github.com/empenguin1186/go-api-sample/domain/model"

type FavoriteRepository interface {
	SelectById(tweetId string) ([]model.Favorite, error)
	InsertFavorite(tweetId string) error
}
