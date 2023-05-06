package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/martinjirku/zasobar/adapters/repository/client"
	"github.com/martinjirku/zasobar/entity"
	"github.com/martinjirku/zasobar/pkg/sqlnull"
)

type StoragePlaceRepository struct {
	ctx     context.Context
	queries *client.Queries
	db      *sql.DB
}

func NewStoragePlaceRepository(ctx context.Context, db *sql.DB) *StoragePlaceRepository {
	queries := client.New(db)
	return &StoragePlaceRepository{ctx, queries, db}
}

func (s *StoragePlaceRepository) Create(storagePlace entity.StoragePlace) (entity.StoragePlace, error) {
	storagePlace.CreatedAt = time.Now()
	storagePlace.UpdatedAt = time.Now()
	storagePlaceId, err := s.queries.CreateStoragePlace(s.ctx, &client.CreateStoragePlaceParams{
		Title: sqlnull.FromString(storagePlace.Title),
		Code:  sqlnull.FromString(storagePlace.Code),
	})
	if err != nil {
		return storagePlace, err
	}
	storagePlace.StoragePlaceId = int32(storagePlaceId)
	return storagePlace, nil
}

func (s *StoragePlaceRepository) Get(storagePlaceId int32) (entity.StoragePlace, error) {
	var storagePlace = entity.StoragePlace{StoragePlaceId: storagePlaceId}
	resp, err := s.queries.GetStoragePlaceById(s.ctx, storagePlaceId)
	if err != nil {
		return storagePlace, err
	}
	storagePlace.Code = resp.Code.String
	storagePlace.Title = resp.Title.String
	storagePlace.CreatedAt = resp.CreatedAt
	storagePlace.UpdatedAt = resp.UpdatedAt
	return storagePlace, err
}

func (s *StoragePlaceRepository) List() ([]entity.StoragePlace, error) {
	results, err := s.queries.ListStoragePlaces(s.ctx)
	if err != nil {
		return nil, err
	}
	storagePlaces := make([]entity.StoragePlace, len(results))
	for i := range results {
		storagePlaces[i] = entity.StoragePlace{
			StoragePlaceId: results[i].StoragePlaceID,
			CreatedAt:      results[i].CreatedAt,
			UpdatedAt:      results[i].UpdatedAt,
			Title:          results[i].Title.String,
			Code:           results[i].Code.String,
		}
	}
	return storagePlaces, nil
}

func (s *StoragePlaceRepository) Update(storagePlace entity.StoragePlace) (entity.StoragePlace, error) {
	err := s.queries.UpdateStoragePlace(s.ctx, &client.UpdateStoragePlaceParams{
		Title:          sqlnull.FromString(storagePlace.Title),
		Code:           sqlnull.FromString(storagePlace.Code),
		StoragePlaceID: storagePlace.StoragePlaceId,
	})
	if err != nil {
		return storagePlace, err
	}
	result, err := s.Get(storagePlace.StoragePlaceId)
	if err != nil {
		return storagePlace, err
	}
	return result, err
}

func (s *StoragePlaceRepository) Delete(storagePlaceId int32) error {
	return s.queries.DeleteStoragePlace(s.ctx, storagePlaceId)
}
