package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/martinjirku/zasobar/entity"
)

type StoragePlaceRepository struct {
	ctx context.Context
	db  *sql.DB
}

func NewStoragePlaceRepository(ctx context.Context, db *sql.DB) *StoragePlaceRepository {
	return &StoragePlaceRepository{ctx, db}
}

func (s *StoragePlaceRepository) Create(storagePlace entity.StoragePlace) (entity.StoragePlace, error) {
	storagePlace.CreatedAt = time.Now()
	storagePlace.UpdatedAt = time.Now()
	result, err := s.db.ExecContext(s.ctx, "INSERT INTO storage_places(created_at, updated_at, title, code) VALUES (?,?,?,?)", storagePlace.CreatedAt, storagePlace.UpdatedAt, storagePlace.Title, storagePlace.Code)
	if err != nil {
		return storagePlace, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return storagePlace, err
	}
	storagePlace.StoragePlaceId = int32(id)
	return storagePlace, nil
}

func (s *StoragePlaceRepository) Get(storagePlaceId int32) (entity.StoragePlace, error) {
	var storagePlace = entity.StoragePlace{StoragePlaceId: storagePlaceId}
	err := s.db.QueryRowContext(s.ctx, "SELECT created_at, updated_at, title, code FROM storage_places WHERE deleted_at IS NULL && storage_place_id=?", storagePlaceId).
		Scan(&storagePlace.CreatedAt, &storagePlace.UpdatedAt, &storagePlace.Title, &storagePlace.Code)
	return storagePlace, err
}

func (s *StoragePlaceRepository) List() ([]entity.StoragePlace, error) {
	storagePlaces := []entity.StoragePlace{}
	row, err := s.db.QueryContext(s.ctx, "SELECT storage_place_id, created_at, updated_at, title, code FROM storage_places WHERE deleted_at IS NULL")
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

func (s *StoragePlaceRepository) Update(storagePlace entity.StoragePlace) (entity.StoragePlace, error) {
	storagePlace.UpdatedAt = time.Now()
	result, err := s.db.ExecContext(s.ctx, "UPDATE storage_places SET updated_at=?,title=?,code=? WHERE storage_place_id=?",
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
	err = s.db.QueryRowContext(s.ctx, "SELECT created_at FROM storage_places WHERE storage_place_id=?", storagePlace.StoragePlaceId).Scan(&storagePlace.CreatedAt)
	return storagePlace, err
}

func (s *StoragePlaceRepository) Delete(storagePlaceId int32) error {
	result, err := s.db.ExecContext(s.ctx, "UPDATE storage_places SET deleted_at=? WHERE storage_place_id=?", time.Now(), storagePlaceId)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if affected == 0 {
		return errors.New("nothing updated")
	}
	return err
}
