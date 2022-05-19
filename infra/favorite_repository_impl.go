package infra

import (
	"database/sql"
	"fmt"
	"github.com/empenguin1186/go-api-sample/domain/model"
	"log"
	"time"
)

type FavoriteRepositoryImpl struct {
	db *sql.DB
	tx *sql.Tx
}

func NewFavoriteRepositoryImpl(db *sql.DB, tx *sql.Tx) *FavoriteRepositoryImpl {
	return &FavoriteRepositoryImpl{
		db: db,
		tx: tx,
	}
}

func (f *FavoriteRepositoryImpl) SelectById(tweetId string) ([]model.Favorite, error) {
	query := fmt.Sprintf("select tweet_id, to_json(registered_at) from favorite where tweet_id = '%s'", tweetId)
	data, err := f.db.Query(query)
	if err != nil {
		return []model.Favorite{}, err
	}
	favorite := model.Favorite{}
	var result []model.Favorite
	for data.Next() {
		err = data.Scan(&favorite.TweetId, &favorite.RegisteredAt)
		if err != nil {
			return []model.Favorite{}, err
		}
		result = append(result, favorite)
	}
	return result, nil
}

func (f *FavoriteRepositoryImpl) InsertFavorite(tweetId string) error {
	now := time.Now().Format("2006-01-02T15:04:05-07:00")
	query, err := f.tx.Prepare("insert into favorite (tweet_id, registered_at) values ($1, $2)")
	if err != nil {
		log.Printf("datastore insert favorite error; %v", err)
		return nil
	}
	defer query.Close()

	_, err = query.Exec(tweetId, now)
	if err != nil {
		return fmt.Errorf("database failed")
	}
	return nil
}
