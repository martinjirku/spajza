package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/martinjirku/zasobar/domain"
	"github.com/martinjirku/zasobar/usecases"
)

type StorageItemRepository struct {
	db *sql.DB
	us *usecases.UnitService
}

func NewStorageItemRepository(db *sql.DB) StorageItemRepository {
	us := usecases.NewUnitService()
	return StorageItemRepository{db: db, us: &us}
}

func (s *StorageItemRepository) Create(ctx context.Context, storageItem domain.NewStorageItem) (domain.StorageItem, error) {
	unit, err := findUnit(s.us.ListAll(), storageItem.Unit)
	if err != nil {
		return domain.StorageItem{}, err
	}

	res := domain.StorageItem{
		Title:          storageItem.Title,
		BaselineAmount: storageItem.Amount,
		CurrentAmount:  storageItem.Amount,
		CategoryId:     storageItem.CategoryId,
		StoragePlaceId: storageItem.StoragePlaceId,
		Quantity:       domain.Quantity(unit.Quantity),
		Unit:           unit.Name,
		ExpirationDate: storageItem.ExpirationDate,
	}

	query := "INSERT INTO storage_items (created_at, updated_at, title," +
		"storage_place_id, category_id, baseline_amount, current_amount," +
		"quantity, unit, expiration_date) VALUES (?,?,?,?,?,?,?,?,?,?)"
	result, err := s.db.ExecContext(ctx, query, time.Now(), time.Now(), res.Title,
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
func (s *StorageItemRepository) UpdateColumn(ctx context.Context, id uint, fieldName string, fieldValue interface{}) error {
	allowedFields := map[string]string{"title": "title", "currentAmount": "current_amount", "unit": "unit", "expirationDate": "expiration_date"}
	fieldToChange := allowedFields[fieldName]
	if fieldToChange == "" {
		return domain.ErrorWrongField
	}
	query := fmt.Sprintf("UPDATE storage_items SET %s=? WHERE storage_item_id=?", fieldToChange)
	result, err := s.db.ExecContext(ctx, query, fieldValue, id)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (s *StorageItemRepository) List(ctx context.Context) ([]domain.StorageItem, error) {
	query := "SELECT storage_item_id, title, storage_place_id, category_id, baseline_amount, current_amount, quantity, unit, expiration_date FROM storage_items"
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp := make([]domain.StorageItem, 0)

	for rows.Next() {
		s := domain.StorageItem{}
		rows.Scan(&s.StorageItemId, &s.Title, &s.StoragePlaceId, &s.CategoryId, &s.BaselineAmount, &s.CurrentAmount, &s.Quantity, &s.Unit, &s.ExpirationDate)
		resp = append(resp, s)
	}
	return resp, nil
}

func (s *StorageItemRepository) GetStorageItemById(ctx context.Context, storageItemId uint) (domain.StorageItem, error) {
	query := "SELECT storage_item_id, title, storage_place_id, category_id, baseline_amount, current_amount, quantity, unit, expiration_date FROM storage_items WHERE storage_item_id=?"
	si := domain.StorageItem{}
	row := s.db.QueryRowContext(ctx, query, storageItemId)
	if row.Err() != nil {
		return si, row.Err()
	}
	row.Scan(&si.StorageItemId, &si.Title, &si.StoragePlaceId, &si.CategoryId,
		&si.BaselineAmount, &si.CurrentAmount, &si.Quantity, &si.Unit, &si.ExpirationDate)
	return si, nil
}

func (s *StorageItemRepository) GetStorageConsumptionById(ctx context.Context, storageItemId uint) ([]domain.StorageItemConsumption, error) {
	query := "SELECT storage_item_consumption_id, normalized_amount, unit, storage_item_id FROM storage_consumptions WHERE storage_item_id=?"
	sic := []domain.StorageItemConsumption{}
	rows, err := s.db.QueryContext(ctx, query, storageItemId)
	if err != nil {
		return sic, err
	}
	defer rows.Close()
	for rows.Next() {
		c := domain.StorageItemConsumption{}
		rows.Scan(&c.StorageItemConsumptionId, &c.NormalizedAmount, &c.Unit, &c.StorageItemId)
		sic = append(sic, c)
	}
	return sic, nil
}

func (s *StorageItemRepository) AddStorageConsumption(ctx context.Context, sc domain.StorageItemConsumption) (domain.StorageItemConsumption, error) {
	query := "INSERT INTO storage_consumptions (created_at, updated_at, normalized_amount, unit, storage_item_id) VALUES (?,?,?,?,?)"
	result, err := s.db.ExecContext(ctx, query, time.Now(), time.Now(), sc.NormalizedAmount, sc.Unit, sc.StorageItemId)
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
