package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/martinjirku/zasobar/entity"
	"github.com/martinjirku/zasobar/usecase"
)

type StorageItemRepository struct {
	ctx context.Context
	db  *sql.DB
	us  *usecase.UnitUsecase
}

func NewStorageItemRepository(ctx context.Context, db *sql.DB) StorageItemRepository {
	us := usecase.NewUnitUsecase()
	return StorageItemRepository{ctx, db, &us}
}

func (s *StorageItemRepository) Create(storageItem entity.NewStorageItem) (entity.StorageItem, error) {
	unit, err := findUnit(s.us.ListAll(), storageItem.Unit)
	if err != nil {
		return entity.StorageItem{}, err
	}

	res := entity.StorageItem{
		Title:          storageItem.Title,
		BaselineAmount: storageItem.Amount,
		CurrentAmount:  storageItem.Amount,
		CategoryId:     storageItem.CategoryId,
		StoragePlaceId: storageItem.StoragePlaceId,
		Quantity:       entity.QuantityType(unit.Quantity),
		Unit:           string(unit.Name),
		ExpirationDate: storageItem.ExpirationDate,
	}

	query := "INSERT INTO storage_items (created_at, updated_at, title," +
		"storage_place_id, category_id, baseline_amount, current_amount," +
		"quantity, unit, expiration_date) VALUES (?,?,?,?,?,?,?,?,?,?)"
	result, err := s.db.ExecContext(s.ctx, query, time.Now(), time.Now(), res.Title,
		res.StoragePlaceId, res.CategoryId, res.BaselineAmount, res.CurrentAmount,
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

// Allows to update specified column with the new value. Only
// allowed columns are:
//   - title
//   - currentAmount
//   - unit
//   - expirationDate
func (s *StorageItemRepository) UpdateColumn(id uint, fieldName string, fieldValue interface{}) error {
	allowedFields := map[string]string{"title": "title", "currentAmount": "current_amount", "unit": "unit", "expirationDate": "expiration_date", "storagePlaceId": "storage_place_id"}
	fieldToChange := allowedFields[fieldName]
	if fieldToChange == "" {
		return entity.ErrWrongField
	}
	query := fmt.Sprintf("UPDATE storage_items SET %s=? WHERE storage_item_id=?", fieldToChange)
	result, err := s.db.ExecContext(s.ctx, query, fieldValue, id)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (s *StorageItemRepository) List() ([]entity.StorageItem, error) {
	query := "SELECT storage_item_id, title, storage_place_id, category_id, baseline_amount, current_amount, quantity, unit, expiration_date FROM storage_items"
	rows, err := s.db.QueryContext(s.ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp := make([]entity.StorageItem, 0)

	for rows.Next() {
		s := entity.StorageItem{}
		rows.Scan(&s.StorageItemId, &s.Title, &s.StoragePlaceId, &s.CategoryId, &s.BaselineAmount, &s.CurrentAmount, &s.Quantity, &s.Unit, &s.ExpirationDate)
		resp = append(resp, s)
	}
	return resp, nil
}

func (s *StorageItemRepository) GetStorageItemById(storageItemId uint) (entity.StorageItem, error) {
	query := "SELECT storage_item_id, title, storage_place_id, category_id, baseline_amount, current_amount, quantity, unit, expiration_date FROM storage_items WHERE storage_item_id=?"
	si := entity.StorageItem{}
	row := s.db.QueryRowContext(s.ctx, query, storageItemId)
	if row.Err() != nil {
		return si, row.Err()
	}
	row.Scan(&si.StorageItemId, &si.Title, &si.StoragePlaceId, &si.CategoryId,
		&si.BaselineAmount, &si.CurrentAmount, &si.Quantity, &si.Unit, &si.ExpirationDate)
	return si, nil
}

func (s *StorageItemRepository) GetStorageConsumptionById(storageItemId uint) ([]entity.StorageItemConsumption, error) {
	query := "SELECT storage_item_consumption_id, normalized_amount, unit, storage_item_id FROM storage_consumptions WHERE storage_item_id=?"
	sic := []entity.StorageItemConsumption{}
	rows, err := s.db.QueryContext(s.ctx, query, storageItemId)
	if err != nil {
		return sic, err
	}
	defer rows.Close()
	for rows.Next() {
		c := entity.StorageItemConsumption{}
		rows.Scan(&c.StorageItemConsumptionId, &c.NormalizedAmount, &c.Unit, &c.StorageItemId)
		sic = append(sic, c)
	}
	return sic, nil
}

func (s *StorageItemRepository) AddStorageConsumption(sc entity.StorageItemConsumption) (entity.StorageItemConsumption, error) {
	query := "INSERT INTO storage_consumptions (created_at, updated_at, normalized_amount, unit, storage_item_id) VALUES (?,?,?,?,?)"
	result, err := s.db.ExecContext(s.ctx, query, time.Now(), time.Now(), sc.NormalizedAmount, sc.Unit, sc.StorageItemId)
	if err != nil {
		return sc, err
	}
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return sc, err
	}
	sc.StorageItemConsumptionId = uint(lastInsertedId)
	return sc, nil
}

func findUnit(units []entity.Unit, unitName string) (entity.Unit, error) {
	var unit entity.Unit
	found := false
	for _, u := range units {
		if string(u.Name) == unitName {
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
