package storage

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type StoragePlacesService struct {
	db *sql.DB
}

func NewStoragePlacesService(db *sql.DB) StoragePlacesService {
	return StoragePlacesService{db}
}

func (s *StoragePlacesService) Create(ctx context.Context, storagePlace StoragePlace) (StoragePlace, error) {
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

func (s *StoragePlacesService) Get(ctx context.Context, storagePlaceId uint) (StoragePlace, error) {
	var storagePlace = StoragePlace{StoragePlaceId: storagePlaceId}
	err := s.db.QueryRowContext(ctx, "SELECT created_at, updated_at, title, code FROM storage_places WHERE deleted_at IS NULL && storage_place_id=?", storagePlaceId).
		Scan(&storagePlace.CreatedAt, &storagePlace.UpdatedAt, &storagePlace.Title, &storagePlace.Code)
	return storagePlace, err
}

func (s *StoragePlacesService) List(ctx context.Context) ([]StoragePlace, error) {
	storagePlaces := []StoragePlace{}
	row, err := s.db.QueryContext(ctx, "SELECT storage_place_id, created_at, updated_at, title, code FROM storage_places WHERE deleted_at IS NULL")
	if err != nil {
		return storagePlaces, err
	}
	defer row.Close()
	for row.Next() {
		storagePlace := StoragePlace{}
		err := row.Scan(&storagePlace.StoragePlaceId, &storagePlace.CreatedAt, &storagePlace.UpdatedAt, &storagePlace.Title, &storagePlace.Code)
		if err != nil {
			return storagePlaces, err
		}
		storagePlaces = append(storagePlaces, storagePlace)
	}
	return storagePlaces, nil
}

func (s *StoragePlacesService) Update(ctx context.Context, storagePlace StoragePlace) (StoragePlace, error) {
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

func (s *StoragePlacesService) Delete(ctx context.Context, storagePlaceId uint) error {
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
