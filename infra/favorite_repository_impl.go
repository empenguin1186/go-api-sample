package infra

import (
	"fmt"
	"github.com/empenguin1186/go-api-sample/domain/model"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type FavoriteRepositoryImpl struct {
	dataSource *DataSource
}

func NewFavoriteRepositoryImpl(dataSource *DataSource) *FavoriteRepositoryImpl {
	return &FavoriteRepositoryImpl{
		dataSource: dataSource,
	}
}

func (f *FavoriteRepositoryImpl) SelectById(tweetId string) ([]model.Favorite, error) {
	data, err := f.dataSource.Db.Query("select tweet_id, registered_at from favorite where tweet_id = $1", tweetId)
	if err != nil {
		log.Printf("datastore get favorite error; %v", err)
		return []model.Favorite{}, err
	}
	favorite := model.Favorite{}
	var result []model.Favorite
	for data.Next() {
		err = data.Scan(&favorite.TweetId, &favorite.RegisteredAt)
		if err != nil {
			log.Printf("datastore parse favorite error; %v", err)
			return []model.Favorite{}, err
		}
		result = append(result, favorite)
	}
	return result, nil
}

func (f *FavoriteRepositoryImpl) InsertFavorite(tweetId string) error {
	f.dataSource.Begin()
	defer f.dataSource.Commit()

	stmt, err := f.dataSource.Tx.Prepare("insert into favorite (tweet_id, registered_at) values ($1, $2)")
	if err != nil {
		log.Printf("failed to create sql statement. error = %v", err)
		return fmt.Errorf("failed to create sql statement. caused by %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(tweetId, time.Now().Format("Mon, 02 Jan 2006 15:04:05 JST"))
	if err != nil {
		log.Printf("failed to execute sql statement. error = %v", err)
		return fmt.Errorf("failed to execute sql statement. caused by %v", err)
	}
	return nil
}
