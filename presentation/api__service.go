/*
 * サンプルAPI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package presentation

import (
	"context"
	"errors"
	"github.com/empenguin1186/go-api-sample/domain/repository"
	"net/http"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
	favoriteRepository repository.FavoriteRepository
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService(favoriteRepository repository.FavoriteRepository) DefaultApiServicer {
	return &DefaultApiService{favoriteRepository}
}

// AdminGetTweetFavoriteList - ツイートごとのお気に入り数を取得
func (s *DefaultApiService) AdminGetTweetFavoriteList(ctx context.Context) (ImplResponse, error) {
	// TODO - update AdminGetTweetFavoriteList with the required logic for this service method.
	// Add api__service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TweetFavoriteList{}) or use other options such as http.Ok ...
	//return Response(200, TweetFavoriteList{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AdminGetTweetFavoriteList method not implemented")
}

// DeleteFavorite - ツイート店舗をお気に入りから削除
func (s *DefaultApiService) DeleteFavorite(ctx context.Context, tweetId int32) (ImplResponse, error) {
	// TODO - update DeleteFavorite with the required logic for this service method.
	// Add api__service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteFavorite method not implemented")
}

// GetFavorite - お気に入りに登録したツイートIDを取得
func (s *DefaultApiService) GetFavorite(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetFavorite with the required logic for this service method.
	// Add api__service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, FavoriteList{}) or use other options such as http.Ok ...
	//return Response(200, FavoriteList{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetFavorite method not implemented")
}

// PostFavorite - ツイート店舗をお気に入りに登録
func (s *DefaultApiService) PostFavorite(ctx context.Context, tweetId TweetId) (ImplResponse, error) {
	// TODO - update PostFavorite with the required logic for this service method.
	// Add api__service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostFavorite method not implemented")
}
