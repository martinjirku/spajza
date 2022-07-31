package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/martinjirku/zasobar/domain"
	"github.com/martinjirku/zasobar/usecases"
)

type StorageService struct {
	db *sql.DB
	us *usecases.UnitService
}

func NewStorageService(db *sql.DB) StorageService {
	us := usecases.NewUnitService()
	return StorageService{db: db, us: &us}
}

func (s *StorageService) Create(ctx context.Context, storageItem domain.NewStorageItemRequest) (domain.StorageItemResponse, error) {
	unit, err := findUnit(s.us.ListAll(), storageItem.Unit)
	if err != nil {
		return domain.StorageItemResponse{}, err
	}

	res := domain.StorageItemResponse{
		Title:          storageItem.Title,
		BaselineAmount: storageItem.Amount,
		CurrentAmount:  storageItem.Amount,
		CategoryId:     storageItem.CategoryId,
		StoragePlaceId: storageItem.StoragePlaceId,
		Quantity:       domain.Quantity(unit.Quantity),
		Unit:           unit.Name,
		ExpirationDate: storageItem.ExpirationDate,
	}

	query := "INSERT INTO storage_items (" +
		"created_at, updated_at, title," +
		"storage_place_id, category_id," +
		"baseline_amount, current_amount," +
		"quantity, unit, expiration_date)" +
		"VALUES (?,?,?,?,?,?,?,?,?,?)"
	result, err := s.db.ExecContext(ctx, query,
		time.Now(), time.Now(), res.Title,
		res.StoragePlaceId, res.CategoryId,
		res.BaselineAmount, res.CurrentAmount,
		res.Quantity, res.Unit, res.ExpirationDate)
	if err != nil {
		return res, err
	}

	storageItemId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}
	res.StorageItemId = uint(storageItemId)
	return res, nil
}

func findUnit(units []domain.Unit, unitName string) (domain.Unit, error) {
	var unit domain.Unit
	found := false
	for _, u := range units {
		if u.Name == unitName {
			unit = u
			found = true
			break
		}
	}
	if !found {
		return unit, errors.New("not valid unit")
	}
	return unit, nil
}
