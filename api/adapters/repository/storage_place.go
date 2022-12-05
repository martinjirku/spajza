package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/martinjirku/zasobar/entity"
)

type StoragePlaceRepository struct {
	db *sql.DB
}

func NewStoragePlaceRepository(db *sql.DB) *StoragePlaceRepository {
	return &StoragePlaceRepository{db}
}

func (s *StoragePlaceRepository) Create(ctx context.Context, storagePlace entity.StoragePlace) (entity.StoragePlace, error) {
	storagePlace.CreatedAt = time.Now()
	storagePlace.UpdatedAt = time.Now()
	result, err := s.db.ExecContext(ctx, "INSERT INTO storage_places(created_at, updated_at, title, code) VALUES (?,?,?,?)", storagePlace.CreatedAt, storagePlace.UpdatedAt, storagePlace.Title, storagePlace.Code)
	if err != nil {
		return storagePlace, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return storagePlace, err
	}
	storagePlace.StoragePlaceId = uint(id)
	return storagePlace, nil
}

func (s *StoragePlaceRepository) Get(ctx context.Context, storagePlaceId uint) (entity.StoragePlace, error) {
	var storagePlace = entity.StoragePlace{StoragePlaceId: storagePlaceId}
	err := s.db.QueryRowContext(ctx, "SELECT created_at, updated_at, title, code FROM storage_places WHERE deleted_at IS NULL && storage_place_id=?", storagePlaceId).
		Scan(&storagePlace.CreatedAt, &storagePlace.UpdatedAt, &storagePlace.Title, &storagePlace.Code)
	return storagePlace, err
}

func (s *StoragePlaceRepository) List(ctx context.Context) ([]entity.StoragePlace, error) {
	storagePlaces := []entity.StoragePlace{}
	row, err := s.db.QueryContext(ctx, "SELECT storage_place_id, created_at, updated_at, title, code FROM storage_places WHERE deleted_at IS NULL")
	if err != nil {
		return storagePlaces, err
	}
	defer row.Close()
	for row.Next() {
		storagePlace := entity.StoragePlace{}
		err := row.Scan(&storagePlace.StoragePlaceId, &storagePlace.CreatedAt, &storagePlace.UpdatedAt, &storagePlace.Title, &storagePlace.Code)
		if err != nil {
			return storagePlaces, err
		}
		storagePlaces = append(storagePlaces, storagePlace)
	}
	return storagePlaces, nil
}

func (s *StoragePlaceRepository) Update(ctx context.Context, storagePlace entity.StoragePlace) (entity.StoragePlace, error) {
	storagePlace.UpdatedAt = time.Now()
	result, err := s.db.ExecContext(ctx, "UPDATE storage_places SET updated_at=?,title=?,code=? WHERE storage_place_id=?",
		storagePlace.UpdatedAt, storagePlace.Title, storagePlace.Code, storagePlace.StoragePlaceId)
	if err != nil {
		return storagePlace, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return storagePlace, err
	}
	if affected != 1 {
		return storagePlace, errors.New("nothing updated")
	}
	err = s.db.QueryRowContext(ctx, "SELECT created_at FROM storage_places WHERE storage_place_id=?", storagePlace.StoragePlaceId).Scan(&storagePlace.CreatedAt)
	return storagePlace, err
}

func (s *StoragePlaceRepository) Delete(ctx context.Context, storagePlaceId uint) error {
	result, err := s.db.ExecContext(ctx, "UPDATE storage_places SET deleted_at=? WHERE storage_place_id=?", time.Now(), storagePlaceId)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if affected == 0 {
		return errors.New("nothing updated")
	}
	return err
}
