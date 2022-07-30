package storage

import (
	"context"
	"database/sql"
	"errors"
	"time"

	goUnits "github.com/bcicen/go-units"
	"github.com/martinjirku/zasobar/units"
)

type StorageService struct {
	db *sql.DB
	us *units.UnitService
}

func NewStorageService(db *sql.DB) StorageService {
	us := units.NewUnitService()
	return StorageService{db: db, us: &us}
}

func (s *StorageService) Create(ctx context.Context, storageItem NewStorageItemRequest) (StorageItemResponse, error) {
	unit, err := findUnit(s.us.ListAll(), storageItem.Unit)
	if err != nil {
		return StorageItemResponse{}, err
	}

	res := StorageItemResponse{
		Title:          storageItem.Title,
		BaselineAmount: storageItem.Amount,
		CurrentAmount:  storageItem.Amount,
		CategoryId:     storageItem.CategoryId,
		StoragePlaceId: storageItem.StoragePlaceId,
		Quantity:       units.Quantity(unit.Quantity),
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

func findUnit(units []goUnits.Unit, unitName string) (goUnits.Unit, error) {
	var unit goUnits.Unit
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
